package escape

import "github.com/mrz1836/go-sanitize"

func Alpha(s string, allowSpaces bool) string {
	return sanitize.Alpha(s, allowSpaces)
}

func AlphaNumeric(s string, allowSpaces bool) string {
	return sanitize.AlphaNumeric(s, allowSpaces)
}

func Password(p string) string {
	return sanitize.Custom(p, `'|=|;|`)
}
