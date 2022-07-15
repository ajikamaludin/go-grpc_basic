package cert

import (
	"crypto/x509"
	"io/ioutil"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/config"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
)

// Config is the struct that used to store the config file
type Config struct {
	FinacleCertPool Cert
}

// Cert is the struct wrapper which contains cert pool and flag to allow skip read the cert or not
type Cert struct {
	AllowSkip bool
	Pool      *x509.CertPool
}

// New init cert config file
func New(config *config.Config) (*Cert, error) {
	//cert config
	if config.Env != constants.EnvProduction {
		return &Cert{
			AllowSkip: true,
		}, nil
	}

	certPool := x509.NewCertPool()

	if config.Cert.Path == "" {
		return nil, nil
	}

	pem, err := ioutil.ReadFile(config.Cert.Path)
	if err != nil {
		return nil, err
	}

	certPool.AppendCertsFromPEM(pem)

	return &Cert{
		AllowSkip: false,
		Pool:      certPool,
	}, nil
}
