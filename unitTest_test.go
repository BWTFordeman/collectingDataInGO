package main

import (
	"testing"
)

//test
func TestGetAPI(t *testing.T) {

}

//testCatchURL testing catchURL
func TestCatchURL(t *testing.T) {
	var input = []int{1, 2, 3, 5}
	var output = []string{"https://api.github.com/repos/Stektpotet/Amazeking", "https://api.github.com/repos/Stektpotet/Amazeking/contributors", "https://api.github.com/repos/Stektpotet/Amazeking/languages", ""}

	for i := range input {
		if catchURL(input[i]) != output[i] {
			t.Errorf("ERROR expected: %s but got: %s", output[i], catchURL(input[i]))
		}
	}
}
