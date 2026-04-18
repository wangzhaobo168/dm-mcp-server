package utils

import "testing"

func TestBuildDataSourceNameEscapesSpecialCharactersInPassword(t *testing.T) {
	got := buildDataSourceName("demo-user", "pass#word:@2025", "127.0.0.1", "5236")
	want := "dm://demo-user:pass%23word%3A%402025@127.0.0.1:5236"

	if got != want {
		t.Fatalf("buildDataSourceName() = %q, want %q", got, want)
	}
}
