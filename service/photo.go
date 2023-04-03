package service

import (
	"my-admin/global"
	"my-admin/model"
)

func CreatePhoto(photo model.Photo) error {
	return insertPhoto(photo)
}

func DeletePhoto(photo model.Photo) error {
	return deletePhoto(photo)
}

func insertPhoto(phote model.Photo) error {
	return global.DBClient.Create(&phote).Error
}

func deletePhoto(photo model.Photo) error {
	return global.DBClient.Delete(photo).Error
}

func isPhotoOutOfCapacity() bool {
	return false
}

func ListPhoto() []model.Photo {
	var photos []model.Photo
	global.DBClient.Find(&photos)
	return photos
}
