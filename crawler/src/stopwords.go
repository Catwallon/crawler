package main

import (
	"bufio"
	"os"
	"strings"
)

type Stopwords struct {
	En []string
	Fr []string
	Es []string
	De []string
}

var stopwords Stopwords

func loadStopwords() {
	loadFromFile(&stopwords.En, "stopwords/en.txt")
	loadFromFile(&stopwords.Fr, "stopwords/fr.txt")
	loadFromFile(&stopwords.Es, "stopwords/es.txt")
	loadFromFile(&stopwords.De, "stopwords/de.txt")
}

func loadFromFile(field *[]string, filename string) {
	file, err := os.Open(filename)
	checkError("Can't open file "+filename, err, true)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		*field = append(*field, strings.Fields(line)...)
	}
	err = scanner.Err()
	checkError("Can't read file "+filename, err, true)
}

func removeWords(list1, list2 []string) []string {
	set := make(map[string]struct{})
	for _, word := range list2 {
		set[word] = struct{}{}
	}
	var result []string
	for _, word := range list1 {
		if _, found := set[word]; !found {
			result = append(result, word)
		}
	}
	return result
}

func removeStopwords(words []string, lang string) []string {
	if strings.HasPrefix(lang, "en") {
		return removeWords(words, stopwords.En)
	} else if strings.HasPrefix(lang, "fr") {
		return removeWords(words, stopwords.Fr)
	} else if strings.HasPrefix(lang, "es") {
		return removeWords(words, stopwords.Es)
	} else if strings.HasPrefix(lang, "de") {
		return removeWords(words, stopwords.De)
	}
	return words
}
