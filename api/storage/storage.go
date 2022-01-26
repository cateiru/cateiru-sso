package storage

import (
	"context"
	"io"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
)

type Storage struct {
	client *storage.Client
	rc     *storage.BucketHandle
}

func NewStorage(ctx context.Context, bucketName string) (*Storage, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	rc := client.Bucket(bucketName)

	return &Storage{
		client: client,
		rc:     rc,
	}, nil
}

func (s *Storage) Close() {
	s.client.Close()
}

// Create Storage object.
// Returns an ObjectHandle that looks like a directory and file to manipulate.
func (s *Storage) Object(dirs []string, fileName string) *storage.ObjectHandle {
	path := append(dirs, fileName)
	return s.rc.Object(strings.Join(path, "/"))
}

// Check if file exists.
// Exist if true, false not.
func (s *Storage) FileExist(ctx context.Context, dirs []string, fileName string) (bool, error) {
	object := s.Object(dirs, fileName)
	_, err := object.Attrs(ctx)
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// Read file.
func (s *Storage) ReadFile(ctx context.Context, dirs []string, fileName string) ([]byte, string, error) {
	object := s.Object(dirs, fileName)
	reader, err := object.NewReader(ctx)
	contentType := reader.ContentType()
	// reader.Attrs.ContentType
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	return b, contentType, nil
}

// Write file
func (s *Storage) WriteFile(ctx context.Context, dirs []string, fileName string, body io.Reader, contentType string) error {
	object := s.Object(dirs, fileName)
	writer := object.NewWriter(ctx)

	writer.ContentType = contentType

	_, err := io.Copy(writer, body)
	if err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

// Delete files
func (s *Storage) Delete(ctx context.Context, dirs []string, fileName string) error {
	object := s.Object(dirs, fileName)
	return object.Delete(ctx)
}

// disable to versioning.
func (s *Storage) DisableVersioning(ctx context.Context) error {
	bucketAttrsToUpdate := storage.BucketAttrsToUpdate{
		VersioningEnabled: false,
	}

	if _, err := s.rc.Update(ctx, bucketAttrsToUpdate); err != nil {
		return err
	}
	return nil
}
