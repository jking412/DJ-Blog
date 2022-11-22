package oss

import (
	"DJ-Blog/pkg/config"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
)

type MinioClient struct {
	client *minio.Client
	ctx    context.Context
}

var (
	internalMinioClient *MinioClient
	endpoint            = config.LoadString("oss.endpoint") // ip:port
	accessKeyID         = config.LoadString("oss.accessKeyId")
	accessKeySecret     = config.LoadString("oss.accessKeySecret")
)

func initMinioClient() {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, accessKeySecret, ""),

		Secure: false, // useSSL
	})
	if err != nil {
		logrus.Error("minio client init error: ", err)
	}

	internalMinioClient = &MinioClient{
		client: minioClient,
		ctx:    context.Background(),
	}
	logrus.Info("minio client init success")
}

func (m *MinioClient) MakeBucket(bucketName string) bool {
	err := m.client.MakeBucket(m.ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		logrus.Error("minio make bucket error: ", err)
		return false
	}
	return true
}

func (m *MinioClient) IsExistBucket(bucketName string) bool {
	exist, err := m.client.BucketExists(m.ctx, bucketName)
	if err != nil {
		logrus.Error("minio check bucket exist error: ", err)
		return false
	}
	return exist
}

func (m *MinioClient) IsExistObject(bucketName, objectName string) bool {
	_, err := m.client.StatObject(m.ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		logrus.Error("minio check file exist error: ", err)
		return false
	}
	return true
}

func (m *MinioClient) UploadObject(bucketName, objectName string, obj interface{}) bool {

	file := obj.(multipart.File)

	_, err := m.client.PutObject(m.ctx, bucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logrus.Error("minio upload file error: ", err)
		return false
	}
	return true
}

func (m *MinioClient) DownloadObject(bucketName, objectName string) (interface{}, bool) {
	ok := m.IsExistObject(bucketName, objectName)
	if !ok {
		return nil, false
	}
	object, err := m.client.GetObject(m.ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		logrus.Error("minio download file error: ", err)
		return nil, false
	}
	bytes := make([]byte, 0)
	var n = 0
	for {
		tmpBytes := make([]byte, 1024)
		nbytes, err := object.Read(tmpBytes)
		n += nbytes
		bytes = append(bytes, tmpBytes[:nbytes]...)
		if err == io.EOF {
			break
		}
	}
	return bytes, true
}

func (m *MinioClient) DeleteObject(bucketName, objectName string) bool {
	err := m.client.RemoveObject(m.ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		logrus.Error("minio delete file error: ", err)
		return false
	}
	return true
}
