package maxchars

import "testing"

func TestMaxChars(t *testing.T) {
	want := "o"
	if got := maxChars("oh how cool"); got != want {
		t.Errorf("Maxchars() = %q, want %q", got, want)
	}
}
