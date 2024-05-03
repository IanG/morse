package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

const space = " "

var morseChars = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	" ": "/",
}

type Command struct {
	Name        string
	Description string
	Run         func(args []string) error
}

var commands = []Command{
	{Name: "morse", Description: "Converts text into morse code", Run: morseFromText},
	{Name: "text", Description: "Converts morse code into text", Run: textFromMorse},
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "morse is a simple utility to work with morse code.")
	_, _ = fmt.Fprintln(os.Stderr, "\nCommands:")
	_, _ = fmt.Fprintln(os.Stderr, "  morse --text \"<text>\"")
	_, _ = fmt.Fprintln(os.Stderr, "  text --morse \"<morse-code>\"")
	_, _ = fmt.Fprintf(os.Stderr, "\nRun `morse <command> -h` to get help for a specific command.\n\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cmd := flag.Arg(0)
	cmdArgs := flag.Args()[1:]

	runCommand(cmd, cmdArgs)
}

func runCommand(name string, args []string) {
	cmdIdx := slices.IndexFunc(commands, func(cmd Command) bool {
		return cmd.Name == name
	})

	if cmdIdx < 0 {
		_, _ = fmt.Fprintf(os.Stderr, "command \"%s\" not found\n\n", name)
		flag.Usage()
		os.Exit(1)
	}

	if err := commands[cmdIdx].Run(args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func morseFromText(args []string) error {
	var text string

	flagSet := flag.NewFlagSet("morse", flag.ExitOnError)
	flagSet.StringVar(&text, "text", "", "The text to convert to morse code")
	flagSet.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "Usage:")
		flagSet.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr)
	}

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	morse, err := toMorse(text)
	if err == nil {
		fmt.Println(morse)
	}

	return err
}

func textFromMorse(args []string) error {
	var morse string

	flagSet := flag.NewFlagSet("text", flag.ExitOnError)
	flagSet.StringVar(&morse, "morse", "", "The morse code to convert to text")
	flagSet.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "Usage:")
		flagSet.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr)
	}

	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	text, err := fromMorse(morse)
	if err == nil {
		fmt.Println(text)
	}

	return err
}

func toMorse(text string) (string, error) {
	var morse strings.Builder

	for _, char := range regexp.MustCompile(`\s+`).ReplaceAllString(strings.ToUpper(text), space) {
		morseVal, charFound := morseChars[string(char)]

		if charFound {
			morse.WriteString(fmt.Sprintf("%s ", morseVal))
		} else {
			return "", errors.New(fmt.Sprintf("Unsupported morse character: %s", string(char)))
		}
	}
	return strings.TrimRight(morse.String(), " "), nil
}

func fromMorse(morse string) (string, error) {
	var text strings.Builder

	var morseWords = getMorseWords(morse)

	for morseWordIdx, morseWord := range morseWords {
		for _, morseCode := range getMorseCodes(morseWord) {
			char, charErr := morseCodeToChar(morseCode)

			if charErr == nil {
				text.WriteString(char)
			} else {
				return "", errors.New(fmt.Sprintf("Unsupported morse sequence: %s", morseCode))
			}
		}

		if morseWordIdx != len(morseWords)-1 {
			text.WriteString(space)
		}
	}

	return text.String(), nil
}

func morseCodeToChar(morseCode string) (string, error) {
	for k, v := range morseChars {
		if v == morseCode {
			return k, nil
		}
	}

	return "", errors.New(fmt.Sprintf("Unsupported morse code: %s", morseCode))
}

func getMorseWords(morsePhrase string) []string {
	var morseWords []string

	wordRegex := regexp.MustCompile("([.\\s\\-])+\\w*")

	for _, wordMatch := range wordRegex.FindAllStringSubmatch(morsePhrase, -1) {
		morseWords = append(morseWords, strings.Trim(wordMatch[0], " "))
	}

	return morseWords
}

func getMorseCodes(morseWord string) []string {
	var morseCodes []string

	codeRegex := regexp.MustCompile("([.\\-])+\\w*")

	for _, codeMatch := range codeRegex.FindAllStringSubmatch(morseWord, -1) {
		morseCodes = append(morseCodes, strings.Trim(codeMatch[0], " "))
	}

	return morseCodes
}
