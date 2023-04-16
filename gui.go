package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/widget"
)

const (
	TXT_SIZE = 15
)

func (d *DockerApi) Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.NewSize(1420, 800))

	tab := d.getDockerTab()

	dockerTab := d.GetDockerTab(tab)
	// toolbar := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
	// 		log.Println("New document")
	// 		d.updateTime(tab)
	// 		// tabs.Refresh()
	// 	}),
	// 	widget.NewToolbarSeparator(),
	// 	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.HelpIcon(), func() {
	// 		log.Println("Display help")
	// 	}),
	// )

	tabs := container.NewAppTabs(
		container.NewTabItem("Docker", dockerTab),
	)
	text1 := canvas.NewText("Admin101", color.Black)
	text1.TextSize = 20

	grid7 := container.NewBorder(text1, nil, tabs, nil, nil)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(grid7)
	myWindow.ShowAndRun()
}

func (d *DockerApi) getTabel() *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	return list
}

func (d *DockerApi) GetList() []container.TabItem {

	dockerList := d.GetDockerContainer()
	var list []container.TabItem

	for _, v := range dockerList {
		Name := canvas.NewText("Image Name: "+v.Names[0], color.Black)
		Id := canvas.NewText("Id: "+v.ID, color.Black)

		Image := canvas.NewText("Image: "+v.Image, color.Black)
		ImageID := canvas.NewText("ImageID: "+v.ImageID, color.Black)
		Status := canvas.NewText("Status: "+v.Status, color.Black)
		State := canvas.NewText("State: "+v.State, color.Black)
		Created := canvas.NewText("Created: "+strconv.FormatInt(int64(v.Created), 10), color.Black)

		Name.TextSize = 20
		Id.TextSize = TXT_SIZE
		Image.TextSize = TXT_SIZE
		ImageID.TextSize = TXT_SIZE
		Status.TextSize = TXT_SIZE
		State.TextSize = TXT_SIZE
		Created.TextSize = TXT_SIZE
		id := binding.NewString()
		id.Set(v.ID)
		list = append(list, container.TabItem{
			Text: v.Names[0],
			Content: container.New(layout.NewGridLayout(1),
				container.New(layout.NewVBoxLayout(),
					container.New(layout.NewGridLayout(4),
						widget.NewButton("Restart", func() {

							log.Println("tapped")
							s, _ := id.Get()
							d.RestartContainerID(s)
						}),
						widget.NewButton("cl", func() {
							log.Println("tapped")
						}),
					),
					Name,
					Id,
					Image,
					ImageID,
					Status,
					State,
					Created,
				),
			),
		})

	}

	return list
}

func (d *DockerApi) GetDockerTab(tab *container.AppTabs) *fyne.Container {

	// list := d.getTabel()
	// text1 := canvas.NewText("1", color.Black)

	// text2 := canvas.NewText("2", color.Black)

	// green2 := color.NRGBA{R: 0, G: 150, B: 0, A: 255}
	// green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	// text4 := canvas.NewRectangle(color.Black)
	// text5 := canvas.NewRectangle(green2)
	// text6 := canvas.NewRectangle(green)
	// text7 := canvas.NewRectangle(green)
	// temp := d.GetList()

	// tabs := container.NewAppTabs()
	// for i, _ := range temp {
	// 	tabs.Append(&temp[i])
	// }

	// tabs.SetTabLocation(container.TabLocationLeading)

	// grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)), grid2)
	// tabs := d.GetDockerTab()
	// grid7 := container.New(layout.NewGridLayout(1), tab)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			log.Println("New document")
			d.updateTime(tab)
			// tabs.Refresh()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	lenContainer := "Docker Container " + strconv.Itoa(len(tab.Items))
	card := widget.NewCard(lenContainer, "W", toolbar)
	grid2 := container.New(layout.NewVBoxLayout(),
		
		card,
	)

	grid7 := container.NewBorder(grid2, nil, tab, nil, nil)
	fmt.Println(len(tab.Items))
	return grid7
}

func (d *DockerApi) getDockerTab() *container.AppTabs {
	temp := d.GetList()

	tabs := container.NewAppTabs()
	for i, _ := range temp {
		tabs.Append(&temp[i])
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	return tabs
}
func (d *DockerApi) updateTime(clock *container.AppTabs) int{
	fmt.Println("restart")
	temp := d.GetList()

	tabs := *container.NewAppTabs()

	for i, _ := range temp {
		tabs.Append(&temp[i])
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	clock.Items = tabs.Items
	return len(temp)

}
