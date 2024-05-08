package main

import (
    "fmt"
    "os"
    "strings"
    "ascii-art/ascii"
)

func main() {
    // Validate command line arguments and determine the filename of the ASCII art file
    manip, fn := ascii.ValidateAndDetermineFilename()
    if !manip {
        // If the arguments are not valid, print an error message and return
        fmt.Println("\x1b[38;5;9m\x1b[40mARGS missing EX : go run . [STRING] [BANNER]\nEX: go run . something standard\x1b[38;5;9m\x1b[40m")
        return
    }
    // Convert escape sequences in the first argument
    param := ascii.ConvertEscapeSequences(os.Args[1])
    // Remove characters before a backspace character in the first argument
    param = ascii.RemoveCharactersBeforeBackspace(param)

    // Check if the first argument contains only printable ASCII characters or whitespace
    if !ascii.ContainsOnlyPrintableOrWhitespace(param) {
        // If it contains any other characters, print an error message and exit
        fmt.Println("Input contains invalid characters from banner")
            os.Exit(0)
    }

    // Split the first argument by newline characters
    input := strings.Split(param, "\\n")

    // Check if the input contains only whitespace or newline characters
    if valid, ifSpace, ind := ascii.CheckForWhitespaceOrNewline(input); valid && ifSpace == "NEWLINE" {
        // If it contains only newline characters, print the appropriate number of newlines and exit
        for i := 0; i < ind-1; i++ {
            fmt.Println()
        }
        os.Exit(0)
    } else if valid, ifSpace, _ := ascii.CheckForWhitespaceOrNewline(input); valid && ifSpace == "SPACE" {
        // If it contains only spaces, exit
        os.Exit(0)
    } else {
        // Load the ASCII art from the file
        new, AsciiSymbol := ascii.LoadAsciiArtFromFile(fn)
        if !new {
            // If there's an error loading the ASCII art, exit
            os.Exit(1)
        }

        // Print each string in the input in ASCII art
        for _, str := range input {
            if str == "" {
                // If the string is empty, print a newline
                fmt.Println()
            } else {
                // Otherwise, print the string in ASCII art
                ascii.PrintWordInAsciiArt(str,AsciiSymbol)
            }
        }
    }
}