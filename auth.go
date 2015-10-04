package alioss
import (
    "time"
    "strings"
    "sort"
    "encoding/base64"
    "crypto/sha1"
    "crypto/hmac"
    "strconv"
    "net/url"
)

type AuthType byte

const (
    AUTH_HEADER AuthType = iota
    AUTH_URL
    AUTH_STS
    AUTH_NONE
)

func (c *Client) InitAuth() *Client {
    switch c.AuthType {
        case AUTH_HEADER:
        c.HeaderAuth()
        case AUTH_URL:
        c.UrlAuth()     // TODO: test not pass
        case AUTH_STS:
        c.StsAuth()
    }
    return c
}

func (c *Client) HeaderAuth() *Client {
    VERB := c.Request.Method + "\n"
    md5 := c.Request.Header.Get("Content-Md5") + "\n"
    cte := c.Request.Header.Get("Content-Type") + "\n"
    if c.Request.Header.Get("Date") == "" {
        c.SetDate(time.Now())
    }
    date := c.Request.Header.Get("Date") + "\n"
    headers := c.CanonicalizedOSSHeaders()
    resources := c.CanonicalizedResource()

    msg := VERB + md5 + cte + date + headers + resources
    auth := "OSS " + c.AccessKey + ":" + c.HmacSha1(msg)
    c.Request.Header.Set("Authorization", auth)
    return c
}

func (c *Client) UrlAuth() *Client {
    VERB := c.Request.Method + "\n"
    md5 := c.Request.Header.Get("Content-Md5") + "\n"
    cte := c.Request.Header.Get("Content-Type") + "\n"
    if _, ok := c.Query["Expires"]; !ok {
        c.SetQuery("Expires", strconv.FormatInt(time.Now().Add(3600e9).Unix(), 10))
    }
    exs := c.Query["Expires"] + "\n"
    headers := c.CanonicalizedOSSHeaders()
    resources := c.CanonicalizedResource()

    msg := VERB + md5 + cte + exs + headers + resources
    c.SetQuery("Signature", url.QueryEscape(c.HmacSha1(msg)))
    c.SetQuery("OSSAccessKeyId", c.AccessKey)
    return c
}

func (c *Client) StsAuth() *Client {
    // TODO
    return c
}

func (c *Client) HmacSha1(msg string) string {
    skey := []byte(c.AccessSecret)
    smsg := []byte(msg)
    h := hmac.New(sha1.New, skey)
    h.Write(smsg)
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *Client) CanonicalizedOSSHeaders() string {
    var t sort.StringSlice
    for k, v := range c.Request.Header {
        k = strings.ToLower(k)
        if strings.HasPrefix(k, "x-oss-") {
            t = append(t, k + ":" + strings.Join(v, ",") + "\n")
        }
    }
    t.Sort()
    return strings.Join(t, "")
}

var SubResources map[string]bool = map[string]bool{
    "acl": true,
    "uploadId": true,
    "partNumber": true,
    "uploads": true,
    "logging": true,
    "website": true,
    "location": true,
    "lifecycle": true,
    "referer": true,
    "cors": true,
    "delete": true,
    "append": true,
    "position": true,
    "security-token": true,
}

func (c *Client) CanonicalizedResource() string {
    t := "/"
    if c.BucketName != "" {
        t += c.BucketName + "/"
    }
    if c.ObjectName != "" {
        t += c.ObjectName
    }
    var sub, query sort.StringSlice
    var s bool
    for k, v := range c.Query {
        s, _ = SubResources[k]
        if s == true {
            if v != "" {
                k += "=" + v
            }
            sub = append(sub, k)
        }
        if strings.HasPrefix(k, "response-") {
            if v != "" {
                k += "=" + v
            }
            query = append(query, k)
        }
    }
    sub.Sort()
    query.Sort()
    if len(sub) > 0 || len(query) > 0 {
        t += "?"
        if len(sub) > 0 {
            t += strings.Join(sub, "&")
        }
        if len(query) > 0 {
            if len(sub) > 0 {
                t += "&"
            }
            t += strings.Join(query, "&")
        }
    }
    return t
}