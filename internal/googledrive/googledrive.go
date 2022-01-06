package googledrive

import (
	"context"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func ShareFile(srv *drive.Service, fileId string, email string) error {
	_, err := srv.Permissions.Create(fileId, &drive.Permission{
		EmailAddress: email,
		Role:         "writer",
		Type:         "user",
	}).Do()
	return err
}

func GetDriveService() (*drive.Service, error) {
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile("./credentials.json"))
	return srv, err
}
