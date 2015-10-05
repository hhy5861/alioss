package alioss
import (
    "testing"
    "bytes"
    "io"
)

func tTestBuffer(t *testing.T) {
    temp := []byte{}
    A(temp)
    t.Log(string(temp), "\n")
}

func A(b []byte) {
    buf := bytes.NewBuffer([]byte{})
    red := bytes.NewReader([]byte("hello world"))
    io.Copy(buf, red)
    // b = buf.Bytes()
    copy(buf.Bytes(), b)
}