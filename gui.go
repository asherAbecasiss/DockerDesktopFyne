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
	swarm := d.GetSwarmTab()
	service := d.GetServicesTab()

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
		container.NewTabItem("Docker ", dockerTab),
		container.NewTabItem("Swarm ", swarm),
		container.NewTabItem("Services ", service),
	)
	text1 := canvas.NewText("Admin101", color.Black)
	text1.TextSize = 20

	grid7 := container.NewBorder(text1, nil, tabs, nil, nil)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(grid7)
	myWindow.ShowAndRun()
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
func (d *DockerApi) updateTime(clock *container.AppTabs) int {
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

//Swarm

func (d *DockerApi) GetSwarmListGui() []container.TabItem {
	dockerList := d.GetSwarmNode()
	var list []container.TabItem

	for _, v := range dockerList {

		Id := canvas.NewText("Id: "+v.ID, color.Black)
		formatted := v.CreatedAt.Format("Created At: 2006-01-02 03:04:05 UTC")
		CreatedAt := canvas.NewText(formatted, color.Black)
		ManagerStatus := canvas.NewText("Worker ", color.Black)
		if v.ManagerStatus != nil {
			ManagerStatus = canvas.NewText("Manager "+v.ManagerStatus.Addr, color.Black)
		}

		Id.TextSize = TXT_SIZE
		ManagerStatus.TextSize = TXT_SIZE

		id := binding.NewString()
		id.Set(v.ID)
		list = append(list, container.TabItem{
			Text: v.Description.Hostname,
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
					ManagerStatus,
					Id,
					CreatedAt,
				),
			),
		})

	}

	return list
}

func (d *DockerApi) GetServicesListGui() []container.TabItem {
	dockerList := d.GetDockerServices()
	var list []container.TabItem

	for _, v := range dockerList {

		Id := canvas.NewText("Id: "+v.ID, color.Black)
		RunningTasks := canvas.NewText("RunningTasks: "+strconv.FormatUint(uint64(v.Version.Index), 10), color.Black)
		Name := canvas.NewText("Service Name: "+v.Spec.Name, color.Black)
		formatted := v.CreatedAt.Format("Created At: 2006-01-02 03:04:05 UTC")
		CreatedAt := canvas.NewText(formatted, color.Black)
		ManagerStatus := canvas.NewText("Worker ", color.Black)
		Image := canvas.NewText("Image: "+v.Spec.TaskTemplate.ContainerSpec.Image, color.Black)

		Id.TextSize = TXT_SIZE
		Image.TextSize = TXT_SIZE
		ManagerStatus.TextSize = TXT_SIZE

		id := binding.NewString()
		id.Set(v.ID)
		list = append(list, container.TabItem{
			Text: v.Spec.Name,
			Content: container.New(layout.NewGridLayout(1),
				container.New(layout.NewVBoxLayout(),
					container.New(layout.NewGridLayout(4),
						widget.NewButton("Remove", func() {

							log.Println("tapped -->")
							s, _ := id.Get()
							d.DockerServicesUpdate(s)
						}),
						widget.NewButton("cl", func() {
							log.Println("tapped")
						}),
					),
					Name,
					ManagerStatus,
					RunningTasks,
					Id,
					CreatedAt,
					Image,
				),
			),
		})

	}

	return list
}

func (d *DockerApi) GetSwarmTab() *fyne.Container {
	temp := d.GetSwarmListGui()

	tabs := container.NewAppTabs()

	for i, _ := range temp {
		tabs.Append(&temp[i])
	}
	tabs.SetTabLocation(container.TabLocationLeading)

	lenContainer := strconv.Itoa(len(tabs.Items)) + " Nodes"
	text1 := canvas.NewText("1", color.Black)
	card := widget.NewCard(lenContainer, "W", text1)
	grid2 := container.New(layout.NewVBoxLayout(),

		card,
	)

	grid7 := container.NewBorder(grid2, nil, tabs, nil, nil)

	return grid7

}

func (d *DockerApi) UpdateServicesTab(clock *container.AppTabs, n *int) int {
	fmt.Println("restart")
	temp := d.GetServicesListGui()

	tabs := *container.NewAppTabs()

	for i, _ := range temp {
		tabs.Append(&temp[i])
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	clock.Items = tabs.Items
	*n = len(tabs.Items)
	return len(clock.Items)

}

func (d *DockerApi) GetServicesTab() *fyne.Container {
	temp := d.GetServicesListGui()

	tabs := container.NewAppTabs()

	for i, _ := range temp {
		tabs.Append(&temp[i])
	}
	var n *int = new(int)

	*n = len(tabs.Items)

	num := binding.NewString()
	num.Set(strconv.Itoa(len(tabs.Items)))
	tabs.SetTabLocation(container.TabLocationLeading)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			log.Println("New document")

			s := d.UpdateServicesTab(tabs, n)
			num.Set(strconv.Itoa(s))
			fmt.Println(s)

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
	s1 := strconv.Itoa(*n)

	lenContainer := s1 + " Services"
	text1 := canvas.NewText("1", color.Black)
	card := widget.NewCard(lenContainer, "W", text1)
	grid2 := container.New(layout.NewVBoxLayout(),
		card,
		toolbar,
	)

	grid7 := container.NewBorder(grid2, nil, tabs, nil, nil)

	return grid7

}
