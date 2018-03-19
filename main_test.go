package main

import "testing"

func TestGetHelloWorldString(t *testing.T) {
	if GetHelloWorldString() != "Hello World!" {
		t.Error("GetHelloWorldString() != \"Hello World!\"")
	}
}