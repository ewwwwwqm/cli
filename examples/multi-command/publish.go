package main

import (
	"github.com/mkideal/cli"
)

var publishCmd = app.Register(&cli.Command{
	Name: "publish",
	Desc: "Publish golang application",
	Argv: func() interface{} { return new(publishT) },
	Fn:   publish,
})

type publishT struct {
	Help   bool   `cli:"h,help" usage:"display help information"`
	Dir    string `cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
	List   bool   `cli:"l,list" usage:"list all sub commands"`
}

func publish(ctx *cli.Context) error {
	argv := ctx.Argv().(*publishT)

	if argv.Help {
		ctx.String(ctx.Usage())
		return nil
	}
	if argv.List {
		ctx.String(ctx.Command().ListChildren("", "    "))
		return nil
	}
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	return nil
}