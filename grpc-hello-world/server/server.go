package server

import "log"

var (
	ServerPort  string
	CertName    string
	CertPemPath string
	CertKeyPath string
)

func Server() (err error) {
	log.Println(ServerPort)
	log.Println(CertName)
	log.Println(CertPemPath)
	log.Println(CertKeyPath)
	return nil
}
