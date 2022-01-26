package storage

import (
	"context"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Storage struct {
	rc *storage.BucketHandle
}

func NewStorage(ctx context.Context, bucketName string) (*Storage, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	rc := client.Bucket(bucketName)

	return &Storage{
		rc: rc,
	}, nil
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
func (s *Storage) ReadFile(ctx context.Context, dirs []string, fileName string) ([]byte, error) {
	object := s.Object(dirs, fileName)
	reader, err := object.NewReader(ctx)
	// reader.Attrs.ContentType
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Write file
func (s *Storage) WriteFile(ctx context.Context, dirs []string, fileName string, body []byte) error {
	object := s.Object(dirs, fileName)
	writer := object.NewWriter(ctx)

	_, err := writer.Write(body)
	if err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

// Delete files
func (s *Storage) Delete(ctx context.Context, prefix string) error {
	objects := s.rc.Objects(ctx, &storage.Query{
		Prefix: prefix,
	})

	for {
		attrs, err := objects.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := s.rc.Object(attrs.Name).Delete(ctx); err != nil {
			return err
		}
	}
	return nil
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
