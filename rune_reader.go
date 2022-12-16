package ckio

import (
	"bufio"
)

type RuneReader struct {
	reader   *bufio.Reader
	nextRune rune
}

func NewRuneReader(reader *bufio.Reader) *RuneReader {
	return &RuneReader{reader: reader}
}

func (runeReader *RuneReader) Peek() (rune, *ReadRuneError) {
	if runeReader.nextRune != 0 {
		return runeReader.nextRune, nil
	}
	runeValue, err := readRuneFromRuneReader(runeReader)
	if err != nil {
		runeReader.nextRune = 0
	} else {
		runeReader.nextRune = runeValue
	}
	return runeReader.nextRune, err
}

func (runeReader *RuneReader) Read() (rune, *ReadRuneError) {
	if runeReader.nextRune != 0 {
		nextRune := runeReader.nextRune
		runeReader.nextRune = 0
		return nextRune, nil
	}
	return readRuneFromRuneReader(runeReader)
}

func readRuneFromRuneReader(runeReader *RuneReader) (rune, *ReadRuneError) {
	runeValue, _, err := runeReader.reader.ReadRune()
	if err != nil {
		return runeValue, NewReadRuneError(err.Error())
	}
	return runeValue, nil
}
