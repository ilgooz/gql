// Copyright 2015 İlker Göktuğ Öztürk <ilkergoktugozturk@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

type Token int

const (
	ILLEGAL Token = iota

	INT
	FLOAT
	STRING

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	DOT    // .

	RPAREN // )
	RBRACK // ]
	RBRACE // }
)

