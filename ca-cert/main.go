package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	caKey  = "www.example.com.key"
	caCert = "www.example.com.cert"
)

func main() {
	dir, err := ioutil.TempDir("", "ca-cert")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = exec.Command("openssl", "req", "-nodes", "-new", "-x509", "-newkey", "rsa:4096", "-keyout", caKey,
		"-out", caCert, "-subj", "/C=JP/ST=Tokyo/L=Chuo/O=My Inc/OU=DevOps/CN=www.example.com").Run()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(caKey)
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(b)
	if block == nil || block.Type != "PRIVATE KEY" {
		log.Println(block.Type)
		log.Fatal(errors.New("Invalid CA Key"))
	}
	_, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	b, err = ioutil.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}
	block, _ = pem.Decode(b)
	if block == nil || block.Type != "CERTIFICATE" {
		log.Println(block.Type)
		log.Fatal(errors.New("Invalid CA Cert"))
	}
	crt, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(crt.Version)
	fmt.Println(crt.Subject.Country)
	fmt.Println(crt.Subject.Organization)
	fmt.Println(crt.Subject.OrganizationalUnit)
	fmt.Println(crt.Subject.Locality)
	fmt.Println(crt.Subject.CommonName)
}
