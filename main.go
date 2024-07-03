package main

import (
    "fmt"
    "github.com/Illedran/go_cross_reproducible/proto"
)

func main() {
    fmt.Printf("Hello world, %+v\n", proto.Msg{Content: "test"})
}
