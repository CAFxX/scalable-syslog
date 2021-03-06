// Code generated by go-bindata.
// sources:
// scalable-syslog-certs/adapter.crt
// scalable-syslog-certs/adapter.csr
// scalable-syslog-certs/adapter.key
// scalable-syslog-certs/scalable-syslog-ca.crl
// scalable-syslog-certs/scalable-syslog-ca.crt
// scalable-syslog-certs/scalable-syslog-ca.key
// scalable-syslog-certs/scheduler.crt
// scalable-syslog-certs/scheduler.csr
// scalable-syslog-certs/scheduler.key
// DO NOT EDIT!

package app_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _adapterCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIEKzCCAhOgAwIBAgIRAOHeuMIFj+29goa2s3DVVKcwDQYJKoZIhvcNAQELBQAw
HTEbMBkGA1UEAxMSc2NhbGFibGUtc3lzbG9nLWNhMB4XDTE3MDIxMzE3MTY0OVoX
DTE5MDIxMzE3MTY0OVowEjEQMA4GA1UEAxMHYWRhcHRlcjCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBANS9/Y4goMdVUe/dtF7Hdn6mfgt5/azLu1ja+bka
ENcCyJG9ajjCWomMMXEu2aFq8AAFt3yjGs9QfN6uwr/AIHj2Ja+Scz1mAAhEj10I
RlERupTV/SQgrRZxWf6slGmdBstG/6Q0bkLrWtMOZ1ZntKu96CMsp2HwtSaR9mkX
UJ0AxeOxKFAeOCialebDM35CkTO/+cUZWo3TOl0Xam4xaShxyyNQlJ9xPq8SnT3g
D8iuvbzijPDIALI9VkFcUu+i8URk4Ddv2iVriPFer+olKUeJx23mxSbYZmWbmDpW
DfThPxGcEN4boDnSW++4yOJp2YH4pJL7JvITb3J9NUpaaM8CAwEAAaNxMG8wDgYD
VR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAdBgNV
HQ4EFgQUzaGjFxXSky/stVvXd+Cc0KxmPo8wHwYDVR0jBBgwFoAUyntbc1IHtXVr
WtUkq0oLeOFqIDcwDQYJKoZIhvcNAQELBQADggIBAFxQ64wI8mLLKoiKikRq9waS
iB/NHJxwQJ2yuhSR6nCSiJGZOqpSs3lySSekzVNMBA4Mh/6x8JaqFgg4ld2WkWMf
EOc6DkdaQoY5v2oMB6ogBRx40aiQveLvL+TNsfdaJ3sybg6RoLfoIC1V+euWTKgx
qCacHRIkbCUE7GCmFiX2atrQbhEYmtKv2P99gHnNMoTXC4K5rGuhsnZ7d9Hkr5ZD
bhSmIv0T/JZ4vLlwhV6z+735eigs8yRhWob0hzRi7VfNc/VG7PZFXbVUgJ0huNd4
6LPpQ/b8eeA7yTw+Cde4ZEruAm1vMNqs/Gi39YyOUPG0mkkd+prUqq22R28QAuw1
fGVHXmjL/Aj7kLifgUHYRm1Fi+7PKLKzvhnx0URxldaJ1+a/X3jetgQn9RJ7upzK
zfj8EB6lvjbLYPZ7xXm4pzs/0Ado1djYV3USsy6JAe6bOIGz2F/h4YRNJbVAGZ+Q
taQYEmqrGpZpNKgRYUJefK6cEK6vr1OZGGVOFdZpqMMmSLYTgoz38q1GXKeOwsR8
xjSVYph6aS2mHAxcNQ3L8UsqgDu6CzzCL1YoaSs4GwIcu/piM0WqlSrsxymwGZUT
vhySioh7iOPTKzJxdrZRfhA/zxJsTzKnZFD4QfYw5tIr0m11rrIt+FSh1HSQ58YI
yuQfnztJjQc9+Wt2+uAU
-----END CERTIFICATE-----
`)

func adapterCrtBytes() ([]byte, error) {
	return _adapterCrt, nil
}

func adapterCrt() (*asset, error) {
	bytes, err := adapterCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "adapter.crt", size: 1505, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _adapterCsr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICVzCCAT8CAQAwEjEQMA4GA1UEAxMHYWRhcHRlcjCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBANS9/Y4goMdVUe/dtF7Hdn6mfgt5/azLu1ja+bkaENcC
yJG9ajjCWomMMXEu2aFq8AAFt3yjGs9QfN6uwr/AIHj2Ja+Scz1mAAhEj10IRlER
upTV/SQgrRZxWf6slGmdBstG/6Q0bkLrWtMOZ1ZntKu96CMsp2HwtSaR9mkXUJ0A
xeOxKFAeOCialebDM35CkTO/+cUZWo3TOl0Xam4xaShxyyNQlJ9xPq8SnT3gD8iu
vbzijPDIALI9VkFcUu+i8URk4Ddv2iVriPFer+olKUeJx23mxSbYZmWbmDpWDfTh
PxGcEN4boDnSW++4yOJp2YH4pJL7JvITb3J9NUpaaM8CAwEAAaAAMA0GCSqGSIb3
DQEBCwUAA4IBAQBWW6jQJq5tPaww0aqdPVDvkKTgZWKxAE/Wm+nfUIS9vgMOcwfc
4liYGI0g0Y/hnIXzAf8q0PUKmU/RURwD3Xw6s1UHjsmVb6usmShJPVpUiVEqX9Pc
aHpOrvkRQTewlwIJ//iZ2Xi3TjfJ9VCE9rAPSFwcvH4gGFVdpV5debcXl25hEgvn
iYgCidvV8dWNc2YgrpWAQjRKj9mf63ySww5jl236ZuWKHtzXAwkhT6hlgGyfSu6T
tdD2+xV0SWX9JI+DJt2GphB4UckFp+wn2KzqbZkhMF1s4t8ZS3VfzhN7hm0F3B7W
08yzBWaUp3GSvhbn9SneKtQ38xkAvZxkDnk8
-----END CERTIFICATE REQUEST-----
`)

