package main

import (
	"github.com/mkideal/cli"
)

type Data struct {
	File []Domain `json:"data"`
}

type jsonFileT struct {
	cli.Helper2
	File Data   `cli:"*d,data"  usage:"文件路径" parser:"jsonfile"`
	Pwd  string `pw:"p,password" usage:"密码" prompt:"请输入密码"`
}

var JsonFileCommand = &cli.Command{
	Name: "json",
	Argv: func() interface{} { return new(jsonFileT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*jsonFileT)

		for _, domain := range argv.File.File {
			md5Str, result, err := domain.Gen(argv.Pwd)
			if err != nil {
				return err
			}

			ctx.String("使用的字符串: %s \nmd5生成的字符串: %s \n结果: %s\n", domain.OriginalStr, md5Str, result)
		}

		return nil
	},
}
