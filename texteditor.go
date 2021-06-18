package main

import (
	"io/ioutil"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Text Editor")
	edit := widget.NewEntry()
	edit.MultiLine = true
	sc := widget.NewScrollContainer(edit)
	inf := widget.NewLabel("information bar.")

	nf := func() {
		dialog.ShowConfirm("Alert", "Create new document?", func(f bool) {
			if f {
				edit.SetText("")
				inf.SetText("create new document.")
			}
		}, w)
	}

	of := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Open file name.", "OK", "Cancel", f, func(b bool) {
			if b {
				fn := f.Text + ".txt"
				ba, er := ioutil.ReadFile(fn)
				if er != nil {
					dialog.ShowError(er, w)
				} else {
					edit.SetText(string(ba))
					inf.SetText("Open from file '" + fn + "'.")
				}
			}
		}, w)
	}

	sf := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Save file name", "OK", "Cancel", f, func(b bool) {
			if b {
				fn := f.Text + ".txt"
				er := ioutil.WriteFile(fn, []byte(edit.Text), os.ModePerm)
				if er != nil {
					dialog.ShowError(er, w)
					return
				}
				inf.SetText("Save to file '" + fn + "'.")

			}
		}, w)
	}

	qf := func() {
		dialog.ShowConfirm("Alert", "Quit application?", func(f bool) {
			if f {
				a.Quit()
			}
		}, w)
	}

	tf := true

	cf := func() {
		if tf {
			a.Settings().SetTheme(theme.LightTheme())
			inf.SetText("change to Light-Theme.")
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			inf.SetText("change to Light-Theme.")
		}
		tf = !tf
	}

	createMenubar := func() *fyne.MainMenu {
		return fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("New", nf),
				fyne.NewMenuItem("Open...", of),
				fyne.NewMenuItem("Save...", sf),
				fyne.NewMenuItem("Change Theme", cf),
				fyne.NewMenuItem("Quit", qf),
			),
			fyne.NewMenu("Edit",
				fyne.NewMenuItem("Cut", func() {
					edit.TypedShortcut(
						&fyne.ShortcutCut{
							Clipboard: w.Clipboard(),
						},
					)
					inf.SetText("Cut text.")
				}),
				fyne.NewMenuItem("Copy", func() {
					edit.TypedShortcut(
						&fyne.ShortcutCopy{
							Clipboard: w.Clipboard(),
						},
					)
					inf.SetText("Copy text.")
				}),
				fyne.NewMenuItem("Paste", func() {
					edit.TypedShortcut(
						&fyne.ShortcutPaste{
							Clipboard: w.Clipboard(),
						},
					)
					inf.SetText("Paste text.")
				}),
			),
		)
	}

	createToolbar := func() *widget.Toolbar {
		return widget.NewToolbar(
			widget.NewToolbarAction(theme.DocumentCreateIcon(), nf),
			widget.NewToolbarAction(theme.FolderOpenIcon(), of),
			widget.NewToolbarAction(theme.DocumentSaveIcon(), sf),
		)
	}

	mb := createMenubar()
	tb := createToolbar()

	w.SetMainMenu(mb)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(tb, inf, nil, nil), tb, inf, sc),
	)

	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
