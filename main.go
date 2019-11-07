package main

import (
	P "github.com/neurons-platform/neurons_script/lib/parser"
	"github.com/neurons-platform/neurons_script/gomod"
	//	"bytes"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// var file_name string
var lua_name string

func init() {
	flag.StringVar(&lua_name, "lua", "", "lua脚本文件")
	// flag.StringVar(&file_name, "f", "", "脚本文件")
	flag.Parse()
	// if file_name == "" {
	// 	flag.Usage()
	// 	os.Exit(2)
	// }
}

func readScript(scriptName string) string {
	r := ""
	fi, err := os.Open(scriptName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return r
	}

	defer fi.Close()

	br := bufio.NewReader(fi)

	count := 0
	for {
		a, _, c := br.ReadLine()
		count = count + 1
		if count < 3 {
			continue
		}
		if c == io.EOF {
			break
		}
		r = r + "\n" + string(a)
	}
	return r
}

func main() {

	l, _ := ioutil.ReadFile(lua_name)
	LuaFun := string(l[:])
	// P.L.PreloadModule("gomod",gomod.Loader)
	P.LuaPools.PreloadModule("gomod",gomod.Loader)
	P.LuaPools.DoString(LuaFun)
	// P.L.Close()
	// P.L.LoadString(LuaFun)

	scriptName := os.Args[3]
	// fmt.Println(arg)

	// b, e := ioutil.ReadFile(file_name)

	// b, e := ioutil.ReadFile(arg)
	// if e != nil {
	// 	fmt.Println("read file error")
	// 	return
	// }
	// s := string(b[:])

	s := readScript(scriptName)
	P.EvalStr(s)
}
