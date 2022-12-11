package ckio

import (
	"bufio"
	"os"
)

type FileRuneReader struct {
	runeReader *RuneReader
	file       *os.File
}

func NewFileRuneReader(fileName string) (*FileRuneReader, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewReader(file)
	return &FileRuneReader{runeReader: NewRuneReader(fileReader), file: file}, nil
}

func (fileRuneReader *FileRuneReader) Close() error {
	return fileRuneReader.file.Close()
}

func (fileRuneReader *FileRuneReader) Peek() (rune, error) {
	return fileRuneReader.runeReader.Peek()
}

func (fileRuneReader *FileRuneReader) Read() (rune, error) {
	return fileRuneReader.runeReader.Read()
}
