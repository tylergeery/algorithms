package main

import (
    "fmt"
    "strconv"
)

type AVL struct {
    value int
    parent, left, right *AVL
}

type AVLTree interface {
    insert(value int) *AVL
    rotateLeft() AVL
    rotateRight() AVL
}

func (tree *AVL) insert(value int) *AVL {
    avl := AVL{value: value}
    comparator := tree

    fmt.Println("New tree value: " + strconv.Itoa(value))
    // sift down the tree
    for comparator.right != nil || comparator.left != nil {
        test := comparator.value

        if value <= test {
            if comparator.left == nil {
                break;
            }

            comparator = comparator.left
        }

        if value > test {
            if comparator.right == nil {
                break;
            }

            comparator = comparator.right
        }
    }

    if value > comparator.value {
        avl.parent = comparator
        comparator.right = &avl
    } else {
        avl.parent = comparator
        comparator.left = &avl
    }

    // check if we need to adjust the tree

    return tree
}

func (tree *AVL) height() int {
    leftHeight, rightHeight := 0, 0

    if tree.left != nil {
        leftHeight = tree.left.height()
    }

    if tree.right != nil {
        rightHeight = tree.right.height()
    }

    if leftHeight > rightHeight {
        return leftHeight
    } else {
        return rightHeight
    }
}

// func (tree *AVL) rotateLeft {
//     fmt.Println("Do something")
//     return tree.right = nil
// }
//
// func (tree *AVL) rotateRight {
//     fmt.Println("Do something")
//     tree.left = nil
// }
