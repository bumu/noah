package schema

import "gorm.io/gorm"

type GatewayCert struct {
	gorm.Model

	CertType         string
	CertKey          string
	CertCSR          string
	Issuer           string
	CertIntermediate string
	CertChain        string
	CertSans         string

	NotBefore uint64
	NotAfter  uint64

	CommonName string
	CertName   string
	CertStatus string
	Comment    string
	Enabled    bool
}

// TableName overrides the table name used by User to `profiles`
func (GatewayCert) TableName() string {
	return "tbl_gateway_cert"
}
