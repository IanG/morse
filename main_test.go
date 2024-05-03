package main

import (
	"errors"
	"slices"
	"testing"
)

func TestToMorseThrowsError(t *testing.T) {
	expectedError := errors.New("Unsupported morse character: %")

	morse, err := toMorse("%ELLO WORLD")

	if err == nil {
		t.Error("toMorse should return error")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("toMorse should return error %q, got %q", expectedError, err)
	}

	if morse != "" {
		t.Errorf("toMorse should have returned empty text")
	}
}

func TestToMorse(t *testing.T) {
	expectedMorse := ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."

	morse, err := toMorse("HELLO WORLD")

	if err != nil {
		t.Error("Error should be nil")
	}

	if morse != expectedMorse {
		t.Errorf("Expected morse to be '%s', got '%s'", expectedMorse, morse)
	}
}

func TestFromMorseThrowsError(t *testing.T) {
	expectedError := errors.New("Unsupported morse sequence: -------")

	text, err := fromMorse("------- . .-.. .-.. --- / .-- --- .-. .-.. -..")

	if err == nil {
		t.Errorf("fromMorse should have returned an error")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("fromMorse should have returned error with '%s', got '%s'", expectedError.Error(), err.Error())
	}

	if text != "" {
		t.Errorf("fromMorse should have returned empty text")
	}
}

func TestFromMorse(t *testing.T) {
	expectedText := "HELLO WORLD"

	text, err := fromMorse(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")

	if err != nil {
		t.Error("Got error while converting to Morse")
	}

	if text != expectedText {
		t.Errorf("Got '%s' instead of '%s'", text, expectedText)
	}
}

func TestMorseCodeToCharThrowsError(t *testing.T) {
	morseCode := "X"
	expectedChar := ""

	char, err := morseCodeToChar(morseCode)

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if char != expectedChar {
		t.Errorf("Expected empty string, got '%s' instead of '%s", char, expectedChar)
	}
}

func TestMorseCodeToChar(t *testing.T) {
	type moreCodeToCharTestCase struct {
		input, expected string
	}

	var testCases = []moreCodeToCharTestCase{
		{".-", "A"},
		{"-...", "B"},
		{"/", " "},
	}

	for _, testCase := range testCases {
		char, err := morseCodeToChar(testCase.input)

		if err != nil {
			t.Errorf("Input %s should not return error", testCase.input)
		}

		if char != testCase.expected {
			t.Errorf("Returned char was '%s' instead of '%s'", char, testCase.expected)
		}
	}
}

func TestGetMorseWords(t *testing.T) {
	expectedMorseWords := []string{".... . .-.. .-.. ---", ".-- --- .-. .-.. -.."}
	morseWords := getMorseWords(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")

	if len(morseWords) != len(expectedMorseWords) {
		t.Errorf("Length of morseWords was %d and not %d", len(morseWords), len(expectedMorseWords))
	}

	if !slices.Equal(morseWords, expectedMorseWords) {
		t.Errorf("Returned morse words do not match expected")
	}
}

func TestGetMorseCodes(t *testing.T) {
	expectedMorseCodes := []string{"....", ".", ".-..", ".-..", "---"}
	morseCodes := getMorseCodes(".... . .-.. .-.. --- ")

	if len(morseCodes) != len(expectedMorseCodes) {
		t.Errorf("Length of morseCodes was %d and not %d", len(morseCodes), len(expectedMorseCodes))
	}

	if !slices.Equal(morseCodes, expectedMorseCodes) {
		t.Errorf("Returned morse codes do not match expected")
	}
}
