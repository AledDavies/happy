package main

import (
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	tables := []struct {
		num    int
		digits []int
	}{
		{5, []int{5}},
		{1111, []int{1, 1, 1, 1}},
		{1234, []int{4, 3, 2, 1}},
	}

	for _, table := range tables {
		digits := Digits(table.num)

		if !reflect.DeepEqual(table.digits, digits) {
			t.Errorf("Digits was incorrect, got: %v, want: %v.", digits, table.digits)
		}
	}
}

func TestHappy(t *testing.T) {
	tables := []struct {
		number int
		happy  bool
	}{
		{1, true},
		{7, true},
		{8, false},
		{10, true},
		{888, true},
	}

	for _, table := range tables {
		happy := Happy(table.number)

		if table.happy != happy {
			t.Errorf("Happy was incorrect, for %v got: %v, want: %v.", table.number, happy, table.happy)
		}
	}
}
