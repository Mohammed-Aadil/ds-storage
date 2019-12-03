package services

import (
	"net/http"

	"github.com/Mohammed-Aadil/ds-storage/config"
	"github.com/Mohammed-Aadil/ds-storage/pkg/utils"
)

//DocToPngServiceUpload upload file to doc_to_png service and get png format
func DocToPngServiceUpload(fileName string, ch chan *utils.HTTPChannelResponse) {
	response, err := http.Get(config.DocToPngServiceURL + fileName + "/")
	if response != nil && response.StatusCode == http.StatusOK {
		ch <- &utils.HTTPChannelResponse{URL: config.DocToPngServiceURL, FileName: fileName, Err: err, Body: response.Body}
	}
	ch <- &utils.HTTPChannelResponse{URL: config.DocToPngServiceURL, FileName: fileName, Err: err}
}
