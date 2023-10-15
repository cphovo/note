package utils

import (
	"sort"
	"strings"

	"github.com/yanyiwu/gojieba"
)

type JieBa struct {
	DictPath      string `json:"dict_path"`
	HmmPath       string `json:"hmm_path"`
	UserDictPath  string `json:"user_dict_path"`
	IdfPath       string `json:"idf_path"`
	StopWordsPath string `json:"stop_words_path"`
}

type WordCount struct {
	Word  string
	Count int
}

// Function to count the content in the `text` by lexicality and sort it in descending order of lexical frequency.
// If no tag is provided, all lexemes are counted.
// The lexical tags can be found in https://github.com/fxsjy/jieba
func (j *JieBa) GetWordCount(text string, tags ...string) []WordCount {
	// Set cutWithoutTag to false if no tags are provided
	cutWithoutTag := false
	if len(tags) == 0 {
		cutWithoutTag = true
	}

	// Create a new Jieba instance with the provided paths
	x := gojieba.NewJieba(
		j.DictPath,
		j.HmmPath,
		j.UserDictPath,
		j.IdfPath,
		j.StopWordsPath,
	)
	defer x.Free()

	wordTags := x.Tag(text)

	tagSet := make(map[string]bool)
	for _, tag := range tags {
		tagSet[tag] = true
	}

	wordCountMap := make(map[string]int)

	for _, wordTag := range wordTags {
		parts := strings.Split(wordTag, "/")
		if len(parts) == 2 {
			word := parts[0]
			tag := parts[1]

			// If cutWithoutTag is true, increment the word count for all words
			if cutWithoutTag {
				wordCountMap[word]++
				continue
			}

			// If cutWithoutTag is false and the tag is in the tag set, increment the word count
			if _, exists := tagSet[tag]; exists {
				wordCountMap[word]++
			}
		}
	}

	var wordCount []WordCount
	for word, count := range wordCountMap {
		wordCount = append(wordCount, WordCount{
			Word:  word,
			Count: count,
		})
	}

	// Sort the word count slice in descending order
	sort.Slice(wordCount, func(i, j int) bool {
		return wordCount[i].Count > wordCount[j].Count
	})

	return wordCount
}
