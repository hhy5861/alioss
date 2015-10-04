package alioss
import (
    "net/http"
    "strconv"
    "encoding/xml"
)

type Error interface {
    GetStatus() int
    GetCode() string
    GetMessage() string
    GetRequestId() string
    GetHostId() string
    Error() string
}
type errors struct {
    XMLName     xml.Name    `xml:"Error"`
    Status      int
    Code        string
    Message     string
    RequestId   string
    HostId      string
}

func (e errors) GetStatus() int {
    return e.Status
}

func (e errors) GetCode() string {
    return e.Code
}

func (e errors) GetMessage() string {
    return e.Message
}

func (e errors) GetRequestId() string {
    return e.RequestId
}

func (e errors) GetHostId() string {
    return e.HostId
}

func (e errors) Error() string {
    return strconv.Itoa(e.Status) + " " + e.Code + ": " + e.Message
}

var (
    E_InvalidUrl    Error = NewError("InvalidUrl")
    E_HttpResp      Error = NewError("HttpResponseParseError")
    E_HttpReq       Error = NewError("HttpRequestError")
    E_ParseXML      Error = NewError("ParseXML")
    E_MarshalXML    Error = NewError("MarshalXMLError")
)

func NewError(c string) errors {
    return errors{Code: c}
}

const NO_CONTENT string = "204 No Content"

func GetReqError(r *http.Response) Error {
    a := NewError("")
    e := UnmarshalXmlResp(r, &a)
    if e != nil {
        return e
    } else {
        a.Status = r.StatusCode
        return a
    }
}