package menus

import (
	"App/tray/locales"

	"github.com/getlantern/systray"
)

type QuitListener struct {
	OnQuit func()
}

type Quit struct {
	Menu     *systray.MenuItem
	chanQuit chan struct{}
}

func (q *Quit) Init() {
	q.Menu = systray.AddMenuItem("", "")
	q.chanQuit = make(chan struct{}, 1)
	q.ResetLanguage()
}

func (q *Quit) ResetLanguage() {
	src := locales.Get()
	q.Menu.SetTitle(src.Quit.String())
	q.Menu.SetTooltip(src.Quit.String())
}

func (q *Quit) Watch(listener *QuitListener) {
	go func() {
		for {
			select {
			case <-q.Menu.ClickedCh:
				listener.OnQuit()
			case <-q.chanQuit:
				return
			}
		}
	}()
}

func (q *Quit) StopWatch() {
	go func() {
		q.chanQuit <- struct{}{}
	}()
}
