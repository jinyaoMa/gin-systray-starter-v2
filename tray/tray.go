package tray

import (
	"App/logger"
	"App/routers"
	"App/tray/locales"
	"App/tray/menus"
	_ "embed"
	"errors"

	"github.com/getlantern/systray"
)

var (
	ErrSetConfigWhileRunning error = errors.New("tray.SetConfig: cannot set config while running")
)

//go:embed icons/icon.ico
var icon []byte

var (
	isRunning bool
	config    *Config
)

var (
	MenuServer   *menus.Server
	MenuLanguage *menus.Language
	MenuQuit     *menus.Quit
)

func init() {
	MenuServer = &menus.Server{}
	MenuLanguage = &menus.Language{}
	MenuQuit = &menus.Quit{}

	isRunning = false
	SetConfig(DefaultConfig())
}

func Start() {
	if isRunning {
		return
	}

	isRunning = true
	systray.Run(onReady, onExit)
}

func Stop() {
	if !isRunning {
		return
	}

	isRunning = false
	systray.Quit()
}

func SetConfig(conf *Config) {
	if isRunning {
		logger.Tray.Fatalln(ErrSetConfigWhileRunning)
		return
	}
	config = conf
}

func onReady() {
	systray.SetIcon(icon)
	resetTooltipLanguage()

	MenuServer.Init()
	MenuServer.Watch(&menus.ServerListener{
		OnStart: func() bool {
			routers.Start(false, false)
			config.StartServer = true
			config.EnableSwag = false
			resetTooltipLanguage()
			return true
		},
		OnStartSwag: func() bool {
			routers.Start(true, false)
			config.StartServer = true
			config.EnableSwag = true
			resetTooltipLanguage()
			return true
		},
		OnStop: func() bool {
			routers.Stop()
			config.StartServer = false
			config.EnableSwag = false
			resetTooltipLanguage()
			return true
		},
	})

	systray.AddSeparator()

	MenuLanguage.Init()
	MenuLanguage.Watch(&menus.LanguageListener{
		OnLanguageChange: func(locale locales.Locale) (ok bool) {
			if ok = locales.Set(locale); ok {
				config.Locale = locale
				resetTooltipLanguage()

				MenuServer.ResetLanguage()
				MenuLanguage.ResetLanguage()
				MenuQuit.ResetLanguage()
			}
			return
		},
	})

	systray.AddSeparator()

	MenuQuit.Init()
	MenuQuit.Watch(&menus.QuitListener{
		OnQuit: func() {
			Stop()
		},
	})

	handleConfig()
	logger.Tray.Println("Tray is running!")
}

func onExit() {
	MenuServer.StopWatch()
	MenuLanguage.StopWatch()
	MenuQuit.StopWatch()
	routers.Stop()

	logger.Tray.Println("Tray is exiting!")
}

func handleConfig() {
	var signal struct{}

	if config.StartServer {
		if config.EnableSwag {
			MenuServer.StartSwag.ClickedCh <- signal
		} else {
			MenuServer.Start.ClickedCh <- signal
		}
	} else {
		MenuServer.Stop.ClickedCh <- signal
	}

	switch config.Locale {
	case locales.En:
		MenuLanguage.En.ClickedCh <- signal
	case locales.Zh:
		MenuLanguage.Zh.ClickedCh <- signal
	}
}

func resetTooltipLanguage() {
	var src = locales.Get()

	var cState string
	if config.StartServer {
		if config.EnableSwag {
			cState = src.ServerStartSwag.String()
		} else {
			cState = src.ServerStart.String()
		}
	} else {
		cState = src.ServerStop.String()
	}

	var cLocale string
	switch config.Locale {
	case locales.En:
		cLocale = src.LanguageEn.String()
	case locales.Zh:
		cLocale = src.LanguageZh.String()
	}

	systray.SetTooltip(src.Tooltip.String(cState, cLocale))
}
