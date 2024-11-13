package osm

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const apiUrl = "https://nominatim.openstreetmap.org/"

func runRequest[T any](ctx context.Context, url string, opts ...optFunc) (*T, error) {
	o := mergeOpts(opts)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", o.Locale)
	req.Header.Set("User-Agent", o.UserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var handled ErrorResult
		if err := json.NewDecoder(resp.Body).Decode(&handled); err == nil {
			return nil, &handled
		}
		return nil, errors.New("unexpected status code: " + resp.Status)
	}
	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
