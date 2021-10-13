package handlers

import (
	"net/http"

	"github.com/yimikao/api-assignment/api/config"
	"github.com/yimikao/api-assignment/api/models"
	"github.com/yimikao/api-assignment/api/services"
)

// Handlers to manage Profile Data
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles := services.GetAllProfiles(config.DB, w, r)
	RespondJSON(w, http.StatusOK, profiles)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	//This is for url-form-encoded FORM
	if err := r.ParseForm(); err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	Name := r.FormValue("name")
	DOB := r.FormValue("dateofbirth")
	checkNameLength(Name, w)
	nameToShort := checkNameLength(Name, w)
	if nameToShort {
		RespondError(w, http.StatusBadRequest, "name can't be less than three")
		return
	}
	profile := models.Profile{Name: Name, DateOfBirth: DOB}

	//This is for JSON
	// profile := models.Profile{}
	// decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&profile); err != nil {
	// 	RespondError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// nameToShort := checkNameLength(profile.Name, w)
	// if nameToShort {
	// 	RespondError(w, http.StatusBadRequest, "name can't be less than three")
	// 	return
	// }

	profile, status, err := services.CreateProfile(profile, config.DB, w, r)
	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, profile)
}

func GetPausedProfiles(w http.ResponseWriter, r *http.Request) {
	profile, status, err := services.GetPausedProfiles(config.DB, w, r)

	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, profile)

}

func PauseProfile(w http.ResponseWriter, r *http.Request) {
	profile, status, err := services.PauseProfile(config.DB, w, r)

	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, profile)

}

func UnPauseProfile(w http.ResponseWriter, r *http.Request) {
	profile, status, err := services.UnPauseProfile(config.DB, w, r)

	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, profile)

}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {

	_, status, err := services.DeleteProfile(config.DB, w, r)
	if err != nil {
		RespondError(w, status, err.Error())
		return
	}
	RespondJSON(w, status, nil)
}
