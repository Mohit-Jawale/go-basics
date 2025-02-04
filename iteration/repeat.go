package iteration

import "strings"

func Repeat(character string, repeatedCount int) string {

	var repeated strings.Builder

	for i := 0; i < repeatedCount; i++ {

		repeated.WriteString(character)
	}

	return repeated.String()

}
