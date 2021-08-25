package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string) string {
	urlHashBytes := sha256Of(initialLink)
	generateNumber := new(big.Int).SetBytes(urlHashBytes)
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	return finalString[:8]
}
