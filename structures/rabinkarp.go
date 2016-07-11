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

    return hashSum
}

/**
 * Compute hash values for all search strings within length
 * of text search
 */
func precomputeHashes() {
    //TODO
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter pattern to search for: ") // abc
    pattern, _ := reader.ReadString('\n')
    fmt.Println(pattern)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text to search within: ") // daewfsaabcasdfsdabc
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    patternHash := hash(pattern)
    patternLength := len(pattern)
    hashes := precomputeHashes(text, patternLength)

    for i, value := range hashes {
        if value == patternHash {
            if pattern == text[i:i+patternLength] {
                // print index of where there is a string match
                fmt.Printf("%d ", i)
            }
        }
    }
}
