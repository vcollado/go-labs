// Bob is a lackadaisical teenager. In conversation, his responses are very limited.

package tasks

import (
	"strings"
)

type sentenceIs func(string) bool

func buildSentenceChecker(token string) sentenceIs {
	return func(toCheck string) bool {
		return strings.LastIndex(toCheck, token) == (len(toCheck) - 1)
	}
}

// Bob will talk with us
func Bob(sentence string) string {

	sentenceIsQuestion := buildSentenceChecker("?")
	sentenceIsYell := buildSentenceChecker("!")
	var bobResponse string

	switch {
	case len(sentence) == 0:
		bobResponse = "Fine. Be that way!'"
		break
	case sentenceIsQuestion(sentence):
		bobResponse = "Sure."
		break
	case sentenceIsYell(sentence):
		bobResponse = "Whoa, chill out!"
		break
	default:
		bobResponse = "Whatever."
		break
	}

	return bobResponse
}
