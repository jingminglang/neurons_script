package main

import (
	P "./lib/parser"
	//	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var file_name string

func init() {
	flag.StringVar(&file_name, "f", "", "脚本文件")
	flag.Parse()
	if file_name == "" {
		flag.Usage()
		os.Exit(2)
	}
}

func main() {

	b, e := ioutil.ReadFile(file_name)
	if e != nil {
		fmt.Println("read file error")
		return
	}
	s := string(b[:])
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(os.Stdin)
	// s := buf.String()
	P.EvalStr(s)
}
