package alioss
import (
    "encoding/xml"
    "time"
)

type PutBucketReq struct {
    XMLName             xml.Name    `xml:"CreateBucketConfiguration"`
    LocationConstraint  string
}

func (c *Client) PutBucket(name, acl, loc string) (err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    req := PutBucketReq{LocationConstraint: loc}
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

func (c *Client) GetBucket(name, loc string, limiter map[string]string) (resp *GetBucketResp, err Error) {
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    c.SetQueries(limiter)
    resp = &GetBucketResp{}
    err = c.DoAll("GET", resp, nil)
    return
}