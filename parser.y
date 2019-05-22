%{
package main

import (
    "fmt"
    "text/scanner"
    "os"
    "strings"
    "strconv"
)

%}

%union{
	num        float64
	str        string
	lval       Lvalue
	expr       Expression
	stmt       Statement
	argList   []Expression
	stmtBlock []Statement
}

%type<expr> program

%token IF THEN ELSE END AND OR NOT
%token <num> NUMBER
%token <str> IDENTIFIER STRING

%left '+' '-'
%left '*' '/'


%type <lval> var
%type <expr> expr condition rel_condition sub_condition
%type <stmt> statement if_stmt assign_stmt
%type <argList> arg_list
%type <stmtBlock> stmt_block program

%%

program	: stmt_block { yylex.(*Lexer).program = $1 }
        ;

stmt_block
	: { $$ = []Statement{} }
	| stmt_block statement { $$ = append($1, $2) }
	;

statement
	: if_stmt { $$ = $1 }
	| assign_stmt { $$ = $1 }
	| error { $$ = nil }
	;

if_stmt	: IF condition THEN stmt_block ELSE stmt_block END IF { $$ = &IfStmt{$2, $4, $6} }
	| IF condition THEN stmt_block END IF { $$ = &IfStmt{$2, $4, nil} }
	;

assign_stmt
	: var '=' expr { $$ = &AssignStmt{$1, $3} }
	;

condition
	: rel_condition { $$ = $1 }
	| sub_condition AND sub_condition { $$ = &LogicExpr{AND, $1, $3} }
	| sub_condition OR sub_condition { $$ = &LogicExpr{OR, $1, $3} }
	| NOT sub_condition { $$ = &LogicExpr{NOT, $2, nil} }


rel_condition
	: expr '=' expr { $$ = &RelExpr{'=', $1, $3} }
	| expr '<' expr { $$ = &RelExpr{'<', $1, $3} }
	| expr '>' expr { $$ = &RelExpr{'>', $1, $3} }
	;

sub_condition
	: rel_condition { $$ = $1 }
	| '(' condition ')' { $$ = $2 }
	;

var	: IDENTIFIER { $$ = Identifier($1) }


expr	: var { $$ = $1 }
	| NUMBER { $$ = Number($1) }
	| STRING { $$ = String($1) }
	| '(' expr ')' { $$ = $2 }
	| expr '+' expr { $$ = &ArithExpr{'+', $1, $3} }
	| expr '-' expr { $$ = &ArithExpr{'-', $1, $3} }
	| expr '*' expr { $$ = &ArithExpr{'*', $1, $3} }
	| expr '/' expr { $$ = &ArithExpr{'/', $1, $3} }
	;


%%

type Lexer struct {
	s         scanner.Scanner
	program   []Statement
	hasErrors bool
}

func NewLexer(r io.Reader) *Lexer {
	l := new(Lexer)
	l.s.Init(r)
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.SkipComments
	return l
}

var lexKeywords = map[string]int{
	"IF":    IF,
	"THEN":  THEN,
	"ELSE":  ELSE,
	"END":   END,
	"WHILE": WHILE,
	"DO":    DO,
	"PRINT": PRINT,
	"AND":   AND,
	"OR":    OR,
	"NOT":   NOT,
}



func main() {
    l := new(Lexer)
    l.Init(strings.NewReader(os.Args[1]))
    yyParse(l)
    fmt.Printf("%#v\n", l.result)
    fmt.Printf("%d\n", Eval(l.result))
}
