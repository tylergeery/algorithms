package main

/**
 * Run with go run buildtree.go avl.go
 */
import (
    "math/rand"
)

func main() {
    i := 25
    avl := new(AVL)

    for i > 0 {
        i--
        r := rand.Intn(100)

        if avl.value == 0 {
            avl.value = r
        } else {
            avl = avl.insert(r)
        }

        avl.printTree()
    }
}
