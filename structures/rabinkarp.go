package main

import (
    "bufio"
    "fmt"
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
func hash(str string) int64 {
    str = reverse(str)
    hashSum, i, xFloated := int64(0), float64(0), float64(HASH_X)

    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        str = str[size:]
        pow := int64(math.Pow(xFloated, i))
        val := int64(r) * pow
        hashSum += val
        i++
    }

    return hashSum
}

/**
* Compute hash values for all search strings within length
* of text search
*/
func precomputeHashes(text string, patternLength int) []int64 {
    indy := 0
    textLength := utf8.RuneCountInString(text) - patternLength
    hashSum := hash(text[:patternLength])
    precomputedHashes := []int64{hashSum % HASH_PRIME}

    for indy < textLength {
        txt := text[patternLength:]
        l, _ := utf8.DecodeRuneInString(txt)
        r, size := utf8.DecodeRuneInString(text)
        text = text[size:]

        /**
        * calculate the new hash value in O(1)
        * e.g. "abra"
        *                HASH_X = 101
        *                ASCII a = 97, b = 98, r = 114.
        * hash("abr") = (97 × 101^2) + (98 × 101^1) + (114 × 101^0) = 999,509
        *                base   old hash    old 'a'         new 'a'
        * hash("bra") = [101 × (999,509 - (97 × 101^2))] + (97 × 101^0) = 1,011,309
        */
        poly := math.Pow(float64(HASH_X), float64(patternLength-1))
        oldA := int64(r) * int64(poly)
        newA := int64(l)
        diff := hashSum - oldA
        hashSum = (int64(HASH_X) * diff) + newA

        precomputedHashes = append(precomputedHashes, (hashSum % HASH_PRIME))
        indy++
    }

    return precomputedHashes
}

/**
* Trims the last new line character from the string
*/
func trim(text string) string {
    return strings.TrimSuffix(text, "\n")
}

/**
* Reverse a string for easy hash computation
*/
func reverse(s string) string {
    o := make([]rune, utf8.RuneCountInString(s));
    i := len(o);
    for _, c := range s {
        i--;
        o[i] = c;
    }
    return string(o);
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter pattern to search for: ") // abc
    pattern, _ := reader.ReadString('\n')
    pattern = trim(pattern)

    newReader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text to search within: ") // daewfsaabcasdfsdabc
    text, _ := newReader.ReadString('\n')
    text = trim(text)

    patternHash := hash(pattern) % HASH_PRIME
    patternLength := utf8.RuneCountInString(pattern) // trims newline character
    hashes := precomputeHashes(text, patternLength)

    for i, value := range hashes {
        if value == patternHash {
            if pattern == text[i:i+patternLength] {

                // print index of where there is a string match
                fmt.Printf("%d ", i)
            }
        }
    }
    fmt.Println()
}
