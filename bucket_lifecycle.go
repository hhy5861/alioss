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

type GetBucketLifeCycleResp PutBucketLifeCycleReq

func (c *Client) GetBucketLifeCycle() (resp *GetBucketLifeCycleResp, err Error) {
    resp = &GetBucketLifeCycleResp{}
    err = c.SetQuery("lifecycle", "").DoAll("GET", resp, nil)
    return
}

// @Forbidden, this will delete the bucket if is empty
func (c *Client) DeleteBucketLifeCycle() Error {
    c.Location = "oss"
    return c.SetQuery("lifecycle", "").DoAll("DELETE", nil, nil)
}