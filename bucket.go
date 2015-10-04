package alioss
import (
    "encoding/xml"
    "fmt"
)

type PutBucketReq struct {
    XMLName             xml.Name    `xml:"CreateBucketConfiguration"`
    LocationConstraint  string
}

func (c *Client) PutBucket(name, acl, loc string) (resp bool, err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    req := PutBucketReq{LocationConstraint: loc}
    b, err := MarshalXmlReqBody(req)
    if err != nil {
        return false, err
    }
    c.Request.Body = b
    r, err := c.Do("PUT")
    _ = fmt.Println
    fmt.Println(c.Request.Header.Get("Authorization"))
    if err != nil {
        return
    }
    defer r.Body.Close()
    if r.StatusCode != 200 {
        return false, GetReqError(r)
    } else {
        return true, nil
    }
}