package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var instance *cloudinary.Cloudinary

func InitCloudinary() bool {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	instance = cld

	if err != nil {
		fmt.Println("Error initializing Cloudinary:", err)
		return false
	}

	instance.Config.URL.Secure = true
	ctx := context.Background()

	return instance != nil && ctx != nil
}

func UploadImage(ctx context.Context, file any, publicID string, folder string) (string, error) {
	uploadParams := uploader.UploadParams{
		PublicID:       publicID,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
		Transformation: "w_250,h_250,c_fill,g_face",
		Folder:         folder,
	}

	resp, err := instance.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %w", err)
	}

	return resp.SecureURL, nil
}
