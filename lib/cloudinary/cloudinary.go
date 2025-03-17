package cloudinary

import (
	"log"
	"my-project-be/config"

	"github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryClient(cfg *config.AppConfig) (*cloudinary.Cloudinary, error) {
	cld ,err := cloudinary.NewFromParams(cfg.CLOUDINARY_CLOUD_NAME, cfg.CLOUDINARY_API_KEY, cfg.CLOUDINARY_API_SECRET)
	if err != nil {
		log.Println("gagal membuat cloudinary client", err)
		return nil, err
	}

	return cld, nil
}