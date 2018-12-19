package utils

import (
	"crypto/md5"
	"io"
	"fmt"
)

type Encrypt struct {

}

func (encrypt *Encrypt) Str2md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}
