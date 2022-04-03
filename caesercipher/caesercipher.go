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

func rotate(r rune, base, delta int32) rune {
	tmp := r - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func cipher(r rune, delta int32) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	} else if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func caesarCipher(s string, k int32) string {
    // Write your code here
	// isAlphaNumeric := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
    chars := strings.Split(s, "")
    var sb strings.Builder
	for _,v := range chars{
		c := []rune(v)[0] 
		sb.WriteRune(cipher(c,k))
	}
	return sb.String()
}


func caesarCipherx(s string, k int32) string {
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
				x := ((c - 'A' + k)%(26)) + 'A'
				sb.WriteRune(x)
			} else {
			x := ((c - 'a' + k)%(26)) + 'a'
			sb.WriteRune(x)
			}	
		}
	}
	return sb.String()
}
