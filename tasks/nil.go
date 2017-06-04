// Nil is a lackadaisical teenager. In conversation, his responses are very limited.
// https://ca.wikipedia.org/wiki/Els_joves

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

// Nil will talk with us
func Nil(sentence string) string {

	sentenceIsQuestion := buildSentenceChecker("?")
	sentenceIsYell := buildSentenceChecker("!")
	var nilResponse string

	switch {
	case len(sentence) == 0:
		nilResponse = "Està bé, enrotlla't tiu"
		break
	case sentenceIsQuestion(sentence):
		nilResponse = "De debò? no ho sabia..."
		break
	case sentenceIsYell(sentence):
		nilResponse = "Ay.. Ay.. M'està montant un malt rotllo a la via pública... m'estic possant paranoic"
		break
	default:
		nilResponse = "El que tu diguis"
		break
	}

	return nilResponse
}
