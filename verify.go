package GameCenter

import (
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	keyLock      sync.Mutex
	publicKeyMap map[string][]byte
)

func getPublicKeyFromRemote(url string) []byte {
	res, _ := http.Get(url)
	buff, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return buff
}
func getPublicKey(url string) []byte {
	keyLock.Lock()
	defer keyLock.Unlock()
	if v, ok := publicKeyMap[url]; ok {
		return v
	}
	tmpBuff := getPublicKeyFromRemote(url)
	publicKeyMap[url] = tmpBuff
	return tmpBuff
}
func Verify(puk, data, sig string) error {
	realData, err1 := base64.StdEncoding.DecodeString(data)
	if err1 != nil {
		return err1
	}
	realSig, err2 := base64.StdEncoding.DecodeString(sig)
	if err2 != nil {
		return err2
	}
	publicKey := getPublicKey(puk)
	certificate, err3 := x509.ParseCertificate(publicKey)
	if err3 != nil {
		return err3
	}
	if err4 := certificate.CheckSignature(certificate.SignatureAlgorithm, realData, realSig); err4 == nil {
		return nil
	} else {
		return err4
	}
}

func init() {
	publicKeyMap = make(map[string][]byte)
}
