package file

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/app"
	"github.com/shanedoc/gohub/pkg/auth"
	"github.com/shanedoc/gohub/pkg/helpers"
)

//package file辅助函数

//将数据存入文件
func Put(data []byte, to string) error {
	if err := os.WriteFile(to, data, 0644); err != nil {
		return err
	}
	return nil
}

//判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	//图片保存目录:没有创建
	publiPath := "public"
	//根据输入格式生成字符串并返回 public/uploads/avatar/date()/uid
	dirPath := fmt.Sprintf("/uploads/avatar/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	//创建文件夹
	os.MkdirAll(dirPath, 0755)
	//保存文件:文件重命名
	fileName := randomNameFromUploadFile(file)
	tPath := publiPath + dirPath + fileName
	if err := c.SaveUploadedFile(file, tPath); err != nil {
		return avatar, err
	}
	//裁剪图片
	img, err := imaging.Open(tPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}
	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarName := randomNameFromUploadFile(file)
	resizeAvatarPath := publiPath + dirPath + resizeAvatarName
	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}
	//删除原文件
	if err = os.Remove(tPath); err != nil {
		return avatar, err
	}

	return dirPath + resizeAvatarName, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
