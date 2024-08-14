package oss

import (
	"context"
	"framework/nacos"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"gopkg.in/yaml.v2"
	"mime/multipart"
)

type Config struct {
	Oss
}

type Oss struct {
	Region   string `yaml:"region"`
	Bucket   string `yaml:"bucket"`
	Endpoint string `yaml:"endpoint"`
	Ak       string `yaml:"ak"`
	Sk       string `yaml:"sk"`
}

func UploadFile(filename string, body multipart.File) (string, error) {
	var conf Config
	config, err := nacos.GetNacosConfig()
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		return "", err
	}

	provider := credentials.NewStaticCredentialsProvider(conf.Oss.Ak, conf.Oss.Sk)

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(conf.Oss.Region).WithEndpoint(conf.Oss.Endpoint)

	client := oss.NewClient(cfg)

	_, err = client.PutObject(context.TODO(), &oss.PutObjectRequest{
		Bucket: oss.Ptr(conf.Oss.Bucket),
		Key:    oss.Ptr(filename),
		Body:   body,
	})

	if err != nil {
		return "", err
	}

	return "https://" + conf.Oss.Bucket + "." + conf.Oss.Endpoint + "/" + filename, nil
}
