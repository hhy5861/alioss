package alioss
import (
    "net/http"
    "time"
)

type Client struct {
    Request         *http.Request       // request
    Client          *http.Client        // http client
    Response        *http.Response      // response, cannot use before call api methods
    AccessKey       string              // access key id
    AccessSecret    string              // access secret
    Location        string              // position, oss-cn-hangzhou, etc.
    // Request.Host = (BucketName + ".") + (Location + ".") + "aliyuncs.com"
    BucketName      string              // bucket
    ObjectName      string              // object, for bucket api, is ""
    Query           map[string]string   // query, for sub-resource
    AuthType        AuthType            // auth type, header, url, sts, none
    // only implement header
}

func NewClient(key, secret, bucket, location string) *Client {
    r, _ := http.NewRequest("GET", "", nil)
    return &Client{
        Request: r,
        Client: &http.Client{},
        Query: map[string]string{},
        AuthType: AUTH_HEADER,
        AccessKey : key,
        AccessSecret: secret,
        BucketName: bucket,
        Location: location,
    }
}

////////////
// all the set/get method is for request
////////////

func (c *Client) SetAuthType(t AuthType) *Client {
    c.AuthType = t
    return c
}

func (c *Client) SetAuthPair(key, secret string) *Client {
    c.AccessKey = key
    c.AccessSecret = secret
    return c
}

func (c *Client) SetDate(t time.Time) *Client {
    c.Request.Header.Set("Date", t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
    return c
}

func (c *Client) SetHeader(k, v string) *Client {
    c.Request.Header.Set(k, v)
    return c
}

func (c *Client) AddHeader(k, v string) *Client {
    c.Request.Header.Add(k, v)
    return c
}

func (c *Client) AddHeaders(h map[string]string) *Client {
    for k, v := range h {
        c.AddHeader(k, v)
    }
    return c
}

func (c *Client) SetHeaders(h map[string]string) *Client {
    for k, v := range h {
        c.SetHeader(k, v)
    }
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
    c.Response = r
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
    if r.StatusCode > 299 {
        return GetReqError(r)
    } else {
        if resp != nil {
            err = UnmarshalXmlResp(r, resp)
        }
        return err
    }
}

func (c *Client) DoBytes(method string, req interface{}) (resp []byte, err Error) {
    if req != nil {
        b, err := MarshalXmlReqBody(req)
        if err != nil {
            return resp, err
        }
        c.Request.Body = b
    }
    r, err := c.Do(method)
    if err != nil {
        return
    }
    defer r.Body.Close()
    if r.StatusCode > 299 {
        err = GetReqError(r)
    } else {
        resp, err = ReadByteResp(r)
    }
    return
}