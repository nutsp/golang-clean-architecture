package repositories

import (
	"context"
	"encoding/json"

	httpClient "github.com/nutsp/golang-clean-architecture/pkg/httpclient"
	"go.uber.org/dig"
)

type IMailerRepository interface {
	CheckEmailAvailability(ctx context.Context, email string) (bool, error)
}

type MailerRepository struct {
	client httpClient.IClient
}

type MailerRepositoryDependencies struct {
	dig.In
	Client httpClient.IClient `name:"HttpClient"`
}

func NewMailerRepository(deps MailerRepositoryDependencies) *MailerRepository {
	return &MailerRepository{
		client: deps.Client,
	}
}

func (r *MailerRepository) CheckEmailAvailability(ctx context.Context, email string) (bool, error) {
	req := &httpClient.Request{
		Method: httpClient.MethodGet,
		URL:    "api.example.com/email-availability?email=" + email,
		Header: httpClient.Header{
			httpClient.ContentType: httpClient.ApplicationJson,
		},
	}

	resp, err := r.client.Do(ctx, req)
	if err != nil {
		return false, err
	}

	if !resp.IsSuccess() {
		return false, nil
	}

	var response struct {
		Available bool `json:"available"`
	}
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return false, err
	}

	return response.Available, nil
}
