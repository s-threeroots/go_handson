package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is Sample widget.")

	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				l.SetText("select 'New' menu item.")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)

	w.SetMainMenu(mm)
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("ok", nil),
		),
	)

	/*
		// toolbar
		tb := widget.NewToolbar(
			widget.NewToolbarAction(theme.HomeIcon(), func() {
				l.SetText("Select Home Icon!")
			}),
			widget.NewToolbarAction(theme.InfoIcon(), func() {
				l.SetText("Select Infomation Icon!")
			}),
		)

		w.SetContent(
			fyne.NewContainerWithLayout(
				layout.NewBorderLayout(nil, tb, nil, nil), l, tb),
		)
	*/
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

/*
// section3
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	//l := widget.NewLabel("Hello Fyne!")


	// layout
	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Bottom", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(bt, bb, bl, br),
			bt, bb, bl, br,
			widget.NewLabel("Center."),
		),
	)


	// tab container
	w.SetContent(
		widget.NewVBox(
			widget.NewTabContainer(
				widget.NewTabItem("First", widget.NewLabel("This is First tab item.")),
				widget.NewTabItem("Second", widget.NewLabel("This is Second tab item.")),
			),
		),
	)


	// group
	ck1 := widget.NewCheck("check 1", nil)
	ck2 := widget.NewCheck("check 2", nil)

	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewGroup("group", ck1, ck2),
			widget.NewButton("OK", func() {
				re := "result: "
				if ck1.Checked {
					re += "Check-1 "
				}
				if ck2.Checked {
					re += "Check-2 "
				}
				l.SetText(re)
			}),
		),
	)



	// form
	ne := widget.NewEntry()
	pe := widget.NewPasswordEntry()

	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewForm(
				widget.NewFormItem("Name", ne),
				widget.NewFormItem("Pass", pe),
			),
			widget.NewButton("OK", func() {
				l.SetText(ne.Text + "&" + pe.Text)
			}),
		),
	)

	w.ShowAndRun()

}
*/

/*
// section 2
func main() {

	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")

	// progress bar
	v := 0.
	p := widget.NewProgressBar()
	b := widget.NewButton("up", func() {
		v += 0.1
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})

	w.SetContent(
		widget.NewVBox(
			l, p, b,
		),
	)


		// radio
		r := widget.NewRadio(
			[]string{"one", "two", "three"},
			func(s string) {
				if s == "" {
					l.SetText("not selected.")
				} else {
					l.SetText("selected: " + s)
				}
			})
		r.SetSelected("one")
		w.SetContent(
			widget.NewVBox(
				l, r,
			),
		)


		// check box
			c := widget.NewCheck("Check!", func(f bool) {
				if f {
					l.SetText("Checked")
				} else {
					l.SetText("unchecked")
				}
			})

			c.SetChecked(true)
			w.SetContent(
				widget.NewVBox(
					l, c,
				),
			)

	w.ShowAndRun()
}
*/

/*
// section 1
func main() {

	a := app.New()

	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")

	w.SetContent(
		widget.NewVBox(
			l, e,
			widget.NewButton("click me!", func() {
				n, _ := strconv.Atoi(e.Text)
				l.SetText("total: " + strconv.Itoa(total(n)))
			}),
		),
	)

	a.Settings().SetTheme(theme.DarkTheme())
	w.ShowAndRun()
}

func total(n int) int {
	t := 0

	for i := 1; i <= n; i++ {
		t += i
	}

	return t
}
*/
