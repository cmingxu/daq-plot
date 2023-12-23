package main

type WidgetsPanel struct {
	Updateable
}

func NewWidgetsPanel() *WidgetsPanel {
	return &WidgetsPanel{}
}

func (p *WidgetsPanel) Update() {}
