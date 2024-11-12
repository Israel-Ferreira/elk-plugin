package plugins

import (
	"github.com/Kong/go-pdk"
)

type ElkPlugin struct {
	ElasticSearchHost     string `json:"elastic_search_host"`
	ElasticSearchUser     string `json:"elastic_search_user"`
	ElasticSearchPassword string `json:"elastic_search_password"`
	ElasticSearchIndex    string `json:"elastic_search_index"`
}

func (elkp ElkPlugin) Response(kong *pdk.PDK) {
	logInfo, err := kong.Log.Serialize()

	if err != nil {
		panic(err)
	}

	kong.Log.Info("LOG PQP: ", logInfo)

}

func New() any {
	return &ElkPlugin{}
}
