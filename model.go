package outline

import (
	"fmt"
)

// Server represents an outline server
type Server struct {
	Name                 string `json:"name"`
	ServerID             string `json:"serverId"`
	MetricsEnabled       bool   `json:"metricsEnabled"`
	CreatedTimestampMs   int    `json:"createdTimestampMs"`
	PortForNewAccessKeys int    `json:"portForNewAccessKeys"`
}

// Hostname represents a hostname for a server
type Hostname struct {
	Hostname string `json:"hostname"`
}

// Name represents a name for a server
type Name struct {
	Name string `json:"name"`
}

// Metrics represents metrics for a server
type Metrics struct {
	MetricsEnabled bool `json:"metricsEnabled"`
}

// AsnMetrics represents ASN metrics for a server
type AsnMetrics struct {
	AsnMetricsEnabled bool `json:"asnMetricsEnabled"`
}

// Port represents a port for an outline server
type Port struct {
	Port int `json:"port"`
}

// Limit represents a limit for an outline server
type Limit struct {
	Limit Bytes `json:"limit"`
}

type Bytes struct {
	Bytes int `json:"bytes"`
}

// NewAccessKey represents a new access key for an outline server
type NewAccessKey struct {
	Name     string     `json:"name,omitempty"`
	Password string     `json:"password,omitempty"`
	Port     int        `json:"port,omitempty"`
	Method   string     `json:"method,omitempty"`
	Limit    *DataLimit `json:"limit,omitempty"`
}

// DataLimit represents a data limit for an access key
type DataLimit struct {
	Bytes int `json:"bytes,omitempty"`
}

// AccessKey represents an access key for an outline server
type AccessKey struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      int    `json:"port"`
	Method    string `json:"method"`
	AccessURL string `json:"accessUrl"`
}

// Error represents an error from outline
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message
func (e *Error) Error() string {
	return fmt.Sprintf("Error %s: %s", e.Code, e.Message)
}
