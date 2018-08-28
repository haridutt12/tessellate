package cert

import (
	"testing"
	"os"
)

func TestMakeCertPool(t *testing.T) {
	certPool, err := MakeCertPool("./testdata/pem-file.pem")
	if err != nil {
		t.Fatalf("Error while making the Cert Pool. %v", err)
	}

	expected := []byte{48, 17, 49, 15, 48, 13, 6, 3, 85, 4, 3, 12, 6, 117, 98, 117, 110, 116, 117}
	for i, x := range certPool.Subjects()[0] {
		if x != expected[i] {
			t.Fatalf("Expected: %v, Got %v", expected, certPool.Subjects()[0])
		}
	}
}

func TestServerCerts(t *testing.T) {
	cert := "./testdata/cert-file.crt"
	key := "./testdata/key-file.key"
	root := "./testdata/pem-file.pem"
	creds, err := ServerCerts(cert, key, root)
	if err != nil {
		t.Fatalf("Could not load server certs.")
	}

	if creds.Info().ServerName != "tessellate-server" {
		t.Fatalf("Server name expected to be tessellate-server, got %v", creds.Info().ServerName)
	}
}

func TestClientCerts(t *testing.T) {
	cert := "./testdata/cert-file.crt"
	key := "./testdata/key-file.key"
	root := "./testdata/pem-file.pem"
	creds, err := ClientCerts(cert, key, root)
	if err != nil {
		t.Fatalf("Could not load server certs.")
	}

	if creds.Info().ServerName != "" {
		t.Fatalf("Server name expected to be empty, got %v", creds.Info().ServerName)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
