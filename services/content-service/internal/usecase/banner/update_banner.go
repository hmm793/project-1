package usecase

import (
	"content-service-v3/services/content-service/domain/dto"
	"content-service-v3/services/content-service/domain/entity"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func (s *serviceBanner) UpdateBanner(id int, input dto.UpdateBannerInput, file *multipart.FileHeader) (entity.BannerEntity, string, error) {
	// Mapper From DTO to Entity
	mappedBanner := entity.Update_BannerDTO_To_BannerEntity(input)

	// Cek Apakah Ada Banner Dengan ID = id
	currBanner, err := s.repository.FindBannerById(id)

	if err != nil {
		return mappedBanner, "", err
	}

	if currBanner.ID == 0 {
		return mappedBanner, "", errors.New("Banner Not Found")
	}

	// Buat Path
	// Inisialisasi Path
	path := `services/content-service/asset`
	
	// Cek Apakah Ada Directory nya? 
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Klo ngga ada buat directory nya
        if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
    }
	
	// Format File := <idUser>-<nameFile>.<ekstensi>
	idUser := strconv.Itoa(currBanner.CreatedById)
	nameFile := strings.Split(file.Filename, ".")[0]
	ekstensi := strings.Split(file.Filename, ".")[1]

	pathForSave := fmt.Sprintf("services/content-service/asset/%s-%s.%s",idUser, nameFile, ekstensi)

	// Periksa Apakah sudah ada yang sama sebelumnya Jika iya maka tidak perlu di buat
	_, err = os.Stat(pathForSave)

	if !os.IsNotExist(err) {
		// Klo sudah ada
		pathForSave = ""
		
		// Update Banner
		currBanner.Link = mappedBanner.Link
		currBanner.Order = mappedBanner.Order
		currBanner.Status = mappedBanner.Status
		currBanner.UpdatedById = mappedBanner.UpdatedById
		currBanner.UpdatedByName = mappedBanner.UpdatedByName
		currBanner.BannerCategoryID = mappedBanner.BannerCategoryID

		// Teruskan Ke Repository
		updatedBanner, err := s.repository.UpdateBanner(currBanner)

		if err != nil {
			return updatedBanner, pathForSave, err
		}
		return updatedBanner, pathForSave, nil
	} else {
		// Klo belum ada
		// Hapus Yang lama 
		pathForDelete := currBanner.FileName

		err := os.Remove(fmt.Sprintf("services/content-service/asset/%s",pathForDelete))
		
		if err != nil {
			return mappedBanner, "", errors.New("Gagal Menghapus File Sebelumnya")
		}
		
		// set mappedBanner file name 
		mappedBanner.FileName = fmt.Sprintf("%s-%s.%s",idUser,nameFile,ekstensi)
		
		// Update Banner
		currBanner.Link = mappedBanner.Link
		currBanner.Order = mappedBanner.Order
		currBanner.Status = mappedBanner.Status
		currBanner.UpdatedById = mappedBanner.UpdatedById
		currBanner.UpdatedByName = mappedBanner.UpdatedByName
		currBanner.BannerCategoryID = mappedBanner.BannerCategoryID
		currBanner.FileName = mappedBanner.FileName

		// Teruskan Ke Repository
		updatedBanner, err := s.repository.UpdateBanner(currBanner)

		if err != nil {
			return updatedBanner, pathForSave, err
		}
		return updatedBanner, pathForSave, nil
	}
}