package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Renatinho")
	if got != "<b>Renatinho</b>" {
		t.Errorf("greeting('Renatinho') = '%s'; want '<b>Renatinho</b>'", got)
	}
}
