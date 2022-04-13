package helpers

import (
	"context"
	"time"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// For reference you can use --> https://www.topcoder.com/thrive/articles/uploading-files-using-golang-gin-and-cloudinary
// --> https://dev.to/hackmamba/robust-media-upload-with-golang-and-cloudinary-gin-gonic-version-54ii

// Helper Function to uploading an image
func CloudinaryImageUploadHelper(input interface{}) (string, error) {
	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Config Instance
	configuration := config.GetConfig()
	// Cloudinary Instance
	cloudinaryInstance := config.GetCloudinaryInstance()

	//upload file
	// uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{})
	uploadResponse, err := cloudinaryInstance.Upload.Upload(ctx, input, uploader.UploadParams{Folder: configuration.Cloudinary.UploadFolder})
	if err != nil {
		return "", err
	}

	return uploadResponse.SecureURL, nil
}
