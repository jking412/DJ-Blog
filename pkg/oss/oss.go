package oss

import (
	"sync"
)

type Oss interface {
	MakeBucket(bucketName string) bool
	IsExistBucket(bucketName string) bool
	UploadFile(bucketName, objectName string, object interface{}) bool
	DownloadFile(bucketName, objectName string) (interface{}, bool)
	DeleteFile(bucketName, objectName string) bool
}

type Client struct {
	Oss
}

var once sync.Once
var client *Client

func init() {
	once.Do(func() {
		// TODO: 根据配置文件初始化不同的Oss
		client = &Client{}
	})
}

// ShouldMakeBucket 会判断 bucket 是否存在，不存在则创建
func ShouldMakeBucket(bucketName string) bool {
	if client.IsExistBucket(bucketName) {
		return true
	}
	return client.MakeBucket(bucketName)
}

func UploadFile(bucketName, objectName string, obj interface{}) bool {
	if !ShouldMakeBucket(bucketName) {
		return false
	}
	return client.UploadFile(bucketName, objectName, obj)
}

func DownloadFile(bucketName, objectName string) (interface{}, bool) {
	return client.DownloadFile(bucketName, objectName)
}

func DeleteFile(bucketName, objectName string) bool {
	return client.DeleteFile(bucketName, objectName)
}
