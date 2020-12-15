package main

import "github.com/marcusolsson/tui-go"

func main() {
	box := ui()
	ui, err := tui.New(box)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
