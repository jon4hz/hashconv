# hashconv

Convert hashrate to a human readable form

## Example 
```go
package main

import (
    "fmt"

    "github.com/jon4hz/hashconv"
)

func main(){
    fmt.Println(hashconv.Format(1000000)) // returns 1 MH
    
    fmt.Println(hashconv.Format(1024)) // returns 1 kH
}
```