package locales

import "strings"

type Dictionary struct {
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

type Text string

func (t *Text) String(params ...string) string {
	text := string(*t)
	for _, param := range params {
		text = strings.Replace(text, "%s", param, 1)
	}
	return text
}
