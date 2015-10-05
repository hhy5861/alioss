package alioss
import "encoding/xml"

type PutBucketReferrerReq struct {
    XMLName             xml.Name        `xml:"RefererConfiguration"`
    AllowEmptyReferrer  bool            `xml:"AllowEmptyReferer"`
    ReferrerList        []string        `xml:"RefererList>Referer"`
}

func (c *Client) PutBucketReferrer(allowEmpty bool, referrerList []string) Error {
    req := &PutBucketReferrerReq{
        AllowEmptyReferrer: allowEmpty,
        ReferrerList      : referrerList,
    }
    return c.SetQuery("referer", "").DoAll("PUT", nil, req)
}

type GetBucketReferrerResp PutBucketReferrerReq

func (c *Client) GetBucketReferrer() (resp *GetBucketReferrerResp, err Error) {
    resp = &GetBucketReferrerResp{}
    err = c.SetQuery("referer", "").DoAll("GET", resp, nil)
    return
}