package main

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"

	"github.com/MapleLeafMakers/tview"
)

const form = `[green]package[white] main

[green]import[white] (
    [red]"github.com/rivo/tview"[white]
)

[green]func[white] [yellow]main[white]() {
    form := tview.[yellow]NewForm[white]().
		[yellow]AddImage[white]([red]"Photo:"[white], img, [red]0[white], [red]12[white], 0[white]).
        [yellow]AddInputField[white]([red]"First name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddInputField[white]([red]"Last name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddDropDown[white]([red]"Role:"[white], [][green]string[white]{
            [red]"Engineer"[white],
            [red]"Manager"[white],
            [red]"Administration"[white],
        }, [red]0[white], nil).
        [yellow]AddCheckbox[white]([red]"On vacation:"[white], false, nil).
        [yellow]AddPasswordField[white]([red]"Password:"[white], [red]""[white], [red]10[white], [red]'*'[white], nil).
        [yellow]AddTextArea[white]([red]"Notes:"[white], [red]""[white], [red]0[white], [red]5[white], [red]0[white], nil).
        [yellow]AddButton[white]([red]"Save"[white], [yellow]func[white]() { [blue]/* Save data */[white] }).
        [yellow]AddButton[white]([red]"Cancel"[white], [yellow]func[white]() { [blue]/* Cancel */[white] })
    tview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](form, true).
        [yellow]Run[white]()
}`

const photo = `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`

// Form demonstrates forms.
func Form(nextSlide func()) (title string, content tview.Primitive) {
	b, _ := base64.StdEncoding.DecodeString(photo)
	img, _ := jpeg.Decode(bytes.NewReader(b))
	f := tview.NewForm().
		AddImage("Photo:", img, 0, 12, 0).
		AddInputField("First name:", "", 0, nil, nil).
		AddInputField("Last name:", "", 0, nil, nil).
		AddDropDown("Role:", []string{"Engineer", "Manager", "Administration"}, 0, nil).
		AddCheckbox("On vacation:", false, nil).
		AddPasswordField("Password:", "", 10, '*', nil).
		AddTextArea("Notes:", "", 0, 2, 0, nil).
		AddButton("Save", nextSlide).
		AddButton("Cancel", nextSlide)
	f.SetBorder(true).SetTitle("Employee Information")
	return "Forms", Code(f, 36, 31, form)
}
