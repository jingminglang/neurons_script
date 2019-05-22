%{
package parser

import (
    "io"
    "fmt"
    "text/scanner"
    "log"
//    "os"
    "strings"
    "strconv"
    Lua "github.com/yuin/gopher-lua"
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

%start program


%token IF THEN ELSE FI AND OR NOT EQ SP PRINT CALL_LUA
%token <num> NUMBER
%token <str> IDENTIFIER STRING


%left '+' '-'
%left '*' '/'


%type <lval> var
%type <expr> expr condition rel_condition sub_condition
%type <stmt> statement if_stmt assign_stmt print_stmt
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
	| if_stmt SP { $$ = $1 }
        | print_stmt { $$ = $1 }
        | print_stmt SP { $$ = $1 }
	| assign_stmt { $$ = $1 }
	| assign_stmt SP { $$ = $1 }
	;


if_stmt
        : IF condition THEN stmt_block ELSE stmt_block FI { $$ = &IfStmt{$2, $4, $6} }
        | IF condition THEN stmt_block FI { $$ = &IfStmt{$2, $4, nil} }
        ;

assign_stmt
	: var '=' expr { $$ = &AssignStmt{$1, $3} }
	;


print_stmt
	: PRINT arg_list { $$ = &PrintStmt{$2} }


arg_list
	: expr { $$ = []Expression{$1} }
	| arg_list ',' expr { $$ = append($1, $3) }


condition
	: rel_condition { $$ = $1 }
	| sub_condition AND sub_condition { $$ = &LogicExpr{AND, $1, $3} }
	| sub_condition OR sub_condition { $$ = &LogicExpr{OR, $1, $3} }
	| NOT sub_condition { $$ = &LogicExpr{NOT, $2, nil} }
        ;

rel_condition
	: expr EQ expr { $$ = &RelExpr{EQ, $1, $3} }
	| expr '<' expr { $$ = &RelExpr{'<', $1, $3} }
	| expr '>' expr { $$ = &RelExpr{'>', $1, $3} }
	;

sub_condition
	: rel_condition { $$ = $1 }
	| '(' condition ')' { $$ = $2 }
	;


var	: IDENTIFIER { $$ = Identifier($1) }


// string : STRING { $$ = String($1) }
//        | ''' STRING  ''' { $$ = String($2) }
//        ;


	;
expr	: var { $$ = $1 }
	| NUMBER { $$ = Number($1) }
	| STRING { $$ = String($1) }
	| '(' expr ')' { $$ = $2 }
	| CALL_LUA '(' expr ',' expr ')' { $$ = &CallLuaExpr{$3,$5} }
	| expr '+' expr { $$ = &ArithExpr{'+', $1, $3} }
	| expr '-' expr { $$ = &ArithExpr{'-', $1, $3} }
	| expr '*' expr { $$ = &ArithExpr{'*', $1, $3} }
	| expr '/' expr { $$ = &ArithExpr{'/', $1, $3} }
	;


%%

var L = Lua.NewState()

type Number float64
type String string
type Identifier string
//type Et string

type Statement interface {
	Execute(ns NS)
}

type Expression interface {
	Evaluate(ns NS) interface{}
}

type Lvalue interface {
	Evaluate(ns NS) interface{}
	Assign(v interface{},ns NS)
}

type AssignStmt struct {
	lval Lvalue
	expr Expression
}


type BinExpr struct {
	op       int
	lhs, rhs Expression
}

type LogicExpr BinExpr
type ArithExpr BinExpr
type RelExpr BinExpr


type IfStmt struct {
	cond                    Expression
	trueClause, falseClause []Statement
}

type CallLuaExpr struct {
        fun  Expression
        arg  Expression
}

type PrintStmt struct {
	argList []Expression
}

type Lexer struct {
	s         scanner.Scanner
	program   []Statement
	hasErrors bool
}


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
//	case scanner.RawString:
//		text := l.s.TokenText()
//                fmt.Println("String: ",text)
//		text = text[1 : len(text)-1]
//		lval.str = text
//		return STRING
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
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats |  scanner.ScanStrings | scanner.ScanComments |scanner.SkipComments
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


//var symTab = map[Identifier]interface{}{}
type NS  map[Identifier]interface{}

func (v Number) Evaluate(ns NS) interface{} {
	return v
}

func (s String) Evaluate(ns NS) interface{} {
	return s
}

func (s *AssignStmt) Execute(ns NS) {
	//s.lval.Assign(s.expr.Evaluate())
	s.lval.Assign(s.expr.Evaluate(ns),ns)
}


func (id Identifier) Assign(val interface{},ns NS) {
	//symTab[id] = val
	ns[id] = val
}

func (id Identifier) Evaluate(ns NS) interface{} {
	//val, ok := symTab[id]
	val, ok := ns[id]
	if !ok {
		log.Println("Identifier.Evaluate: symbol", id, "undefined")
	}
	return val
}


func (e *ArithExpr) Evaluate(ns NS) interface{} {
	lhs := e.lhs.Evaluate(ns)
	rhs := e.rhs.Evaluate(ns)

	if e.op == '+' {
		s1, ok1 := lhs.(String)
		s2, ok2 := rhs.(String)
		if ok1 && ok2 {
			return s1 + s2
		}
	}
	{
		lhs := lhs.(Number)
		rhs := rhs.(Number)
		switch e.op {
		case '+':
			return lhs + rhs
		case '-':
			return lhs - rhs
		case '*':
			return lhs * rhs
		case '/':
			return lhs / rhs
		default:
			panic("unreached")
		}
	}
}

func (e *RelExpr) Evaluate(ns NS) interface{} {
	lhs := e.lhs.Evaluate(ns)
	rhs := e.rhs.Evaluate(ns)

	if lhs, ok := lhs.(String); ok {
		rhs := rhs.(String)
		switch e.op {
		case '<':
			return lhs < rhs
		case '>':
			return lhs > rhs
		case EQ:
			return lhs == rhs
		default:
			panic("unreached")
		}
	}
	{
		lhs := lhs.(Number)
		rhs := rhs.(Number)
		switch e.op {
		case '<':
			return lhs < rhs
		case '>':
			return lhs > rhs
		case EQ:
			return lhs == rhs
		default:
			panic("unreached")
		}
	}
}

func (e *LogicExpr) Evaluate(ns NS) interface{} {
	lhs := e.lhs.Evaluate(ns).(bool)

	switch e.op {
	case AND:
		return lhs && e.rhs.Evaluate(ns).(bool)
	case OR:
		return lhs || e.rhs.Evaluate(ns).(bool)
	case NOT:
		return !lhs
	default:
		panic("unreached")
	}
}



func callLua(fname string,farg string) string {
	//先获取lua中定义的函数
	fn := L.GetGlobal(fname)
	if err := L.CallByParam(Lua.P{
		Fn: fn,
		NRet: 1,
		Protect: true,
	}, Lua.LString(farg)); err != nil {
		panic(err)
	}
	//这里获取函数返回值
	ret := L.Get(-1)
	return ret.String()
}


func (s *CallLuaExpr) Evaluate(ns NS) interface{} {
	funName := s.fun.Evaluate(ns).(String)
	argStr := s.arg.Evaluate(ns).(String)

	var luaAddFun = `
function add(str)
     return "test add"..str
end
`
	if err := L.DoString(luaAddFun); err != nil {
		panic(err)
	}

	r := callLua(fmt.Sprintf("%s", funName), fmt.Sprintf("%s", argStr))
        return String(r)
}

func (s *IfStmt) Execute(ns NS) {
	if s.cond.Evaluate(ns).(bool) {
		RunStmtBlock(s.trueClause,ns)
	} else {
		RunStmtBlock(s.falseClause,ns)
	}
}


func (s *PrintStmt) Execute(ns NS) {
	args := make([]interface{}, len(s.argList))
	for i, expr := range s.argList {
		args[i] = expr.Evaluate(ns)
	}
	fmt.Print(args...)
	fmt.Println()
}




func RunStmtBlock(block []Statement,ns NS) {
	for _, stmt := range block {
		stmt.Execute(ns)
	}
}


var lexKeywords = map[string]int{
	"call_lua":    CALL_LUA,
	"CALL_LUA":    CALL_LUA,
	"IF":    IF,
	"if":    IF,
	"THEN":  THEN,
	"then":  THEN,
	"ELSE":  ELSE,
	"else":  ELSE,
//	"END":   END,
	"FI":   FI,
	"fi":   FI,
	"AND":   AND,
	"and":   AND,
	"OR":    OR,
	"or":    OR,
	"NOT":   NOT,
	"not":   NOT,
	"EQ":   EQ,
	"eq":   EQ,
	"PRINT": PRINT,
	"print": PRINT,
	";":   SP,
}

func Eval(str string,mp map[string]string) {

    var my NS = make(NS)
    //my[Identifier("a")] = String("dd1")
    //my[Identifier("b")] = String("dd4")
    //read := strings.NewReader(os.Args[1])

    for k, v := range mp {
        my[Identifier(k)] = String(v)
    }
    read := strings.NewReader(str)
    lexer := NewLexer(read)
    yyParse(lexer)
    prog := lexer.Program()
    //fmt.Printf("%+v", prog[0])
    //fmt.Println("%+v", prog[0])
    //fmt.Println("%#v", prog)
    RunStmtBlock(prog,my)
    for k, v := range my {
        mp[string(k)] = fmt.Sprintf("%v", v)
    }
    //fmt.Println("%#v", my)
}


func EvalStr(str string) {
    var my NS = make(NS)
    read := strings.NewReader(str)
    lexer := NewLexer(read)
    yyParse(lexer)
    prog := lexer.Program()
    RunStmtBlock(prog,my)
}



// https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/
