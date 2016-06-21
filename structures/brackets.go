package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

// declare custom bracket type
type Bracket struct {
    pos int
    char string
}

// go through string to find if all brackets are properly closed
// examples of valid brackets include
//   "{}", "[]", "()"
func parse(path string) (index int, err error) {
    // construct stack for tracking valid brackets
    var brackets []Bracket
    var currentCharacter = 0
    var tmp Bracket

    // open file and defer closure
    file, err := os.Open(path)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    // bufio for reading line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // scanner.Text() has value
        txt := scanner.Text()
        for _, r := range txt {
            c := fmt.Sprintf("%c", r)
            // 1-based index result
            currentCharacter++

            switch c {
            case "[","{","(":
                // create and assign new bracket obj
                // append to stack
                brackets = append(brackets, Bracket{currentCharacter, c})
            case "]","}",")":
                // return unmatched closing bracket
                if len(brackets) == 0 {
                    return currentCharacter, scanner.Err()
                }

                // get first value of stack
                tmp, brackets = brackets[len(brackets)-1], brackets[:len(brackets)-1]
                log.Println("TMP: " + tmp.char)

                if (tmp.char == "[" && c == "]" ||
                    tmp.char == "{" && c == "}" ||
                    tmp.char == "(" && c == ")") {
                    // this is valid
                } else {
                    return tmp.pos, scanner.Err()
                }
            }
        }
    }

    // check for unclosed characters
    // find first occurence of that character if found
    if len(brackets) > 0 {
        res := brackets[0]

        return res.pos, scanner.Err()
    }


    // success
    return 0, scanner.Err()
}

func main() {
    // call parse
    indy, _ := parse("/tmp/brackets.txt")

    if indy == 0 {
        log.Println("Success")
    } else {
        log.Println(indy)
    }
}
