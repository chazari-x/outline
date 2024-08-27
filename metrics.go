package outline

import (
	"context"
	"encoding/json"
	"net/http"
)

// ReturnsTheDataTransferredPerAccessKey returns the data transferred per access key
func (o *Outline) ReturnsTheDataTransferredPerAccessKey(ctx context.Context) (*map[string]int, error) {
	body, err := o.doRequest(ctx, http.MethodGet, "metrics/transfer", nil)
	if err != nil {
		return nil, err
	}

	var data struct {
		BytesTransferredByUserID map[string]int `json:"bytesTransferredByUserId"`
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data.BytesTransferredByUserID, nil
}

// ReturnsWhetherMetricsIsBeingShared returns whether metrics is being shared
func (o *Outline) ReturnsWhetherMetricsIsBeingShared(ctx context.Context) (*Metrics, error) {
	body, err := o.doRequest(ctx, http.MethodGet, "metrics/enabled", nil)
	if err != nil {
		return nil, err
	}

	var data Metrics
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// EnablesOrDisablesSharingOfMetrics enables or disables sharing of metrics
func (o *Outline) EnablesOrDisablesSharingOfMetrics(ctx context.Context, metricsEnabled bool) error {
	_, err := o.doRequest(ctx, http.MethodPut, "metrics/enabled", Metrics{metricsEnabled})
	return err
}

// AnnotatesPrometheusDataMetricsWithAutonomous annotates Prometheus data metrics with autonomous
func (o *Outline) AnnotatesPrometheusDataMetricsWithAutonomous(ctx context.Context, asnMetricsEnabled bool) error {
	_, err := o.doRequest(ctx, http.MethodPut, "experimental/asn-metrics/enabled", AsnMetrics{asnMetricsEnabled})
	return err
}
