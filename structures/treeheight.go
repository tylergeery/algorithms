package main

import (
  "bufio"
  "fmt"
  "strconv"
  "os"
  "strings"
)

// most basic tree
type Tree struct {
    head int
}

// simple node structure
// will use index from file as node id
type Node struct {
    id int
    children []int
}

// recursively count the height of a node by the count of its children
func recursiveCount(nodes []Node, n Node) int {
    if hasChildren(n) {
        // get max of all childrens heights
        max := 1

        // loop through children to find max height
        for _, child := range n.children {
            count := recursiveCount(nodes, nodes[child])
            if count >= max {
                max = count
            }
        }

        // height for current node + childrens max height
        return 1 + max
    }

    // height for current node w/ no children
    return 1
}

func hasChildren(n Node) bool {
    if (len(n.children) > 0) {
        return true
    }

    return false
}
// parse file
func main() {
    // variables needed to parse file
    var nodes []Node
    tree := new(Tree)

    // open file and defer closure
    file, err := os.Open("/tmp/treeheight.txt")
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

        if index == 0 {
            // first line contains node count
        } else {
            // second line contains nodes as representations of parent values
            // indexes are node id

            // split on spaces
            tmpNodes := strings.Fields(txt)
            nodeChildList := make([][]int, len(tmpNodes))

            // create nodes list
            for i, n := range tmpNodes {
                parentIndex, _ := strconv.Atoi(n)
                nodes = append(nodes, Node{i, make([]int, 0)})

                if parentIndex == -1 {
                    // head of tree
                    tree.head = i
                } else {
                    // add as children of parent
                    nodeChildList[parentIndex] = append(nodeChildList[parentIndex], i)
                }
            }

            // move child nodes to nodes objects
            for i, _ := range tmpNodes {
                nodes[i].children = nodeChildList[i]
            }
        }

        index++
    }

    // call recursive function with tree head
    treeHeight := recursiveCount(nodes, nodes[tree.head])

    fmt.Println("Tree height:", treeHeight)
}
