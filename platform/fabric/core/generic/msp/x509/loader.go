/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package x509

import (
	x5092 "crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"path/filepath"

	pkcs112 "github.com/hyperledger-labs/fabric-smart-client/integration/nwo/common/pkcs11"

	"github.com/hyperledger-labs/fabric-smart-client/pkg/utils/proto"
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/config"
	msp2 "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/pkcs11"
	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/hyperledger/fabric/msp"
	"github.com/pkg/errors"
)

// Identity refers to the creator of a tx;
type Identity interface {
	Serialize() ([]byte, error)

	Verify(msg []byte, sig []byte) error
}

// SigningIdentity defines the functions necessary to sign an
// array of bytes; it is needed to sign the commands transmitted to
// the prover peer service.
type SigningIdentity interface {
	Identity //extends Identity

	Sign(msg []byte) ([]byte, error)
}

// GetSigningIdentity retrieves a signing identity from the passed arguments
func GetSigningIdentity(mspConfigPath, mspID string, bccspConfig *config.BCCSP) (SigningIdentity, error) {
	mspInstance, err := LoadLocalMSPAt(mspConfigPath, mspID, "bccsp", bccspConfig)
	if err != nil {
		return nil, err
	}

	signingIdentity, err := mspInstance.GetDefaultSigningIdentity()
	if err != nil {
		return nil, err
	}

	return signingIdentity, nil
}

// LoadLocalMSPAt loads an MSP whose configuration is stored at 'dir', and whose
// id and type are the passed as arguments.
func LoadLocalMSPAt(dir, id, mspType string, bccspConfig *config.BCCSP) (msp.MSP, error) {
	if mspType != "bccsp" {
		return nil, errors.Errorf("invalid msp type, expected 'bccsp', got %s", mspType)
	}
	conf, err := msp.GetLocalMspConfig(dir, nil, id)
	if err != nil {
		return nil, errors.WithMessagef(err, "could not get msp config from dir [%s]", dir)
	}

	cp, _, err := GetBCCSPFromConf(dir, bccspConfig)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get bccsp from config [%v]", bccspConfig)
	}

	mspOpts := &msp.BCCSPNewOpts{
		NewBaseOpts: msp.NewBaseOpts{
			Version: msp.MSPv1_0,
		},
	}
	thisMSP, err := msp.New(mspOpts, cp)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to create new BCCSPMSP instance at [%s]", dir)
	}
	err = thisMSP.Setup(conf)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to setup the new BCCSPMSP instance at [%s]", dir)
	}
	return thisMSP, nil
}

// GetBCCSPFromConf returns a BCCSP instance and its relative key store from the passed configuration.
// If no configuration is passed, the default one is used, namely the `SW` provider.
func GetBCCSPFromConf(dir string, conf *config.BCCSP) (bccsp.BCCSP, bccsp.KeyStore, error) {
	if conf == nil {
		return GetSWBCCSP(dir)
	}

	switch conf.Default {
	case "SW":
		return GetSWBCCSP(dir)
	case "PKCS11":
		return GetPKCS11BCCSP(conf)
	default:
		return nil, nil, errors.Errorf("invalid config.BCCSP.Default.%s", conf.Default)
	}
}

// GetPKCS11BCCSP returns a new instance of the HSM-based BCCSP
func GetPKCS11BCCSP(conf *config.BCCSP) (bccsp.BCCSP, bccsp.KeyStore, error) {
	if conf.PKCS11 == nil {
		return nil, nil, errors.New("invalid config.BCCSP.PKCS11. missing configuration")
	}

	p11Opts := *conf.PKCS11
	ks := sw.NewDummyKeyStore()
	mapper := skiMapper(p11Opts)
	csp, err := pkcs11.New(*pkcs112.ToPKCS11Opts(&p11Opts), ks, pkcs11.WithKeyMapper(mapper))
	if err != nil {
		return nil, nil, errors.WithMessagef(err, "Failed initializing PKCS11 library with config [%+v]", p11Opts)
	}
	return csp, ks, nil
}

func skiMapper(p11Opts config.PKCS11) func([]byte) []byte {
	keyMap := map[string]string{}
	for _, k := range p11Opts.KeyIDs {
		keyMap[k.SKI] = k.ID
	}

	return func(ski []byte) []byte {
		keyID := hex.EncodeToString(ski)
		if id, ok := keyMap[keyID]; ok {
			return []byte(id)
		}
		if p11Opts.AltID != "" {
			return []byte(p11Opts.AltID)
		}
		return ski
	}
}

// GetSWBCCSP returns a new instance of the software-based BCCSP
func GetSWBCCSP(dir string) (bccsp.BCCSP, bccsp.KeyStore, error) {
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	if err != nil {
		return nil, nil, err
	}
	cryptoProvider, err := sw.NewDefaultSecurityLevelWithKeystore(ks)
	if err != nil {
		return nil, nil, err
	}
	return cryptoProvider, ks, nil
}

func GetEnrollmentID(id []byte) (string, error) {
	si := &msp2.SerializedIdentity{}
	err := proto.Unmarshal(id, si)
	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshal to msp.SerializedIdentity{}")
	}
	block, _ := pem.Decode(si.IdBytes)
	if block == nil {
		return "", errors.New("bytes are not PEM encoded")
	}
	switch block.Type {
	case "CERTIFICATE":
		cert, err := x5092.ParseCertificate(block.Bytes)
		if err != nil {
			return "", errors.WithMessage(err, "pem bytes are not cert encoded ")
		}
		return cert.Subject.CommonName, nil
	default:
		return "", errors.Errorf("bad block type %s, expected CERTIFICATE", block.Type)
	}
}
