# Morse

## Introduction

This is a small utility which can be used to convert text into morse code and vice versa. 

This utility was built as my first usage of the [Go](https://go.dev/) programming language and was an opportunity to
understand some of its key concepts and commonly used practices such as:

- Range
- Slices
- Maps
- Function passing
- Command line parameter parsing
- Regular expressions and capture groups

I wanted to do all this by only using what is available in the [Go standard library](https://pkg.go.dev/std) even
though there are no doubt numerous other packages available.  Its also worth pointing out that a 3rd party library
[gSpera/morse](https://github.com/gSpera/morse) already exists to do text/morse transformations - however - by building 
the utility it was an opportunity to understand some of Go's `range`, `slice` and `map` operations.

## Tests

You can run the associated tests with:

`go test morse`

## Usage

You can run `go run main.go -h` to get help on how to use the utility

### Convert Text To Morse Code
Run

`go run main.go morse --text "Hello World"`

Which will return:

`.... . .-.. .-.. --- / .-- --- .-. .-.. -..`

### Convert Morse Code To Text
Run
`go run main.go text --morse ".... . .-.. .-.. --- / .-- --- .-. .-.. -.."`

Which will return:

`HELLO WORLD`

## Observations 
- The Go `regex` library does not support named capture groups at present and is discussed under
[#64108](https://github.com/golang/go/issues/64108). Only a single repeating capture group was required for this so
not being able to name it was not a serious issue.
- The `flag` library is a little clunky - especially for sub-command processing.

## Tools used
- A zsh shell
- [JetBrains GoLand IDE](https://www.jetbrains.com/go/)
