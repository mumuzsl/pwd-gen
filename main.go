package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
	"strings"
)

type argT struct {
	cli.Helper2
	RawStrArray []string `cli:"*s" usage:"原始字符串数组"`
	Pwd         string   `pw:"p,password" usage:"密码"`
	Separate    string   `cli:"S" usage:"拼接字符" dft:"_"`
	Prefix      string   `cli:"P" usage:"前缀，截取md5的字符串中包含小写字母和数字，有些密码可能还需要特殊字符和大写字母，所以加上这样的前缀。" dft:"0X_"`
	Point       int      `usage:"截取md5字符串起点" dft:"17"`
	Length      int      `usage:"截取md5字符串长度" dft:"8"`
}

var root = &cli.Command{
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		originalStr := strings.Join(argv.RawStrArray, argv.Separate)

		md5Str, result, err := Gen(originalStr, argv.Prefix, argv.Point, argv.Length)
		if err != nil {
			return err
		}

		ctx.String("使用的字符串: %s \nmd5生成的字符串: %s \n结果: %s", originalStr, md5Str, result)

		return nil
	},
}

func main() {
	err := cli.Root(root, cli.Tree(JsonFileCommand)).Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
