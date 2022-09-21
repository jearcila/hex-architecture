package format

import (
	"encoding/json"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// GetRawJSONAsString parses json into a string
func GetRawJSONAsString(inputJSON interface{}) (string, error) {
	// Marshaling
	responseBytes, err := json.Marshal(inputJSON)
	if err != nil {
		return "", err
	}

	// Converting into string
	str := string(responseBytes)

	return str, nil
}

func GetNormalizedASCII(original string) (string, error) {
	destinationBytes := make([]byte, len(original))

	// Function returning all non-ascii runes
	isASCII := func(r rune) bool { return r >= unicode.MaxASCII }

	// Chain NFD (Canonical Decomposition) and then eliminate non-ascii runes
	// This will get a more 'likely' ascii string
	t := transform.Chain(norm.NFD, runes.Remove(runes.Predicate(isASCII)))

	nDst, _ /*nSrc*/, err := t.Transform(destinationBytes, []byte(original), true)
	if err != nil {
		return "", err
	}

	dstWithoutEscapeChars := RemoveEscapeChars(string(destinationBytes)[:nDst])

	return dstWithoutEscapeChars, nil
}

func RemoveEscapeChars(original string) string {
	replacer := strings.NewReplacer("\a", "", "\b", "", "\\", "", "\t", "", "\n", "", "\f", "", "\r", "", "\v", "", "\"", "")
	return replacer.Replace(original)
}

func PadRight(value string, char string, maxLength int) (ret string) {
	ret = value

	for len(ret) < maxLength {
		ret += char
	}

	return
}
