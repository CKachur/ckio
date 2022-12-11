package ckio

import (
	"bufio"
	"os"
)

type FileRuneReader struct {
	fileReader *bufio.Reader
	nextRune   rune
}

func NewFileRuneReader(fileName string) (*FileRuneReader, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewReader(file)
	return &FileRuneReader{fileReader: fileReader, nextRune: 0}, nil
}

func (fileRuneReader *FileRuneReader) Peek() (rune, error) {
	if fileRuneReader.nextRune != 0 {
		return fileRuneReader.nextRune, nil
	}
	runeValue, err := readRuneFromFileRuneReader(fileRuneReader)
	if err != nil {
		fileRuneReader.nextRune = 0
	} else {
		fileRuneReader.nextRune = runeValue
	}
	return fileRuneReader.nextRune, err
}

func (fileRuneReader *FileRuneReader) Read() (rune, error) {
	if fileRuneReader.nextRune != 0 {
		nextRune := fileRuneReader.nextRune
		fileRuneReader.nextRune = 0
		return nextRune, nil
	}
	return readRuneFromFileRuneReader(fileRuneReader)
}

func readRuneFromFileRuneReader(fileRuneReader *FileRuneReader) (rune, error) {
	runeValue, _, err := fileRuneReader.fileReader.ReadRune()
	return runeValue, err
}
