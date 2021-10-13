package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/yimikao/api-assignment/api/models"
)

func GetAllProfiles(db *gorm.DB, w http.ResponseWriter, r *http.Request) []models.Profile {
	profiles := []models.Profile{}
	db.Find(&profiles)
	return profiles
}

func CreateProfile(p models.Profile, db *gorm.DB, w http.ResponseWriter, r *http.Request) (models.Profile, int, error) {

	if err := db.Save(&p).Error; err != nil {
		return p, http.StatusInternalServerError, err
	}
	return p, http.StatusCreated, nil
}

func GetPausedProfiles(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*[]models.Profile, int, error) {

	query := r.URL.Query().Get("status")

	profile, status, err := getProfiles(db, "status", query, w, r)
	return profile, status, err
}

func PauseProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*models.Profile, int, error) {
	vars := mux.Vars(r)

	id := vars["id"]

	profile, status, err := getProfile(db, "id", id, w, r)

	if err != nil {
		return profile, status, err
	}

	profile.Status = "PAUSED"

	if err := db.Save(&profile).Error; err != nil {
		return profile, http.StatusInternalServerError, err
	}
	return profile, http.StatusOK, nil
}
func UnPauseProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*models.Profile, int, error) {
	vars := mux.Vars(r)

	id := vars["id"]

	profile, status, err := getProfile(db, "id", id, w, r)

	if err != nil {
		return profile, status, err
	}

	profile.Status = "ACTIVE"

	if err := db.Save(&profile).Error; err != nil {
		return profile, http.StatusInternalServerError, err
	}
	return profile, http.StatusOK, nil
}

func DeleteProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*models.Profile, int, error) {
	vars := mux.Vars(r)

	id := vars["id"]
	profile, status, err := getProfile(db, "id", id, w, r)

	if err != nil {
		return profile, status, err
	}

	if err := db.Delete(&profile).Error; err != nil {
		return profile, http.StatusInternalServerError, err
	}
	return profile, http.StatusNoContent, err
}

func getProfile(db *gorm.DB, search string, finder string, w http.ResponseWriter, r *http.Request) (*models.Profile, int, error) {
	profile := models.Profile{}
	if err := db.Where(search+" = ?", finder).Find(&profile).Error; err != nil {
		return &profile, http.StatusNotFound, err
	}
	return &profile, http.StatusOK, nil
}

func getProfiles(db *gorm.DB, search string, finder string, w http.ResponseWriter, r *http.Request) (*[]models.Profile, int, error) {
	profiles := []models.Profile{}

	if err := db.Where(search+" = ?", finder).Find(&profiles).Error; err != nil {
		return &profiles, http.StatusNotFound, err
	}
	return &profiles, http.StatusOK, nil
}
