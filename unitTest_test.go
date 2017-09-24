package main

import (
	"testing"
)

//TestGrabAndDecode goes through most functions to see if they all work together
func TestGrabAndDecode(t *testing.T) {

	var p Repository
	var c []Contributors
	var v Lang
	grabAndDecode(&p, &c, &v)

	if p.Owner.Username != "Stektpotet" {
		t.Errorf("ERROR expected: %s but got: %s", "Stektpotet", p.Owner.Username)
	}
	if p.FullName != "Stektpotet/Amazeking" {
		t.Errorf("ERROR expected: %s but got: %s", "Stektpotet/Amazeking", p.FullName)
	}
	if c[0].Contributions != 107 {
		t.Errorf("ERROR expected: %d but got: %d", 107, c[0].Contributions)
	}

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
