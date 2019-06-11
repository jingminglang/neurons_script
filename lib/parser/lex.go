package parser

import (
	"io"
	"text/scanner"
	"log"
	//    "os"
	// "fmt"
	"strconv"
)



type Lexer struct {
	s         scanner.Scanner
	program   []Statement
	hasErrors bool
}


// 词法分析
func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.s.Scan()
	switch tok {
	case scanner.EOF:
                //fmt.Println("EOF: ")
		return 0
	case scanner.Int, scanner.Float:
		lval.num, _ = strconv.ParseFloat(l.s.TokenText(), 64)
                //fmt.Println("number: ",lval.num)
		return NUMBER
	case scanner.Ident:
		ident := l.s.TokenText()
                // fmt.Println(ident)
		keyword, isKeyword := lexKeywords[ident]
		if isKeyword {
                 //       fmt.Println("keyword: ",ident)
			return keyword
		}
                //fmt.Println("ident: ",ident)
		lval.str = ident
		return IDENTIFIER
	case scanner.String:
		text := l.s.TokenText()
                // fmt.Println("String: ",text)
		text = text[1 : len(text)-1]
		lval.str = text
		return STRING
	case scanner.RawString:
		text := l.s.TokenText()
                // fmt.Println("String: ",text)
		text = text[1 : len(text)-1]
		lval.str = text
		return STRING
	//case scanner.Char:
	//	text := l.s.TokenText()
        //        fmt.Println("char: ",text)
	//	return int(tok)
	default:
                if int(tok) == 59 {
		   return SP
                }
                // fmt.Println("tok: ",tok)
		return int(tok)
	}
}



func NewLexer(r io.Reader) *Lexer {
	l := new(Lexer)
	l.s.Init(r)
	//l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanChars |  scanner.ScanStrings | scanner.SkipComments
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats |  scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments |scanner.SkipComments
	return l
}


func (l *Lexer) Error(s string) {
	log.Println("Parser:", s)
	l.hasErrors = true
}


func (l *Lexer) Program() []Statement {
	if l.hasErrors {
		return nil
	}
	return l.program
}
