package exercise09

import "strings"

const prefix = "$"

// expand replaces each substring "$foo" within s by the text returned by f("foo")
func expand(s string, f func(string) string) string {
	if f == nil {
		return s
	}

	ss := strings.Split(s, " ")
	for i := range ss {
		if strings.HasPrefix(ss[i], prefix) {
			ss[i] = f(strings.TrimPrefix(ss[i], prefix))
		}
	}
	return strings.Join(ss, " ")
}
