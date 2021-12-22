package menus

import (
	"App/tray/locales"

	"github.com/getlantern/systray"
)

type LanguageListener struct {
	OnLanguageChange func(locale locales.Locale) (ok bool)
}

type Language struct {
	Menu     *systray.MenuItem
	En       *systray.MenuItem
	Zh       *systray.MenuItem
	chanQuit chan struct{}
}

func (l *Language) Init() {
	l.Menu = systray.AddMenuItem("", "")
	l.En = l.Menu.AddSubMenuItemCheckbox("", "", true)
	l.Zh = l.Menu.AddSubMenuItemCheckbox("", "", false)
	l.chanQuit = make(chan struct{}, 1)
	l.ResetLanguage()
}

func (l *Language) ResetLanguage() {
	src := locales.Get()
	l.Menu.SetTitle(src.Language.String())
	l.Menu.SetTooltip(src.Language.String())
	l.En.SetTitle(src.LanguageEn.String())
	l.En.SetTitle(src.LanguageEn.String())
	l.Zh.SetTitle(src.LanguageZh.String())
	l.Zh.SetTitle(src.LanguageZh.String())
}

func (l *Language) Watch(listener *LanguageListener) {
	go func() {
		for {
			select {
			case <-l.En.ClickedCh:
				if listener.OnLanguageChange(locales.En) {
					l.En.Check()
					l.Zh.Uncheck()
				}
			case <-l.Zh.ClickedCh:
				if listener.OnLanguageChange(locales.Zh) {
					l.En.Uncheck()
					l.Zh.Check()
				}
			case <-l.chanQuit:
				return
			}
		}
	}()
}

func (l *Language) StopWatch() {
	go func() {
		l.chanQuit <- struct{}{}
	}()
}
