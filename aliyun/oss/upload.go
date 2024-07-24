package oss

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/spf13/viper"
)

func UploadFile(filename, path string) (string, error) {
	region := viper.GetString("oss.region")
	bucket := viper.GetString("oss.bucket")
	endpoint := viper.GetString("oss.endpoint")
	ak := viper.GetString("oss.ak")
	sk := viper.GetString("oss.sk")
	provider := credentials.NewStaticCredentialsProvider(ak, sk)

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(region).WithEndpoint(endpoint)

	client := oss.NewClient(cfg)

	_, err := client.PutObjectFromFile(context.TODO(), &oss.PutObjectRequest{
		Bucket: oss.Ptr(bucket),
		Key:    oss.Ptr(filename),
	}, path)

	if err != nil {
		return "", err
	}

	return "https://" + bucket + "." + endpoint + "/" + filename, nil
}
