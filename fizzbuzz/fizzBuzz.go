package fizzBuzz

import "strconv"

func say(number int) string {
	if isFizzNumber(number) && isBuzzNumber(number) {
		return "FizzBuzz"
	}

	if isFizzNumber(number) {
		return "Fizz"
	}

	if isBuzzNumber(number) {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func isFizzNumber(number int) bool {
	return number%3 == 0
}

func isBuzzNumber(number int) bool {
	return number%5 == 0
}
