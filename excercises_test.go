package main

import (
	"reflect"
	"testing"
)

func TestAbbreviate(t *testing.T) {

	result := Abbreviate("Work From Home")
	if result != "WFH" {
		t.Errorf("Expected WFH but got %s\n", result)
	}

	result = Abbreviate("College of Engineering Trivandrum")
	if result != "CET" {
		t.Errorf("Expected CET but got %s\n", result)
	}
}

func TestPangram(t *testing.T) {
	result := Pangram("abcdefghijklmnopqrstuvwxyz")
	if !result {
		t.Errorf("Expected true but got false")
	}

	result = Pangram("Hello World")
	if result {
		t.Errorf("Expected true but got false")
	}

}

func TestPalindrome(t *testing.T) {

	result := Palindrome("radar")
	if !result {
		t.Errorf("Expected true but got false")
	}

	result = Palindrome("hello")
	if result {
		t.Errorf("Expected true but got false")
	}

}

func TestFrequency(t *testing.T) {

	result := Frequency("Okey")
	expected := map[string]int{
		"O": 1,
		"k": 1,
		"e": 1,
		"y": 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	result = Frequency("Noon")
	expected = map[string]int{
		"N": 1,
		"o": 2,
		"n": 1,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestFizzBizz(t *testing.T) {

	result := Fizzbiz(16)
	expected := []string{"1", "2", "fizz", "4", "bizz", "fizz", "7", "8", "fizz", "bizz", "11", "fizz", "13", "14", "fizzbizz", "16"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	result = Fizzbiz(15)
	expected = []string{"1", "2", "fizz", "4", "bizz", "fizz", "7", "8", "fizz", "bizz", "11", "fizz", "13", "14", "fizzbizz"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

}
