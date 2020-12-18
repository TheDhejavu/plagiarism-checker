package main

import (
	"fmt"
	"math"
	"strings"
)

type RabinKarp struct {
	Base        int
	Text        string
	PatternSize int
	Start       int
	End         int
	Mod         int
	Hash        int
}

func NewRabinKarp(text string, patternSize int) *RabinKarp {
	rb := &RabinKarp{
		Base:        26,
		PatternSize: patternSize,
		Start:       0,
		End:         0,
		Mod:         5807,
		Text:        text,
	}

	rb.GetHash()

	return rb
}

//GetHash creates hash for the first window
func (rb *RabinKarp) GetHash() {
	hashValue := 0
	for n := 0; n < rb.PatternSize; n++ {
		runeText := []rune(rb.Text)
		value := int(runeText[n])
		mathpower := int(math.Pow(float64(rb.Base), float64(rb.PatternSize-n-1)))
		hashValue = (hashValue + (value-96)*(mathpower)) % rb.Mod
	}
	rb.Start = 0
	rb.End = rb.PatternSize

	rb.Hash = hashValue
}

// NextWindow returns boolean and create new hash for the next window
// using rolling hash technique
func (rb *RabinKarp) NextWindow() bool {
	if rb.End <= len(rb.Text)-1 {
		mathpower := int(math.Pow(float64(rb.Base), float64(rb.PatternSize-1)))
		rb.Hash -= (int([]rune(rb.Text)[rb.Start]) - 96) * mathpower
		rb.Hash *= rb.Base
		rb.Hash += int([]rune(rb.Text)[rb.End] - 96)
		rb.Hash %= rb.Mod
		rb.Start++
		rb.End++
		return true
	}
	return false
}

// CurrentWindowText return the current window text
func (rb *RabinKarp) CurrentWindowText() string {
	return rb.Text[rb.Start:rb.End]
}

func Checker(text, pattern string) string {
	textRolling := NewRabinKarp(strings.ToLower(text), len(pattern))
	patternRolling := NewRabinKarp(strings.ToLower(pattern), len(pattern))

	for i := 0; i <= len(text)-len(pattern)+1; i++ {
		fmt.Println(textRolling.Hash, patternRolling.Hash)
		if textRolling.Hash == patternRolling.Hash {
			return "Found"
		}
		textRolling.NextWindow()
	}
	return "Not Found"
}

// func main() {
// 	text := "ABDCCEAG"
// 	pattern := "ag"
// 	fmt.Println(Checker(text, pattern))
// }
