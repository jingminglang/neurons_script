package parser

import (
	"fmt"
	"github.com/BixData/gluabit32"
	"github.com/BixData/gluasocket"
	"github.com/cjoudrey/gluahttp"
	Lua "github.com/neurons-platform/gopher-lua"
	"github.com/zhu327/gluadb"
	LuaJson "layeh.com/gopher-json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
	fun     Expression
	argList []Expression
}

type ToNumExpr struct {
	v Expression
}

type PrintStmt struct {
	argList []Expression
}

type WhileStmt struct {
	cond Expression
	body []Statement
}

func (s *WhileStmt) Execute(ns NS) {
	for s.cond.Evaluate(ns).(bool) {
		RunStmtBlock(s.body, ns)
	}
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

func Float64toStr(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func Str2Float64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return i
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
		if ok1 && !ok2 {
			rhs := rhs.(Number)
			return s1 + String(Float64toStr(float64(rhs)))
		}

		if !ok1 && ok2 {
			lhs := lhs.(Number)
			return String(Float64toStr(float64(lhs))) + s2
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

func callLua(fname string, fargs ...string) string {
	//先获取lua中定义的函数
	ls := []Lua.LValue{}
	for _, fg := range fargs {
		ls = append(ls, Lua.LString(fg))
		// ls = append(ls,fg)
	}
	fn := L.GetGlobal(fname)
	if err := L.CallByParam(Lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, ls...); err != nil {
		panic(err)
	}
	//这里获取函数返回值
	ret := L.Get(-1)
	return ret.String()
}

func (s *ToNumExpr) Evaluate(ns NS) interface{} {
	v := s.v.Evaluate(ns).(String)
	r := Str2Float64(string(v))
	return Number(r)
}

var LuaFun = `
function test(str)
     return "test add"..str
end
`

func (s *CallLuaExpr) Evaluate(ns NS) interface{} {
	funName := s.fun.Evaluate(ns).(String)
	// argStr := s.arg.Evaluate(ns).(String)

	args := make([]interface{}, len(s.argList))
	for i, expr := range s.argList {
		args[i] = expr.Evaluate(ns)
	}

	if err := L.DoString(LuaFun); err != nil {
		panic(err)
		// return String("null")

	}
	argList := []string{}
	for _, v := range args {
		vs := fmt.Sprintf("%s", v)
		argList = append(argList, vs)
	}

	// r := callLua(fmt.Sprintf("%s", funName), fmt.Sprintf("%s", argStr))
	// r := callLua(fmt.Sprintf("%s", funName), fmt.Sprintf("%s", argStr))
	r := callLua(fmt.Sprintf("%s", funName), argList...)
	return String(r)
}

func (s *CallLuaExpr) Execute(ns NS) {
	funName := s.fun.Evaluate(ns).(String)
	// argStr := s.arg.Evaluate(ns).(String)

	args := make([]interface{}, len(s.argList))
	for i, expr := range s.argList {
		args[i] = expr.Evaluate(ns)
	}

	if err := L.DoString(LuaFun); err != nil {
		panic(err)
		// return String("null")
	}
	argList := []string{}
	for _, v := range args {
		vs := fmt.Sprintf("%s", v)
		argList = append(argList, vs)
	}

	callLua(fmt.Sprintf("%s", funName), argList...)
}

// func (s *CallLuaExpr) Execute(ns NS)  {
// 	funName := s.fun.Evaluate(ns).(String)
// 	argStr := s.arg.Evaluate(ns).(String)
//
// 	var luaAddFun = `
// function add(str)
//      return "test add"..str
// end
// `
// 	if err := L.DoString(luaAddFun); err != nil {
// 		panic(err)
// 	}
//
//         callLua(fmt.Sprintf("%s", funName), fmt.Sprintf("%s", argStr))
// }

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
	"_lua":  CALL_LUA,
	"_LUA":  CALL_LUA,
	"IF":    IF,
	"if":    IF,
	"THEN":  THEN,
	"then":  THEN,
	"ELSE":  ELSE,
	"else":  ELSE,
	"FI":    FI,
	"fi":    FI,
	"WHILE": WHILE,
	"while": WHILE,
	"DO":    DO,
	"do":    DO,
	"DONE":  DONE,
	"done":  DONE,
	"NUM":   TO_NUM,
	"num":   TO_NUM,
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

var L = Lua.NewState()

func init() {
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	LuaJson.Preload(L)
	gluasocket.Preload(L)
	gluabit32.Preload(L)
	gluadb.Preload(L)
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
	// fmt.Printf("%+v",lexer)
	yyParse(lexer)
	prog := lexer.Program()
	RunStmtBlock(prog, my)
}
