package pivit

import (
	"encoding/pem"
	"fmt"

	"github.com/cashapp/pivit/pkg/pivit/utils"
	"github.com/cashapp/pivit/pkg/pivit/yubikey"
	"github.com/pkg/errors"
)

// PrintCertificate exports the certificate.
func PrintCertificate(slot string) error {
	yk, err := yubikey.GetSigner(slot)
	if err != nil {
		return err
	}

	defer func() {
		_ = yk.Close()
	}()

	cert, err := yk.Certificate(utils.GetSlot(slot))
	if err != nil {
		return errors.Wrap(err, "get PIV certificate")
	}

	certBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	fingerprint := utils.CertHexFingerprint(cert)
	_, _ = fmt.Printf("%s\n%s", fingerprint, string(certBytes))

	return nil
}