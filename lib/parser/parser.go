//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"io"
	"log"
	"text/scanner"
	//    "os"
	Lua "github.com/yuin/gopher-lua"
	"strconv"
	"strings"
)

//line parser.y:17
type yySymType struct {
	yys       int
	num       float64
	str       string
	lval      Lvalue
	expr      Expression
	stmt      Statement
	argList   []Expression
	stmtBlock []Statement
}

const IF = 57346
const THEN = 57347
const ELSE = 57348
const FI = 57349
const AND = 57350
const OR = 57351
const NOT = 57352
const EQ = 57353
const SP = 57354
const PRINT = 57355
const CALL_LUA = 57356
const NUMBER = 57357
const IDENTIFIER = 57358
const STRING = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IF",
	"THEN",
	"ELSE",
	"FI",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"SP",
	"PRINT",
	"CALL_LUA",
	"NUMBER",
	"IDENTIFIER",
	"STRING",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'='",
	"','",
	"'<'",
	"'>'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:124

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
	Assign(v interface{}, ns NS)
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
	fun Expression
	arg Expression
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
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanComments | scanner.SkipComments
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
type NS map[Identifier]interface{}

func (v Number) Evaluate(ns NS) interface{} {
	return v
}

func (s String) Evaluate(ns NS) interface{} {
	return s
}

func (s *AssignStmt) Execute(ns NS) {
	//s.lval.Assign(s.expr.Evaluate())
	s.lval.Assign(s.expr.Evaluate(ns), ns)
}

