# AAA Test Format — Canonical Reference

**Source**: User-provided screenshot (2026-04-06)

## Structure

Every test function follows strict **Arrange → Act → Assert** with `args.Map`:

```go
func Test_Cov5_IsEmptyPtr_Nil(t *testing.T) {
    // Arrange — (nothing to arrange when input is nil)

    // Act — call function, wrap result in args.Map
    actual := args.Map{"result": stringutil.IsEmptyPtr(nil)}

    // Assert — build expected args.Map, call ShouldBeEqual
    expected := args.Map{"result": true}
    expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns true -- nil", actual)
}
```

```go
func Test_Cov5_IsEmptyPtr_Empty(t *testing.T) {
    // Arrange
    s := ""

    // Act
    actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}

    // Assert
    expected := args.Map{"result": true}
    expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns true -- empty string", actual)
}
```

```go
func Test_Cov5_IsContainsPtrSimple_CaseInsensitive(t *testing.T) {
    // Arrange
    lines := []string{"Hello"}

    // Act
    actual := args.Map{"result": stringutil.IsContainsPtrSimple(&lines, "hello", 0, false)}

    // Assert
    expected := args.Map{"result": true}
    expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns true -- case insensitive", actual)
}
```

## Rules

1. **Always use `args.Map`** for both `actual` and `expected` — never use raw `if` checks or `t.Errorf`
2. **Key naming**: use `"result"` for single-return functions; use descriptive keys for multi-value returns
3. **`ShouldBeEqual` signature**: `expected.ShouldBeEqual(t, index, "description", actual)`
   - `t` — test context
   - `index` — always `0` for single-case tests
   - description — human-readable, includes function name and scenario (e.g., `"IsEmptyPtr returns true -- nil"`)
   - `actual` — the actual args.Map
4. **Arrange section**: declare inputs/setup variables; skip if trivial (e.g., nil input)
5. **Act section**: single line calling the function under test, result wrapped in `args.Map`
6. **Assert section**: build expected `args.Map`, then call assertion method
7. **No comments needed** for AAA sections in simple tests — the structure itself is self-documenting
8. **Description format**: `"FunctionName returns/does X -- scenario"` with double-dash separator for scenario
