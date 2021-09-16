package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "betul"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("betul")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("betul") = %q, %v, want watch for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
