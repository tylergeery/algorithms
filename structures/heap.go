package main

import (
  "bufio"
  "fmt"
  "strconv"
  "os"
  "strings"
  "math"
)

/*
 * Simple interface for a min-heap
 *
 * Will be implemented using a slice
 *
 * Lowest values will be stored at the top of the tree
 * All values are expected to be integers
 */
type MinHeap interface {
    extractMin() int
    insert(i int)
    siftUp(pos int)
}

/*
 * Basic tree will implement MinHeap interface
 */
type Tree struct {
    values []int
}
/*
 * Implementation of heap methods
 */
func (t *Tree) extractMin() int {
    min, t.values := t.values[0], t.values[1:]

    return min
}

func (t *Tree) insert(i int) {
    t.values = append(t.values, i)
}

func (t *Tree) siftUp(pos int) {
    if pos == 0 {
        return
    }

    parentPos := math.Floor(float64(pos-1) / float64(2))

    if t.values[pos] <= t.values[parentPos] {
        parentValue := t.values[parentPos]
        t.values[parentPos] = t.values[pos]
        t.values[pos] = parentValue

        // continue recursive call
        t.siftUp(parentPos)
    }
}

func main() {
    // open file and defer closure
    file, err := os.Open("/tmp/heap.txt")
    if err != nil {
        fmt.Println("Error: %s", err)
    }
    defer file.Close()

    buildHeap()
}
