package locales

const En Locale = "en"

func init() {
	locales[En] = &Locales{
		Title:           "gin-systray-starter-v2",
		Tooltip:         "Server: {1}\nLanguage: {2}",
		Quit:            "Quit",
		Server:          "Server",
		ServerStart:     "Start",
		ServerStartSwag: "Start with Swagger",
		ServerStop:      "Stop",
		Language:        "Language",
		LanguageEn:      "English",
		LanguageZh:      "中文",
	}
}
