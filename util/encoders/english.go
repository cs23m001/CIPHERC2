package encoders

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	insecureRand "math/rand"
	"strings"
)

var dictionary map[int][]string

// English Encoder - An ASCIIEncoder for binary to english text
type English struct{}

// Encode - Binary => English
func (e English) Encode(data []byte) ([]byte, error) {
	if dictionary == nil {
		buildDictionary()
	}
	words := []string{}
	for _, b := range data {
		possibleWords := dictionary[int(b)]
		index := insecureRand.Intn(len(possibleWords))
		words = append(words, possibleWords[index])
	}
	return []byte(strings.Join(words, " ")), nil
}

// Decode - English => Binary
func (e English) Decode(words []byte) ([]byte, error) {
	wordList := strings.Split(string(words), " ")
	data := []byte{}
	for _, word := range wordList {
		word = strings.TrimSpace(word)
		if len(word) == 0 {
			continue
		}
		byteValue := SumWord(word)
		data = append(data, byte(byteValue))
	}
	return data, nil
}

var rawEnglishDictionary []string

func SetEnglishDictionary(dictionary []string) {
	rawEnglishDictionary = dictionary
}

func getEnglishDictionary() []string {
	return rawEnglishDictionary
}

func buildDictionary() {
	dictionary = map[int][]string{}
	for _, word := range getEnglishDictionary() {
		word = strings.TrimSpace(word)
		sum := SumWord(word)
		dictionary[sum] = append(dictionary[sum], word)
	}
}

// SumWord - Sum the ASCII values of a word
func SumWord(word string) int {
	sum := 0
	for _, char := range word {
		sum += int(char)
	}
	return sum % 256
}
