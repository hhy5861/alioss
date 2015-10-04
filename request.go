package alioss
import (
    "net/url"
    "strings"
)

func (c *Client) InitRequest() error {
    h, u, err := c.GetUrl()
    c.Request.Host = h
    c.Request.URL = u
    return err
}

func (c *Client) GetUrl() (h string, u *url.URL, err error) {
    s := "http://"
    if c.BucketName != "" {
        h += c.BucketName + "."
    }
    h += c.Host
    s += h + "/" + c.ObjectName
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
    return h, u, err
}