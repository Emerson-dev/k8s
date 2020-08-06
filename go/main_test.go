package main

import (
	"testing"
)

type ChallengeTable struct {
	Param string
}

func TestGreeting(t *testing.T) {

	data := ChallengeTable{"Code.education Rocks!"}

	if data == (ChallengeTable{}) {
		t.Error("Resultado invalido")
	}
}
