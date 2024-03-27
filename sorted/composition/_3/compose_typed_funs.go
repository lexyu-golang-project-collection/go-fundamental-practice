package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) error {

	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("New filename:", newFilename)

	return nil
}

func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

// sha1, prefix, hmac
func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func QQPrefixFilename(filename string) string {
	return "QQ_" + filename
}

func main() {

	str := QQPrefixFilename("cool_pic.jpg")
	println(str)

	h := hashFilename("cool_pic.jpg")
	println(h)

	s := &Server{
		filenameTransformFunc: prefixFilename("TT_"),
	}

	s.handleRequest("cool_pic.jpg")

}
