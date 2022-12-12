package ckio

type PeekableRuneReader interface {
	Peek() (rune, error)
	Read() (rune, error)
}
