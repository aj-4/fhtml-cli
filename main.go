package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	htmlStr := ""
	for scanner.Scan() {
		htmlStr += scanner.Text()
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

		// fmt.Println("called", inClosingTag, inOpeningTag)

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

		if nextChar == "<" && nextNextChar != "/" {
			// close opening tag
			if char == ">" {
				printHTMLLine(lineBuffer, indentIdx, true)
				isSelfClosingTag := strings.Contains(lineBuffer, "meta") || strings.Contains(lineBuffer, "img")
				if !isSelfClosingTag {
					indentIdx++
				}
				inClosingTag = false
			} else if lineBuffer != "" {
				printHTMLLine(lineBuffer, indentIdx, false)
			}
			lineBuffer = ""
			inOpeningTag = true
			continue
		}

		if inOpeningTag && char == ">" {
			isSelfClosingTag := strings.Contains(lineBuffer, "meta") || strings.Contains(lineBuffer, "<img")
			printHTMLLine(lineBuffer, indentIdx, true)
			lineBuffer = ""
			if !isSelfClosingTag {
				indentIdx++
			}
			inOpeningTag = false
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
		if strings.Contains(lineStr, "HTML") {
			color.Red(lineStr)
		} else if strings.Contains(lineStr, "HEAD") {
			color.Magenta(lineStr)
		} else if strings.Contains(lineStr, "BODY") {
			color.Green(lineStr)
		} else {
			color.Cyan(lineStr)
		}

	} else {
		fmt.Println(lineStr)
	}
}
