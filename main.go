package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var buf = bytes.Buffer{}

func contains_in_string_slice(s []string, e string) bool {
	for _, s := range s {
		if s == e {
			return true
		}
	}
	return false
}

func contains_in_rune_slice(s []rune, e rune) bool {
	for _, s := range s {
		if s == e {
			return true
		}
	}
	return false
}

func replaceVowelWithTone(tone Tones, vowel bytes.Buffer) string {
	copyVowel := bytes.NewBuffer(vowel.Bytes())
	firstVowel, _ := copyVowel.ReadByte()
	var result = ""
	switch tone {
	case Acute:
		result = addAcuteOnVowel(string(firstVowel))
	case Grave:
		result = addGraveOnVowel(string(firstVowel))
	case Circumflex:
		result = addCircumflexOnVowel(string(firstVowel))
	case Macron:
		result = addMacronOnVowel(string(firstVowel))
	case VerticalLineAbove:
		result = addVerticalLineAboveOnVowel(string(firstVowel))
	}
	newBuffer := bytes.Buffer{}
	newBuffer.WriteString(result)
	newBuffer.WriteString(copyVowel.String())
	return newBuffer.String()
}

func bufferJudge() {
	bufHead := bytes.Buffer{}
	bufVowels := bytes.Buffer{}
	bufTail := bytes.Buffer{}
	var tone Tones
	findVowels := false
	for {
		judgeCharacter, err := buf.ReadByte()
		if nil != err {
			break
		}

		// Check for tone word
		if contains_in_rune_slice(TonesKeyMap, rune(judgeCharacter)) {
			switch rune(judgeCharacter) {
			case 'x':
				tone = Acute
			case 'd':
				tone = Grave
			case 'f':
				tone = VerticalLineBelow
			case 'y':
				tone = Circumflex
			case 'q':
				tone = Macron
			case 'w':
				tone = VerticalLineAbove
			}
			continue
		}

		// check if Vowels appear
		// consider all previous characters are used to initial a word if a vowel has appeared.
		if contains_in_string_slice(AllowedVowels, string(judgeCharacter)) {
			bufVowels.WriteByte(judgeCharacter)
			findVowels = true
			continue
		}
		if !findVowels {
			bufHead.WriteByte(judgeCharacter)
		} else {
			bufTail.WriteByte(judgeCharacter)
		}
	}
	headValue, _ := bufHead.ReadString('\n')
	//vowelValue, _ := bufVowels.ReadString('\n')
	tailValue, _ := bufTail.ReadString('\n')
	fmt.Printf("Head: %s \n", headValue)
	fmt.Printf("vowel: %s \n", replaceVowelWithTone(tone, bufVowels))
	fmt.Printf("tail: %s \n", tailValue)
	fmt.Printf("tone: %d \n", tone)
}

func parser(word string) {
	buf.Reset()
	buf.WriteString(word)
	bufferJudge()
}

func main() {
	for true {
		buf := bufio.NewReader(os.Stdin)
		input, _ := buf.ReadString('\n')
		parser(input)
	}
}
