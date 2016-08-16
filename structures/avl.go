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
        avl.parent = &comparator.right
        comparator.right = &avl
    } else {
        avl.parent = &comparator.left
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

    fmt.Println("")
    fmt.Println("")
}

func (tree *AVL) height() int {
    leftHeight, rightHeight := 0, 0

    if tree.left != nil {
        leftHeight = tree.left.height()
    }

    if tree.right != nil {
        rightHeight = tree.rigth.height()
    }

    if leftHeight > rightHeight {
        return leftHeight
    } else {
        return rightHeight
    }
}

func (tree *AVL) rotateLeft {
    // do something
}

func (tree *AVL) rotateRight {
    // do something
}
