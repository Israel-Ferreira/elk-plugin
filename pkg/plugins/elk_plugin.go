package plugins

import (
	"log"

	"github.com/Kong/go-pdk"
)

type ElkPlugin struct {
	ElasticSearchHost     string `json:"elastic_search_host"`
	ElasticSearchUser     string `json:"elastic_search_user"`
	ElasticSearchPassword string `json:"elastic_search_password"`
	ElasticSearchIndex    string `json:"elastic_search_index"`
}

func (elkp ElkPlugin) Log(kong *pdk.PDK) {
	logInfo, err := kong.Log.Serialize()

	if err != nil {
		panic(err)
	}

	if err := kong.Log.Info("LOG PQP: ", logInfo); err != nil {
		kong.Log.Alert(err)
		log.Println(err)
	}

}

func New() any {
	return &ElkPlugin{}
}
