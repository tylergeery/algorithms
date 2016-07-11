package main

import (
  "bufio"
  "fmt"
  "strconv"
  "os"
  "strings"
  "unicode/utf8"
  "math"
)

const HASH_X = 263
const HASH_PRIME = 1000000007

// basic map interface
// will be implemented with HashMap type
type Hash interface {
    get(entry string) string
    set(entry string)
    remove(entry string)
}

type HashMap struct {
    bucketsCount int
    buckets [][]string
}

/**
 * Implements a string hashing function
 */
func (h HashMap) hash(str string) int {
    hashSum, i, xFloated := 0, float64(0), float64(HASH_X)

    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        str = str[size:]
        hashSum += (int(r) * int(math.Pow(xFloated, i))) % HASH_PRIME
        i++
    }

    return hashSum % h.bucketsCount
}

/**
 * Check if entry exists in hash list
 * hash lists contain elements with duplicate hash values
 */
func (h HashMap) entryExists(entry string, possibles []string) bool {
    for _, str := range possibles {
        if entry == str {
            return true
        }
    }

    return false
}

// get a given map value
func (h HashMap) get(entry string) string {
    ha := h.hash(entry)
    possibles := h.buckets[ha]
    found := h.entryExists(entry, possibles)

    if found {
        return entry
    }
    return ""
}

// set a given map value
func (h HashMap) set(entry string) {
    ha := h.hash(entry)
    possibles := h.buckets[ha]
    found := h.entryExists(entry, possibles)

    if !found {
        h.buckets[ha] = append(h.buckets[ha], entry)
    }
}

// set slice index back to nil
func (h HashMap) remove(entry string) {
    ha := h.hash(entry)
    possibles := h.buckets[ha]

    for i, str := range possibles {
        if ha == h.hash(str) {
            // remove element from this hash list
            h.buckets[ha] = append(h.buckets[ha][:i], h.buckets[ha][i+1:]...)
            break;
        }
    }

}

func process(book HashMap, txt string) {
    instructions := strings.Fields(txt)

    switch instructions[0] {
    case "find":
        tmp := book.get(instructions[1])

        if tmp != "" {
            fmt.Println("yes")
        } else {
            fmt.Println("no")
        }
    case "check":
        indy, _ := strconv.Atoi(instructions[1])
        fmt.Println(book.buckets[indy])
    case "add":
        book.set(instructions[1])
    case "del":
        book.remove(instructions[1])
    }
}

/**
 * This script will parse a list of available actions that could be performed on a hash
 *
 * It will use a chained-list hash expecting any utf8 strings
 * e.g. file
    5
    12
    add world
    add HellO
    check 4
    find World
    find world
    del world
    check 4
    del HellO
    add luck
    add GooD
    check 2
    del good
 */
func main() {
    // variables needed to parse file
    var book HashMap

    // open file and defer closure
    file, err := os.Open("/tmp/hashbook.txt")
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
            // this will contain the number of buckets
            book.bucketsCount, _ = strconv.Atoi(txt)
            book.buckets = make([][]string, book.bucketsCount)
        } else if index > 1 {
            // process the requested actions
            process(book, txt)
        }

        index++
    }
}
