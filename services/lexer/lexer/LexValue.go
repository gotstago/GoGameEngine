package lexer

import (
	"strings"

	"github.com/gotstago/GoGameEngine/services/lexer/errors"
	"github.com/gotstago/GoGameEngine/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_VALUE with the value to be assigned
to a key.
*/
func LexValue(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEWLINE) {
			lexer.Emit(lexertoken.TOKEN_VALUE)
			return LexBegin
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}
