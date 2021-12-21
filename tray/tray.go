package tray

import (
	"App/tray/locales"
	"log"

	"github.com/getlantern/systray"
)

type Listener struct {
	OnReady  func()
	OnExit   func()
	OnServer func(run bool, withSwag bool) (ok bool)
}

//go:embed icons/icon.ico
var icon []byte

var (
	currentConfig   *Config
	currentListener *Listener

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

func SetConfig(config *Config) {
	currentConfig = config
}

func Run(listener *Listener) {
	currentListener = listener
	systray.Run(onReady, onExit)
}

func onReady() {
	defer currentListener.OnReady()

	makeTray()
	resetLanguage()
	resetStatus()

	go watchClicks()

	log.Println("Tray starting!")
}

func onExit() {
	defer currentListener.OnExit()

	if currentConfig.StartServer {
		currentListener.OnServer(false, false)
	}

	log.Println("Tray exiting!")
}

func watchClicks() {
	for {
		isPass := false
		select {
		case <-serverStart.ClickedCh:
			isPass = currentListener.OnServer(true, false)
			if isPass {
				currentConfig.StartServer = true
				currentConfig.EnableSwagger = false
			}
		case <-serverStartSwag.ClickedCh:
			isPass = currentListener.OnServer(true, true)
			if isPass {
				currentConfig.StartServer = true
				currentConfig.EnableSwagger = true
			}
		case <-serverStop.ClickedCh:
			isPass = currentListener.OnServer(false, false)
			if isPass {
				currentConfig.StartServer = false
				currentConfig.EnableSwagger = false
			}
		case <-languageEn.ClickedCh:
			currentConfig.Locale = locales.En
			languageEn.Check()
			languageZh.Uncheck()
			resetLanguage()
			isPass = true
		case <-languageZh.ClickedCh:
			currentConfig.Locale = locales.Zh
			languageEn.Uncheck()
			languageZh.Check()
			resetLanguage()
			isPass = true
		case <-quit.ClickedCh:
			systray.Quit()
			return
		}
		if isPass {
			resetStatus()
		}
	}
}

func resetStatus() {
	var serverMessage string
	var languageMessage string

	if currentConfig.StartServer {
		if currentConfig.EnableSwagger {
			serverMessage = currentConfig.Locale.ServerStartSwag.String()
		} else {
			serverMessage = currentConfig.Locale.ServerStart.String()
		}
	} else {
		serverMessage = currentConfig.Locale.ServerStop.String()
	}

	if currentConfig.Locale == locales.En {
		languageMessage = currentConfig.Locale.LanguageEn.String()
	} else {
		languageMessage = currentConfig.Locale.LanguageZh.String()
	}

	systray.SetTooltip(currentConfig.Locale.Tooltip.String(serverMessage, languageMessage))
}

func resetLanguage() {
	systray.SetTitle(currentConfig.Locale.Title.String())

	serverMenu.SetTitle(currentConfig.Locale.Server.String())
	serverMenu.SetTooltip(currentConfig.Locale.Server.String())
	serverStart.SetTitle(currentConfig.Locale.ServerStart.String())
	serverStart.SetTooltip(currentConfig.Locale.ServerStart.String())
	serverStartSwag.SetTitle(currentConfig.Locale.ServerStartSwag.String())
	serverStartSwag.SetTooltip(currentConfig.Locale.ServerStartSwag.String())
	serverStop.SetTitle(currentConfig.Locale.ServerStop.String())
	serverStop.SetTooltip(currentConfig.Locale.ServerStop.String())

	languageMenu.SetTitle(currentConfig.Locale.Language.String())
	languageMenu.SetTooltip(currentConfig.Locale.Language.String())
	languageEn.SetTitle(currentConfig.Locale.LanguageEn.String())
	languageEn.SetTooltip(currentConfig.Locale.LanguageEn.String())
	languageZh.SetTitle(currentConfig.Locale.LanguageZh.String())
	languageZh.SetTooltip(currentConfig.Locale.LanguageZh.String())

	quit.SetTitle(currentConfig.Locale.Quit.String())
	quit.SetTooltip(currentConfig.Locale.Quit.String())
}

func makeTray() {
	systray.SetIcon(icon)

	serverMenu = systray.AddMenuItem("", "")
	serverStart = serverMenu.AddSubMenuItemCheckbox("", "", currentConfig.StartServer && !currentConfig.EnableSwagger)
	serverStartSwag = serverMenu.AddSubMenuItemCheckbox("", "", currentConfig.StartServer && currentConfig.EnableSwagger)
	serverStop = serverMenu.AddSubMenuItemCheckbox("", "", false)
	if currentConfig.StartServer {
		serverStart.Disable()
		serverStartSwag.Disable()
		currentListener.OnServer(currentConfig.StartServer, currentConfig.EnableSwagger)
	} else {
		serverStop.Disable()
	}

	systray.AddSeparator()

	languageMenu = systray.AddMenuItem("", "")
	languageEn = languageMenu.AddSubMenuItemCheckbox("", "", currentConfig.Locale == locales.En)
	languageZh = languageMenu.AddSubMenuItemCheckbox("", "", currentConfig.Locale == locales.Zh)

	systray.AddSeparator()

	quit = systray.AddMenuItem("", "")
}
