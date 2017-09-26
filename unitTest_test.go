package main

import (
	"testing"
)

//TestGrabAndDecode goes through most functions to see if they all work together
func TestGrabAndDecode(t *testing.T) {

	var p Repository
	var c []Contributors
	var v Lang
	grabAndDecode("apache/kafka", &p, &c, &v)

	if p.Owner.Username != "apache" {
		t.Errorf("ERROR expected: %s but got: %s", "apache", p.Owner.Username)
	}
	if p.FullName != "apache/kafka" {
		t.Errorf("ERROR expected: %s but got: %s", "apache/kafka", p.FullName)
	}
	if c[0].Contributions != 315 {
		t.Errorf("ERROR expected: %d but got: %d", 315, c[0].Contributions)
	}

}

//testCatchURL testing catchURL
func TestCatchURL(t *testing.T) {
	var input = []int{1, 2, 3, 5}
	var output = []string{"https://api.github.com/repos/apache/kafka", "https://api.github.com/repos/apache/kafka/contributors", "https://api.github.com/repos/apache/kafka/languages", ""}

	for i := range input {
		if catchURL(input[i], "apache/kafka") != output[i] {
			t.Errorf("ERROR expected: %s but got: %s", output[i], catchURL(input[i], "apache/kafka"))
		}
	}
}
