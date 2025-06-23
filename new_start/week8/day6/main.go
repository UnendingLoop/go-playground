package main

import (
	"fmt"
	"strings"
)

// DecodeBits - translates bits of raw data into dots and dashes (00100111... => .--...--...)
func DecodeBits(bits string) string {
	bits = strings.Trim(bits, "0")     //get rid of excessive pre- and post-zeros
	multiplier := findMultiplier(bits) //calculating transmission rate
	fmt.Println("Multiplier is:", multiplier)

	dash := strings.Repeat("111", multiplier)
	dot := strings.Repeat("1", multiplier)
	//interDashDotGap := strings.Repeat("0", multiplier)
	interCharGap := strings.Repeat("000", multiplier)
	interWordGap := strings.Repeat("0000000", multiplier)

	morseLine := ""
	for i, j := 0, 0; i < len(bits); i = j {
		for j = i; j < len(bits) && bits[i] == bits[j]; {
			j++
		}
		switch {
		case bits[i:j] == dash:
			morseLine += "-"
		case bits[i:j] == dot:
			morseLine += "."
		case string(bits[i]) == "1" && (j-i+1 < len(dash) || j-i+1 > len(dot)):
			morseLine += "."
		case bits[i:j] == interCharGap:
			morseLine += " "
		case bits[i:j] == interWordGap:
			morseLine += "   "
		}
	}
	return morseLine
}

func findMultiplier(bits string) int {
	minCount := len(bits)
	for i, j := 0, 0; i < len(bits); i = j {
		for j = i; j < len(bits) && bits[i] == bits[j]; {
			j++
		}
		if j-i < minCount {
			minCount = j - i
		}
	}
	return minCount
}

// DecodeMorse - translates dots and dashes into readable text (.--.-.-... => ABCD...)
func DecodeMorse(morseCode string) string {
	listOfWords := strings.Split(morseCode, "   ")
	finalText := ""
	for i, w := range listOfWords {
		chars := strings.Split(w, " ")
		for _, c := range chars {
			c = c //finalText += MORSE_CODE[c]
		}
		if i != len(listOfWords)-1 {
			finalText += " "
		}
	}
	return finalText
}
func main() {
	input := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	input1 := "11110000111100001111000011110000000000001111000000000000111111111111000011110000111111111111000011111111111100000000000000000000000000001111000011111111111100001111111111110000111111111111000000000000111100001111000011111111111100000000000011111111111100001111000011110000000000001111"
	inputM := "110011"
	inputI := "11111100111111"
	inputEE := "1100000011"
	/*       "·     ·     ·     ·     _    ·   _      −        ·      −         −      _   __         ·     −         −         −      _      ·     ·     −      _      −         ·     ·     _    ·"
	//"|····| |·| |−·−−|   |·−−− ··− −·· ·|"
	|11001100110011| 000000 |11| 000000 |111111 00 11 00 111111 00 111111|   00000000000000 |11 00 111111 00 111111 00 111111| 000000 |11 00 11 00 111111| 000000 |111111 00 11 00 11| 000000 |11|
	    H                     E                           Y                       _                          J                                U                            D                    E
	dot  =          11
	wordspace =     00000000000000 - 14 '0's
	charspace =     000000 - 6 '0's
	dash&dotspace = 00
	dash =          111111 - 6 '1's

	···· · −·−−   ·−−− ··− −·· ·
	.... . -.--   .--- ..- -.. .

	1) сначала разбить на слова
	2) разбить слова на буквы
	*/
	decodedBits := DecodeBits(input)
	fmt.Printf("The result is: '%s'\n", decodedBits)
	decodedBits1 := DecodeBits(input1)
	fmt.Printf("The result is: '%s'\n", decodedBits1)

	decodedBitsM := DecodeBits(inputM)
	fmt.Printf("The result for M is: '%s'\n", decodedBitsM)
	decodedBitsI := DecodeBits(inputI)
	fmt.Printf("The result for I is: '%s'\n", decodedBitsI)
	decodedBitsEE := DecodeBits(inputEE)
	fmt.Printf("The result for EE is: '%s'\n", decodedBitsEE)
}
