package locales

import (
	"fmt"
	"strings"
)

type Locales struct {
	Title           Text
	Tooltip         Text
	Quit            Text
	Server          Text
	ServerStart     Text
	ServerStartSwag Text
	ServerStop      Text
	Language        Text
	LanguageEn      Text
	LanguageZh      Text
}

type Locale string

type Text string

func (t *Text) String(params ...string) string {
	text := string(*t)
	for i, param := range params {
		text = strings.Replace(text, fmt.Sprintf("{%d}", i+1), param, 1)
	}
	return text
}

var (
	locale  Locale              = En
	locales map[Locale]*Locales = make(map[Locale]*Locales, 2)
)

func Get() Locales {
	if strings, ok := locales[locale]; ok {
		return *strings
	}
	return Locales{}
}

func Set(lang Locale) (ok bool) {
	if _, ok = locales[lang]; ok {
		locale = lang
	}
	return
}
