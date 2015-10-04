package alioss
import (
    "testing"
)

func tTestMarshalXML(t *testing.T) {
    c := PutBucketReq{LocationConstraint: "hello world"}
    MarshalXmlReqBody(c)
}

func tTestPutBucket(t *testing.T) {
    c := NewTestClient()
    err := c.PutBucket("acrazing-test", "public-read", "oss-cn-qingdao")
    if err != nil {
        t.Errorf("put bucket error: %s\n", err.Error())
    }
}

func tTestGetBucket(t *testing.T) {
    c := NewTestClient()
    r, e := c.GetBucket("acrazing-test", "oss-cn-qingdao", EmptyStringMap)
    if e != nil {
        t.Errorf("get bucket error %s\n", e.Error())
    }
    t.Log(MarshalJson(r))
}

func tTestDeleteBucket(t *testing.T) {
    e := NewTestClient().DeleteBucket("acrazing-test2", "oss-cn-qingdao")
    if e != nil {
        t.Errorf("delete bucket error %s\n", e)
    }
}