package ckio

import (
	"bufio"
	"os"
)

type FileRuneReader struct {
	runeReader *RuneReader
	file       *os.File
}

func NewFileRuneReader(fileName string) (*FileRuneReader, *FileOpenError) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, NewFileOpenError(err.Error())
	}
	fileReader := bufio.NewReader(file)
	return &FileRuneReader{runeReader: NewRuneReader(fileReader), file: file}, nil
}

func (fileRuneReader *FileRuneReader) Close() *FileCloseError {
	err := fileRuneReader.file.Close()
	if err != nil {
		return NewFileCloseError(err.Error())
	}
	return nil
}

func (fileRuneReader *FileRuneReader) Peek() (rune, *ReadRuneError) {
	return fileRuneReader.runeReader.Peek()
}

func (fileRuneReader *FileRuneReader) Read() (rune, *ReadRuneError) {
	return fileRuneReader.runeReader.Read()
}
