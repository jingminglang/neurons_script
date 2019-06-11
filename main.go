package main

import (
	P "neurons_script/lib/parser"
	//	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var file_name string
var lua_name string

func init() {
	flag.StringVar(&lua_name, "l", "", "lua脚本文件")
	flag.StringVar(&file_name, "f", "", "脚本文件")
	flag.Parse()
	if file_name == "" {
		flag.Usage()
		os.Exit(2)
	}
}

func main() {


	l, _ := ioutil.ReadFile(lua_name)
	b, e := ioutil.ReadFile(file_name)
	if e != nil {
		fmt.Println("read file error")
		return
	}
	LuaFun := string(l[:])
	P.LuaFun = LuaFun
	s := string(b[:])
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(os.Stdin)
	// s := buf.String()
	P.EvalStr(s)
}
