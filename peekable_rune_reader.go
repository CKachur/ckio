package ckio

type PeekableRuneReader interface {
	Peek() (rune, *ReadRuneError)
	Read() (rune, *ReadRuneError)
}
