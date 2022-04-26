package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"content-service-v3/services/content-service/internal/usecase/banner/formatter"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func (s *serviceBanner) CreateBanner(input dto.CreateBannerInput, file *multipart.FileHeader) (formatter.CreateBannerResponseFormatter, string, error) {
	// Mapper From DTO to Entity
	mappedBanner := entity.Create_BannerDTO_To_BannerEntity(input)

	// Buat Path
	// Inisialisasi Path
	path := fmt.Sprintf("content-service-v3/services/content-service/asset")
	
	// Cek Apakah Ada Directory nya? 
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Klo ngga ada buat directory nya
        if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
    }
	
	// Format File := <idUser>-<nameFile>.<ekstensi>
	idUser := strconv.Itoa(mappedBanner.CreatedById)
	nameFile := strings.Split(file.Filename, ".")[0]
	ekstensi := strings.Split(file.Filename, ".")[1]

	pathForSave := fmt.Sprintf("services/content-service/asset/%s-%s.%s",idUser, nameFile, ekstensi)

	// set mappedBanner file name 
	mappedBanner.FileName = fmt.Sprintf("%s-%s.%s",idUser,nameFile,ekstensi)
	
	// Teruskan Ke Repository
	newBanner, err := s.repository.SaveBanner(mappedBanner)

	// Formatter
	formattedBanner := formatter.FormatCreateBannerResponse(newBanner)
	
	// Klo Error
	if err != nil {
		return formattedBanner, "", err
	}

	// Tidak Error
	return formattedBanner, pathForSave, nil
}