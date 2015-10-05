package alioss
import "encoding/xml"

type PutBucketLoggingReq struct {
    XMLName         xml.Name                `xml:"BucketLoggingStatus"`
    LoggingEnabled  putBucketLoggingEnabled
}

type putBucketLoggingEnabled struct {
    TargetBucket    string
    TargetPrefix    string
}

func (c *Client) PutBucketLogging(bucket, prefix string) Error {
    req := &PutBucketLoggingReq{LoggingEnabled: putBucketLoggingEnabled{bucket, prefix}}
    return c.SetQuery("logging", "").DoAll("PUT", nil, req)
}

type GetBucketLoggingResp PutBucketLoggingReq

func (c *Client) GetBucketLogging() (resp *GetBucketLoggingResp, err Error) {
    resp = &GetBucketLoggingResp{}
    err = c.SetQuery("logging", "").DoAll("GET", resp, nil)
    return
}

func (c *Client) DeleteBucketLogging() Error {
    return c.SetQuery("logging", "").DoAll("DELETE", nil, nil)
}