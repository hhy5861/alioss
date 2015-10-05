package alioss
import (
    "testing"
    "bytes"
    "io"
    "os"
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

func tTestPutObject(t *testing.T) {
    file, e := os.Open("./object.go")
    if e != nil {
        t.Errorf("open file error: %s\n", e)
    } else {
        stat, e := file.Stat()
        if e != nil {
            t.Errorf("get file stat error: %s\n", e)
        } else {
            err := NewTestClient().PutObject("object.go", file, stat.Size(), EmptyStringMap)
            if err != nil {
                t.Errorf("put object error: %s\n", err)
            }
        }
    }
}