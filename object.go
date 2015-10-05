package alioss
import (
    "io"
    "io/ioutil"
    "strconv"
    "encoding/xml"
    "time"
)

func (c *Client) PutObject(object string, body io.Reader, size int64, headers map[string]string) Error {
    c.ObjectName = object
    b, ok := body.(io.ReadCloser);
    if !ok {
        b = ioutil.NopCloser(body)
    }
    c.Request.Body = b
    if _, ok := headers["Content-Length"]; !ok {
        headers["Content-Length"] = strconv.FormatInt(size, 10)
    }
    c.SetHeaders(headers)
    return c.DoAll("PUT", nil, nil)
}

type CopyObjectResp struct {
    XMLName         xml.Name        `xml:"CopyObjectResult"`
    LastModified    time.Time
    ETag            string
}

func (c *Client) CopyObject(source, target string, headers map[string]string) (resp *CopyObjectResp, err Error) {
    headers["x-oss-copy-source"] = source
    c.SetHeaders(headers)
    c.ObjectName = target
    resp = &CopyObjectResp{}
    err = c.DoAll("PUT", resp, nil)
    return
}

func (c *Client) GetObject(object string, headers map[string]string, override map[string]string) (resp []byte, err Error) {
    c.ObjectName = object
    c.SetHeaders(headers)
    c.SetQueries(override)
    resp, err = c.DoBytes("GET", nil)
    return
}

type AppendObjectRespHeader struct {
    HashCrc64ecma       int64
    NextAppendPosition  int64
}

func (c *Client) AppendObject(object string, position int64, body io.Reader, size int64, headers map[string]string) (resp *AppendObjectRespHeader, err Error) {
    c.ObjectName = object
    c.SetQuery("append", "").SetQuery("position", strconv.FormatInt(position, 10))
    b, ok := body.(io.ReadCloser);
    if !ok {
        b = ioutil.NopCloser(body)
    }
    c.Request.Body = b
    if _, ok := headers["Content-Length"]; !ok {
        headers["Content-Length"] = strconv.FormatInt(size, 10)
    }
    c.SetHeaders(headers)
    err = c.DoAll("POST", nil, nil)
    if err != nil {
        return (*AppendObjectRespHeader)(nil), err
    }
    resp = &AppendObjectRespHeader{
        HashCrc64ecma: c.Response.Header.Get("x-oss-hash-crc64ecma"),
        NextAppendPosition: c.Response.Header.Get("x-oss-next-append-position"),
    }
    return resp, nil
}

func (c *Client) DeleteObject(object string) Error {
    c.ObjectName = object
    return c.DoAll("DELETE", nil, nil)
}

type DeleteMultiObjectReq struct {
    XMLName         xml.Name        `xml:"Delete"`
    Quiet           bool
    Objects         []string        `xml:"Object>Key"`
    EncodingType    string          `xml:"encoding-type"`
}

type DeleteMultiObjectResp struct {
    XMLName         xml.Name        `xml:"DeleteResult"`
    Deleted         []string        `xml:"Deleted>Key"`
    EncodingType    string          `xml:"encoding-type"`
}

func (c *Client) DeleteMultiObject(objects []string, quiet bool, encodingType string) (resp *DeleteMultiObjectResp, err Error) {
    req := &DeleteMultiObjectReq{
        Quiet: quiet,
        Objects: objects,
        EncodingType: encodingType,
    }
    resp = &DeleteMultiObjectResp{}
    c.SetQuery("delete", "")
    err = c.DoAll("POST", resp, req)
    return
}

// please use c.Response.Header.Get to get elements
func (c *Client) HeadObject(object string, headers map[string]string) Error {
    c.SetHeaders(headers)
    c.ObjectName = object
    return c.DoAll("HEAD", nil, nil)
}

func (c *Client) PostObject(object, body io.Reader, formData map[string]string) Error {
    // TODO
    return nil
}