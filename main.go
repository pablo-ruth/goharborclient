package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/pablo-ruth/goharborclient/client/products"
	"github.com/pablo-ruth/goharborclient/models"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/pablo-ruth/goharborclient/client"
)

func main() {
	client := apiclient.New(httptransport.New(os.Getenv("HARBOR_ADDR"), "api", nil), strfmt.Default)
	basicAuth := httptransport.BasicAuth(os.Getenv("HARBOR_USER"), os.Getenv("HARBOR_PASSWORD"))

	projects, err := client.Products.GetProjects(products.NewGetProjectsParams(), basicAuth)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, project := range projects.Payload {
		fmt.Printf("%d %s\n", project.ProjectID, project.Name)

		repositories, err := getRepositories(client, basicAuth, project.ProjectID)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repository := range repositories {
			fmt.Printf("  %d %s\n", repository.ID, repository.Name)

			tags, err := getTags(client, basicAuth, repository.Name)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, tag := range tags {
				created, err := time.Parse(time.RFC3339Nano, tag.Created)
				if err != nil {
					fmt.Println(err)
					return
				}

				manifest, err := getTagManifest(client, basicAuth, repository.Name, tag.Name)
				if err != nil {
					fmt.Println(err)
					return
				}

				labels, err := getLabels(manifest.Config)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("    %s %v %v %v\n", tag.Name, created, len(tag.Labels), labels)
			}
		}
	}
}

func getRepositories(client *apiclient.Harbor, auth runtime.ClientAuthInfoWriter, projectID int32) ([]*models.Repository, error) {

	params := &products.GetRepositoriesParams{ProjectID: projectID}
	params.SetTimeout(10 * time.Second)

	repositories, err := client.Products.GetRepositories(params, auth)
	if err != nil {
		return nil, err
	}

	return repositories.Payload, nil
}

func getTags(client *apiclient.Harbor, auth runtime.ClientAuthInfoWriter, repoName string) ([]*models.DetailedTag, error) {

	params := &products.GetRepositoriesRepoNameTagsParams{RepoName: repoName}
	params.SetTimeout(10 * time.Second)

	tags, err := client.Products.GetRepositoriesRepoNameTags(params, auth)
	if err != nil {
		return nil, err
	}

	return tags.Payload, nil
}

func getTagManifest(client *apiclient.Harbor, auth runtime.ClientAuthInfoWriter, repoName, tag string) (*models.Manifest, error) {

	params := &products.GetRepositoriesRepoNameTagsTagManifestParams{RepoName: repoName, Tag: tag}
	params.SetTimeout(10 * time.Second)

	manifest, err := client.Products.GetRepositoriesRepoNameTagsTagManifest(params, auth)
	if err != nil {
		return nil, err
	}

	return manifest.Payload, nil
}

func deleteTag(client *apiclient.Harbor, auth runtime.ClientAuthInfoWriter, repoName, tag string) error {

	params := &products.DeleteRepositoriesRepoNameTagsTagParams{RepoName: repoName, Tag: tag}
	params.SetTimeout(10 * time.Second)

	_, err := client.Products.DeleteRepositoriesRepoNameTagsTag(params, auth)

	return err
}

func getLabels(jsonData string) (map[string]string, error) {

	config := ImageConfig{}
	if err := json.Unmarshal([]byte(jsonData), &config); err != nil {
		return map[string]string{}, err
	}

	return config.Config.Labels, nil
}

type ImageConfig struct {
	Config struct {
		Labels map[string]string `json:"Labels"`
	} `json:"config"`
}
