package core

import (
	"DJ-Blog/internal/conf"
	"DJ-Blog/pkg/oss"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
)

type IOSSController interface {
	UploadAvatarImg(c *gin.Context)
	GetAvatarImg(c *gin.Context)
}

func UploadAvatarImg(username string, avatarImg *multipart.FileHeader) bool {

	ok := oss.ShouldMakeBucket(username)
	if !ok {
		logrus.Error("创建bucket失败")
		return false
	}

	exist := oss.IsExistFile(username, conf.OssAvatarImgObjectName)
	if exist {
		ok = oss.DeleteFile(username, conf.OssAvatarImgObjectName)
		if !ok {
			logrus.Warn("删除原头像失败")
			return false
		}
	}

	ok = oss.UploadFile(username, conf.OssAvatarImgObjectName, avatarImg)
	if !ok {
		logrus.Error("上传头像失败")
		return false
	}

	return true
}

func GetAvatarImg(username string) ([]byte, bool) {

	file, ok := oss.DownloadFile(username, conf.OssAvatarImgObjectName)
	if !ok {
		logrus.Warn("下载头像失败")
		return nil, false
	}

	fileReader := file.(io.Reader)

	var buf []byte
	var err error
	buf,err = io.ReadAll(fileReader)
	if err != nil{
		logrus.Warn("读取头像失败",err)
	}

	return buf, true
}
