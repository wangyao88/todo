package test

import (
	"testing"
	"crypto/md5"
	"io"
	"fmt"
)

func TestStr2md5(t *testing.T) {
	str := "123111"
	w := md5.New()
	io.WriteString(w, str)
	result := fmt.Sprintf("%x", w.Sum(nil))
	println(result)
}