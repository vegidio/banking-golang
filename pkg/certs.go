package pkg

import (
	"crypto/ecdsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"log"
)

//go:embed certs/access-token-private.pem
var accessTokenPrivate []byte

//go:embed certs/access-token-public.pem
var accessTokenPublic []byte

//go:embed certs/refresh-token-private.pem
var refreshTokenPrivate []byte

//go:embed certs/refresh-token-public.pem
var refreshTokenPublic []byte

type certs struct {
	AccessTokenPrivate  *ecdsa.PrivateKey
	AccessTokenPublic   *ecdsa.PublicKey
	RefreshTokenPrivate *ecdsa.PrivateKey
	RefreshTokenPublic  *ecdsa.PublicKey
}

type keyType int

const (
	Private keyType = iota
	Public
)

func parseEcKey(keyType keyType, pemKey []byte) interface{} {
	block, _ := pem.Decode(pemKey)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the key")
	}

	var key interface{}
	var err error

	if keyType == Private {
		key, err = x509.ParseECPrivateKey(block.Bytes)
	} else {
		key, err = x509.ParsePKIXPublicKey(block.Bytes)
	}

	if err != nil {
		log.Fatal("failed to parse PEM block containing the key", err)
	}

	return key
}

var Certs = certs{
	AccessTokenPrivate:  parseEcKey(Private, accessTokenPrivate).(*ecdsa.PrivateKey),
	AccessTokenPublic:   parseEcKey(Public, accessTokenPublic).(*ecdsa.PublicKey),
	RefreshTokenPrivate: parseEcKey(Private, refreshTokenPrivate).(*ecdsa.PrivateKey),
	RefreshTokenPublic:  parseEcKey(Public, refreshTokenPublic).(*ecdsa.PublicKey),
}
