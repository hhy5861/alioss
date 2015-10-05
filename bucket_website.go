package alioss
import "encoding/xml"

type PutBucketWebsiteReq struct {
    XMLName         xml.Name                        `xml:"WebsiteConfiguration"`
    IndexDocument   putBucketWebsiteIndexDocument
    ErrorDocument   putBucketWebsiteErrorDocument
}

type putBucketWebsiteIndexDocument struct {
    Suffix      string
}

type putBucketWebsiteErrorDocument struct {
    Key         string
}

func (c *Client) PutBucketWebsite(indexObject, errorObject string) Error {
    req := &PutBucketWebsiteReq{
        IndexDocument:putBucketWebsiteIndexDocument{indexObject},
        ErrorDocument: putBucketWebsiteErrorDocument{errorObject},
    }
    return c.SetQuery("website", "").DoAll("PUT", nil, req)
}

type GetBucketWebsiteResp PutBucketWebsiteReq

func (c *Client) GetBucketWebsite() (resp *GetBucketWebsiteResp, err Error) {
    resp = &GetBucketWebsiteResp{}
    err = c.SetQuery("website", "").DoAll("GET", resp, nil)
    return
}

func (c *Client) DeleteBucketWebsite() Error {
    return c.SetQuery("website", "").DoAll("Delete", nil, nil)
}