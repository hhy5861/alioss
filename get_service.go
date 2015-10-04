package alioss
import (
    "encoding/xml"
    "time"
)

type GetServiceResp struct {
    XMLName     xml.Name    `xml:"ListAllMyBucketsResult"`
    Prefix      string
    Marker      string
    MaxKeys     int
    IsTruncated bool
    NextMarker  string
    Owner       getServiceOwner
    Buckets     []getServiceBucket      `xml:"Buckets>Bucket"`
}

type getServiceOwner struct {
    ID          string
    DisplayName string
}

type getServiceBucket struct {
    Name            string
    CreationDate    time.Time
    Location        string
}

func (c *Client) GetService() (resp *GetServiceResp, err Error) {
    c.Host = "oss.aliyuncs.com"
    r, err := c.Do("GET")
    if err != nil {
        return
    }
    defer r.Body.Close()
    if r.StatusCode != 200 {
        return resp, GetReqError(r)
    } else {
        resp = &GetServiceResp{}
        err = UnmarshalXmlResp(r, resp)
        return resp, err
    }
}