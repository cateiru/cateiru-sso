package lib

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/storage"
)

type CloudStorageInterface interface {
	Read(ctx context.Context, path string) ([]byte, string, error)
	Write(ctx context.Context, path string, data io.Reader, contentType string) error
	Delete(ctx context.Context, path string) error
}

type CloudStorage struct {
	BucketName string
}

func NewCloudStorage(bucketName string) CloudStorageInterface {

	return &CloudStorage{
		BucketName: bucketName,
	}
}

// クラウドストレージから取得する
func (c *CloudStorage) Read(ctx context.Context, path string) ([]byte, string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, "", err
	}
	defer client.Close()
	rc := client.Bucket(c.BucketName)
	object := rc.Object(path)

	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	contentType := reader.Attrs.ContentType

	return b, contentType, nil
}

// ファイル書き出し
func (c *CloudStorage) Write(ctx context.Context, path string, data io.Reader, contentType string) error {
	if contentType == "" {
		return errors.New("Content-Type is empty")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	rc := client.Bucket(c.BucketName)
	object := rc.Object(path)

	writer := object.NewWriter(ctx)
	defer writer.Close()
	writer.ContentType = contentType

	_, err = io.Copy(writer, data)
	if err != nil {
		return err
	}

	return nil
}

// クラウドストレージから削除する
func (c *CloudStorage) Delete(ctx context.Context, path string) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	rc := client.Bucket(c.BucketName)
	object := rc.Object(path)

	return object.Delete(ctx)
}
