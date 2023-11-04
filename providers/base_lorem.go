package providers

import (
	"errors"
	"github.com/harmlessprince/goFaker/constants"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"strings"
)

type BaseLorem struct {
	BaseProvider
	wordList []string
}

type LoremInterface interface {
	Word() (string, error)
	WordsAsText(numberOfWords int) (string, error)
	WordsAsList(numberOfWords int) ([]string, error)
	Sentence(numberOfWords int, variableNumberOfWords bool) (string, error)
	SentencesAsText(numberOfSentences int, variableNumberOfWords bool) (string, error)
	SentencesAsList(numberOfSentences int, variableNumberOfWords bool) ([]string, error)
	Paragraph(numberOfSentences int, variableNumberOfSentences bool) (string, error)
	ParagraphsAsList(numberOfParagraphs ...int) ([]string, error)
	ParagraphsAsText(numberOfParagraphs ...int) (string, error)
	Text(maxNumberOfCharacters ...int) (string, error)
}

func (bl *BaseLorem) GetWordList() []string {
	bl.wordList = constants.LoremWordList
	return bl.wordList
}

// Word generates a random word.
//
// Returns:
// - A string representing a randomly generated word.
func (bl *BaseLorem) Word() (string, error) {
	word, err := bl.RandomElementFromStringSlice(bl.GetWordList())
	if err != nil {
		return "", err
	}
	return word, nil
}

// WordsAsText generates a string of random words.
//
// Parameters:
// - numberOfWords: The number of words to generate.
//
// Returns:
// - A string representing a space-separated list of randomly generated words.
func (bl *BaseLorem) WordsAsText(numberOfWords int) (string, error) {
	var words []string
	for i := 0; i < numberOfWords; i++ {
		word, _ := bl.Word()
		words = append(words, word)
	}
	return strings.Join(words, " "), nil
}

// WordsAsList generates a list of random words.
//
// Parameters:
// - numberOfWords: The number of words to generate.
//
// Returns:
// - A slice of strings representing randomly generated words.
func (bl *BaseLorem) WordsAsList(numberOfWords int) ([]string, error) {
	var words []string
	for i := 0; i < numberOfWords; i++ {
		word, _ := bl.Word()
		words = append(words, word)
	}
	return words, nil
}

// Sentence generates a random sentence.
//
// Parameters:
// - numberOfWords: The number of words in the sentence.
// - variableNumberOfWords: Specify if the number of words can vary.
//
// Returns:
// - A string representing a randomly generated sentence.
func (bl *BaseLorem) Sentence(numberOfWords int, variableNumberOfWords bool) (string, error) {
	if numberOfWords <= 0 {
		return "", errors.New("numberOfWords should be greater than 0")
	}
	if variableNumberOfWords {
		numberOfWords = bl.randomizeNbElements(numberOfWords)
	}
	words, err := bl.WordsAsList(numberOfWords)
	if err != nil {
		return "", err
	}
	words[0] = cases.Title(language.Und).String(words[0])
	return strings.Join(words, " ") + ".", nil
}

// SentencesAsText generates a string of random sentences.
//
// Parameters:
// - numberOfSentences: The number of sentences to generate.
// - variableNumberOfWords: Specify if the number of words can vary in each sentence.
//
// Returns:
// - A string representing a space-separated list of randomly generated sentences.
func (bl *BaseLorem) SentencesAsText(numberOfSentences int, variableNumberOfWords bool) (string, error) {
	if numberOfSentences <= 0 {
		return "", errors.New("numberOfSentences should be greater than 0")
	}
	var sentences []string
	for i := 0; i < numberOfSentences; i++ {
		sentence, err := bl.Sentence(6, variableNumberOfWords)
		if err != nil {
			return "", err
		}
		sentences = append(sentences, sentence)
	}
	return strings.Join(sentences, " "), nil
}

