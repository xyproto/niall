package niall

import "testing"

func TestHello(t *testing.T) {
	const in1 = "hello there!"
	const in2 = "I can say hello several times"
	const in3 = "How are you?"
	const out = "Hello several times."
	Learn(in1)
	Learn(in2)
	Learn(in3)
	if x := Talk(); x != out {
		t.Errorf("Reply was %v, want %v", x, out)
	}
}
