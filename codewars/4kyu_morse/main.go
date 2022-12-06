package main

import (
	"fmt"
	"strings"
)

func unitLen(bits string, bit string) int {
	unit_len := 0
	first := true
	for _, unit := range strings.Split(bits, bit) {
		unit_curlen := len(unit)
		if (unit_curlen > 0) && first {
			unit_len = unit_curlen
			first = false
		} else {
			if (unit_curlen > 0) && (unit_curlen < unit_len) {
				unit_len = unit_curlen
				break
			}
		}
	}
	return unit_len
}

func DecodeBits(bits string) string {
	bits = strings.Trim(bits, "0")
	unit_len := unitLen(bits, "0")
	unit_len1 := unitLen(bits, "1")
	if unit_len1 > 0 && unit_len1 < unit_len {
		unit_len = unit_len1
	}
	bits = strings.Replace(bits, strings.Repeat("0", 7*unit_len), "   ", -1)
	bits = strings.Replace(bits, strings.Repeat("0", 3*unit_len), " ", -1)
	bits = strings.Replace(bits, strings.Repeat("1", 3*unit_len), "-", -1)
	bits = strings.Replace(bits, strings.Repeat("1", unit_len), ".", -1)
	bits = strings.Replace(bits, strings.Repeat("0", unit_len), "", -1)
	return bits
}

func unitLenAdvanced(bits string, split string) (int, int) {
	unit_min := 0
	unit_max := 0
	first := true
	for _, unit := range strings.Split(bits, split) {
		unit_curlen := len(unit)
		if (unit_curlen > 0) && first {
			unit_min = unit_curlen
			unit_max = unit_curlen
			first = false
		} else {
			if unit_curlen > 0 {
				if unit_curlen < unit_min {
					unit_min = unit_curlen
				}
				if unit_curlen > unit_max {
					unit_max = unit_curlen
				}
			}
		}
	}
	return unit_min, unit_max
}

func decodeBitsAdvanced(bits string) string {
	bits = strings.Trim(bits, "0")
	pause_min, pause_max := unitLenAdvanced(bits, "1")
	signal_min, signal_max := unitLenAdvanced(bits, "0")
	pause_split := float32(pause_max-pause_min)/3 + float32(pause_min)
	signal_split := float32(signal_max-signal_min)/3 + float32(signal_min)

	fmt.Println(pause_min, pause_max, pause_split, signal_min, signal_max, signal_split)
	return ""
}

func DecodeMorse(morseCode string) string {
	MORSE_CODE := map[string]string{
		"-----": "0", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
		".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",

		".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
		"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
		"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
		".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
		"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y", "--..": "Z",

		"---...": ":", "--..--": ",", "-.--.": "(", "-.--.-": ")", "-.-.--": "!",
		"-.-.-.": ";", "-..-.": "/", "-...-": "=", "-....-": "-", ".-.-.": "+",
		".--.-.": "@", ".----.": "'", ".-.-.-": ".", ".-..-.": "\"", ".-...": "&",
		"..--.-": "_", "..--..": "?", "...-..-": "$",

		"...---...": "SOS",
	}
	var result string
	morseCode = strings.TrimLeft(morseCode, " ")
	morseCode = strings.Replace(morseCode, "   ", "  ", -1)
	for _, code := range strings.Split(morseCode, " ") {
		if code == "" {
			result += " "
		} else {
			result += MORSE_CODE[code]
		}
	}
	return result
}

func main() {
	fmt.Println(decodeBitsAdvanced("0000000011011010011100000110000001111110100111110011111100000000000111011111111011111011111000000101100011111100000111110011101100000100000"))
}
