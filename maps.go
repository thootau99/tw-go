package main

var AllowedHeads = []string{"t", "p", "k", "b", "h", "s", "n", "j", "l", "th", "ph", "kh"}
var AllowedVowels = []string{"a", "i", "u", "e", "o"}
var TonesKeyMap = []rune{'x', 'd', 'f', 'y', 'q'}

var AcuteVowels = map[string]string{
	"a":  "á",
	"i":  "í",
	"u":  "ú",
	"e":  "é",
	"o":  "ó",
	"oo": "ó͘",
	// UPPERCASE
	"A":  "Á",
	"I":  "Í",
	"U":  "Ú",
	"E":  "É",
	"O":  "Ó",
	"OO": "Ó͘",
}

var GraveVowels = map[string]string{
	"a":  "à",
	"i":  "ì",
	"u":  "ù",
	"e":  "è",
	"o":  "ò",
	"oo": "ò͘",
	// UPPERCASE
	"A":  "À",
	"I":  "Ì",
	"U":  "Ù",
	"E":  "È",
	"O":  "Ò",
	"OO": "Ò͘",
}

var CircumflexVowels = map[string]string{
	"a":  "â",
	"i":  "î",
	"u":  "û",
	"e":  "ê",
	"o":  "ô",
	"oo": "ô͘",
	// UPPERCASE
	"A":  "Â",
	"I":  "Î",
	"U":  "Û",
	"E":  "Ê",
	"O":  "Ô",
	"OO": "Ô͘",
}

var MacronVowels = map[string]string{
	"a":  "ā",
	"i":  "ī",
	"u":  "ū",
	"e":  "ē",
	"o":  "ō",
	"oo": "ō͘",
	// UPPERCASE
	"A":  "Ā",
	"I":  "Ī",
	"U":  "Ū",
	"E":  "Ē",
	"O":  "Ō",
	"OO": "Ō͘",
}

var VerticalLineAboveVowels = map[string]string{
	"a":  "a̍",
	"i":  "i̍",
	"u":  "u̍",
	"e":  "e̍",
	"o":  "o̍",
	"oo": "o̍͘",
	// UPPERCASE
	"A":  "A̍",
	"I":  "I̍",
	"U":  "U̍",
	"E":  "E̍",
	"O":  "O̍",
	"OO": "O̍͘",
}

type Tones int

const (
	Acute             = 2
	Grave             = 3
	VerticalLineBelow = 4
	Circumflex        = 5
	Macron            = 7
	VerticalLineAbove = 8
)
