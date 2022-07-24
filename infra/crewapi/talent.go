package crewapi

import (
	"fmt"
	"github.com/Ventilateur/crew-interview/domain/seed"
	"github.com/goccy/go-json"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
)

const (
	maxRetries = 10
)

type talentAPI struct {
	url        string
	httpClient *retryablehttp.Client
}

func NewTalentAPIClient(url string) *talentAPI {
	apiClient := &talentAPI{
		url:        url,
		httpClient: retryablehttp.NewClient(),
	}
	apiClient.httpClient.RetryMax = maxRetries
	return apiClient
}

func (t *talentAPI) ListTalents(page int, limit int) ([]seed.Talent, error) {
	resp, err := t.httpClient.Get(fmt.Sprintf("%s?page=%d&limit=%d", t.url, page, limit))
	if err != nil {
		return nil, fmt.Errorf("failed to query for talents: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query for talents: status code is %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var ret []seed.Talent
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return ret, nil
}
