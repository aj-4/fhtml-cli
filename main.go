package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need exactly 2 args")
		return
	}

	htmlStr := os.Args[1]

	if htmlStr[0] != '<' {
		fmt.Println("not valid html")
		return
	}

	printHTML(htmlStr)
}

func printHTML(htmlStr string) {
	indentIdx := 0
	lineBuffer := ""
	firstChar := true
	inClosingTag := false
	inOpeningTag := true

	for i := range htmlStr {
		char := string(htmlStr[i])
		var nextChar, nextNextChar string

		if i == len(htmlStr)-1 {
			nextChar = ""
		} else {
			nextChar = string(htmlStr[i+1])
		}

		if i >= len(htmlStr)-2 {
			nextNextChar = ""
		} else {
			nextNextChar = string(htmlStr[i+2])
		}

		lineBuffer += char

		if firstChar {
			firstChar = false
			continue
		}

		if nextChar == "<" && nextNextChar != "/" {
			if char == ">" {
				printHTMLLine(lineBuffer, indentIdx, true)
				lineBuffer = ""
				indentIdx++
			} else if lineBuffer != "" {
				printHTMLLine(lineBuffer, indentIdx, false)
			}
			inOpeningTag = true
			continue
		}

		if inOpeningTag && char == ">" {
			printHTMLLine(lineBuffer, indentIdx, true)
			lineBuffer = ""
			indentIdx++
			inOpeningTag = false
			continue
		}

		if nextChar == "<" && nextNextChar == "/" {
			if inClosingTag || inOpeningTag {
				indentIdx--
				printHTMLLine(lineBuffer, indentIdx, true)
			} else {
				printHTMLLine(lineBuffer, indentIdx, false)
			}
			lineBuffer = ""
			inClosingTag = true
			continue
		}
		if inClosingTag && char == ">" {
			indentIdx--
			printHTMLLine(lineBuffer, indentIdx, true)
			lineBuffer = ""
			inClosingTag = false
			continue
		}
	}
}

func printHTMLLine(buffer string, whiteSpace int, useColor bool) {
	padding := ""
	for i := 0; i < whiteSpace; i++ {
		padding += "  "
	}
	lineStr := fmt.Sprintf("%s%s", padding, buffer)
	if useColor {
		color.Cyan(lineStr)
	} else {
		fmt.Println(lineStr)
	}
}
