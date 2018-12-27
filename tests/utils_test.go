package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

func TestStr2md5(t *testing.T) {
	str := "root111"
	w := md5.New()
	io.WriteString(w, str)
	result := fmt.Sprintf("%x", w.Sum(nil))
	println(result)
}
