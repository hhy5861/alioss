package alioss
import "encoding/xml"

func (c *Client) PutBucketAcl(acl string) (err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.SetQuery("acl", "")
    return c.DoAll("PUT", nil, nil)
}

type GetBucketAclResp struct {
    XMLName             xml.Name        `xml:"AccessControlPolicy"`
    Owner               getServiceOwner
    AccessControlList   []string        `xml:"AccessControlList>Grant"`
}

func (c *Client) GetBucketAcl() (resp *GetBucketAclResp,err Error) {
    c.SetQuery("acl", "")
    resp = &GetBucketAclResp{}
    err = c.DoAll("GET", resp, nil)
    return
}