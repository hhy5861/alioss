package alioss
import (
    "encoding/xml"
)

type PutBucketLifeCycleReq struct {
    XMLName         xml.Name                `xml:"LifecycleConfiguration"`
    Rules           []BucketLifeCycleRule   `xml:"Rule"`
}

type BucketLifeCycleRule struct {
    ID          string
    Prefix      string
    Status      string
    Expiration  LifeCycleExpiration
}

type LifeCycleExpiration struct {
    Date        *ISO8601Time    `xml:",omitempty"`
    Days        uint            `xml:",omitempty"`
}

func (c *Client) PutBucketLifeCycle(rules []BucketLifeCycleRule) Error {
    req := &PutBucketLifeCycleReq{Rules: rules}
    return c.SetQuery("lifecycle", "").DoAll("PUT", nil, req)
}

func (c *Client) GetBucketLifeCycle() (resp *PutBucketLifeCycleReq, err Error) {
    resp = &PutBucketLifeCycleReq{}
    err = c.SetQuery("lifecycle", "").DoAll("GET", resp, nil)
    return
}

// @Forbidden, this will delete the bucket if is empty
func (c *Client) DeleteBucketLifeCycle() Error {
    c.Location = "oss"
    return c.SetQuery("lifecycle", "").DoAll("DELETE", nil, nil)
}