package token

import "testing"

func TestGetTokenType(t *testing.T) {
	tests := []struct {
		literal string
		want    string
	}{
		{"if", "IF"},
		{"else", "ELSE"},
		{"print", "PRINT"},
		{"true", "TRUE"},
		{"false", "FALSE"},
		{"42", "INT"},
		{"123i32", "I32"},
		{"999i64", "I64"},
		{"myVar", "IDENT"},
		{"!", "ILLEGAL"},
	}

	for _, tt := range tests {
		got := GetTokenType(tt.literal)
		if got != tt.want {
			t.Errorf("GetTokenType(%q) = %q; want %q", tt.literal, got, tt.want)
		}
	}
}
