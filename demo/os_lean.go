package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var b = flag.Bool("b", false, "bool类型参数")
var s = flag.String("s", "", "string类型参数")

func main() {
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
	for idx, args := range os.Args[1:] {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Println(os.Args[1:])

	flag.Parse()
	fmt.Println("-b:", *b)
	fmt.Println("-s:", *s)
	fmt.Println("其他参数：", flag.Args())

}
