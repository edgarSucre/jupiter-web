package escape

import (
	"strings"

	"github.com/mrz1836/go-sanitize"
)

func Alpha(s string, allowSpaces bool) string {
	ss := sanitize.Alpha(s, allowSpaces)
	return strings.TrimSpace(ss)
}

func AlphaNumeric(s string, allowSpaces bool) string {
	ss := sanitize.AlphaNumeric(s, allowSpaces)
	return strings.TrimSpace(ss)
}

func Email(s string) string {
	ss := sanitize.Email(s, true)
	return strings.TrimSpace(ss)
}

func Password(p string) string {
	ss := sanitize.Custom(p, `'|=|;|`)
	return strings.TrimSpace(ss)
}
