package ascii

import (
	"fmt"
	"os"
	"strings"
)

// ContainsOnlyPrintableOrWhitespace checks if the input string contains only printable ASCII characters or whitespace.
// It returns false and prints an error message if the string contains any other characters.
func ContainsOnlyPrintableOrWhitespace(args string) bool {
	for _, ch := range args {
		if !((ch >= 32 && ch <= 126) || (ch >= 8 && ch <= 13)) {
			fmt.Println("Error: Input contains invalid characters")
			return false
		}
	}
	return true
}

// ConvertEscapeSequences replaces escape sequences in the input string with their corresponding characters.
// For example, it replaces "\\t" with four spaces.
func ConvertEscapeSequences(s string) string {
	replacer := strings.NewReplacer(
		"\\v", "\\n\\n\\n\\n",
		"\n", "\\n",
		"\\t", "    ",
		"\\b", "\b",
		"\\r", "\r",
		"\\f", "\f",
	)
	return replacer.Replace(s)
}

// CheckForWhitespaceOrNewline checks if the input string array contains only whitespace or newline characters.
// It returns a boolean indicating whether the array contains only whitespace or newline characters,
// a string indicating the type of whitespace, and the count of whitespace characters.
func CheckForWhitespaceOrNewline(ar []string) (bool, string, int) {
	count := 0
	for _, ch := range ar {
		if ch != "" {
			return false, "", count
		} else {
			count++
		}
	}

	if count == 1 {
		return true, "SPACE", count
	} else {
		return true, "NEWLINE", count
	}
}

// PrintWordInAsciiArt prints a word in ASCII art.
// It takes a string and a map of runes to ASCII art strings as input.
func PrintWordInAsciiArt(word string, AsciiSymbol map[rune][]string) {
	for i := 0; i < GetCharacterHeight(); i++ {
		lineOutput := ""
		beforeCarriage := ""
		for _, l := range word {
			switch l {
			case '\n':
				fmt.Println(lineOutput)
				lineOutput = ""
			case '\r':
				beforeCarriage = lineOutput
				lineOutput = ""
			default:
				if i < len(AsciiSymbol[l]) {
					lineOutput += AsciiSymbol[l][i]
				}
			}
		}
		if len(beforeCarriage) > len(lineOutput) {
			lineOutput += beforeCarriage[len(lineOutput):]
		}
		fmt.Println(lineOutput)
	}
}

// ValidateAndDetermineFilename validates the command line arguments and determines the filename of the ASCII art file.
// It returns a boolean indicating whether the arguments are valid and the filename.
func ValidateAndDetermineFilename() (bool, string) {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		return false, ""
	}

	filename := "standard.txt"
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard", "standard.txt":
			filename = "standard.txt"
		case "shadow", "shadow.txt":
			filename = "shadow.txt"
		case "thinkertoy", "thinkertoy.txt":
			filename = "thinkertoy.txt"
		default:
			fmt.Printf("Could not create the alphabet. Are you sure %s exists and is a valid ascii file?\n", os.Args[2])
			os.Exit(1)
			return false, ""
		}
	}
	return true, filename
}

// RemoveCharactersBeforeBackspace removes characters before a backspace character in the input string.
func RemoveCharactersBeforeBackspace(s string) string {
	temp := ""
	for _, ch := range s {
		if ch != '\b' {
			temp += string(ch)
		} else if len(temp) > 0 {
			temp = temp[:len(temp)-1]
		}
	}
	return temp
}
