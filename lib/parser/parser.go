//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2
//line parser.y:7
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
const NUMBER = 57360
const IDENTIFIER = 57361
const STRING = 57362

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

//line parser.y:134

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

const yyNprod = 41
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 160

var yyAct = [...]int{

	24, 26, 31, 2, 12, 34, 41, 52, 81, 52,
	79, 66, 32, 44, 45, 46, 47, 52, 50, 67,
	35, 42, 43, 41, 51, 22, 49, 19, 21, 18,
	44, 45, 46, 47, 53, 54, 55, 17, 42, 43,
	56, 16, 59, 60, 61, 62, 63, 64, 65, 39,
	15, 68, 40, 70, 36, 69, 46, 47, 12, 44,
	45, 46, 47, 57, 58, 67, 40, 40, 37, 38,
	8, 12, 32, 6, 76, 32, 77, 78, 7, 12,
	44, 45, 46, 47, 5, 74, 44, 45, 46, 47,
	23, 71, 4, 3, 29, 27, 14, 28, 29, 27,
	14, 28, 25, 20, 1, 0, 33, 29, 27, 14,
	28, 0, 0, 0, 30, 25, 44, 45, 46, 47,
	9, 0, 72, 73, 10, 0, 0, 9, 0, 48,
	80, 10, 11, 13, 0, 14, 0, 9, 0, 11,
	13, 10, 14, 75, 9, 0, 0, 0, 10, 11,
	13, 0, 14, 0, 0, 0, 11, 13, 0, 14,
}
var yyPact = [...]int{

	-1000, -1000, 140, -1000, 35, 26, 22, 14, 12, 77,
	77, 81, -23, -5, -1000, -1000, -1000, -1000, -1000, -1000,
	49, -1000, 57, 90, 9, 77, -1000, -1000, -1000, -7,
	15, -9, 95, 81, 81, 81, -1000, 90, 90, -1000,
	-1000, 81, 81, 81, 81, 81, 81, 81, -16, -8,
	81, -1000, 81, 38, 95, 65, 116, -1000, -1000, 95,
	95, 95, 33, 33, -1000, -1000, -1000, -1000, 59, 133,
	95, 81, -1000, -1000, 81, -1000, -17, 123, -19, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 104, 1, 0, 103, 28, 25, 93, 92, 84,
	78, 73, 70, 2, 3,
}
var yyR1 = [...]int{

	0, 1, 14, 14, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 12, 8, 8, 9, 10, 11,
	13, 13, 4, 4, 4, 4, 5, 5, 5, 6,
	6, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 1, 2, 1, 2, 1, 2,
	1, 2, 1, 2, 6, 7, 5, 5, 3, 2,
	1, 3, 1, 3, 3, 2, 3, 3, 3, 1,
	3, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	6,
}
var yyChk = [...]int{

	-1000, -1, -14, -7, -8, -9, -11, -10, -12, 4,
	8, 16, -2, 17, 19, 15, 15, 15, 15, 15,
	-4, -5, -6, 13, -3, 25, -2, 18, 20, 17,
	-4, -13, -3, 25, 28, 25, 5, 11, 12, -6,
	-5, 14, 29, 30, 21, 22, 23, 24, -4, -3,
	25, 9, 26, -3, -3, -3, -14, -6, -6, -3,
	-3, -3, -3, -3, -3, -3, 27, 27, -3, -14,
	-3, 26, 6, 7, 26, 10, -13, -14, -13, 27,
	7, 27,
}
var yyDef = [...]int{

	2, -2, 1, 3, 4, 6, 8, 10, 12, 0,
	0, 0, 0, 0, 31, 5, 7, 9, 11, 13,
	0, -2, 0, 0, 0, 0, 32, 33, 34, 0,
	0, 19, 20, 0, 0, 0, 2, 0, 0, 25,
	29, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 2, 0, 0, 18, 0, 0, 23, 24, 26,
	27, 28, 36, 37, 38, 39, 30, 35, 0, 0,
	21, 0, 2, 16, 0, 17, 0, 0, 0, 14,
	15, 40,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	25, 27, 23, 21, 26, 22, 3, 24, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	29, 28, 30,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
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
		//line parser.y:38
		{
			yylex.(*Lexer).program = yyDollar[1].stmtBlock
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:42
		{
			yyVAL.stmtBlock = []Statement{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:43
		{
			yyVAL.stmtBlock = append(yyDollar[1].stmtBlock, yyDollar[2].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:47
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:48
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:49
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:50
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:51
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:52
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:53
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:54
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:55
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:56
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:59
		{
			yyVAL.stmt = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:64
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, yyDollar[6].stmtBlock}
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:65
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[4].stmtBlock, nil}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:69
		{
			yyVAL.stmt = &WhileStmt{yyDollar[2].expr, yyDollar[4].stmtBlock}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:74
		{
			yyVAL.stmt = &AssignStmt{yyDollar[1].lval, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:79
		{
			yyVAL.stmt = &PrintStmt{yyDollar[2].argList}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:84
		{
			yyVAL.argList = []Expression{yyDollar[1].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:85
		{
			yyVAL.argList = append(yyDollar[1].argList, yyDollar[3].expr)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:89
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:90
		{
			yyVAL.expr = &LogicExpr{AND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:91
		{
			yyVAL.expr = &LogicExpr{OR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:92
		{
			yyVAL.expr = &LogicExpr{NOT, yyDollar[2].expr, nil}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:96
		{
			yyVAL.expr = &RelExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:97
		{
			yyVAL.expr = &RelExpr{'<', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:98
		{
			yyVAL.expr = &RelExpr{'>', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:102
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:103
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:107
		{
			yyVAL.lval = Identifier(yyDollar[1].str)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:120
		{
			yyVAL.expr = yyDollar[1].lval
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:121
		{
			yyVAL.expr = Number(yyDollar[1].num)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.expr = String(yyDollar[1].str)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:124
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:125
		{
			yyVAL.expr = &ArithExpr{'+', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:126
		{
			yyVAL.expr = &ArithExpr{'-', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:127
		{
			yyVAL.expr = &ArithExpr{'*', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:128
		{
			yyVAL.expr = &ArithExpr{'/', yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:130
		{
			yyVAL.expr = &CallLuaExpr{yyDollar[3].expr, yyDollar[5].argList}
		}
	}
	goto yystack /* stack new state and value */
}
