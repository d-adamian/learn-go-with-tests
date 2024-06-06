package iteration

import "bytes"

func Repeat(char string, repeatCount int) string {
	var buf bytes.Buffer
	for i := 0; i < repeatCount; i++ {
		buf.WriteString(char)
	}

	return buf.String()
}
