package gomod

import (
	U "github.com/neurons-platform/gotools/utils"
	"github.com/yuin/gopher-lua"
	"regexp"
)

func matchRegex(str, re string) bool {
	rp := regexp.MustCompile(re)
	return rp.MatchString(str)
}

func _matchRegex(L *lua.LState) int {
	str := L.ToString(1)
	re := L.ToString(2)
	v := matchRegex(str, re)
	lv := lua.LBool(v)
	L.Push(lv)
	return 1
}

func _md5(L *lua.LState) int {
	str := L.ToString(1)
	v := U.Md5(str)
	lv := lua.LString(v)
	L.Push(lv)
	return 1
}

func _uuid(L *lua.LState) int {
	v := U.GetUUID()
	lv := lua.LString(v)
	L.Push(lv)
	return 1
}

var exports = map[string]lua.LGFunction{
	"matchRegex": _matchRegex,
	"md5":        _md5,
	"uuid":       _uuid,
}

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("gomod"))
	// returns the module
	L.Push(mod)
	return 1
}
