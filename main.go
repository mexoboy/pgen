package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"golang.org/x/term"
	"math/big"
	"os"
	"strings"
)

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars    = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?'`~"
)

type Config struct {
	lowercase    bool
	chars        bool
	numbers      bool
	specialChars bool
	length       int
	count        int
}

func main() {
	config := parseFlags()

	charset := buildCharset(config)
	if charset == "" {
		fmt.Println("Error: No character types selected for password generation")
		os.Exit(1)
	}

	passwords := generatePasswords(config.count, config.length, charset)
	printPasswordsInColumns(passwords)
}

func parseFlags() Config {
	config := Config{}

	flag.BoolVar(&config.lowercase, "lowercase", false, "use only lowercase characters")
	flag.BoolVar(&config.chars, "chars", true, "use alphabetic characters")
	flag.BoolVar(&config.numbers, "numbers", true, "use numeric characters")
	flag.BoolVar(&config.specialChars, "special-chars", false, "use special characters")
	flag.IntVar(&config.length, "length", 16, "password length")
	flag.IntVar(&config.count, "count", 20, "count of generated passwords")

	flag.Parse()

	return config
}

func buildCharset(config Config) string {
	var charset strings.Builder

	if config.chars {
		if config.lowercase {
			charset.WriteString(lowercaseChars)
		} else {
			charset.WriteString(lowercaseChars)
			charset.WriteString(uppercaseChars)
		}
	}

	if config.numbers {
		charset.WriteString(numberChars)
	}

	if config.specialChars {
		charset.WriteString(specialChars)
	}

	return charset.String()
}

func generatePasswords(count, length int, charset string) []string {
	passwords := make([]string, count)
	charsetRunes := []rune(charset)
	charsetLen := big.NewInt(int64(len(charsetRunes)))

	for i := 0; i < count; i++ {
		password := make([]rune, length)
		for j := 0; j < length; j++ {
			randomIndex, err := rand.Int(rand.Reader, charsetLen)
			if err != nil {
				fmt.Printf("Random number generation error: %v\n", err)
				os.Exit(1)
			}
			password[j] = charsetRunes[randomIndex.Int64()]
		}
		passwords[i] = string(password)
	}

	return passwords
}

func printPasswordsInColumns(passwords []string) {
	if len(passwords) == 0 {
		return
	}

	terminalWidth := getTerminalWidth()
	passwordLength := len(passwords[0])
	columnsCount := terminalWidth / (passwordLength + 2)

	if columnsCount < 1 {
		columnsCount = 1
	}

	for i := 0; i < len(passwords); i += columnsCount {
		var row strings.Builder

		for j := 0; j < columnsCount && i+j < len(passwords); j++ {
			if j > 0 {
				row.WriteString("  ")
			}
			row.WriteString(passwords[i+j])
		}

		fmt.Println(row.String())
	}
}

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80
	}
	return width
}
