package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gookit/color"
)

func generateHashes(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("error in reading the file")
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return errors.New("error returning the File Info")
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)

	if err != nil && err != io.EOF {
		return errors.New("error in parsing the file")
	}

	color.Cyan.Println("MD5:")
	fmt.Printf("%x\n\n", md5.Sum(bs))

	color.Cyan.Println("Sha1:")
	fmt.Printf("%x\n\n", sha1.Sum(bs))

	color.Cyan.Println("Sha224:")
	fmt.Printf("%x\n\n", sha256.Sum224(bs))

	color.Cyan.Println("Sha256:")
	fmt.Printf("%x\n\n", sha256.Sum256(bs))

	color.Cyan.Println("Sum384:")
	fmt.Printf("%x\n\n", sha512.Sum384(bs))

	color.Cyan.Println("Sha512:")
	fmt.Printf("%x\n\n", sha512.Sum512(bs))

	color.Cyan.Println("SHA512_224:")
	fmt.Printf("%x\n\n", sha512.Sum512_224(bs))

	color.Cyan.Println("SHA512_256:")
	fmt.Printf("%x\n\n", sha512.Sum512_256(bs))

	return nil
}
