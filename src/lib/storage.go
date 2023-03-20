package lib

type CloudStorageInterface interface {
	Get(path string) ([]byte, error)
	Insert(path string, data []byte) error
	Update(path string, data []byte) error
	Delete(path string) error
}

type CloudStorage struct {
}

func NewCloudStorage() CloudStorageInterface {
	return &CloudStorage{}
}

// クラウドストレージから取得する
func (c *CloudStorage) Get(path string) ([]byte, error) {
	return []byte{}, nil
}

func (c *CloudStorage) Insert(path string, data []byte) error {
	return nil
}

func (c *CloudStorage) Update(path string, data []byte) error {
	return nil
}

// クラウドストレージから削除する
func (c *CloudStorage) Delete(path string) error {
	return nil
}
