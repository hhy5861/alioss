package alioss
import (
    "net/http"
    "time"
    "io"
)

type Client struct {
    Request         *http.Request
    Client          *http.Client
    Body            io.Reader
    AccessKey       string
    AccessSecret    string
    Host            string
    BucketName      string
    ObjectName      string
    Query           map[string]string
    AuthType        AuthType
}

func NewClient(authType AuthType, key, secret string) *Client {
    r, _ := http.NewRequest("GET", "", nil)
    return &Client{
        Request: r,
        Client: &http.Client{},
        Query: map[string]string{},
        AuthType: authType,
        AccessKey : key,
        AccessSecret: secret,
    }
}

func (c *Client) SetAuthPair(key, secret string) *Client {
    c.AccessKey = key
    c.AccessSecret = secret
    return c
}

func (c *Client) SetDate(t time.Time) *Client {
    t = t.Add(-8*3600e9)
    c.Request.Header.Set("Date", t.Format("Mon, 02 Jan 2006 15:04:05 GMT"))
    return c
}

func (c *Client) SetQuery(k, v string) *Client {
    c.Query[k] = v;
    return c
}

func (c *Client) SetQueries(q map[string]string) *Client {
    for k, v := range q {
        c.Query[k] = v
    }
    return c
}

func (c *Client) Do(method string) (*http.Response, Error) {
    c.Request.Method = method
    c.InitAuth()
    if c.InitRequest() != nil {
        return nil, E_InvalidUrl
    }
    r, e := c.Client.Do(c.Request)
    if e != nil {
        return r, E_HttpReq
    }
    return r, nil
}

func (c *Client) DoAll(method string, resp interface{}, req interface{}) (err Error) {
    if req != nil {
        b, err := MarshalXmlReqBody(req)
        if err != nil {
            return err
        }
        c.Request.Body = b
    }
    r, err := c.Do(method)
    if err != nil {
        return
    }
    defer r.Body.Close()
    if r.StatusCode != 200 {
        return GetReqError(r)
    } else {
        if resp != nil {
            err = UnmarshalXmlResp(r, resp)
        }
        return err
    }
}