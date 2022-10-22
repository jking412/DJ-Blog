package main

import (
	"DJ-Blog/boot"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Initialize()
	m.Run()
}

//func TestTemp(t *testing.T) {
//	query, err := search.ZincCli.Query("posts", "test测试")
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Println(query)
//}

//func TestTemp(t *testing.T) {
//	post := &model.Post{}
//	database.DB.First(&post)
//	dataMap := map[string]interface{}{
//		"id":         post.Id,
//		"title":      "我也不知道中文测试自己，来吧，ok",
//		"tag":        post.Tag,
//		"content":    post.Content,
//		"created_at": post.CreatedAt,
//		"updated_at": post.UpdatedAt,
//		"likes":      post.Likes,
//		"stared":     post.Stared,
//		"user_id":    post.UserId,
//	}
//	search.ZincCli.PutDoc("posts", 1, dataMap)
//}

//func TestMarkdown(t *testing.T) {
//	file, _ := os.OpenFile("test.md", os.O_RDONLY, os.ModePerm)
//	defer file.Close()
//	reader := bufio.NewReader(file)
//	var content string
//	for {
//		line, _, err := reader.ReadLine()
//		if err != nil {
//			break
//		}
//		content += string(line) + "\n"
//	}
//	parseContent := markdown.ParseContent(content)
//	saveFile, _ := os.OpenFile("test.html", os.O_CREATE|os.O_WRONLY, os.ModePerm)
//	defer saveFile.Close()
//	saveFile.WriteString(parseContent)
//}
