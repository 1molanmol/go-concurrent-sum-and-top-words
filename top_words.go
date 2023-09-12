package cos418_hw1_1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
//
//	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
//
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuation and other characters
// are removed, e.g. "don't" becomes "dont".

// topWords reads a text document from the specified path, tokenizes it into words,
// removes punctuation and converts to lowercase, and returns the top 'numWords'
// words that meet a minimum character threshold 'charThreshold'.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// Read the text document
	content, err := ioutil.ReadFile(path)
	checkError(err)

	// Tokenize the text into words, remove punctuation, and convert to lowercase
	words := strings.Fields(string(content))
	wordCounts := make(map[string]int)

	// Define a regular expression to remove non-alphanumeric characters
	regex := regexp.MustCompile("[^0-9a-zA-Z]+")

	// Loop through the words and count occurrences
	for _, word := range words {
		word = strings.ToLower(word)
		word = regex.ReplaceAllString(word, "") // Remove non-alphanumeric characters
		if len(word) >= charThreshold {
			wordCounts[word]++
		}
	}

	// Convert word counts to WordCount structs
	wordCountList := make([]WordCount, 0, len(wordCounts))
	for word, count := range wordCounts {
		wordCountList = append(wordCountList, WordCount{Word: word, Count: count})
	}

	// Sort the word counts
	sortWordCounts(wordCountList)

	// Take the top numWords word counts
	if numWords > len(wordCountList) {
		numWords = len(wordCountList)
	}
	return wordCountList[:numWords]
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
