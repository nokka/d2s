package d2s

import "strings"

type name [16]byte

func (n name) String() string {
	return strings.Trim(string(n[:]), "\x00")
}
