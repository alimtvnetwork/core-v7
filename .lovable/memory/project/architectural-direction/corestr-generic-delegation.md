# corestr → coregeneric Delegation

## Direction

Methods in `corestr` should delegate their logic to `coregeneric` to reduce duplication. The exported `corestr` method signatures must remain unchanged — only the internal implementation should be refactored.

## Example Pattern

```go
// Keep this exported method in corestr
func (c *Collection) ListPtr() *[]string {
    // Delegate to coregeneric internally
    return coregeneric.ToListPtr(c.items)
}
```

## Rules

- ✅ Keep all existing exported `corestr` signatures
- ✅ Move shared logic into `coregeneric`
- ✅ Have `corestr` methods call `coregeneric` internally
- ❌ Never remove `corestr` exported methods
