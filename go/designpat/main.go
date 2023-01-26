package main

import "fmt"

func main() {
	done := make(chan struct{})
	result := generator(done)
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
	close(done)
}

type Mediator interface {
	createColleagues()
	colleagueChanged()
}

type LoginFrame struct {
	checkBoxGroup      ColleagueButton
	colleagueTextField ColleagueTextField
}

func (l LoginFrame) createColleagues() {
}
func (l LoginFrame) colleagueChanged() {
}
func newLoginFrame(title string) {
	lf := LoginFrame{
		checkBoxGroup:      ColleagueButton{},
		colleagueTextField: ColleagueTextField{},
	}
	lf.
}

type Colleague interface {
	setMediator(mediator Mediator)
	setColleagueEnabled(enabled bool)
}

type ColleagueButton struct {
	mediator Mediator
}

func (c ColleagueButton) setMediator(mediator Mediator) {
	c.mediator = mediator
}
func (c ColleagueButton) setColleagueEnabled(enabled bool) {
	setEnabled(enabled)
}
func setEnabled(enabled bool) {}

type ColleagueTextField struct {
	mediator Mediator
}

func (c ColleagueTextField) setMediator(mediator Mediator) {
	c.mediator = mediator
}
func (c ColleagueTextField) setColleagueEnabled(enabled bool) {
	setEnabled(enabled)
}
