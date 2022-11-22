package oss

import (
	"fmt"
	"os"
	"testing"
)

func TestShouldMakeBucket(t *testing.T) {
	ok := ShouldMakeBucket("test")
	if !ok {
		t.Error("创建bucket失败")
	}
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("test.txt")
	ok := UploadFile("test", "test.txt", file)
	if !ok {
		t.Error("上传文件失败")
	}
}

func TestIsExistFile(t *testing.T) {
	ok := IsExistFile("test", "test.txt")
	if !ok {
		t.Log("文件不存在")
	} else {
		t.Log("文件存在")
	}
}

func TestDownloadFile(t *testing.T) {
	fileBytes, ok := DownloadFile("test", "test.txt")
	if !ok {
		t.Error("下载文件失败")
	}
	fmt.Println(string(fileBytes.([]byte)))
}

func TestDeleteFile(t *testing.T) {
	ok := DeleteFile("test", "test.txt")
	if !ok {
		t.Error("删除文件失败")
	}
}
