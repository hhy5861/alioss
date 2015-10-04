package alioss
import (
    "bytes"
    "net/http"
    "io"
    "encoding/xml"
    "fmt"
    "io/ioutil"
)

func UnmarshalXmlResp(r *http.Response, c interface{}) Error {
    buf := bytes.NewBuffer([]byte{})
    _, e := io.Copy(buf, r.Body)
    if e != nil {
        return E_HttpResp
    }

    _ = fmt.Println
    // fmt.Println(string(buf.Bytes()))
    e = xml.Unmarshal(buf.Bytes(), c)
    if e != nil {
        return E_ParseXML
    }
    return nil
}

func MarshalXmlReqBody(c interface{}) (r io.ReadCloser, err Error) {
    buf, e := xml.Marshal(c)
    if e != nil {
        return nil, E_MarshalXML
    }
    return ioutil.NopCloser(bytes.NewReader(buf)), nil
}