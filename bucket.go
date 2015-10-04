package alioss
import (
    "encoding/xml"
    "time"
)

type PutBucketReq struct {
    XMLName             xml.Name    `xml:"CreateBucketConfiguration"`
    LocationConstraint  string
}

func (c *Client) PutBucket(acl string) (err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    req := PutBucketReq{LocationConstraint: c.Location}
    err = c.DoAll("PUT", nil, req)
    return
}

type GetBucketResp struct {
    XMLName         xml.Name            `xml:"ListBucketResult"`
    Contents        []getBucketContent  `xml:"Contents"`
    CommonPrefixes  string
    Delimiter       string
    EncodingType    string              `xml:"encoding-type"`
    IsTruncated     bool
    Marker          string
    MaxKeys         int
    Name            string
    Owner           getServiceOwner
    Prefix          string
}

type getBucketContent struct {
    Owner           getServiceOwner
    ETag            string
    Key             string
    LastModified    time.Time
    Size            int
    StorageClass    string
}

func (c *Client) GetBucket(limiter map[string]string) (resp *GetBucketResp, err Error) {
    c.SetQueries(limiter)
    resp = &GetBucketResp{}
    err = c.DoAll("GET", resp, nil)
    return
}

func (c *Client) DeleteBucket() (err Error) {
    return c.DoAll("DELETE", nil, nil)
}