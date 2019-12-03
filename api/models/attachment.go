package models

import (
	"encoding/json"
	"log"

	"github.com/Mohammed-Aadil/ds-storage/api/services"
	"github.com/Mohammed-Aadil/ds-storage/config"
	"github.com/Mohammed-Aadil/ds-storage/pkg/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Attachment attachment model struct
type Attachment struct {
	gorm.Model
	GroupID uuid.UUID `json:"groupID"`
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Num     int       `gorm:"AUTO_INCREMENT"`
	Meta    postgres.Jsonb
}

//GetPath IAttachment interface method
func (a Attachment) GetPath() string {
	return a.Path
}

//GetName IAttachment interface method
func (a Attachment) GetName() string {
	return a.Name
}

//Init Attachment model Init
func (a *Attachment) Init() {
	db, err := utils.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(a)
}

//AfterCreate attachment after save hook
func (a *Attachment) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(a).Update("URL", config.StorageBaseURL+string(a.ID))
	return
}

//CreateAttachment create the attachment entry in db.
//it should be call after uploading file to server
func CreateAttachment(fileName string, mimeType string, size int64, groupID uuid.UUID) {
	db, _ := utils.GetDBConnection()
	ch := make(chan *utils.HTTPChannelResponse)
	go services.DocToPngServiceUpload(fileName, ch)
	response := <-ch
	if response.Err == nil {
		defer response.Body.Close()
		attachments := &[]Attachment{}
		json.NewDecoder(response.Body).Decode(attachments)
		meta, err := json.Marshal(map[string]interface{}{
			"Size":     size,
			"MimeType": mimeType,
		})
		if err != nil {
			log.Println(err)
		}
		for _, attachment := range *attachments {
			attachment.Meta = postgres.Jsonb{RawMessage: meta}
			attachment.GroupID = groupID
			db.Create(&attachment)
		}
	} else {
		log.Println(response.URL)
		log.Println(response.FileName)
		log.Println(response.Err)
	}
}
