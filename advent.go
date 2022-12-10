package advent

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAsString(filepath string) string {
	dat, err := os.ReadFile(filepath)
	check(err)
	return string(dat)
}

func SplitTwoIntsByString(line, splitBy string) (int, int) {
	splitString := strings.Split(line, splitBy)
	num1, err := strconv.Atoi(splitString[0])
	check(err)
	num2, err := strconv.Atoi(splitString[1])
	check(err)
	return num1, num2
}