package hosttech

import "strings"

func RemoveTrailingDot(input string) string {
	return strings.TrimRight(input, ".")
}
