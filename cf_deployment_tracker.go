package cf_deployment_tracker

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

var deploymentTrackerURL = "https://deployment-tracker.mybluemix.net/api/v1/track"

type Repository struct {
	Url string
}

type Package struct {
	Name       string
	Version    string
	Repository Repository
}

type Event struct {
	DateSent           string   `json:date_sent"`
	CodeVersion        string   `json:code_version"`
	RepositoryURL      string   `json:repository_url"`
	ApplicationURL     string   `json:application_name"`
	SpaceID            string   `json:space_id"`
	ApplicationVersion string   `json:application_version"`
	ApplicatonURIs     []string `json:application_uris"`
}

func Track() {
	content, err := ioutil.ReadFile("package.json")
	//exit early if we cant read the file
	if err != nil {
		return
	}

	var info Package
	err = json.Unmarshal(content, &info)
	//exit early if we can't parse the file
	if err != nil {
		return
	}

	request := gorequest.New()
	event := Event{RepositoryURL: info.Repository.Url}
	_, _, errs := request.Post(deploymentTrackerURL).
		Send(event).
		End()

	if errs != nil {
		panic(err)
	}
}