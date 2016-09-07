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
    adjust()
    height() int
    heights() (int, int)
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
    tree = tree.adjust()

    return tree
}

func (tree *AVL) heights() (int, int) {
    leftHeight, rightHeight := 0, 0

    if tree.left != nil {
        leftHeight = 1 + tree.left.height()
    }

    if tree.right != nil {
        rightHeight = 1 + tree.right.height()
    }

    return leftHeight, rightHeight
}

func (tree *AVL) height() int {
    leftHeight, rightHeight := tree.heights()

    if leftHeight > rightHeight {
        return leftHeight
    } else {
        return rightHeight
    }
}

func (tree *AVL) adjust() *AVL {
    leftHeight, rightHeight := tree.heights()
    fmt.Println("Heights:", leftHeight, rightHeight)

    if (leftHeight - rightHeight) > 1 {
        tree = tree.rotateRight()
    }

    if (rightHeight - leftHeight) > 1 {
        tree = tree.rotateLeft()
    }

    return tree
}

func (tree *AVL) rotateLeft() *AVL {
    fmt.Println("Rotating left")
    leftHeight, rightHeight := tree.right.heights()

    if (leftHeight > rightHeight) {
        // double shift, go right first
        tree.right = tree.right.rotateRight()
    }

    tmp := tree
    tree = tmp.right
    tmp.right = tree.left
    tree.left = tmp

    return tree
}

func (tree *AVL) rotateRight() *AVL {
    fmt.Println("Rotating right")
    leftHeight, rightHeight := tree.left.heights()

    if (rightHeight > leftHeight) {
        // double shift, go left first
        tree.left = tree.left.rotateLeft()
    }

    tmp := tree
    tree = tmp.left
    tmp.left = tree.right
    tree.right = tmp

    return tree
}
