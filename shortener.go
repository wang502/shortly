package main

import (
    "math"
)

const BASE = 62
const UPPER_OFFSET = 39
const LOWER_OFFSET = 97
const DIGIT_OFFSET = 4

func RealNum(r rune) int {
    if '0' <= r && r <= '9' {
          return int(r) + DIGIT_OFFSET
    } else if 'A' <= r && r <= 'Z' {
          return int(r) - UPPER_OFFSET
    } else if 'a' <= r && r <= 'z' {
          return int(r) - LOWER_OFFSET
    } else {
          return -1
    }
}

func RealRune(num int) rune {
    if 0 <= num && num <= 25{
        // a-z
        return rune(num + LOWER_OFFSET)
    } else if 26 <= num && num <= 51 {
        // A-Z
        return rune(num + UPPER_OFFSET)
    } else if 52 <= num && num <= 61{
        // 0-9
        return rune(num - DIGIT_OFFSET)
    } else {
        return rune(1)
    }
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

/*
  0  → a
  1  → b
  ...
  25 → z
  ...
  52 → 0
  61 → 9
*/

func Encode(id int) string {
    if id <= 0 {
        return ""
    }
    id += 1000
    str := make([]rune, 0)
    var rem int
    for id > 0{
        rem = id % BASE
        str = append(str, RealRune(rem))
        id = id / BASE
    }
    return Reverse(string(str))
}

func Decode(str string) int {
    res := 0
    rev_str := Reverse(str)
    for i, r := range rev_str {
        res += RealNum(r) * int(math.Pow(BASE, float64(i)))
    }
    return res - 1000
}

/*
func main() {
  for i:=0; i<126; i++ {
        en := Encode(i)
        fmt.Printf("Encoded %d as %s\n", i, en)
        de := Decode(en)
        fmt.Printf("Decoded %s as %d\n", en, de)
        if i != de {
            fmt.Printf("Encode() Decode() Error\n")
        }
  }
}
*/
