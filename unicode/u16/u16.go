package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	/* uint16 δΈ rune θ½¬ζ’ */
	words := []rune{'π', 'π'}
	// func Encode(s []rune) []uint16
	u16 := utf16.Encode(words)
	fmt.Println(u16)
	// func Decode(s []uint16) []rune
	fmt.Println(utf16.Decode(u16))
	// func EncodeRune(r rune) (r1, r2 rune)
	r1, r2 := utf16.EncodeRune('π')
	fmt.Println(r1, r2)
	// func DecodeRune(r1, r2 rune) rune
	fmt.Println(utf16.DecodeRune(r1, r2))
	fmt.Println(string(utf16.DecodeRune(r1, r2)))
	// func IsSurrogate(r rune) bool ζ―ε¦δΈΊζζδ»£ηε―Ή
	fmt.Println(utf16.IsSurrogate(r1))
	fmt.Println(utf16.IsSurrogate(r2))
	fmt.Println(utf16.IsSurrogate(123))
}
