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
    resp = &GetServiceResp{}
    err = c.DoAll("GET", resp, nil)
    return
}