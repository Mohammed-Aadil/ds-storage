package initconfig

import (
	"log"
	"net/http"

	"github.com/Mohammed-Aadil/ds-storage/api/models"
	"github.com/Mohammed-Aadil/ds-storage/init/routers"
)

//ModelInit Function to init models config
func modelInit() {
	log.Println("Init Attachment model")
	attachment := models.Attachment{}
	attachment.Init()
}

//RouterInit Function to init router config
func routerInit() http.Handler {
	log.Println("Initiating routes")
	return routers.Init()
}

//Init init app
func Init() http.Handler {
	modelInit()
	return routerInit()
}
