package proxy

import "testing"

func TestProxy(t *testing.T) {
	var sub Subject
	sub = NewProxy("hello", "world")

	res := sub.Do()

	if res != "pre:helloworld:after" {
		t.Fail()
	}
}
