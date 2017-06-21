package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// HSMSafenetClientInfo h s m safenet client info
// swagger:model HSMSafenetClientInfo
type HSMSafenetClientInfo struct {

	// Generated File - Chrystoki.conf .
	// Read Only: true
	ChrystokiConf string `json:"chrystoki_conf,omitempty"`

	// Client Certificate generated by createCert.
	ClientCert string `json:"client_cert,omitempty"`

	// Name prepended to client key and certificate filename.
	// Required: true
	ClientIP string `json:"client_ip"`

	// Client Private Key generated by createCert.
	ClientPrivKey string `json:"client_priv_key,omitempty"`

	// Major number of the sesseion.
	SessionMajorNumber int32 `json:"session_major_number,omitempty"`

	// Minor number of the sesseion.
	SessionMinorNumber int32 `json:"session_minor_number,omitempty"`
}