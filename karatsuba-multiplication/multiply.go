package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	x := userInput("Enter first number:")
	y := userInput("Enter second number:")

	product := karatsuba(x, y)

	fmt.Println("Product:", product)
}

func karatsuba(x, y string) string {
	maxLen := getMaxLen(x, y)

	if maxLen > 1 && maxLen%2 > 0 {
		maxLen++
	}

	x = padLeft(x, "0", maxLen)
	y = padLeft(y, "0", maxLen)

	if maxLen < 2 {
		n1, _ := strconv.ParseInt(x, 10, 0)
		n2, _ := strconv.ParseInt(y, 10, 0)
		res := strconv.FormatInt(n1*n2, 10)
		return res
	}

	xL := x[:maxLen/2]
	xR := x[maxLen/2 : len(x)]
	yL := y[:maxLen/2]
	yR := y[maxLen/2 : len(y)]

	prod1 := karatsuba(xL, yL)
	prod2 := karatsuba(xR, yR)
	prod3 := karatsuba(addLong(xL, xR), addLong(yL, yR))
	sub := subtractLong(subtractLong(prod3, prod1), prod2)

	res := addLong(addLong(addZeros(prod1, maxLen), addZeros(sub, maxLen/2)), prod2)
	return strings.TrimLeft(res, "0")
}

func addLong(x, y string) string {
	maxLen := getMaxLen(x, y)

	var carry int
	var result = []string{}

	x = padLeft(x, "0", maxLen)
	y = padLeft(y, "0", maxLen)

	for i := maxLen - 1; i >= 0; i-- {
		num1, _ := strconv.Atoi(string(x[i]))
		num2, _ := strconv.Atoi(string(y[i]))
		sum := num1 + num2
		if carry > 0 {
			sum += carry
			carry = 0
		}
		if sum > 9 {
			carry = 1
			result = append([]string{strconv.Itoa(sum % 10)}, result...)
		} else {
			result = append([]string{strconv.Itoa(sum)}, result...)
		}
	}

	if carry > 0 {
		result = append([]string{strconv.Itoa(carry)}, result...)
	}

	res := strings.Join(result, "")

	return res
}

func subtractLong(x, y string) string {
	maxLen := getMaxLen(x, y)

	var carry bool
	var result = []string{}

	x = padLeft(x, "0", maxLen)
	y = padLeft(y, "0", maxLen)

	for i := maxLen - 1; i >= 0; i-- {
		num1, _ := strconv.Atoi(string(x[i]))
		num2, _ := strconv.Atoi(string(y[i]))

		if carry {
			num1--
			carry = false
		}

		if num1 < num2 {
			num1 += 10
			carry = true
		}

		sub := num1 - num2

		result = append([]string{strconv.Itoa(sub)}, result...)
	}

	res := strings.Join(result, "")

	return strings.TrimLeft(res, "0")
}

func padLeft(str, pad string, lenght int) string {
	for len(str) < lenght {
		str = pad + str
	}
	return str
}

func addZeros(str string, length int) string {
	for i := 0; i < length; i++ {
		str = str + "0"
	}
	return str
}

func getMaxLen(x, y string) int {
	xLen := len(x)
	yLen := len(y)
	if xLen > yLen {
		return xLen
	} else {
		return yLen
	}
}

func userInput(msg string) string {
	fmt.Println(msg)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	matched, _ := regexp.MatchString("^\\d+$", input)

	if !matched {
		log.Fatal("non-digit input")
	}

	return input
}
