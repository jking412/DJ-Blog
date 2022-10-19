package main

import (
	"DJ-Blog/boot"
	"DJ-Blog/pkg/markdown"
	"bufio"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Initialize()
	m.Run()
}

func TestMarkdown(t *testing.T) {
	file, _ := os.OpenFile("test.md", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	reader := bufio.NewReader(file)
	var content string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		content += string(line) + "\n"
	}
	parseContent := markdown.ParseContent(content)
	saveFile, _ := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer saveFile.Close()
	saveFile.WriteString(parseContent)
}
