package outline

import (
	"context"
	"encoding/json"
	"net/http"
)

// ReturnsInformationAboutTheServer returns information about the server
func (o *Outline) ReturnsInformationAboutTheServer(ctx context.Context) (*Server, error) {
	body, err := o.doRequest(ctx, http.MethodGet, "server", nil)
	if err != nil {
		return nil, err
	}

	var data Server
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// RenamesTheServer renames the server
func (o *Outline) RenamesTheServer(ctx context.Context, name string) error {
	_, err := o.doRequest(ctx, http.MethodPut, "name", Name{name})
	return err
}

// ChangesTheDefaultPortForNewlyCreatedAccess changes the default port for newly created access
func (o *Outline) ChangesTheDefaultPortForNewlyCreatedAccess(ctx context.Context, port int) error {
	_, err := o.doRequest(ctx, http.MethodPut, "server/port-for-new-access-keys", Port{port})
	return err
}

// SetDataTransferLimitForAllAccessKeys sets a data transfer limit for all access keys
func (o *Outline) SetDataTransferLimitForAllAccessKeys(ctx context.Context, limit int) error {
	_, err := o.doRequest(ctx, http.MethodPut, "server/access-key-data-limit", Limit{Bytes{limit}})
	return err
}

// RemovesTheAccessKeyDataLimit removes the access key data limit, lifting data transfer restrictions on all access keys.
func (o *Outline) RemovesTheAccessKeyDataLimit(ctx context.Context) error {
	_, err := o.doRequest(ctx, http.MethodDelete, "server/access-key-data-limit", nil)
	return err
}

// ChangesTheHostnameForAccessKeys changes the hostname for access keys
func (o *Outline) ChangesTheHostnameForAccessKeys(ctx context.Context, hostname string) error {
	_, err := o.doRequest(ctx, http.MethodPut, "server/hostname-for-access-keys", Hostname{hostname})
	return err
}
