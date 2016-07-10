package main

import (
  "bufio"
  "fmt"
  "strconv"
  "os"
  "strings"
)

// minimum allowable phone number
const HASH_MAX_SIZE = 9999999

// basic map interface
// will be implemented with HashMap type
type Hash interface {
    get(ph int) string
    set(ph int, name string)
    remove(ph int)
}

type HashMap []string

// get a given map value
func (h HashMap) get(ph int) string {
    return h[ph]
}

// set a given map value
func (h HashMap) set(ph int, name string) {
    h[ph] = name
}

// set slice index back to nil
func (h HashMap) remove(ph int) {
    h[ph] = ""
}

func process(phonebook HashMap, txt string) {
    instructions := strings.Fields(txt)
    number, _ := strconv.Atoi(instructions[1])

    switch instructions[0] {
    case "find":
        tmp := phonebook.get(number)

        if tmp != "" {
            fmt.Println(tmp)
        } else {
            fmt.Println("not found")
        }
    case "add":
        phonebook.set(number, instructions[2])
    case "del":
        phonebook.remove(number)
    }
}

/**
 * This script will parse a list of available actions that could be performed on a phone book
 *
 * It will use a direct-addresses hash expecting valid phone numbers of 7 integers
 */
func main() {
    // variables needed to parse file
    var phonebook HashMap
    phonebook = make([]string, HASH_MAX_SIZE)

    // open file and defer closure
    file, err := os.Open("/tmp/phonebook_direct.txt")
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

        if index != 0 {
            // process the requested actions
            process(phonebook, txt)
        }

        index++
    }
}
