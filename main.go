package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	var file string
	var line string

	fmt.Print("Input a filename:")
	fmt.Scan(&file)
	data, err := scanWords(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Print("Input a word:")
		fmt.Scan(&line)
		fmt.Println(wordCencor(data, line))

		if line == "exit" {
			fmt.Println("Bye!")
			break
		}
	}
}

// scanWords - opens a reads the input file words.
func scanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Print(cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

// wordCencor - replace the obscene word in a sentence.
func wordCencor(data []string, text string) string {

	sentence := text
	reg := regexp.MustCompile(`[,. ]`)
	str := reg.Split(text, -1)
	for _, value := range data {
		for _, word := range str {

			if compareStrings := strings.EqualFold(value, word); compareStrings && word != "" {
				correctWord := strings.Repeat("*", len(value))
				sentence = strings.ReplaceAll(sentence, word, correctWord)
			}
		}
	}
	if ok := strings.ContainsAny(sentence, "*"); !ok {
		return text
	}

	return sentence
}