// SentencesAsList generates a list of random sentences.
//
// Parameters:
// - numberOfSentences: The number of sentences to generate.
// - variableNumberOfWords: Specify if the number of words can vary in each sentence.
//
// Returns:
// - A slice of strings representing randomly generated sentences.
func (bl *BaseLorem) SentencesAsList(numberOfSentences int, variableNumberOfWords bool) ([]string, error) {
	if numberOfSentences <= 0 {
		return nil, errors.New("numberOfSentences should be greater than 0")
	}
	var sentences []string
	for i := 0; i < numberOfSentences; i++ {
		sentence, err := bl.Sentence(6, variableNumberOfWords)
		if err != nil {
			return nil, err
		}
		sentences = append(sentences, sentence)
	}
	return sentences, nil
}

// Paragraph generates a random paragraph.
//
// Parameters:
// - numberOfSentences: The number of sentences in the paragraph.
// - variableNumberOfSentences: Specify if the number of sentences can vary.
//
// Returns:
// - A string representing a randomly generated paragraph.
func (bl *BaseLorem) Paragraph(numberOfSentences int, variableNumberOfSentences bool) (string, error) {
	if numberOfSentences <= 0 {
		return "", errors.New("numberOfSentences should be greater than 0")
	}
	if variableNumberOfSentences {
		numberOfSentences = bl.randomizeNbElements(numberOfSentences)
	}
	sentences, err := bl.SentencesAsList(numberOfSentences, false)
	if err != nil {
		return "", err
	}
	return strings.Join(sentences, " "), nil
}

// ParagraphsAsList generates a list of random paragraphs.
//
// Parameters:
// - numberOfParagraphs: The number of paragraphs to generate.
//
// Returns:
// - A slice of strings representing randomly generated paragraphs.
func (bl *BaseLorem) ParagraphsAsList(numberOfParagraphs ...int) ([]string, error) {
	n := 3
	if len(numberOfParagraphs) > 0 {
		n = numberOfParagraphs[0]
	}

	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraph, err := bl.Paragraph(3, true)
		if err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, paragraph)
	}
	return paragraphs, nil
}

func (bl *BaseLorem) ParagraphsAsText(numberOfParagraphs ...int) (string, error) {
	n := 3
	if len(numberOfParagraphs) > 0 {
		n = numberOfParagraphs[0]
	}
	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraph, err := bl.Paragraph(3, true)
		if err != nil {
			return "", err
		}
		paragraphs = append(paragraphs, paragraph)
	}
	return strings.Join(paragraphs, "\n\n"), nil
}

// Text generates a random text with a specified maximum number of characters.
//
// Parameters:
// - maxNumberOfCharacters (optional): The maximum number of characters for the generated text.
//
// Returns:
// - A string representing randomly generated text.
func (bl *BaseLorem) Text(maxNumberOfCharacters ...int) (string, error) {
	maxN := 200
	if len(maxNumberOfCharacters) > 0 {
		maxN = maxNumberOfCharacters[0]
		if maxN < 5 {
			return "", errors.New("Text() can only generate text of at least 5 characters")
		}
	}

	texType := "paragraph"

	if maxN < 100 {
		texType = "sentence"
	}
	if maxN < 25 {
		texType = "word"
	}
	var text []string

	for len(text) == 0 {
		size := 0
		// until maxN is reached
		for size < maxN {
			word := "" + texType
			if size > 0 {
				word = " " + texType
			}
			text = append(text, word)
			size += len(word)
		}
		text = text[:len(text)-1]
	}

	if texType == "word" {
		// capitalize first letter
		text[0] = cases.Title(language.Und).String(text[0])
		// end sentence with full stop
		text[len(text)-1] += "."
	}
	return strings.Join(text, ""), nil
}
func (bl *BaseLorem) randomizeNbElements(numberOfElements int) int {
	numberBetween, err := bl.NumberBetween(60, 140)
	if err != nil {
		log.Fatal(err.Error())
	}

	return (numberOfElements * numberBetween / 100) + 1
}
