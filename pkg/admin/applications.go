package admin

import (
	"fmt"
	"hexa/pkg/web_support"
	"log"
	"net/http"
)

type Application struct {
	ID            string
	IntegrationId string
	ObjectId      string
	Name          string
	Description   string
}

type applicationsHandler struct {
	orchestratorUrl string
	client          Client
}

func NewApplicationsHandler(orchestratorUrl string, client Client) applicationsHandler {
	return applicationsHandler{orchestratorUrl, client}
}

func (p applicationsHandler) List(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%v/applications", p.orchestratorUrl)
	applications, err := p.client.Applications(url)
	if err != nil {
		log.Println(err)
	}
	model := web_support.Model{Map: map[string]interface{}{"resource": "applications", "applications": applications}}
	_ = web_support.ModelAndView(w, "applications", model)
}
