package main

import (

	"strings"
	"regexp"
	"unicode"
)


func main() {
	input := "ahello'-abcd123[a-z-A-Z]"
	println(input)
	println(caesarCipher(input, 1))
}

func caesarCipher(s string, k int32) string {
    // Write your code here
	isAlphaNumeric := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
    chars := strings.Split(s, "")
    var sb strings.Builder
	for _,v := range chars{
		// println(v)
		if !isAlphaNumeric(v) {
			sb.WriteString(v)
		} else {
			//get char
			c := []rune(v)[0] 
			if unicode.IsUpper(c) {				
				x := ((c - 'A' + k)%(1+'Z'-'A')) + 'A'
				sb.WriteRune(x)
			} else {
			x := ((c - 'a' + k)%(1+'z'-'a')) + 'a'
			sb.WriteRune(x)
			}	
		}
	}
	return sb.String()
}
