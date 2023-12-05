package fiz

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		s    string
		want string
		err  error
	}{
		{"Hello", "olleH", nil},
		//{"Hello, World!", "!dlroW ,olleH", nil},
		//{"", "", nil},
		//{"Hello, 世界", "界世 ,olleH", nil},
		//{string([]byte{130}), "", errors.New("input is not valid UTF-8")},
	}

	for _, tt := range tests {
		got, err := Reverse(tt.s)
		if (err != nil) != (tt.err != nil) {
			t.Errorf("Reverse(%q) returned error %v, want error %v", tt.s, err, tt.err)
			continue
		}
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("Reverse(%q) returned error %q, want error %q", tt.s, err, tt.err)
		}
		if got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
