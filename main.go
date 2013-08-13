package main

import "fmt"

func main() {
    go runWSServer()
    go serveWeb()

    fmt.Scanln(new(string))
}
