package alioss
import "encoding/xml"

func (c *Client) PutBucketAcl(name, acl, loc string) (err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    c.SetQuery("acl", "")
    return c.DoAll("PUT", nil, nil)
}

type GetBucketAclResp struct {
    XMLName             xml.Name        `xml:"AccessControlPolicy"`
    Owner               getServiceOwner
    AccessControlList   []string        `xml:"AccessControlList>Grant"`
}

func (c *Client) GetBucketAcl(name, loc string) (resp *GetBucketAclResp,err Error) {
    c.SetQuery("acl", "")
    resp = &GetBucketAclResp{}
    err = c.SetBucketHost(name, loc).DoAll("GET", resp, nil)
    return
}