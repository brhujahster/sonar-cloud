package main

import "testing"

func TestSum(t *testing.T) {
	result := somar(2, 2)
	
	if result != 4 {
		t.Error("result must be 4")
	}
}