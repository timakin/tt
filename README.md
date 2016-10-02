TT
==========
Bench marker for local function in Golang source code.

Usage
----------
```
import (
 ...
 "github.com/timakin/tt"
)

func main() {
    defer tt.Untrace(tt.Trace())
}
```
