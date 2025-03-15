package main

import "fmt"

func genap(n int) int {
    if n == 0 {
        return 0
    }
    var i int
    if n%2 == 0 {
        i = 1
    }
    return i + genap(n/10)
}

func main() {
    var masukan int
    fmt.Scan(&masukan)
    fmt.Print(genap(masukan))
}
