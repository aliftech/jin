package function

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/aliftech/jin/src/dto"
)

// fetchSSL fetches SSL/TLS certificate details
func FetchSSL(host string) (dto.SSLData, error) {
	data := dto.SSLData{}
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", host+":443", nil)
	if err != nil {
		return data, fmt.Errorf("failed to fetch SSL certificate: %v", err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	data.Issuer = cert.Issuer.String()
	data.Subject = cert.Subject.String()
	data.Expiration = cert.NotAfter.Format("2006-01-02 15:04:05")
	data.SANs = cert.DNSNames

	return data, nil
}
