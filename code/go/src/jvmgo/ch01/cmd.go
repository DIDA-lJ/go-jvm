package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd 定义一个结构体来表示命令行选项和参数/*
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

// os 包中定义了一个 Args 变量，其中存放传递给命令行的全部参数
// 如果直接处理 os.Args 变量，需要写很多的代码，不过 go 有许多内置的 flag 包，这个包可以帮助我们处理命令行选项
func parseCmd() *Cmd {
	cmd := &Cmd{}
	// 设置 flag.Usage 变量，然后将 printUsage 函数赋值给它
	flag.Usage = printUsage
	// 调用 flag 包来解析各种 Var（）函数设置需要解析的选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	// 调用 Parse() 函数解析选项，如果 Parse() 函数解析失败，它就调用 printUsage（）函数把命令的用法打印到控制台。
	flag.Parse()

	//如果解析成功，调用flag.Args（）函数可以捕获其他没有被解析 的参数。其中第一个参数就是主类名，剩下的是要传递给主类的参
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

//如果Parse（）函数解析失败，它就调用  printUsage（）函数把命令的用法打印到控制台
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
