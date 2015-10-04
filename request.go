package alioss
import (
    "net/url"
    "strings"
)

func (c *Client) InitRequest() error {
    u, err := c.GetUrl()
    c.Request.URL = u
    if err == nil {
        c.Request.Host = u.Host
    }
    c.Request.Header.Add("User-Agent", "HTTP/1.1 Client")
    return err
}

func (c *Client) GetUrl() (u *url.URL, err error) {
    s := "http://"
    if c.BucketName != "" {
        s += c.BucketName + "."
    }
    if c.Location != "" {
        s += c.Location + "."
    }
    s += "aliyuncs.com/" + c.ObjectName
    if len(c.Query) > 0 {
        s += "?"
        q := []string{}
        for k, v := range c.Query {
            if v != "" {
                k += "=" + v
            }
            q = append(q, k)
        }
        s += strings.Join(q, "&")
    }
    u, err = url.Parse(s)
    return u, err
}