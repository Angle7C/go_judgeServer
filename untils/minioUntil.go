package untils

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyId     string `yaml:"AccessKeyId"`
	SecretAccessKey string `yaml:"SecretAccessKey"`
	Secure          string `yaml:"Secure"`
	TimeZone        string `yaml:"TimeZone"`
}

var (
	client *minio.Client
	config *MinioConfig
	err    error
	ctx    context.Context = context.Background()
)

func (config MinioConfig) Init() {
	client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyId, config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("MinIO链接错误")
	} else {
		log.Println("MinIO链接成功")
	}
}

func CreateBucket(buckName string) bool {
	err := client.MakeBucket(ctx, buckName, minio.MakeBucketOptions{Region: config.TimeZone})
	if err != nil {
		log.Println("创建Bucket失败")
		return false
	} else {
		return true
	}
}
func UploadImage(bucketName, imageName string, reader io.Reader, size int64) bool {
	_, err := client.PutObject(ctx, bucketName, imageName, reader, size, minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		log.Fatalf("上传文件%v失败", imageName)
		return false
	} else {
		return true
	}

}
func DownloadFile(bucketName, objectName string) io.Reader {
	object, err := client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal("下载文件失败")
	}
	return object
}
