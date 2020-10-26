
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(strings.Join(os.Agrs[1:], " "))
}
