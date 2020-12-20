package main

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	l, values := &linkedList{}, []string{"alfa", "bravo", "charlie"}
	for _, val := range values {
		l.prepend(&node{name: val})
	}

	if l.length != 3 {
		t.Error("length should be 3; values not being added.")
	}
}

func TestPrettyPrint(t *testing.T) {
	l, values := &linkedList{}, []string{"alfa", "bravo", "charlie"}
	for _, val := range values {
		l.prepend(&node{name: val})
	}

	out := string(l.prettyPrint())
	expected := `name: charlie
name: bravo
name: alfa
`

	if out != expected {
		t.Errorf("\nexpected\n%s\ngot\n%s", expected, out)
	}
}

func TestExtractValues(t *testing.T) {
	l, values := &linkedList{}, []string{"alfa", "bravo", "charlie"}
	for _, val := range values {
		l.prepend(&node{name: val})
	}

	out := string(l.extractValues(","))

	if out != "charlie,bravo,alfa" {
		t.Errorf("expected 'charlie,bravo,alfa'. got %s", out)
	}
}

func TestPopWithValue_Empty(t *testing.T) {
	l := &linkedList{}
	out := l.popWithValue("hello")

	if out != nil {
		t.Error("the linked list should be empty")
	}
}

func TestPopWithValue_Full(t *testing.T) {
	l, values := &linkedList{}, []string{"alfa", "bravo", "charlie"}
	for _, val := range values {
		l.prepend(&node{name: val})
	}

	out1, out2, out3, out4 := l.popWithValue("alfa"), l.popWithValue("delta"), l.popWithValue("alfa"), l.popWithValue("charlie")

	if out1 == nil {
		t.Error("alfa should have been found")
	}

	if out1.name != "alfa" {
		t.Errorf("alfa should have returned 'alfa', got %s", out1.name)
	}

	if out2 != nil {
		t.Error("delta should not have been found")
	}

	if out3 != nil {
		t.Error("alfa should not have been found")
	}

	if out4 == nil {
		t.Error("charlie should have been found")
	}

	if out4.name != "charlie" {
		t.Errorf("charlie should have returned 'charlie', got %s", out1.name)
	}
}
