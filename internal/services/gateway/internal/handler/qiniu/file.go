package qiniu

import (
	"gateway/internal"
	"gateway/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/net/context"
	"mime/multipart"
	"net/http"
)

type QiNiuHandler struct {
	accessKey string
	secretKey string
	bucket    string
	prefix    string
}

func NewQiNiuHandler(accessKey, secretKey, bucket, prefix string) *QiNiuHandler {
	return &QiNiuHandler{
		accessKey: accessKey,
		secretKey: secretKey,
		bucket:    bucket,
		prefix:    prefix,
	}
}

func (q *QiNiuHandler) uploadFile(file multipart.File, size int64) (string, error) {
	policy := storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.accessKey,
		q.secretKey,
	)
	
	token := policy.UploadToken(mac)
	extra := storage.PutExtra{}
	
	uploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false, // 使用自定义域名CDN加速
		UseHTTPS:      false,
	})
	ret := storage.PutRet{}
	
	if err := uploader.PutWithoutKey(context.TODO(), &ret, token, file, size, &extra); err != nil {
		return "", err
	}
	
	return q.prefix + ret.Key, nil
}

func (q *QiNiuHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			internal.Error(c, errno.ErrNoFile)
			return
		}
		internal.ServerError(c, err.Error())
		return
	}
	
	url := ""
	if header != nil {
		url, err = q.uploadFile(file, header.Size)
		if err != nil {
			internal.Error(c, errno.ParamDataError)
			return
		}
	} else {
		internal.Error(c, errno.HeaderNotExist)
		return
	}
	
	internal.Success(c, url)
}
