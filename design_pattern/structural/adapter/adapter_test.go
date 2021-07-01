package adapter

import "testing"

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	customAdaptee := NewCustomAdaptee()
	target := NewAdapter(customAdaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
