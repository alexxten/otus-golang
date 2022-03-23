package hw02unpackstring

import (
    "errors"
	"strings"
	"unicode"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func IsStringValid(inputString string) (bool) {
    var initialString = []rune(inputString)
    if unicode.IsDigit(initialString[0]) {return false}

    for i:=0;i<len(initialString)-1;i++ {
        if unicode.IsDigit(initialString[i]) && unicode.IsDigit(initialString[i+1]) {
            return false
        }
    }
    return true
}

func Unpack(inputString string) (string, error) {
	var initialString = inputString

	// если пустая строка
	if len(initialString) == 0 {
		return "", nil
	}

	// валидируем строку
	var isValid = IsStringValid(initialString)
	if !isValid {
		return "", ErrInvalidString
	}

	var formattedStr strings.Builder

	// формируем массив строк, где элемент - символ или символ+цифра
	var arr = make([]string, 0)
	var initialStrRunes = []rune(initialString)
	for i, v := range initialStrRunes {
		switch {
		case i == len(initialStrRunes)-1:
			if !unicode.IsDigit(v) {
				arr = append(arr, string(v))
			}
		case unicode.IsDigit(v):
			continue
		case unicode.IsDigit(initialStrRunes[i+1]):
			arr = append(arr, string(initialStrRunes[i:i+2]))

		default:
			arr = append(arr, string(v))
		}
	}

	for _, v := range arr {
		runeVal := []rune(v)
		switch {
		case len(runeVal) == 2: // если символ+цифра
			repeat, _ := strconv.Atoi(string(runeVal[1]))
			formattedStr.WriteString(strings.Repeat(string(runeVal[0]), repeat))
		default:
			formattedStr.WriteString(v)
		}
	}
	return formattedStr.String(), nil
}
