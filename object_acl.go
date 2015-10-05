package alioss

func (c *Client) PubObjectAcl(object, acl string) Error {
    c.ObjectName = object
    c.SetQuery("acl", "")
    c.SetHeader("x-oss-object-acl", acl)
    return c.DoAll("PUT", nil, nil)
}

type GetObjectAclResp GetBucketAclResp

func (c *Client) GetObjectAcl(object string) (resp *GetObjectAclResp,err Error) {
    c.ObjectName = object
    c.SetQuery("acl", "")
    resp = &GetObjectAclResp{}
    err = c.DoAll("GET", resp, nil)
    return
}