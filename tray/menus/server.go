package menus

import (
	"App/tray/locales"

	"github.com/getlantern/systray"
)

type ServerListener struct {
	OnStart     func() (ok bool)
	OnStartSwag func() (ok bool)
	OnStop      func() (ok bool)
}

type Server struct {
	Menu      *systray.MenuItem
	Start     *systray.MenuItem
	StartSwag *systray.MenuItem
	Stop      *systray.MenuItem
	chanQuit  chan struct{}
}

func (s *Server) Init() {
	s.Menu = systray.AddMenuItem("", "")
	s.Start = s.Menu.AddSubMenuItemCheckbox("", "", false)
	s.StartSwag = s.Menu.AddSubMenuItemCheckbox("", "", false)
	s.Stop = s.Menu.AddSubMenuItemCheckbox("", "", false)
	s.Stop.Disable()
	s.chanQuit = make(chan struct{}, 1)
	s.ResetLanguage()
}

func (s *Server) ResetLanguage() {
	src := locales.Get()
	s.Menu.SetTitle(src.Server.String())
	s.Menu.SetTooltip(src.Server.String())
	s.Start.SetTitle(src.ServerStart.String())
	s.Start.SetTitle(src.ServerStart.String())
	s.StartSwag.SetTitle(src.ServerStartSwag.String())
	s.StartSwag.SetTitle(src.ServerStartSwag.String())
	s.Stop.SetTitle(src.ServerStop.String())
	s.Stop.SetTitle(src.ServerStop.String())
}

func (s *Server) Watch(listener *ServerListener) {
	go func() {
		for {
			select {
			case <-s.Start.ClickedCh:
				if listener.OnStart() {
					s.Start.Check()
					s.Start.Disable()
					s.StartSwag.Uncheck()
					s.StartSwag.Disable()
					s.Stop.Uncheck()
					s.Stop.Enable()
				}
			case <-s.StartSwag.ClickedCh:
				if listener.OnStartSwag() {
					s.Start.Uncheck()
					s.Start.Disable()
					s.StartSwag.Check()
					s.StartSwag.Disable()
					s.Stop.Uncheck()
					s.Stop.Enable()
				}
			case <-s.Stop.ClickedCh:
				if listener.OnStop() {
					s.Start.Uncheck()
					s.Start.Enable()
					s.StartSwag.Uncheck()
					s.StartSwag.Enable()
					s.Stop.Check()
					s.Stop.Disable()
				}
			case <-s.chanQuit:
				return
			}
		}
	}()
}

func (s *Server) StopWatch() {
	go func() {
		s.chanQuit <- struct{}{}
	}()
}
