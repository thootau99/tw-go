package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var AllowedHeads = []string{"t", "p", "k", "th", "ph", "kh"}
var AllowedVowels = []string{"a", "i", "u", "e", "o"}
var buf = bytes.Buffer{}

func contains(s []string, e string) bool {
	for _, s := range s {
		if s == e {
			return true
		}
	}
	return false
}

func bufferJudge() {
	bufHead := bytes.Buffer{}
	bufVowels := bytes.Buffer{}
	bufTail := bytes.Buffer{}
	findVowels := false
	for {
		judgeCharacter, err := buf.ReadByte()
		if nil != err {
			break
		}
		if contains(AllowedVowels, string(judgeCharacter)) {
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
	vowelValue, _ := bufVowels.ReadString('\n')
	tailValue, _ := bufTail.ReadString('\n')
	fmt.Printf("Head: %s \n", headValue)
	fmt.Printf("vowel: %s \n", vowelValue)
	fmt.Printf("tail: %s \n", tailValue)
}

func parser(word string) {
	buf.Reset()
	buf.WriteString(word)
	bufferJudge()
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	input, _ := buf.ReadString('\n')
	parser(input)
}
