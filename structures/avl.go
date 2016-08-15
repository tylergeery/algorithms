package main

import (
    "fmt"
    "strconv"
)

type AVL struct {
    value int
    height int
    left *AVL
    right *AVL
}

type AVLTree interface {
    insert(value int) *AVL
    rotateLeft() AVL
    rotateRight() AVL
}

func (tree *AVL) insert(value int) *AVL {
    avl := AVL{value: value, height: 0}
    comparator := tree

    fmt.Println("New tree value: " + strconv.Itoa(value))
    // sift down the tree
    for comparator.right != nil && comparator.left != nil {
        comparator.height++
        test := comparator.value

        if value <= test {
            comparator = comparator.left
            fmt.Println("Tree.left")
        }

        if value > test {
            comparator = comparator.right
            fmt.Println("Tree.right")
        }
    }

    if value > comparator.value {
        comparator.right = &avl
    } else {
        comparator.left = &avl
    }

    // check if we need to adjust the tree

    return tree
}

func (tree *AVL) printTree() {
    var oldValues []*AVL
    oldValues = append(oldValues, tree)

    for len(oldValues) > 0 {
        iterator := oldValues
        oldValues = make([]*AVL, 0)

        for _,t := range iterator {
            fmt.Printf("%d ", t.value)

            if t.left != nil {
                oldValues = append(oldValues, t.left)
            }

            if t.right != nil {
                oldValues = append(oldValues, t.right)
            }
        }

        fmt.Println("")
    }
}

// func (tree *AVL) rotateLeft {
//     // do something
// }
//
// func (tree *AVL) rotateRight {
//     // do something
// }
