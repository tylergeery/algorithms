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
 * Will build a heap using basic go slice
 *
 * Records swap output necessary to create valid heap
 */
func main() {
    var heap []int
    var swaps [][]int

    // open file and defer closure
    file, err := os.Open("/tmp/buildheap.txt")
    if err != nil {
        fmt.Println("Error: %s", err)
    }
    defer file.Close()

    // parse file
    // bufio for reading line by line
    index := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // scanner.Text() has value
        txt := scanner.Text()

        // ingore first line
        if index == 1 {
            // split on spaces
            values := strings.Fields(txt)

            // add values to unsorted heap
            for _, stringVal := range values {
                val, _ := strconv.Atoi(stringVal)
                heap = append(heap, val)
            }

            if len(heap) > 0 {
                // start in bottom left, start organizing heap
                log := math.Log2(float64(len(heap)))
                heapRowStart := int(math.Max(float64(0), math.Ceil(log)))
                heapRowIndex := int(math.Max(float64(heapRowStart), float64(len(heap)-1)))


                for heapRowIndex >= heapRowStart {
                    // check if node is less than parent
                    pos, parentPos := heapRowIndex, int(math.Floor(float64(heapRowIndex-1) / float64(2)))

                    for heap[pos] < heap[parentPos] {
                        // swap is neccessary
                        tmp := heap[parentPos]
                        heap[parentPos] = heap[pos]
                        heap[pos] = tmp

                        // record the swap for output
                        swaps = append(swaps, []int{parentPos, pos})

                        if parentPos != 0 {
                            pos, parentPos = parentPos, int(math.Floor(float64(parentPos-1) / float64(2)))
                        } else {
                            pos, parentPos = 0, 0
                        }
                    }

                    heapRowIndex--
                }
            }

            // output all the swaps
            fmt.Println(len(swaps))
            for _, swap := range swaps {
                fmt.Println(swap)
            }
        }

        index++
    }
}
