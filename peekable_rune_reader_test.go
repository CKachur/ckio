package ckio

import (
	"bufio"
	"compress/gzip"
	"os"
	"strings"
	"testing"
)

func TestRuneReaderFromFile(t *testing.T) {
	testdataFile := openTestdataFile()
	defer testdataFile.Close()
	gzipReader := getGzipReader(testdataFile)
	runeReader := NewRuneReader(bufio.NewReader(gzipReader))
	peekTwiceThenRead(runeReader, t)
}

func openTestdataFile() *os.File {
	testdataFile, err := os.Open("testdata/testdata.json.gz")
	if err != nil {
		panic(err)
	}
	return testdataFile
}

func getGzipReader(file *os.File) *gzip.Reader {
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	return gzipReader
}

func TestRuneReaderFromString(t *testing.T) {
	testString := "This cat is on fire!!!!!!!!!!!!!!1 🔥😼🔥\n"
	runeReader := NewRuneReader(bufio.NewReader(strings.NewReader(testString)))
	peekTwiceThenRead(runeReader, t)
}

func TestFileRuneReader(t *testing.T) {
	fileRuneReader, err := NewFileRuneReader("testdata/smalltest.json")
	if err != nil {
		panic(err)
	}
	defer fileRuneReader.Close()

	peekTwiceThenRead(fileRuneReader, t)
}

func peekTwiceThenRead(runeReader PeekableRuneReader, t *testing.T) {
	numberOfRunesRead := 0
	for {
		firstPeekValue, _ := runeReader.Peek()
		if firstPeekValue == 0 {
			break
		}
		secondPeekValue, _ := runeReader.Peek()
		readValue, _ := runeReader.Read()
		if firstPeekValue != secondPeekValue {
			t.Fatalf(`consecutive Peek() values do not match: '%c' and '%c'`, firstPeekValue, secondPeekValue)
		}
		if firstPeekValue != readValue {
			t.Fatalf(`Peek() value followed by Read() value does not match: '%c' and '%c'`, firstPeekValue, readValue)
		}
		numberOfRunesRead++
	}
	if numberOfRunesRead == 0 {
		t.Fatalf("did not read contents of file")
	}
}
