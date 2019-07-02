//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2
//line parser.y:8
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
const IDENTIFIER = 57362
const STRING = 57363

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

//line parser.y:140

// https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 21,
	11, 29,
	12, 29,
	-2, 22,
}

const yyNprod = 42
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 173

var yyAct = [...]int{

	24, 26, 32, 2, 12, 42, 45, 46, 47, 48,
	35, 77, 33, 45, 46, 47, 48, 54, 85, 69,
	68, 43, 44, 42, 54, 83, 50, 22, 54, 52,
	21, 45, 46, 47, 48, 55, 56, 57, 51, 43,
	44, 58, 36, 61, 62, 63, 64, 65, 66, 67,
	20, 40, 70, 71, 41, 73, 19, 72, 47, 48,
	12, 31, 45, 46, 47, 48, 59, 60, 78, 41,
	41, 18, 17, 16, 12, 33, 49, 80, 33, 81,
	82, 15, 53, 12, 45, 46, 47, 48, 23, 37,
	69, 8, 29, 30, 27, 14, 28, 38, 39, 6,
	7, 25, 29, 30, 27, 14, 28, 5, 4, 3,
	1, 34, 45, 46, 47, 48, 0, 74, 29, 30,
	27, 14, 28, 0, 0, 0, 0, 25, 45, 46,
	47, 48, 9, 0, 75, 76, 10, 9, 0, 0,
	84, 10, 0, 0, 11, 13, 0, 0, 14, 11,
	13, 9, 0, 14, 0, 10, 9, 79, 0, 0,
	10, 0, 0, 11, 13, 0, 0, 14, 11, 13,
	0, 0, 14,
}
var yyPact = [...]int{

	-1000, -1000, 152, -1000, 66, 58, 57, 56, 41, 75,
	75, 85, -19, 16, -1000, -1000, -1000, -1000, -1000, -1000,
	84, -1000, 86, 101, 9, 75, -1000, -1000, -1000, 12,
	3, 73, 1, 106, 85, 85, 85, -1000, 101, 101,
	-1000, -1000, 85, 85, 85, 85, 85, 85, 85, -8,
	-9, 85, 85, -1000, 85, 62, 106, 90, 128, -1000,
	-1000, 106, 106, 106, 34, 34, -1000, -1000, -1000, -1000,
	-16, 40, 147, 106, 85, -1000, -1000, 85, -1000, -1000,
	-3, 133, -10, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 110, 1, 0, 50, 30, 27, 109, 108, 107,
	100, 99, 91, 2, 3,
}
var yyR1 = [...]int{

	0, 1, 14, 14, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 12, 8, 8, 9, 10, 11,
	13, 13, 4, 4, 4, 4, 5, 5, 5, 6,
	6, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 1, 2, 1, 2, 1, 2,
	1, 2, 1, 2, 6, 7, 5, 5, 3, 2,
	1, 3, 1, 3, 3, 2, 3, 3, 3, 1,
	3, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	6, 4,
}
var yyChk = [...]int{

	-1000, -1, -14, -7, -8, -9, -11, -10, -12, 4,
	8, 16, -2, 17, 20, 15, 15, 15, 15, 15,
	-4, -5, -6, 13, -3, 26, -2, 19, 21, 17,
	18, -4, -13, -3, 26, 29, 26, 5, 11, 12,
	-6, -5, 14, 30, 31, 22, 23, 24, 25, -4,
	-3, 26, 26, 9, 27, -3, -3, -3, -14, -6,
	-6, -3, -3, -3, -3, -3, -3, -3, 28, 28,
	-3, -3, -14, -3, 27, 6, 7, 27, 28, 10,
	-13, -14, -13, 28, 7, 28,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 6, 8, 10, 12, 0,
	0, 0, 0, 0, 31, 5, 7, 9, 11, 13,
	0, -2, 0, 0, 0, 0, 32, 33, 34, 0,
	0, 0, 19, 20, 0, 0, 0, 2, 0, 0,
	25, 29, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 2, 0, 0, 18, 0, 0, 23,
	24, 26, 27, 28, 36, 37, 38, 39, 30, 35,
	0, 0, 0, 21, 0, 2, 16, 0, 41, 17,
	0, 0, 0, 14, 15, 40,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	26, 28, 24, 22, 27, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 29, 31,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
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
		//line parser.y:39
		{
			yylex.(*Lexer).program = yyDollar[1].stmtBlock
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:43
		{
			yyVAL.stmtBlock = []Statement{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:44
		{
			yyVAL.stmtBlock = append(yyDollar[1].stmtBlock, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:49
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:50
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:51
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:52
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:53
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:54
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:55
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:56
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:57
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:58
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:63
		{
			yyVAL.stmt = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:68
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, yyDollar[6].stmtBlock}
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:69
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, nil}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:73
		{
			yyVAL.stmt = &WhileStmt{yyDollar[2].expr, yyDollar[4].stmtBlock}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:78
		{
			yyVAL.stmt = &AssignStmt{yyDollar[1].lval, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:83
		{
			yyVAL.stmt = &PrintStmt{yyDollar[2].argList}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:88
		{
			yyVAL.argList = []Expression{yyDollar[1].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:89
		{
			yyVAL.argList = append(yyDollar[1].argList, yyDollar[3].expr)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:93
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:94
		{
			yyVAL.expr = &LogicExpr{AND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:95
		{
			yyVAL.expr = &LogicExpr{OR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:96
		{
			yyVAL.expr = &LogicExpr{NOT, yyDollar[2].expr, nil}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:100
		{
			yyVAL.expr = &RelExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:101
		{
			yyVAL.expr = &RelExpr{'<', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:102
		{
			yyVAL.expr = &RelExpr{'>', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:106
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:107
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:111
		{
			yyVAL.lval = Identifier(yyDollar[1].str)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:125
		{
			yyVAL.expr = yyDollar[1].lval
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:126
		{
			yyVAL.expr = Number(yyDollar[1].num)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:128
		{
			yyVAL.expr = String(yyDollar[1].str)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:129
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:130
		{
			yyVAL.expr = &ArithExpr{'+', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:131
		{
			yyVAL.expr = &ArithExpr{'-', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:132
		{
			yyVAL.expr = &ArithExpr{'*', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:133
		{
			yyVAL.expr = &ArithExpr{'/', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:135
		{
			yyVAL.expr = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:136
		{
			yyVAL.expr = &ToNumExpr{yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
