package cloudinary

import (
	"log"
	"my-project-be/config"

	"github.com/cloudinary/cloudinary-go/v2"
)

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file: ", err)
// 	}

// 	cld, err := cloudinary.New()
// 	if err != nil {
// 		log.Fatal("failed to create cloudinary client: ", err)
// 	}

// 	img, err := cld.Image("cld-sample")
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	url, err := img.String()
// 	if err != nil {
// 		fmt.Println("error")
// 	}else{
// 		fmt.Println("image URL:", url)
// 	}
// }

func GetCloudinaryClient(cfg *config.AppConfig) (*cloudinary.Cloudinary, error) {
	cld ,err := cloudinary.NewFromParams(cfg.CLOUDINARY_CLOUD_NAME, cfg.CLOUDINARY_API_KEY, cfg.CLOUDINARY_API_SECRET)
	if err != nil {
		log.Println("gagal membuat cloudinary client", err)
		return nil, err
	}

	return cld, nil
}