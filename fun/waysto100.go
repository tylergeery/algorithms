package main

import (
    "fmt"
    "strconv"
)

var d = []int{1,2,3,4,5,6,7,8,9}

func main() {
	solve(d, []string{})
}

func solve(seq []int, path []string) {
    options := []string{"+","-",""}

    for i, len := 0, len(options); i < len; i++ {
        // shift next number off sequence
        _, new_seq := seq[0], seq[1:]
        len := cap(new_seq)

        if (len > 0) {
            // copy path to new value
            // add current step
            new_path := append([]string(nil), path...)
            new_path = append(new_path, options[i])

            solve(new_seq, new_path)
        } else {
            if(calculate(path) == 100) {
                printAnswer(path)
            }
        }
    }
}

func calculate(path []string) int64 {
    // always starts with 1
    var result int64 = 1

    // keep last operator for future use
    var lastOp string = ""

    // carry will lag with number up until operator
    var carry string = "";

    for j, len := 0, len(path) - 1; j < len; j++ {
        val := d[j+1]

        switch path[j] {
        case "+", "-":
            fmt.Println(result)
            result = adjustResult(lastOp, result, carry);
            lastOp, carry = "", "" // reset values
            lastOp = path[j]
            break;
        }

        carry = setCarryValue(val, carry)
    }

    result = adjustResult(lastOp, result, carry);
    fmt.Println(path, result)
    return result
}

func adjustResult(op string, result int64, carry string) int64 {
    if op == "+" {
        result += getCarryValue(carry);
    } else if op == "-" {
        result -= getCarryValue(carry);
    }

    return result
}

func getCarryValue(carry string) int64 {
    val, _ := strconv.ParseInt(carry, 10, 0)

    return val
}

func setCarryValue(val int, existing string) string {
    str := strconv.Itoa(val)

    if (existing != "") {
        str = existing + str
    }

    return str
}

func printAnswer(path []string) {
    answer := "1"
    for k, len := 0, len(path); k < len; k++ {
        answer += path[k] + strconv.Itoa(d[k+1])
    }

    fmt.Println(answer)
}
