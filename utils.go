package alioss
import (
    "bytes"
    "net/http"
    "io"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

var (
    EmptyStringMap      map[string]string = map[string]string{}
    EmptyStringSlice    []string          = []string{}
)

func UnmarshalXmlResp(r *http.Response, c interface{}) Error {
    buf := bytes.NewBuffer([]byte{})
    _, e := io.Copy(buf, r.Body)
    if e != nil {
        return E_HttpResp
    }
    e = xml.Unmarshal(buf.Bytes(), c)
    if e != nil {
        return E_ParseXML
    }
    return nil
}

func ReadByteResp(r *http.Response) (resp []byte, err Error) {
    buf := bytes.NewBuffer([]byte{})
    _, e := io.Copy(buf, r.Body)    // maybe use r.Header["Content-Length"] & io.ReadFull()?
    if e != nil {
        return resp, E_HttpResp
    }
    return buf.Bytes(), nil
}

func MarshalXmlReqBody(c interface{}) (r io.ReadCloser, err Error) {
    buf, e := xml.Marshal(c)
    if e != nil {
        return nil, E_MarshalXML
    }
    return ioutil.NopCloser(bytes.NewReader(buf)), nil
}

func MarshalJson(r interface{}) string {
    _ = fmt.Println
    b, e := json.Marshal(r)
    if e != nil {
        fmt.Println(e)
        return ""
    }
    return string(b)
}

func MarshalXml(r interface{}) string {
    _ = fmt.Println
    b, e := xml.Marshal(r)
    if e != nil {
        fmt.Println(e)
        return ""
    }
    return string(b)
}