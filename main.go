package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"./regex"
)

var (
	expression = ""
	debug      = false
	help       = false
)

func init() {
	flag.BoolVar(&help, "h", false, "帮助信息")
	flag.BoolVar(&debug, "d", false, "调试信息")
	flag.StringVar(&expression, "e", "(a|b|c)*h|l", "正则式")
}

func usage() {
	fmt.Fprintf(os.Stderr, " 正则引擎\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if help {
		usage()
		return
	}
	if len(expression) == 0 {
		fmt.Println("-e参数输入正则式")
		return
	}
	regex := regex.NewRegex(expression, debug)
	fmt.Println("正则式:" + expression)
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("输入匹配字符串")
		match, _ := rd.ReadString('\n')
		match = match[:len(match)-1]
		fmt.Printf("%s=>%t \n", append([]byte(match), 0), regex.Match(match))
	}
}
