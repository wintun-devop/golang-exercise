### Go Variable Rules
- Use var keyword for explicit declaration:
```
var age int = 30
```
- Type inference with := shorthand (only inside functions):
```
name := "Win"
```
- Variables must be used—Go doesn’t allow unused variables.
- Zero values: If you declare a variable without initializing, Go assigns a default:
 - int → 0
 - string → ""
 - bool → false
 - pointer → nil

### Constant Rules
 - Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
 - Constants can be declared both inside and outside of a function
### Constant Types
- Typed constants
- Untyped constants