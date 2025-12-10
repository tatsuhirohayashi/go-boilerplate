package pointer

// String は文字列をポインタに変換するヘルパー関数です
func String(s string) *string {
	return &s
}

// Int は整数をポインタに変換するヘルパー関数です
func Int(i int) *int {
	return &i
}

// Bool はブール値をポインタに変換するヘルパー関数です
func Bool(b bool) *bool {
	return &b
}