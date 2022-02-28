package main

import (
	"encoding/json"
	"fmt"
	"github.com/mkideal/cli"
	"os"
	"path/filepath"
	"strings"
)

type argT struct {
	cli.Helper2
	RawStrArray  []string `cli:"*s" usage:"原始字符串数组"`
	Pwd          string   `pw:"p,password" usage:"密码"`
	Separate     string   `cli:"S" usage:"拼接字符" dft:"_"`
	Prefix       string   `cli:"P" usage:"前缀，截取md5的字符串中包含小写字母和数字，有些密码可能还需要特殊字符和大写字母，所以加上这样的前缀。" dft:"0X_"`
	Point        int      `usage:"截取md5字符串起点" dft:"25"`
	Length       int      `usage:"截取md5字符串长度" dft:"8"`
	AccountIndex int      `cli:"a" usage:"账号索引"`
}

type Config struct {
	Account   []string `json:"account"`
	CommonPwd string   `json:"common_pwd"`
}

var root = &cli.Command{
	Argv: func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		config, err := getConfig()
		if err != nil {
			return err
		}

		argv := ctx.Argv().(*argT)

		index := 1

		rear := append([]string{}, argv.RawStrArray[index:]...)
		argv.RawStrArray = append(argv.RawStrArray[0:index], config.Account[argv.AccountIndex])
		argv.RawStrArray = append(argv.RawStrArray, rear...)
		argv.RawStrArray = append(argv.RawStrArray, config.CommonPwd)

		originalStr := strings.Join(argv.RawStrArray, argv.Separate)

		md5Str, result, err := Gen(originalStr, argv.Prefix, argv.Point, argv.Length)
		if err != nil {
			return err
		}

		ctx.String("使用的字符串: %s \nmd5生成的字符串: %s \n结果: %s \n", originalStr, md5Str, result)

		return nil
	},
}

func getConfig() (*Config, error) {
	var config *Config

	executable, err := os.Executable()
	if err != nil {
		return config, err
	}
	dir := filepath.Dir(executable)

	filePtr, err := os.Open(dir + "/config.json")
	if err != nil {
		fmt.Println("未能打开配置文件", err.Error())
		return nil, err
	}
	defer filePtr.Close()

	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
		return config, err
	}

	return config, nil
}

func main() {
	err := cli.Root(root, cli.Tree(JsonFileCommand)).Run(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
