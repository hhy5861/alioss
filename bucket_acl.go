package alioss

func (c *Client) PutBucketAcl(name, acl, loc string) (err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    c.SetQuery("acl", "")
    return c.DoAll("PUT", nil, nil)
}