#!/usr/bin/env python3
"""
Fix 'no new variables on left side of :=' errors across all test files.

Usage:
    python3 scripts/fix_redeclaration.py [--dry-run]

Parses blocked-packages output to find exact file:line locations,
then replaces ':=' with '=' on those lines.
"""

import os
import re
import sys
import glob

def find_errors_from_blocked(blocked_path):
    """Parse blocked-packages file to extract file:line pairs."""
    fixes = {}  # {filepath: set(line_numbers)}
    pattern = re.compile(r'^(.+?\.go):(\d+):\d+: no new variables on left side of :=$')
    
    with open(blocked_path, 'r') as f:
        for line in f:
            line = line.strip()
            m = pattern.match(line)
            if m:
                filepath = m.group(1)
                lineno = int(m.group(2))
                if filepath not in fixes:
                    fixes[filepath] = set()
                fixes[filepath].add(lineno)
    
    return fixes

def scan_all_test_files(root_dir):
    """Scan all _test.go files for := that redeclares existing variables."""
    fixes = {}
    
    for dirpath, dirnames, filenames in os.walk(root_dir):
        for fname in filenames:
            if not fname.endswith('_test.go'):
                continue
            
            filepath = os.path.join(dirpath, fname)
            try:
                with open(filepath, 'r') as f:
                    lines = f.readlines()
            except:
                continue
            
            # Track declared variables per function scope
            in_func = False
            func_vars = set()
            
            for i, line in enumerate(lines, 1):
                stripped = line.strip()
                
                # Simple function boundary detection
                if stripped.startswith('func ') or stripped.startswith('func('):
                    in_func = True
                    func_vars = set()
                
                # Check for := pattern
                m = re.match(r'^\s*(\w+)\s*:=\s*', stripped)
                if m:
                    var_name = m.group(1)
                    if var_name in func_vars:
                        # This is a redeclaration
                        if filepath not in fixes:
                            fixes[filepath] = set()
                        fixes[filepath].add(i)
                    else:
                        func_vars.add(var_name)
                
                # Also track = assignments
                m2 = re.match(r'^\s*(\w+)\s*=\s*', stripped)
                if m2:
                    func_vars.add(m2.group(1))
    
    return fixes

def apply_fixes(fixes, dry_run=False):
    """Replace ':=' with '=' on the specified lines."""
    total_fixed = 0
    files_fixed = 0
    
    for filepath, line_numbers in sorted(fixes.items()):
        if not os.path.exists(filepath):
            print(f"  SKIP (not found): {filepath}")
            continue
        
        with open(filepath, 'r') as f:
            lines = f.readlines()
        
        changed = False
        for lineno in sorted(line_numbers):
            idx = lineno - 1
            if idx < 0 or idx >= len(lines):
                continue
            
            old_line = lines[idx]
            # Replace first ':=' with '='
            new_line = old_line.replace(':=', '=', 1)
            if new_line != old_line:
                lines[idx] = new_line
                changed = True
                total_fixed += 1
        
        if changed:
            files_fixed += 1
            if dry_run:
                print(f"  WOULD FIX: {filepath} ({len(line_numbers)} lines)")
            else:
                with open(filepath, 'w') as f:
                    f.writelines(lines)
                print(f"  FIXED: {filepath} ({len(line_numbers)} lines)")
    
    return files_fixed, total_fixed

def main():
    dry_run = '--dry-run' in sys.argv
    
    # Try to find blocked-packages file
    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    
    blocked_files = sorted(glob.glob(os.path.join(root, 'data', 'blocked-packages*.txt')))
    if not blocked_files:
        # Also check current dir
        blocked_files = sorted(glob.glob('blocked-packages*.txt'))
    
    if blocked_files:
        blocked_path = blocked_files[-1]  # latest
        print(f"Using blocked-packages file: {blocked_path}")
        fixes = find_errors_from_blocked(blocked_path)
    else:
        print("No blocked-packages file found. Scanning all test files...")
        fixes = scan_all_test_files(os.path.join(root, 'tests'))
    
    if not fixes:
        print("No ':=' redeclaration errors found.")
        return
    
    total_lines = sum(len(v) for v in fixes.values())
    print(f"Found {total_lines} ':=' errors across {len(fixes)} files")
    
    if dry_run:
        print("\n[DRY RUN - no changes will be made]\n")
    
    files_fixed, lines_fixed = apply_fixes(fixes, dry_run)
    
    mode = "Would fix" if dry_run else "Fixed"
    print(f"\n{mode} {lines_fixed} lines across {files_fixed} files")

if __name__ == '__main__':
    main()
