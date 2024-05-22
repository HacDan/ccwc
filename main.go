package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name        string
	description string
	callback    func([]byte) string
}

func main() {
	var stdin []byte
	commands := getCommands()
	if len(os.Args) > 2 || len(os.Args) < 1 {
		cmdHelp(stdin)
		os.Exit(1)
	}
	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	} else {
		stdin, err := os.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	if len(os.Args) > 2 || len(os.Args) < 1 {
		cmdHelp(stdin)
		os.Exit(1)
	}

	var byteFlag = flag.Bool("c", false, commands["bytes"].description)
	var charFlag = flag.Bool("m", false, commands["chars"].description)
	var lineFlag = flag.Bool("l", false, commands["lines"].description)
	var wordFlag = flag.Bool("w", false, commands["words"].description)

	var helpFlag = flag.Bool("help", false, commands["help"].description)
	var versionFlag = flag.Bool("version", false, commands["version"].description)
	flag.Parse()

	if *helpFlag {
		cmdHelp()
		os.Exit(0)
	}
	if *versionFlag {
		cmdVersion()
		os.Exit(0)
	}
	if *byteFlag {
		cmdBytes()
		os.Exit(0)
	}
	if *charFlag {
		cmdChars()
		os.Exit(0)
	}
	if *lineFlag {
		cmdLines()
		os.Exit(0)
	}
	if *wordFlag {
		cmdWords()
		os.Exit(0)
	}

	lineSep := []byte{'\n'}

	wordCount := cmdWordCount(&testFile)
	byteCount := len(testFile)
	lineCount := bytes.Count(testFile, lineSep)

	fmt.Println("Wordcount: ", wordCount)
	fmt.Println("Bytecount: ", byteCount)
	fmt.Println("Linecount: ", lineCount)
}

func cmdWordCount(f *[]byte) int {
	return len(strings.Fields(string(*f)))
}

func getCommands() map[string]Command {
	return map[string]Command{
		"bytes": {
			name:        "c",
			description: "print the byte counts",
			callback:    cmdBytes,
		},
		"chars": {
			name:        "m",
			description: "print the character counts",
			callback:    cmdChars,
		},
		"lines": {
			name:        "l",
			description: "print the newLine counts",
			callback:    cmdLines,
		},
		"words": {
			name:        "w",
			description: "print the word counts",
			callback:    cmdWords,
		},
		"help": {
			name:        "help",
			description: "prints the help text",
			callback:    cmdHelp,
		},
		"version": {
			name:        "version",
			description: "prints the version information",
			callback:    cmdVersion,
		},
	}
}
