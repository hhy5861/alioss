package alioss

func (c *Client) PutBucketAcl(name, acl, loc string) (resp bool, err Error) {
    c.Request.Header.Set("x-oss-acl", acl)
    c.BucketName = name
    c.Host = loc + ".aliyuncs.com"
    c.SetQuery("acl", "")
    r, err := c.Do("PUT")
    if err != nil {
        return
    }
    defer r.Body.Close()
    if r.StatusCode != 200 {
        return false, GetReqError(r)
    } else {
        return true, nil
    }
}