package hw02unpackstring

import (
    "errors"
    "fmt"
	"strings"
	"unicode"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input_string string) (string, error) {
    var initial_str = input_string
    var formatted_str strings.Builder

    var arr = make([]string, 0)
    for i, v := range initial_str {
        switch {
        case i == len(initial_str)-1:
            if !unicode.IsDigit(v) {
                arr = append(arr, string(v))
            }
        case unicode.IsDigit(v):
            continue
        case unicode.IsDigit(rune(initial_str[i+1])):
            arr = append(arr, initial_str[i:i+2])

        default:
            arr = append(arr, string(v))
        }

    }
    fmt.Println(arr)
    //switch {
    //case unicode.IsDigit(rune(initial_str[0])):
      //  fmt.Println("некорректно")
    //}
    //case unicode.IsDigit(v) && unicode.IsDigit(rune(initial_str[i+1])):
    //		fmt.Println("некорректно")

    for _, v := range arr {
        switch {
        case len(v) == 2:
            repeat, _ := strconv.Atoi(string(v[1]))
            formatted_str.WriteString(strings.Repeat(string(v[0]), repeat))
        default:
            formatted_str.WriteString(string(v))
        }
    }

	return formatted_str.String(), nil
}
