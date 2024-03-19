package supabase

import (
	"mime/multipart"
	"os"

	// "os"

	// supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	storage_go "github.com/supabase-community/storage-go"
)

type Interface interface {
	Upload(file *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type SupabaseStorage struct {
	client *storage_go.Client
}

func Init() Interface {
	storageClient := storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), nil)

	return &SupabaseStorage{
		client: storageClient,
	}
}

func (s *SupabaseStorage) Upload(file *multipart.FileHeader) (string, error) {
	buff, err := file.Open()
	if err != nil {
		return "", err
	}
	defer buff.Close()

	_, err = s.client.UploadFile(os.Getenv("SUPABASE_BUCKET"), file.Filename, buff)
	if err != nil {
		return "", err
	}
	link := s.client.GetPublicUrl(os.Getenv("SUPABASE_BUCKET"), file.Filename).SignedURL
	return link, nil
}

func (s *SupabaseStorage) Delete(link string) error {
	_, err := s.client.RemoveFile(os.Getenv("SUPABASE_BUCKET"), []string{link})
	if err != nil {
		return err
	}
	return nil
}
