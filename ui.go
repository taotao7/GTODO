package main

import "github.com/marcusolsson/tui-go"

func ui() *tui.Box {
	todoLists := getListDatas()[0]
	completeLists := getListDatas()[1]

	form := tui.NewTable(0, 0)
	form.SetFocused(false)
	form.SetColumnStretch(0, 1)
	form.SetColumnStretch(1, 1)
	form.AppendRow(
		tui.NewLabel("🚅TODOS"),
		tui.NewLabel("🔥COMPLETE"),
	)

	for _, value := range todoLists {
		form.AppendRow(tui.NewLabel(value))
	}

	for _, value := range completeLists {
		form.AppendRow(tui.NewLabel(value))
	}

	box := tui.NewVBox(
		form,
	)
	box.SetTitle("TODOAPP🚀")
	box.SetBorder(true)
	box.SetSizePolicy(tui.Preferred, tui.Maximum)

	return box
}
