package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	defer ui.Close()

	Hello := widgets.NewParagraph()
	Hello.Text = "Hello World!"
	Hello.Title = "Text"
	Hello.ColorModel()
	Hello.SetRect(0, 0, len(Hello.Text) + 2, 3)

	ui.Render(Hello)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