func (id Identifier) Assign(val interface{}, ns NS) {
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

func callLua(fname string, farg string) string {
	//先获取lua中定义的函数
	fn := L.GetGlobal(fname)
	if err := L.CallByParam(Lua.P{
		Fn:      fn,
		NRet:    1,
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
		RunStmtBlock(s.trueClause, ns)
	} else {
		RunStmtBlock(s.falseClause, ns)
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

func RunStmtBlock(block []Statement, ns NS) {
	for _, stmt := range block {
		stmt.Execute(ns)
	}
}

var lexKeywords = map[string]int{
	"call_lua": CALL_LUA,
	"CALL_LUA": CALL_LUA,
	"IF":       IF,
	"if":       IF,
	"THEN":     THEN,
	"then":     THEN,
	"ELSE":     ELSE,
	"else":     ELSE,
	//	"END":   END,
	"FI":    FI,
	"fi":    FI,
	"AND":   AND,
	"and":   AND,
	"OR":    OR,
	"or":    OR,
	"NOT":   NOT,
	"not":   NOT,
	"EQ":    EQ,
	"eq":    EQ,
	"PRINT": PRINT,
	"print": PRINT,
	";":     SP,
}

func Eval(str string, mp map[string]string) {

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
	RunStmtBlock(prog, my)
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
	RunStmtBlock(prog, my)
}

// https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 15,
	8, 23,
	9, 23,
	-2, 16,
}

const yyNprod = 35
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 117

var yyAct = [...]int{

	18, 20, 56, 2, 9, 33, 42, 43, 27, 25,
	14, 16, 36, 37, 38, 39, 38, 39, 34, 35,
	41, 57, 36, 37, 38, 39, 13, 44, 45, 31,
	40, 66, 46, 12, 49, 50, 51, 52, 53, 54,
	55, 47, 48, 58, 59, 7, 11, 28, 9, 36,
	37, 38, 39, 24, 8, 17, 5, 10, 57, 23,
	21, 10, 22, 64, 63, 9, 33, 36, 37, 38,
	39, 19, 62, 36, 37, 38, 39, 29, 30, 34,
	35, 23, 21, 10, 22, 23, 21, 10, 22, 6,
	4, 3, 1, 26, 32, 0, 0, 19, 36, 37,
	38, 39, 15, 7, 7, 60, 61, 65, 0, 0,
	0, 0, 8, 8, 15, 10, 10,
}
var yyPact = [...]int{

	-1000, -1000, 41, -1000, 34, 21, 14, 45, 67, -14,
	-1000, -1000, -1000, -1000, 42, -1000, 69, 71, 55, 45,
	-1000, -1000, -1000, -20, -16, 80, 67, 67, -1000, 71,
	71, -1000, -1000, 67, 67, 67, 67, 67, 67, 67,
	-25, -6, 67, 67, 31, 80, 99, -1000, -1000, 80,
	80, 80, -4, -4, -1000, -1000, -1000, -1000, 49, 80,
	-1000, -1000, 67, 100, 4, -1000, -1000,
}
var yyPgo = [...]int{

	0, 92, 1, 0, 10, 94, 11, 91, 90, 89,
	56, 53, 3,
}
var yyR1 = [...]int{

	0, 1, 12, 12, 7, 7, 7, 7, 7, 7,
	8, 8, 9, 10, 11, 11, 4, 4, 4, 4,
	5, 5, 5, 6, 6, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 1, 2, 1, 2, 1, 2,
	7, 5, 3, 2, 1, 3, 1, 3, 3, 2,
	3, 3, 3, 1, 3, 1, 1, 1, 1, 3,
	6, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -12, -7, -8, -10, -9, 4, 13, -2,
	16, 12, 12, 12, -4, -5, -6, 10, -3, 26,
	-2, 15, 17, 14, -11, -3, 26, 22, 5, 8,
	9, -6, -5, 11, 24, 25, 18, 19, 20, 21,
	-4, -3, 26, 23, -3, -3, -12, -6, -6, -3,
	-3, -3, -3, -3, -3, -3, 27, 27, -3, -3,
	6, 7, 23, -12, -3, 7, 27,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 6, 8, 0, 0, 0,
	25, 5, 7, 9, 0, -2, 0, 0, 0, 0,
	26, 27, 28, 0, 13, 14, 0, 0, 2, 0,
	0, 19, 23, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 12, 0, 17, 18, 20,
	21, 22, 31, 32, 33, 34, 24, 29, 0, 15,
	2, 11, 0, 0, 0, 10, 30,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	26, 27, 20, 18, 23, 19, 3, 21, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 22, 25,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:47
		{
			yylex.(*Lexer).program = yyDollar[1].stmtBlock
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:51
		{
			yyVAL.stmtBlock = []Statement{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:52
		{
			yyVAL.stmtBlock = append(yyDollar[1].stmtBlock, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:56
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:57
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:58
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:59
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:60
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:61
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 10:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:66
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, yyDollar[6].stmtBlock}
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:67
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, nil}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:71
		{
			yyVAL.stmt = &AssignStmt{yyDollar[1].lval, yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:76
		{
			yyVAL.stmt = &PrintStmt{yyDollar[2].argList}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:80
		{
			yyVAL.argList = []Expression{yyDollar[1].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:81
		{
			yyVAL.argList = append(yyDollar[1].argList, yyDollar[3].expr)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:85
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:86
		{
			yyVAL.expr = &LogicExpr{AND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:87
		{
			yyVAL.expr = &LogicExpr{OR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:88
		{
			yyVAL.expr = &LogicExpr{NOT, yyDollar[2].expr, nil}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:92
		{
			yyVAL.expr = &RelExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:93
		{
			yyVAL.expr = &RelExpr{'<', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:94
		{
			yyVAL.expr = &RelExpr{'>', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:98
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:99
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:103
		{
			yyVAL.lval = Identifier(yyDollar[1].str)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:112
		{
			yyVAL.expr = yyDollar[1].lval
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:113
		{
			yyVAL.expr = Number(yyDollar[1].num)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:114
		{
			yyVAL.expr = String(yyDollar[1].str)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:116
		{
			yyVAL.expr = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:117
		{
			yyVAL.expr = &ArithExpr{'+', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:118
		{
			yyVAL.expr = &ArithExpr{'-', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:119
		{
			yyVAL.expr = &ArithExpr{'*', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:120
		{
			yyVAL.expr = &ArithExpr{'/', yyDollar[1].expr, yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
