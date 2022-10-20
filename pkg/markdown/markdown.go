package markdown

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var GoldMark goldmark.Markdown

func InitMarkDownParser() {
	GoldMark = goldmark.New(
		goldmark.WithExtensions(extension.Table, extension.GFM),
	)
	logrus.Info("Markdown parser initialized")
}

func ParseContent(content string) string {
	var buffer bytes.Buffer
	GoldMark.Convert([]byte(content), &buffer)
	return buffer.String()
}