func adapterCsrBytes() ([]byte, error) {
	return _adapterCsr, nil
}

func adapterCsr() (*asset, error) {
	bytes, err := adapterCsrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "adapter.csr", size: 887, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _adapterKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1L39jiCgx1VR7920Xsd2fqZ+C3n9rMu7WNr5uRoQ1wLIkb1q
OMJaiYwxcS7ZoWrwAAW3fKMaz1B83q7Cv8AgePYlr5JzPWYACESPXQhGURG6lNX9
JCCtFnFZ/qyUaZ0Gy0b/pDRuQuta0w5nVme0q73oIyynYfC1JpH2aRdQnQDF47Eo
UB44KJqV5sMzfkKRM7/5xRlajdM6XRdqbjFpKHHLI1CUn3E+rxKdPeAPyK69vOKM
8MgAsj1WQVxS76LxRGTgN2/aJWuI8V6v6iUpR4nHbebFJthmZZuYOlYN9OE/EZwQ
3hugOdJb77jI4mnZgfikkvsm8hNvcn01SlpozwIDAQABAoIBADb/00BeQRKnhSJd
D9d8+65sfqOa6zE0DoA/RFZNXdMoXhAHhyGbZWqv+sxHD26Cxf9Lew9W34R/S3kK
d0C5e+upvz5vvyk9aH66stGbZQpqafE2jVY8uGLX9Ss5FLllJi+BcG/UOjlEvVtZ
GvAkqXnJA+2sCq3wW/TXwXuT0nx0EMEwRyRjzscO62AxtkygxOJZ/+JoAk9GoKIO
FKzZXehsCzgkypKCv70LAxvTgWSrZzg0BKTK0eOmSX0Uxd77xtqOiA+dWTPaOFvr
7obh3C/oUduBWhuU3GwQ+Mf3cKWeLxzFg4iDP87De/X9MnkytqmoOneNkcJt/v1U
3xrLFQECgYEA+f2VFZkNsATua9bovIYI6d4KzRNkWiZni3e/Gq84DGgzOGqIqGPn
URkLuDtMODTUb+rPFttGCBVWQ+y49OJN9R1fSThcMiQNKwZ0dgKg7mN0lNt7p5Mv
YpQj2HKhp9/Oo0t8QMdCdxz2Ut6ayid6SFsotGGORBsU6Zr3027lRoECgYEA2dsv
XOheGfKpdAGLGi1ZYDZuzl/n6OOquKlRvjKH4O6SIk3tL+Nmejpv33uBKdiWwEvH
ghlSyf44W8mRBnNyeyw+b/8hi24miLxbtO24fiG07Gf5A7m4stFAEq/QMOGFjOgo
DtqI1W+rkX+UiX1gf1T5L2337xf3/TtpdA3cJ08CgYEAvr9lACPWs9Yvbg1/bEoB
hyErsKr5SWhAXcSzBtNnut/PJV70gzgnilPIA30LfV+iMPtDpLcPSp+tQQrhXclH
np4Y1K+yeXfUrF2yg1EiJoOwstx+D7FY2KkcaM2e148IBsVUO8FOz8BDm5vLFDDn
N0qHVouRmBnp7Q2Xx7a4LwECgYAq1CBXj4cQathNahBCsS+k7o6SG+CntDpX21Gq
ppx35+7qt48tnvMCjIJ52PnnlCDu5Pbv6LKR61yBDvtn3UJgXK5ZrjqWAq2oq/8L
bsCaa136K+aUiOp6nRCPm+i981gh/3IAmY3VCAMes4osDW3vGnFylZfsdTXXntp4
alEhmQKBgQDWbRVfiXpfK/uwbCc0y4Qt77lJ5JzBbnnBt7JgMjKKf7rMXchuUN66
mY/2Z4pFA2owxMwDiFbWkx74FYVTD/4U/kYC4Im6dmTAyCuCWCdWJ/EhCglrtU6V
aR3ds6oRnzEYdchvNM15ek4MkxspfzIoHw8ajRbkBO3AsZCPIV9cjg==
-----END RSA PRIVATE KEY-----
`)

func adapterKeyBytes() ([]byte, error) {
	return _adapterKey, nil
}

func adapterKey() (*asset, error) {
	bytes, err := adapterKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "adapter.key", size: 1679, mode: os.FileMode(288), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scalableSyslogCaCrl = []byte(`-----BEGIN X509 CRL-----
MIICjDB2AgEBMA0GCSqGSIb3DQEBCwUAMB0xGzAZBgNVBAMTEnNjYWxhYmxlLXN5
c2xvZy1jYRcNMTcwMjEzMTcxNjQ5WhcNMjcwMjEzMTcxNjQ5WjAAoCMwITAfBgNV
HSMEGDAWgBTKe1tzUge1dWta1SSrSgt44WogNzANBgkqhkiG9w0BAQsFAAOCAgEA
IfsY2l/eyNTBI2TEVyIrU4/S6FmfsQHR0/ajbB3XfISdYMMyyp4XNBronu0qhODU
vh/4afrRIR8mLZ6fEMR4j7vyVTSuUpRoWF+jhIjLs+QN8gcKIacAFTX5OjGugS1M
sFsf5mP1qh7EfYRTptGdepQUd5Rly3obcY1FCCnuqQKIkcr2B/FGewnnsYfF4Ps+
2yvmCCOpPW/glFSFgAm6bbsLEK4jr8nKbAJkEoukEhGBouLWT+yxX2lNPTJDfXln
1LE73L1sirTwRRHzHBcU2EXmfJutzGP8aWyFe3tBgq0iK9rQlC7LofTGXdiVvJo8
Ag2bMnpI0/w7fa/pQOWNewjGoNuNUJhsbjzyLRzHbOuxbx8BG0+/dmYQnPrqLi+0
1l9d3ke/EjbF2l/XGdGBtYK3X//XNhJNr1gmL02E5taahIRo1vypnPdZ7H8gPYxi
G7E5hZyVCSVfee/rBIC1hLNGjjuJLr1IyBEwLPpc5EJtFoYnepsMJAc52lP4UIhq
sydyHbBfMaBAqj0DrrZkjFXoq70/9Rb/chAFgxwVvNnCeZWaQXVsE3BIqTnHGb4/
UwOLadGdiZW3Hnwoio5cY9NFS08K+nX1DRoSgP401CnwafdAXWa/7dfYfX7uR8OQ
tuoUZ64+6ojOM1ZHh71TrZ/y1K3//0w0W0KBMzbCJSo=
-----END X509 CRL-----
`)

func scalableSyslogCaCrlBytes() ([]byte, error) {
	return _scalableSyslogCaCrl, nil
}

func scalableSyslogCaCrl() (*asset, error) {
	bytes, err := scalableSyslogCaCrlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scalable-syslog-ca.crl", size: 938, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scalableSyslogCaCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIE+jCCAuKgAwIBAgIBATANBgkqhkiG9w0BAQsFADAdMRswGQYDVQQDExJzY2Fs
YWJsZS1zeXNsb2ctY2EwHhcNMTcwMjEzMTcxNjQ2WhcNMjcwMjEzMTcxNjQ5WjAd
MRswGQYDVQQDExJzY2FsYWJsZS1zeXNsb2ctY2EwggIiMA0GCSqGSIb3DQEBAQUA
A4ICDwAwggIKAoICAQC220SlKA6W95Eqn3oj35K2J01Mas8iL52SjkxAf7hOd05p
PjfSqz55/hC2b4CG+K4Y4vxXsr8NOvuf2XVlVgibUG3R4MAS0dizlbQOm5+0rZ8G
W4ZeopbwuEbfGVt3SplgL7YS+5OVYJX934kQe5+QDwrkrnGmY+8T0oA41f/0i/kX
PVLzkypesNKKmUeswYkO27rWcTJZ2hckicYQHAqBRswV6LyJYY4itX6ZI2BSYaNG
QHMe75t1NWbvD5rGSVGg/XBKeuVOCB3jkSN2phoKWCTo/LPPf2yD4br5wtLxi9lV
dAIAAZGlGO/glq/dGFyQZkkANiefIZTBQjtOvBgIbKq7XpMAuQNNVjdzZRh8eOhq
c8ajdlBJISrE5t/8qzUaZupPfBc+ueusBcL9AyXvfiSjhDX1BAAWmNy2QB2cqb6t
EIptpCKscNlibtqqaAiqCpl7FZSRmoII42DgshArTl4LkKXQanf1RmH0iyaVDYz9
eaz/ouCL2kKN21qFOY6x591cgCVWeUpEVjZOnzFH/Si35uZMl7vWXi81A1xNakQg
fIaNbMpRRD+Hkit2OEnF0kVfx8CqfrYWosoG0wb6kSRdi2/Kl+LgqcznFJk9FOBP
HiyZjRgtSYCw5pXS2SoE+tRblmG3FFJzfG+PqdC1UJ7C+esD1p7UPwjFlHNIMwID
AQABo0UwQzAOBgNVHQ8BAf8EBAMCAQYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNV
HQ4EFgQUyntbc1IHtXVrWtUkq0oLeOFqIDcwDQYJKoZIhvcNAQELBQADggIBAEtu
rb7rGjfSoyCivUuYe0kv52vOQ2sC+aiucklmk1KQhHaknAhbl8YorL8uZQKytYdQ
B7Jp8eDfw/H8b/k3K7Hk/S4CXyKbbnaCUyaUPC6DQx6Bt9N0vECcyb3D1lsE0+/p
Qzs166Vcs/YunC+EXpXVhv9WlNwH+EFSXH6ulVMM8iSQClqqMl8a9Ihup5nWY7+X
U6qf0ZvLQSzr5sH2yMkqdzOqI9b674JyGzXQsSAFoOIeD1MRGoF6msHM6ZfeYc4G
gHqGqq27LebZKOznNPO3r30UkhohpusUiK04Qqu4RohtQnPrzJ4tVwU1qTH1utnA
DgWm/drm+s7oxwwBS+MIrhL/+5zI0XIi0/PbBqufOOiePGjd3v0wbeE42I6QITL1
RD/XNzaxBs9MSNwu0WIrwkal/szXrZEjXvsVhNfA7XNzMkDDtM6XDduo6bCU7k1r
2gRluUsqnV3XUoo+LisNtPxBjJdSepEMloR+9zOeH9CQXMIl/45EqXRyHslFstB0
wYNQqv5KgQg74EI2fljneXcOXvTvxxYRXVhRiV+P3F3oOOT6qc6CZCYnr5jhQW87
OvLr2HhepUTrH9r+Eo0KgKLroprSWFeCt9vRLJTj+gH/2R9lR8qNB/OCGvYzpjH2
PCh9NVk9+g9qXDk9W5ofTte6JryKx0xNxiA3Z89+
-----END CERTIFICATE-----
`)

