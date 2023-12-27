package types

import (
	"crypto/x509"
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
)

type VerifyInput struct {
	SignatureBytes    []byte
	DataBytes         []byte
	IsDetached        bool
	MustCheckCertTime bool
}

type ValidateType string

const (
	ValidateOCSP    ValidateType = "OCSP"
	ValidateCRL     ValidateType = "CRL"
	ValidateNothing ValidateType = "Nothing"
)

type ValidateCertInput struct {
	Certificate   *x509.Certificate
	CheckCertTime bool
	ValidateType  ValidateType
	OCSPUrl       string
	CRLPath       string
}

// Kalkan - это обертка над методами KalkanCrypt.
type Kalkan interface {
	LoadKeyStore(path, password string) (err error)
	LoadKeyStoreFromBytes(key []byte, password string) (err error)
	X509ExportCertificateFromStore(outputPEM bool) (result string, err error)

	Sign(data []byte, isDetached, withTSP bool) (signature []byte, err error)
	SignXML(xml string, withTSP bool) (signedXML string, err error)
	SignWSSE(xml, id string) (signedXML string, err error)
	SignHash(algo ckalkan.HashAlgo, inHash []byte, isDetached, withTSP bool) (signedHash []byte, err error)

	Verify(input *VerifyInput) (string, error)
	VerifyXML(signedXML string, mustCheckCertTime bool) (result string, err error)
	VerifyDetached(signature, data []byte) (string, error)

	ValidateCert(input *ValidateCertInput) (string, error)

	HashSHA256(data []byte) ([]byte, error)
	HashGOST95(data []byte) ([]byte, error)

	SetProxyOn(proxyURL string) error
	SetProxyOff(proxyURL string) error

	GetCertFromCMS(cms []byte) ([]*x509.Certificate, error)
	GetCertFromXML(xml string) ([]*x509.Certificate, error)

	X509CertificateGetInfo(inCert string, fields []string) (string, error)
	GetTimeFromSig(cmsDer []byte) (time.Time, error)
	GetSigAlgFromXML(xml string) (string, error)

	Close() error
}
