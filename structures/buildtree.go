package main

import (
    "math/rand"
)

func main() {
    i := 7
    avl := new(AVL)

    for i > 0 {
        i--
        r := rand.Intn(100)

        if avl.value == 0 {
            avl.value = r
            avl.height = 0
        } else {
            avl = avl.insert(r)
        }

        avl.printTree()
    }
}
