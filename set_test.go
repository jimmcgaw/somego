package main

import (
	"reflect"
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	if len(set) != 2 {
		t.Fatalf("#Add: set does not contain added items")
	}
}

func TestSetHas(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(4)
	if set.Has(1) != true {
		t.Fatalf("#Has: Set does not note having added value")
	}
	if set.Has(3) {
		t.Fatalf("#Has: Set claims to contain item it does not")
	}
}

func TestSetValues(t *testing.T) {
	set := IntSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(set.Values(), expected) {
		t.Fatalf("#Values: Set returns incorrect values")
	}

}

func TestAddMultipleValues(t *testing.T) {
	set := IntSet{}
	expected := []int{5, 6, 7, 8}
	set.Update(expected)
	if !reflect.DeepEqual(set.Values(), expected) {
		t.Fatalf("#Update: Set does not bulk add values properly")
	}
}
