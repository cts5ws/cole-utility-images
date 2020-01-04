package ratelim

import (
	"net"
	"testing"
)

func TestIsValidRequest(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"192.0.2.1", false},
		{"192.0.2.1", true},
		{"2001:db8::1", false},
		{"2001:db8::1", true},
	}

	for index, test := range tests {
		if got := ShouldRateLimit(net.ParseIP(test.input)); got != test.want {
			t.Errorf("case(%d) - ShouldRateLimit(%s) = %v", index, test.input, got)
		}
	}
}
