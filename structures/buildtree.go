package main

/**
 * Run with go run buildtree.go avl.go
 */
import (
    "fmt"
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

// print tree for visual debugging
func (tree *AVL) printTree() {
    var oldValues []*AVL
    oldValues = append(oldValues, tree)
    spaces := 30

    for len(oldValues) > 0 {
        iterator := oldValues
        oldValues = make([]*AVL, 0)
        printSpaces(spaces)

        for _,t := range iterator {
            if t.parent != nil && t.value > t.parent.value {
                printSpaces(1)
            }

            fmt.Printf("%d ", t.value)

            if t.left != nil {
                oldValues = append(oldValues, t.left)
            }

            if t.right != nil {
                oldValues = append(oldValues, t.right)
            }
        }

        fmt.Println("")
        spaces -= 1
    }

    fmt.Println("")
    fmt.Println("")
}

// debugging helper method
func printSpaces(i int) {
    inc := 0
    str := ""

    for inc < i {
        inc++
        str += " "
    }

    fmt.Print(str)
}
