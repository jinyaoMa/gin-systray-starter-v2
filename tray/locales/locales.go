package locales

import (
	"fmt"
	"strings"
)

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
	for i, param := range params {
		text = strings.Replace(text, fmt.Sprintf("{%d}", i+1), param, 1)
	}
	return text
}
