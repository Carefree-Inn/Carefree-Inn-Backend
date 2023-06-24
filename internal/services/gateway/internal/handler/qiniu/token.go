package qiniu

import (
	"gateway/internal"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuHandler struct {
	accessKey string
	secretKey string
	bucket    string
}

func NewQiNiuHandler(accessKey, secretKey, bucket string) *QiNiuHandler {
	return &QiNiuHandler{
		accessKey: accessKey,
		secretKey: secretKey,
		bucket:    bucket,
	}
}

func (q *QiNiuHandler) GetToken(c *gin.Context) {
	policy := storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.accessKey,
		q.secretKey,
	)
	
	internal.Success(c, policy.UploadToken(mac))
}
