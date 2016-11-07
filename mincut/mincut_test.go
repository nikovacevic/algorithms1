package mincut

import (
	"fmt"
	"testing"
)

func TestMincut(t *testing.T) {
	path := "test/small.txt"
	g, err := NewGraphFromFile(path)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	act := g.Mincut()
	if act != 2 {
		t.Errorf("Mincut() expected %v, actual %v", 2, act)
	}

	path = "test/trial1.txt"
	g, err = NewGraphFromFile(path)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("\nTRIAL1: %v\n", g.Mincut())

	path = "test/mincut.txt"
	g, err = NewGraphFromFile(path)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("\nCOURSERA: %v\n", g.Mincut())
}
