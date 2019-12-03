package apihandlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/Mohammed-Aadil/ds-storage/api/models"
	"github.com/Mohammed-Aadil/ds-storage/pkg/utils"
)

//Upload uploading the attachment
func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handle, err := r.FormFile("file")
	if err != nil {
		utils.HTTPErrorResponse(w, err)
		return
	}
	groupID := uuid.New()
	if err := utils.SaveFile(file, handle); err != nil {
		utils.HTTPErrorResponse(w, err)
		return
	} else {
		go models.CreateAttachment(
			handle.Filename,
			handle.Header.Get("Content-Type"),
			handle.Size,
			groupID,
		)
	}
	utils.HTTPSuccessResponse(
		w,
		map[string]string{
			"msg":     "File uploaded successfully. Please wait for 2 mins to generate.",
			"groupID": groupID.String(),
		},
		200,
	)
}

//ListAll list all the attachment
// not using it for now
// I should not allow everyone to access anyone's file
func ListAll(w http.ResponseWriter, r *http.Request) {
	db, _ := utils.GetDBConnection()
	attachments := &[]models.Attachment{}
	db.Find(&attachments)
	utils.HTTPSuccessResponse(w, attachments, 200)
}

//DownloadFiles it will download all files for same groupID documents
func DownloadFiles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := uuid.Parse(vars["id"]); err != nil {
		utils.HTTPErrorResponse(w, err)
		return
	}
	db, _ := utils.GetDBConnection()
	attachments := []models.Attachment{}
	// getting attachments
	db.Where(fmt.Sprintf("group_id='%s'", vars["id"])).Find(&attachments)
	//converting to iAttachment so we can use FilesToGzip as standard method for all
	var iAttachments []utils.IAttachment
	for _, v := range attachments {
		iAttachments = append(iAttachments, v)
	}
	if path, err := utils.FilesToGzip(iAttachments); err != nil {
		utils.HTTPErrorResponse(w, err)
	} else {
		// downloading files
		utils.HTTPFileSuccessResponse(w, r, path)
	}
}

//ListOne list one attachment
func ListOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := uuid.Parse(vars["id"]); err != nil {
		utils.HTTPErrorResponse(w, err)
		return
	}
	db, _ := utils.GetDBConnection()
	attachments := &[]models.Attachment{}
	db.Where(fmt.Sprintf("group_id='%s'", vars["id"])).Find(&attachments)
	utils.HTTPSuccessResponse(w, attachments, 200)
}