func scalableSyslogCaCrtBytes() ([]byte, error) {
	return _scalableSyslogCaCrt, nil
}

func scalableSyslogCaCrt() (*asset, error) {
	bytes, err := scalableSyslogCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scalable-syslog-ca.crt", size: 1785, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scalableSyslogCaKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func scalableSyslogCaKeyBytes() ([]byte, error) {
	return _scalableSyslogCaKey, nil
}

func scalableSyslogCaKey() (*asset, error) {
	bytes, err := scalableSyslogCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scalable-syslog-ca.key", size: 3243, mode: os.FileMode(288), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schedulerCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIELDCCAhSgAwIBAgIQM1vslGwPDfABVwBK1nVpRDANBgkqhkiG9w0BAQsFADAd
MRswGQYDVQQDExJzY2FsYWJsZS1zeXNsb2ctY2EwHhcNMTcwMjEzMTcxNjQ5WhcN
MTkwMjEzMTcxNjQ5WjAUMRIwEAYDVQQDEwlzY2hlZHVsZXIwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDGyif/W2A0qC1lTsMsZq9YmW9bKSc4DqoGiz2B
+w/whP9xZSf7XFrY82NBOH5Ay16hMl0yPjJNpJX1qTQwSjftzBQoDUTUDk9Zhijk
m6p60oqgvFGeOU950uEXmhEz+4R4Nm2DsU++ru/dyutSpOTQ6IeYmgdFLFCEpj98
hNqNoRkA8veobi5i10mkGv1jvaBRDV9nHsdZb1TtMg+ctyWe4AtVrro6k2PZkjw3
NxVNGuoxlI/O+J46CSNCmkbG6PKfCgboIXq7kxCL8ayyaTTTc4vKdmU9qX5u6FTq
M1tGq3lZPpaSaU5J66lmGEVDN7ow7HiD9tbq/mrES5xoM30RAgMBAAGjcTBvMA4G
A1UdDwEB/wQEAwIDuDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHQYD
VR0OBBYEFLXsVc22saq8oiMN+u//mKPFwJXHMB8GA1UdIwQYMBaAFMp7W3NSB7V1
a1rVJKtKC3jhaiA3MA0GCSqGSIb3DQEBCwUAA4ICAQA6vAthBfivcgZNwE9wU0t+
0BdOZ3RvY7sLS5L6rjrDpzpWo5jxk6fMzVjzLvMA3CuEOQRafjRExQPDNzKzug5o
W1AoptGte+tNrJ7/fCWBIGl8FQgvym26QKSM7pS8Or86R2SeEMuOvmDdt0HV3tL7
vG4ve7Gy/j6X/OjbtZ2JubIWs+8leE3OaYImXE178yY49we57HZOdbXm4dtQKgV2
Xlkat7MmvNYT0nFU2eXWtcJ6DPd5BlzGhZu0gyQC9SeJpDxHcW/BYYTAh0dFCrXt
CqISQnQwexCMT+dRTZv8z4KzZnb7cpCEGjekx5vhyE95kxjUjlQSBZ7P06A+UGJp
Dre06juclTFNA50wIX1PhZyurexJPeGFPK+e0iDv5s/wDA0tt02XRfVINmi7icBA
pLWhmcQEvSBKp5VeG4U9b88t3215t3ZyXAi0qqaoJSNrtI10KAJ+OXyj6+AGKFci
8mfFrjlk/bBZoK8A9HuI7TET82wPk+VCYt8+0GzzqWNpx81mLcK+XkfJtizMphp7
BWWczAK86LuwqsKVKQMW5LutyRhxsaSmJQqaPt2TiV6E6bznYCV66YmyHVRJbdXq
1yH6YFDpJM55O3hERPWsASk8akOthK4mmwd2wENTtg65gSWW7DjOLZ/omc1xmQEV
mvg0V0/lmbzs6zBnot/f/Q==
-----END CERTIFICATE-----
`)

func schedulerCrtBytes() ([]byte, error) {
	return _schedulerCrt, nil
}

func schedulerCrt() (*asset, error) {
	bytes, err := schedulerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scheduler.crt", size: 1509, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schedulerCsr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICWTCCAUECAQAwFDESMBAGA1UEAxMJc2NoZWR1bGVyMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEAxson/1tgNKgtZU7DLGavWJlvWyknOA6qBos9gfsP
8IT/cWUn+1xa2PNjQTh+QMteoTJdMj4yTaSV9ak0MEo37cwUKA1E1A5PWYYo5Juq
etKKoLxRnjlPedLhF5oRM/uEeDZtg7FPvq7v3crrUqTk0OiHmJoHRSxQhKY/fITa
jaEZAPL3qG4uYtdJpBr9Y72gUQ1fZx7HWW9U7TIPnLclnuALVa66OpNj2ZI8NzcV
TRrqMZSPzvieOgkjQppGxujynwoG6CF6u5MQi/Gssmk003OLynZlPal+buhU6jNb
Rqt5WT6WkmlOSeupZhhFQze6MOx4g/bW6v5qxEucaDN9EQIDAQABoAAwDQYJKoZI
hvcNAQELBQADggEBADdUy9QEB+aMWGJq7luy9G0fDTT71WcpdH5ifuRKDY+vZW7z
Eo4VklsvDlASZJyKvzlPBwE/KeLqXg1VM+I4NFMkzjoVZFOsLjXXG3NR4Z3KyMtK
gc6pR2EXbwtNlLxChonbSg75f046KwA2FElX0Crk10IoQufM+I5DHT9JWXL5w+D/
pkdAzjMQ/gdf08YwHHSpkfMVfQfdmLb2XAG7zeV76hTwOgVb+eQf4hbvKFap49ba
lVlK1Ih/BXRYvnL9lbDDuJlxarrQCuPevu62UKE6u8j/QmaL0y9CRaA8XUQpD70I
gZcMNtukU6IhKTeea++8spiIC+zlguhIKz7sDJ0=
-----END CERTIFICATE REQUEST-----
`)

