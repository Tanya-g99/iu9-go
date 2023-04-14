package main

import (
	"fmt"
)

func encode(utf32 []rune) []byte {
	var utf8 []byte
	for i := 0; i < len(utf32); i++ {
		if utf32[i] < 128 {
			utf8 = append(utf8, byte(utf32[i]))
		} else if utf32[i] < 2048 {
			b2 := utf32[i] % 0x40
			b1 := utf32[i] / 0x40
			utf8 = append(utf8, 0xC0+byte(b1), 0x80+byte(b2))
		} else if utf32[i] < 65536 {
			b3 := utf32[i] % 0x40
			b2 := (utf32[i] / 0x40) % 0x40
			b1 := utf32[i] / 0x1000
			utf8 = append(utf8, 0xE0+byte(b1), 0x80+byte(b2), 0x80+byte(b3))
		} else {
			b1 := byte(utf32[i] / 262144)
			b2 := byte(utf32[i] / 0x1000 % 0x40)
			b3 := byte(utf32[i] / 0x40 % 0x40)
			b4 := byte(utf32[i] % 0x40)
			utf8 = append(utf8, 0xF0+b1, 0x80+b2, 0x80+b3, 0x80+b4)
		}
	}
	return utf8
}

func decode(utf8 []byte) []rune {
	var utf32 []rune
	for i := 0; i < len(utf8); i++ {
		if utf8[i] < 128 {
			utf32 = append(utf32, rune(utf8[i]))
		} else if utf8[i] < 224 {
			utf32 = append(utf32, rune(utf8[i]-192)*0x40+rune(utf8[i+1]-128))
			i++
		} else if utf8[i] < 240 {
			utf32 = append(utf32, rune(utf8[i]-224)*0x1000+rune(utf8[i+1]-128)*0x40+rune(utf8[i+2]-128))
			i += 2
		} else {
			utf32 = append(utf32, rune(utf8[i]-240)*262144+rune(utf8[i+1]-128)*0x1000+rune(utf8[i+2]-128)*0x40+rune(utf8[i+3]-128))
			i += 3
		}
	}
	return utf32
}

func main() {
	sources := [][]rune{
		{},
		{'\u002c', '\u002d'},
		{'\U000002ac', '\U000002a1'},
		{'\U000022ac', '\U00002aa1'},
		{'\U001022ac', '\U000102a1'},
	}
	for _, source := range sources {
		utf8 := encode(source)
		for _, v := range utf8 {
			fmt.Printf("0x%x ", v)
		}
		fmt.Printf("\n")
		utf32 := decode(utf8)
		for _, v := range utf32 {
			fmt.Printf("\\U%x ", v)
		}
		fmt.Printf("\n")
	}
}
