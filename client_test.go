package alioss

func NewTestClient() *Client {
    c := NewClient(AccessKey, AccessSecret, "acrazing-test", "oss-cn-qingdao")
    return c
}