func schedulerCsrBytes() ([]byte, error) {
	return _schedulerCsr, nil
}

func schedulerCsr() (*asset, error) {
	bytes, err := schedulerCsrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scheduler.csr", size: 891, mode: os.FileMode(292), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schedulerKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxson/1tgNKgtZU7DLGavWJlvWyknOA6qBos9gfsP8IT/cWUn
+1xa2PNjQTh+QMteoTJdMj4yTaSV9ak0MEo37cwUKA1E1A5PWYYo5JuqetKKoLxR
njlPedLhF5oRM/uEeDZtg7FPvq7v3crrUqTk0OiHmJoHRSxQhKY/fITajaEZAPL3
qG4uYtdJpBr9Y72gUQ1fZx7HWW9U7TIPnLclnuALVa66OpNj2ZI8NzcVTRrqMZSP
zvieOgkjQppGxujynwoG6CF6u5MQi/Gssmk003OLynZlPal+buhU6jNbRqt5WT6W
kmlOSeupZhhFQze6MOx4g/bW6v5qxEucaDN9EQIDAQABAoIBAQCtkwapqZevYsmA
k+1Hb8Hbkt3ws1ZTVDqRzDLAs/+O/BIvaSwoBtQsaxhXcoSK/wh5qAhIPpadxQBL
qJqxL5z2PiO7dRqMLpKKkOJT+rg41m3FO6nohhInRULV4Oj0gsGAVcX6H3CKByXa
e7xORTymTONZE+UVOUR1DtSQXVCYKG+kX2SC6prGsyDBLp5kIBd+MB/tlvxPjTJe
FVmtDXn/Rq4jN4bNe2bpUYpNU+b1KHhFdDw61g6JYhXuZLa98WFgm/nvjDsvo4RJ
M+a1K6QNbUZ1IaE/kJdMe35Ty8CuEYMBxspLh6H4P5FsWs0DcHG2mxVZVwEe8ezT
/V7kXpIBAoGBANVhFa6mKXVg7gNawPfX6gK93UHjg3l7oHmb1KhUmnQS3iQcpswN
nIfEeYqBupAzMpZ0eVQSQRteV66YYhnNvwPiKg5CkdHRGMl5fBWsJ4GzQqU50QYu
qRadCZPaOC/Kb1VcRWPYtw4jv+wQaZKkOsvm4dLAlkdG0p83HMi/1/5JAoGBAO5/
DM8sEvxtrowFNjwBSyKcbFteNjoeCvlUjSW0hXPljIa6Bo+0oZcYdWhhyJarP/x+
dIUSx5MJknch2c/7Kt26octJO7jTJh6lH1/MBGMq/lArCkStN6IiKybHSLR6Zl3z
siyYVQfswhUbpjWnMfKEP8bZMGnkk/XmQJni2SiJAoGAKwyJOVEgsCO8IUUgBWC0
P5VanD5wegPtHlvIDaXz+1MLKjH1nNlqKSIEe2Ms5obNv/kDuWhtZj4kZn8r25Tc
4sWwQmw/yxrqa0ttM1omJk6qNXdyqlKjnoJUOlYeW9X1nD0fasOAOwQPZmrxQbqM
ejK8gK6GIx8wwxp1XjMGOrECgYEAo9cIF7KxgFRhGrPh5CsHUMmg7suz+Qz2gP6C
0KpJG6tQdKWLKGHwuHoG/iRexpLyrCS7K0gRboOe6NSKa4SimBxEVgEDVB9KCrdo
7EidovrAJpbKwQ82Lt3GQyeYXNPRSy3E8znbAEy3sASElEs91trfbV17EuQCeWUq
cA0VHfECgYAbkPMD+GznOFMEOrzfqeycMaVsgLs4UcRX4yaQBXwQn0ANIAtfNRu8
1XULV3Neeq8KHAqAvjzcUFCMbJJfj5Kbxi272ymAc2vTCYJXpe0DHBFQq6EffdQ0
e5Nv9vaO64r05F9L6Dvn93971n4AkYEJQUUKRTnlRCmAnEIsM3eA3A==
-----END RSA PRIVATE KEY-----
`)

func schedulerKeyBytes() ([]byte, error) {
	return _schedulerKey, nil
}

func schedulerKey() (*asset, error) {
	bytes, err := schedulerKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scheduler.key", size: 1679, mode: os.FileMode(288), modTime: time.Unix(1487006209, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"adapter.crt":            adapterCrt,
	"adapter.csr":            adapterCsr,
	"adapter.key":            adapterKey,
	"scalable-syslog-ca.crl": scalableSyslogCaCrl,
	"scalable-syslog-ca.crt": scalableSyslogCaCrt,
	"scalable-syslog-ca.key": scalableSyslogCaKey,
	"scheduler.crt":          schedulerCrt,
	"scheduler.csr":          schedulerCsr,
	"scheduler.key":          schedulerKey,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"adapter.crt":            &bintree{adapterCrt, map[string]*bintree{}},
	"adapter.csr":            &bintree{adapterCsr, map[string]*bintree{}},
	"adapter.key":            &bintree{adapterKey, map[string]*bintree{}},
	"scalable-syslog-ca.crl": &bintree{scalableSyslogCaCrl, map[string]*bintree{}},
	"scalable-syslog-ca.crt": &bintree{scalableSyslogCaCrt, map[string]*bintree{}},
	"scalable-syslog-ca.key": &bintree{scalableSyslogCaKey, map[string]*bintree{}},
	"scheduler.crt":          &bintree{schedulerCrt, map[string]*bintree{}},
	"scheduler.csr":          &bintree{schedulerCsr, map[string]*bintree{}},
	"scheduler.key":          &bintree{schedulerKey, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
