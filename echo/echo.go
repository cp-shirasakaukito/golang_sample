//unixのechoコマンドを実装

package echo

import (
	"os"
	"flag"
)

//flagから取得できるのはポインタ
var omitNewline = flag.Bool("n",false, "don't print final newline")

const (
	Space = " "
	Newline = "\n"
)

func Echo() {
	flag.Parse()
	var s string = "" //表示用文字列

	//文字列を作成
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	//改行
	if !*omitNewline {
		s += Newline
	}

	//標準出力
	os.Stdout.WriteString(s)
}