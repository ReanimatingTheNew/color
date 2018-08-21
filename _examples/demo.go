package main

import (
	"fmt"
	"github.com/gookit/color"
)

// go run ./_examples/demo.go
func main() {
	colorUsage()

}

func colorUsage() {
	// simple usage
	color.Cyan.Printf("Simple to use %s\n", "color")

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	color.New(color.FgMagenta, color.BgBlack).Println("custom color style")
	// can also:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

	// internal theme/style:
	color.Info.Tips("message")
	color.Info.Prompt("message")
	color.Info.Println("message")
	color.Warn.Println("message")
	color.Error.Println("message")

	// use defined color tag
	color.Print("use color tag: <suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// use custom color tag
	color.Print("custom color tag: <fg=yellow;bg=black;op=underscore;>hello, welcome</>\n")

	// set a color tag
	color.Tag("info").Println("info style message")

	// tips
	color.Tips("info").Print("tips style message")
	color.Tips("warn").Print("tips style message")

	// lite tips
	color.LiteTips("info").Print("lite tips style message")
	color.LiteTips("warn").Print("lite tips style message")

	i := 0
	fmt.Print("\n- All Available color Tags: \n\n")

	for tag, _ := range color.GetColorTags() {
		i++
		color.Tag(tag).Print(tag)

		if i%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
