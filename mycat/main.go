package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type sliceString []string

func (s *sliceString) String() string { return fmt.Sprintf("%s", *s) }
func (s *sliceString) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("-i フラグは複数使用できません")
	}
	for _, v := range strings.Split(value, ",") {
		*s = append(*s, v)
	}
	return nil
}

var impf sliceString

var num = flag.Bool("n", false, "行番号を各行につける（累計）\n")

var i int = 1

func init() {
	flag.Var(&impf, "i", "表示するファイル（複数可）\n記法: -i value1,value2...\n")
}

func main() {
	flag.Parse()
	flagArgs()

}

func flagArgs() {
	if len(impf) == 0 {
		fmt.Fprintln(os.Stderr, "ファイルが指定されていません")
		return
	}
	for j := 0; j < len(impf); j++ {
		readLine(impf[j])
	}
}

func readLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました\n", err)
		return
	}
	s := bufio.NewScanner(f)
	var t string
	for s.Scan() {
		t = s.Text()
		switch *num {
		case true:
			fmt.Printf("%3d: %s\n", i, t)
		case false:
			fmt.Printf("     %s\n", t)
		}
		i++
	}
	if s.Err() != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました\n", err)
	}

	defer f.Close()
}
