package oss

import (
	"sync"
)

// TODO: 测试 Oss功能

type Oss interface {
	MakeBucket(bucketName string) bool
	IsExistBucket(bucketName string) bool
	UploadObject(bucketName, objectName string, object interface{}) bool
	DownloadObject(bucketName, objectName string) (interface{}, bool)
	DeleteObject(bucketName, objectName string) bool
	IsExistObject(bucketName, objectName string) bool
}

type Client struct {
	Oss
}

var once sync.Once
var client *Client

func init() {
	once.Do(func() {
		// TODO: 根据配置文件初始化不同的Oss
		initMinioClient()
		client = &Client{
			Oss: internalMinioClient,
		}
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
	return client.UploadObject(bucketName, objectName, obj)
}

func DownloadFile(bucketName, objectName string) (interface{}, bool) {
	return client.DownloadObject(bucketName, objectName)
}

func DeleteFile(bucketName, objectName string) bool {
	return client.DeleteObject(bucketName, objectName)
}

func IsExistFile(bucketName, objectName string) bool {
	return client.IsExistObject(bucketName, objectName)
}
