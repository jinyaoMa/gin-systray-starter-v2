package tray

import (
	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

var (
	config *Config

	serverMenu      *systray.MenuItem
	serverStart     *systray.MenuItem
	serverStartSwag *systray.MenuItem
	serverStop      *systray.MenuItem
	languageMenu    *systray.MenuItem
	languageEn      *systray.MenuItem
	languageZh      *systray.MenuItem
	quit            *systray.MenuItem
)

func init() {
	SetConfig(DefaultConfig())
}

func SetConfig(config Config) {
	config = config
}

func Run(callback func()) {
	systray.Run(onReady(callback), onExit)
}

func onReady(callback func()) func() {
	return func() {
		defer callback()

		makeTray()
		resetLanguage()

	}
}

func onExit() {

}

func resetLanguage() {
	systray.SetTitle(config.Locale.Title.String())

	serverMenu.SetTitle(config.Locale.Server.String())
	serverMenu.SetTooltip(config.Locale.Server.String())
	serverStart.SetTitle(config.Locale.ServerStart.String())
	serverStart.SetTooltip(config.Locale.ServerStart.String())
	serverStartSwag.SetTitle(config.Locale.ServerStartSwag.String())
	serverStartSwag.SetTooltip(config.Locale.ServerStartSwag.String())
	serverStop.SetTitle(config.Locale.ServerStop.String())
	serverStop.SetTooltip(config.Locale.ServerStop.String())

	languageMenu.SetTitle(config.Locale.Language.String())
	languageMenu.SetTooltip(config.Locale.Language.String())
	languageEn.SetTitle(config.Locale.LanguageEn.String())
	languageEn.SetTooltip(config.Locale.LanguageEn.String())
	languageZh.SetTitle(config.Locale.LanguageZh.String())
	languageZh.SetTooltip(config.Locale.LanguageZh.String())

	quit.SetTitle(config.Locale.Quit.String())
	quit.SetTooltip(config.Locale.Quit.String())
}

func makeTray() {
	systray.SetIcon(icon)

	serverMenu = systray.AddMenuItem("", "")
	serverStart = serverMenu.AddSubMenuItemCheckbox("", "", false)
	serverStartSwag = serverMenu.AddSubMenuItemCheckbox("", "", false)
	serverStop = serverMenu.AddSubMenuItemCheckbox("", "", false)

	systray.AddSeparator()

	languageMenu = systray.AddMenuItem("", "")
	languageEn = languageMenu.AddSubMenuItemCheckbox("", "", false)
	languageZh = languageMenu.AddSubMenuItemCheckbox("", "", false)

	systray.AddSeparator()

	quit = systray.AddMenuItem("", "")
}
