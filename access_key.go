package outline

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateAccessKey creates a new access key
func (o *Outline) CreateAccessKey(ctx context.Context, key NewAccessKey) (*AccessKey, error) {
	body, err := o.doRequest(ctx, http.MethodPost, "access-keys", key)
	if err != nil {
		return nil, err
	}

	var data AccessKey
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// ListAllAccessKeys lists all access keys
func (o *Outline) ListAllAccessKeys(ctx context.Context) (*[]AccessKey, error) {
	body, err := o.doRequest(ctx, http.MethodGet, "access-keys", nil)
	if err != nil {
		return nil, err
	}

	var data struct {
		AccessKeys []AccessKey `json:"accessKeys"`
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.AccessKeys, nil
}

// CreateAccessKeyWithSpecificIdentifier creates a new access key with a specific identifier
func (o *Outline) CreateAccessKeyWithSpecificIdentifier(ctx context.Context, id string, key NewAccessKey) (*AccessKey, error) {
	body, err := o.doRequest(ctx, http.MethodPut, "access-keys/"+id, key)
	if err != nil {
		return nil, err
	}

	var data AccessKey
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// GetAnAccessKey get an access key
func (o *Outline) GetAnAccessKey(ctx context.Context, id string) (*AccessKey, error) {
	body, err := o.doRequest(ctx, http.MethodGet, "access-keys/"+id, nil)
	if err != nil {
		return nil, err
	}

	var data AccessKey
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteAnAccessKey deletes an access key
func (o *Outline) DeleteAnAccessKey(ctx context.Context, id string) error {
	_, err := o.doRequest(ctx, http.MethodDelete, "access-keys/"+id, nil)
	return err
}

// RenameAnAccessKey renames an access key
func (o *Outline) RenameAnAccessKey(ctx context.Context, id, name string) error {
	_, err := o.doRequest(ctx, http.MethodPut, "access-keys/"+id+"/name", Name{name})
	return err
}

// SetDataLimitForTheAccessKey sets a data limit for the given access key
func (o *Outline) SetDataLimitForTheAccessKey(ctx context.Context, id string, limit int) error {
	_, err := o.doRequest(ctx, http.MethodPut, "access-keys/"+id+"/data-limit", Limit{Bytes{limit}})
	return err
}

// RemoveDataLimitForTheAccessKey removes the data limit on the given access key.
func (o *Outline) RemoveDataLimitForTheAccessKey(ctx context.Context, id string) error {
	_, err := o.doRequest(ctx, http.MethodDelete, "access-keys/"+id+"/data-limit", nil)
	return err
}
