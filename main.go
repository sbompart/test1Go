package main

import (
	"errors"
	"strconv"
)

//const strIn = "A0511AB398765UJ1N230200c11"
var values []string

func getType(str string) (string, error) {
	if len(str) < 3 {
		return "", errors.New("GetType - String must contain 3 chars min")
	}
	if  str[:1] != "N" && str[:1] != "A" {
		return "", errors.New("GetType - First leter is not N or A")
	}
	return str[3:], nil
}

func getLenght(str string) (string, string, error) {
	if len(str) < 2 {
		return "", "", errors.New("GetLength - String must contain 2 chars min")
	}
	return str[:2], str[2:], nil
}

func getValue(str string, lgt int) (string, string, error) {
	if len(str) < lgt {
		return "", "", errors.New("GetValue - The value contains fewer characters than necessary.")
	}
	return str[:lgt], str[lgt:], nil
}

func loopForStr(str string) ([]string, error) {
	if len(str) < 1 {
		return nil, errors.New("Erorr empty string")
	}

	str, err := getType(str)
	if err != nil {
		return nil, err
	}
	l, str, err := getLenght(str)
	if err != nil {
		values = values[:0]
		return nil, err
	}

	lgt, _ := strconv.Atoi(l)
	v, str, err := getValue(str, lgt)
	if err != nil {
		values = values[:0]
		return nil, err
	}
	values = append(values, v)
	if len(str) == 0 {
		return values, nil
	}
	return loopForStr(str)
}

func main() {
	/*value, err := loopForStr(strIn)
	fmt.Println("Result: ", value, err)*/
}
