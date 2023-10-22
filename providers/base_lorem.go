package providers

import (
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

func (bl *BaseLorem) GetWordList() []string {
	bl.wordList = constants.LoremWordList
	return bl.wordList
}

func (bl *BaseLorem) Word() string {
	word, err := bl.RandomElementFromStringSlice(bl.GetWordList())
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return word
}

func (bl *BaseLorem) WordsAsText(numberOfWords int) string {
	var words []string
	for i := 0; i < numberOfWords; i++ {
		words = append(words, bl.Word())
	}
	return strings.Join(words, " ")
}

func (bl *BaseLorem) WordsAsList(numberOfWords int) []string {
	var words []string
	for i := 0; i < numberOfWords; i++ {
		words = append(words, bl.Word())
	}
	return words
}

func (bl *BaseLorem) Sentence(numberOfWords int, variableNumberOfWords bool) string {
	if numberOfWords <= 0 {
		return ""
	}
	if variableNumberOfWords {
		numberOfWords = bl.randomizeNbElements(numberOfWords)
	}
	words := bl.WordsAsList(numberOfWords)
	words[0] = cases.Title(language.Und).String(words[0])
	return strings.Join(words, " ") + "."
}

func (bl *BaseLorem) SentencesAsText(numberOfSentences int, variableNumberOfWords bool) string {
	if numberOfSentences <= 0 {
		return ""
	}
	var sentences []string
	for i := 0; i < numberOfSentences; i++ {
		sentences = append(sentences, bl.Sentence(6, variableNumberOfWords))
	}
	return strings.Join(sentences, " ")
}
func (bl *BaseLorem) SentencesAsList(numberOfSentences int, variableNumberOfWords bool) []string {
	if numberOfSentences <= 0 {
		return []string{}
	}
	var sentences []string
	for i := 0; i < numberOfSentences; i++ {
		sentences = append(sentences, bl.Sentence(6, variableNumberOfWords))
	}
	return sentences
}

func (bl *BaseLorem) Paragraph(numberOfSentences int, variableNumberOfSentences bool) string {
	if numberOfSentences <= 0 {
		return ""
	}
	if variableNumberOfSentences {
		numberOfSentences = bl.randomizeNbElements(numberOfSentences)
	}
	sentences := bl.SentencesAsList(numberOfSentences, false)
	return strings.Join(sentences, " ")
}

func (bl *BaseLorem) ParagraphsAsList(numberOfParagraphs ...int) []string {
	n := 3
	if len(numberOfParagraphs) > 0 {
		n = numberOfParagraphs[0]
	}

	var paragraphs []string
	for i := 0; i < n; i++ {
		sentences := bl.Paragraph(3, true)
		paragraphs = append(paragraphs, sentences)
	}
	return paragraphs
}

func (bl *BaseLorem) ParagraphsAsText(numberOfParagraphs ...int) string {
	n := 3
	if len(numberOfParagraphs) > 0 {
		n = numberOfParagraphs[0]
	}
	var paragraphs []string
	for i := 0; i < n; i++ {
		sentences := bl.Paragraph(3, true)
		paragraphs = append(paragraphs, sentences)
	}
	return strings.Join(paragraphs, "\n\n")
}
func (bl *BaseLorem) Text(maxNumberOfCharacters ...int) string {
	maxN := 200
	if len(maxNumberOfCharacters) > 0 {
		maxN = maxNumberOfCharacters[0]
		if maxN < 5 {
			log.Fatal("Text() can only generate text of at least 5 characters")
			return ""
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
	return strings.Join(text, "")
}
func (bl *BaseLorem) randomizeNbElements(numberOfElements int) int {
	numberBetween, err := bl.NumberBetween(60, 140)
	if err != nil {
		log.Fatal(err.Error())
	}

	return (numberOfElements * numberBetween / 100) + 1
}
