//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2
//line parser.y:8
type yySymType struct {
	yys       int
	num       float64
	bl        bool
	str       string
	table     map[interface{}]interface{}
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
const WHILE = 57350
const DO = 57351
const DONE = 57352
const AND = 57353
const OR = 57354
const NOT = 57355
const EQ = 57356
const SP = 57357
const PRINT = 57358
const CALL_LUA = 57359
const TO_NUM = 57360
const NUMBER = 57361
const BOOL = 57362
const IDENTIFIER = 57363
const STRING = 57364

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IF",
	"THEN",
	"ELSE",
	"FI",
	"WHILE",
	"DO",
	"DONE",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"SP",
	"PRINT",
	"CALL_LUA",
	"TO_NUM",
	"NUMBER",
	"BOOL",
	"IDENTIFIER",
	"STRING",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"','",
	"')'",
	"'='",
	"'<'",
	"'>'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:146

// https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 21,
	11, 30,
	12, 30,
	-2, 22,
}

const yyNprod = 44
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 168

var yyAct = [...]int{

	24, 26, 2, 36, 12, 43, 33, 46, 47, 48,
	49, 22, 34, 79, 46, 47, 48, 49, 21, 69,
	70, 55, 44, 45, 53, 23, 51, 55, 86, 30,
	31, 28, 27, 14, 29, 41, 56, 57, 58, 25,
	52, 59, 42, 37, 62, 63, 64, 65, 66, 67,
	68, 60, 61, 71, 72, 43, 74, 73, 42, 42,
	54, 12, 55, 84, 46, 47, 48, 49, 46, 47,
	48, 49, 44, 45, 70, 12, 34, 48, 49, 34,
	82, 19, 81, 18, 12, 83, 30, 31, 28, 27,
	14, 29, 20, 17, 16, 38, 35, 30, 31, 28,
	27, 14, 29, 32, 15, 39, 40, 25, 46, 47,
	48, 49, 8, 78, 46, 47, 48, 49, 50, 75,
	46, 47, 48, 49, 9, 6, 77, 76, 10, 7,
	9, 5, 4, 85, 10, 3, 11, 13, 1, 0,
	9, 14, 11, 13, 10, 0, 80, 14, 0, 0,
	9, 0, 11, 13, 10, 0, 0, 14, 0, 0,
	0, 0, 11, 13, 0, 0, 0, 14,
}
var yyPact = [...]int{

	-1000, -1000, 146, -1000, 89, 79, 78, 68, 66, 12,
	12, 69, -27, 16, -1000, -1000, -1000, -1000, -1000, -1000,
	90, -1000, 94, 80, 41, 12, -1000, -1000, -1000, -1000,
	13, -3, 51, -7, 97, 69, 69, 69, -1000, 80,
	80, -1000, -1000, 69, 69, 69, 69, 69, 69, 69,
	-10, -9, 69, 69, -1000, 69, 45, 97, 91, 120,
	-1000, -1000, 97, 97, 97, 52, 52, -1000, -1000, -1000,
	-1000, 85, -16, 136, 97, 69, -1000, -1000, 69, -1000,
	-1000, 34, 126, -1, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 138, 1, 0, 92, 18, 11, 135, 132, 131,
	129, 125, 112, 6, 2,
}
var yyR1 = [...]int{

	0, 1, 14, 14, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 12, 8, 8, 9, 10, 11,
	13, 13, 4, 4, 4, 4, 5, 5, 5, 5,
	6, 6, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 1, 2, 1, 2, 1, 2,
	1, 2, 1, 2, 6, 5, 7, 5, 3, 2,
	1, 3, 1, 3, 3, 2, 3, 3, 3, 1,
	1, 3, 1, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 6, 4,
}
var yyChk = [...]int{

	-1000, -1, -14, -7, -8, -9, -11, -10, -12, 4,
	8, 16, -2, 17, 21, 15, 15, 15, 15, 15,
	-4, -5, -6, 13, -3, 27, -2, 20, 19, 22,
	17, 18, -4, -13, -3, 27, 30, 27, 5, 11,
	12, -6, -5, 14, 31, 32, 23, 24, 25, 26,
	-4, -3, 27, 27, 9, 28, -3, -3, -3, -14,
	-6, -6, -3, -3, -3, -3, -3, -3, -3, 29,
	29, -3, -3, -14, -3, 28, 7, 6, 28, 29,
	10, -13, -14, -13, 29, 7, 29,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 6, 8, 10, 12, 0,
	0, 0, 0, 0, 32, 5, 7, 9, 11, 13,
	0, -2, 0, 0, 29, 0, 33, 34, 35, 36,
	0, 0, 0, 19, 20, 0, 0, 0, 2, 0,
	0, 25, 30, 0, 0, 0, 0, 0, 0, 0,
	0, 29, 0, 0, 2, 0, 0, 18, 0, 0,
	23, 24, 26, 27, 28, 38, 39, 40, 41, 31,
	37, 0, 0, 0, 21, 0, 15, 2, 0, 43,
	17, 0, 0, 0, 14, 16, 42,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 29, 25, 23, 28, 24, 3, 26, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	31, 30, 32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22,
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
		//line parser.y:42
		{
			yylex.(*Lexer).program = yyDollar[1].stmtBlock
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:46
		{
			yyVAL.stmtBlock = []Statement{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:47
		{
			yyVAL.stmtBlock = append(yyDollar[1].stmtBlock, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:52
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:53
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:54
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:55
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:56
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:57
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:58
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:59
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:60
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:61
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:66
		{
			yyVAL.stmt = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:71
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, nil}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:72
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, yyDollar[6].stmtBlock}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:76
		{
			yyVAL.stmt = &WhileStmt{yyDollar[2].expr, yyDollar[4].stmtBlock}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:81
		{
			yyVAL.stmt = &AssignStmt{yyDollar[1].lval, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:86
		{
			yyVAL.stmt = &PrintStmt{yyDollar[2].argList}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:91
		{
			yyVAL.argList = []Expression{yyDollar[1].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:92
		{
			yyVAL.argList = append(yyDollar[1].argList, yyDollar[3].expr)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:96
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:97
		{
			yyVAL.expr = &LogicExpr{AND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:98
		{
			yyVAL.expr = &LogicExpr{OR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:99
		{
			yyVAL.expr = &LogicExpr{NOT, yyDollar[2].expr, nil}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:104
		{
			yyVAL.expr = &RelExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:105
		{
			yyVAL.expr = &RelExpr{'<', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:106
		{
			yyVAL.expr = &RelExpr{'>', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:107
		{
			yyVAL.expr = &RelExpr{EQ, Bool(true), yyDollar[1].expr}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:111
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:112
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:116
		{
			yyVAL.lval = Identifier(yyDollar[1].str)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:130
		{
			yyVAL.expr = yyDollar[1].lval
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:131
		{
			yyVAL.expr = Bool(yyDollar[1].bl)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:132
		{
			yyVAL.expr = Number(yyDollar[1].num)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:134
		{
			yyVAL.expr = String(yyDollar[1].str)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:135
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:136
		{
			yyVAL.expr = &ArithExpr{'+', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:137
		{
			yyVAL.expr = &ArithExpr{'-', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:138
		{
			yyVAL.expr = &ArithExpr{'*', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:139
		{
			yyVAL.expr = &ArithExpr{'/', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:141
		{
			yyVAL.expr = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:142
		{
			yyVAL.expr = &ToNumExpr{yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
