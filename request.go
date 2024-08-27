package outline

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
)

type FingerprintTransport struct {
	http.RoundTripper
	Fingerprint string
}

func NewFingerprintTransport(fingerprint string, rt http.RoundTripper) *FingerprintTransport {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &FingerprintTransport{
		RoundTripper: rt,
		Fingerprint:  fingerprint,
	}
}

func (f *FingerprintTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Perform the request
	resp, err := f.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Get the certificate chain from the response
	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		// Calculate the SHA-256 fingerprint of the first certificate in the chain
		cert := resp.TLS.PeerCertificates[0]
		hash := sha256.Sum256(cert.Raw)
		certFingerprint := hex.EncodeToString(hash[:])

		// Compare the calculated fingerprint with the expected fingerprint
		if certFingerprint != f.Fingerprint {
			return nil, fmt.Errorf("certificate fingerprint does not match: expected %s, got %s", f.Fingerprint, certFingerprint)
		}
	} else {
		return nil, fmt.Errorf("no TLS certificate found")
	}

	return resp, nil
}

// doRequest makes a request to the server and returns the response body or an error
func (o *Outline) doRequest(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var marshal []byte
	var err error
	if body != nil {
		marshal, err = json.Marshal(body)
	}

	href := fmt.Sprintf("%s/%s", o.apiUrl, endpoint)
	request, err := http.NewRequestWithContext(ctx, method, href, io.NopCloser(bytes.NewReader(marshal)))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: NewFingerprintTransport(o.certSha256, &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}),
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if !slices.Contains([]int{http.StatusOK, http.StatusCreated, http.StatusNoContent}, response.StatusCode) {
		var responseError Error
		if err = json.Unmarshal(responseBody, &responseError); err == nil {
			return nil, fmt.Errorf(responseError.Error())
		}

		return nil, fmt.Errorf("error response status: %s", response.Status)
	}

	return responseBody, err
}
