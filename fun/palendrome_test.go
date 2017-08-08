package main

import (
    "testing"
)

func TestSolve(t *testing.T) {
    test_expectations := map[string]string{
        "aaaa": "aaaa",
        "abba": "abba",
        "abcdc": "abcdcba",
        "abcddc": "abcddcba",
        "abbdddb": "abbdddbba",
        "racec": "racecar",
        "raceca": "racecar",
        "amanaplanacanal": "amanaplanacanalpanama",
        "amanaplanacanall": "amanaplanacanallanacanalpanama",
    }

    for key, exp_value := range test_expectations {
        solve_value := Solve(key)

        if exp_value != solve_value {
            t.Error("Expected:", exp_value, "got: ", solve_value)
        }
    }
}
