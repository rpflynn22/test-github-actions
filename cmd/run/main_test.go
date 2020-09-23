package main

import "testing"

func Test_a(t *testing.T) {
	if a() != 1 {
		t.Error("fail")
	}
}
