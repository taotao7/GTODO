package main

import "github.com/marcusolsson/tui-go"

func ui() *tui.Box {
	user := tui.NewEntry()
	user.SetFocused(true)
	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("🙍用户"), tui.NewLabel("🔑密码"))
	form.AppendRow(user, password)

	box := tui.NewVBox(
		tui.NewPadder(1, 0, form),
	)
	box.SetTitle("TODO🚀")
	box.SetBorder(true)
	box.SetSizePolicy(tui.Preferred, tui.Maximum)

	tui.DefaultFocusChain.Set(user, password)
	return box
}
