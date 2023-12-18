package main

import (
	"fmt"
	"os"
)
func main() {
    
    fmt.Println("hello world")

    var args = os.Args
    var stringMap = make(map[string]int)

    for _, arg := range args {
        stringMap[arg]++
    }

fmt.Print(stringMap)
}
