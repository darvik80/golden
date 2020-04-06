package widgets

import (
	"fmt"
	"net/http"
)

type FormWidget struct {
	Action string
	Method string
	Widget IWidget
}

func (self *FormWidget) SetAction(s string) *FormWidget {
	self.Action = s
	return self
}

func (self *FormWidget) SetMethod(s string) *FormWidget {
	self.Method = s
	return self
}

func (self *FormWidget) SetWidget(w IWidget) *FormWidget {
	self.Widget = w
	return self
}

func NewFormWidget() *FormWidget {
	fw:= new(FormWidget)
	return fw
}

func (self *FormWidget) Render(w http.ResponseWriter) error {
	fmt.Fprintf(w, "<form action=\"%s\" method=\"%s\">\n", self.Action, self.Method)
	self.Widget.Render(w)
	fmt.Fprintf(w, "</form>\n")
	return nil
}