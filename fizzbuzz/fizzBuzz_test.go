package fizzBuzz

import "testing"

func Test_input_1_say_1(t *testing.T) {
	actual := say(1)

	if actual != "1" {
		t.Error("expect 1 but got", actual)
	}
}

func Test_input_2_say_2(t *testing.T) {
	actual := say(2)

	if actual != "2" {
		t.Error("expect 2 but got", actual)
	}
}

func Test_input_3_say_Fizz(t *testing.T) {
	input := 3
	actual := say(input)

	if actual != "Fizz" {
		t.Error("expect Fizz but got", actual)
	}
}

func Test_input_5_say_Buzz(t *testing.T) {
	input := 5
	actual := say(input)

	if actual != "Buzz" {
		t.Error("expect Buzz but got", actual)
	}
}

func Test_input_6_say_Fizz(t *testing.T) {
	input := 6
	actual := say(input)

	if actual != "Fizz" {
		t.Error("expect Fizz but got", actual)
	}
}

func Test_input_10_say_Buzz(t *testing.T) {
	input := 10
	actual := say(input)

	if actual != "Buzz" {
		t.Error("expect Buzz but got", actual)
	}
}

func Test_input_15_say_Buzz(t *testing.T) {
	input := 15
	actual := say(input)

	if actual != "FizzBuzz" {
		t.Error("expect FizzBuzz but got", actual)
	}
}
