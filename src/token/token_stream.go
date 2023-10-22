package token

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

type Stream struct {
	tokens []*Token
	index  int
}

func (s *Stream) Tokens() []*Token {
	return s.tokens
}

func NewStream(src string) *Stream {
	s := newScanner(src)
	tokens := []*Token{}

	for {
		s.Scan()
		kind := getKind(s.TokenText())
		tokens = append(tokens, &Token{
			Text:     s.TokenText(),
			Kind:     kind,
			Position: s.Position,
		})
		if kind == EOF {
			break
		}
	}

	return &Stream{
		tokens: tokens,
		index:  0,
	}
}

func (s *Stream) Peek() *Token {
	return s.tokens[s.index]
}

func (s *Stream) Next() *Token {
	token := s.tokens[s.index]
	s.index++
	return token
}

func (s *Stream) String() string {
	var builder strings.Builder

	for index, token := range s.tokens {
		builder.WriteString(fmt.Sprintf("%q {kind: %s, position: %s}", token.Text, token.Kind.String(), token.Position.String()))
		if index+1 < len(s.tokens) {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func newScanner(src string) scanner.Scanner {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"
	s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch)
	}
	return s
}
