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
    err := c.PutBucket("public-read")
    if err != nil {
        t.Errorf("put bucket error: %s\n", err.Error())
    }
}

func tTestGetBucket(t *testing.T) {
    c := NewTestClient()
    r, e := c.GetBucket(EmptyStringMap)
    if e != nil {
        t.Errorf("get bucket error %s\n", e.Error())
    }
    t.Log(r.Name)
}

func tTestDeleteBucket(t *testing.T) {
    e := NewTestClient().DeleteBucket()
    if e != nil {
        t.Errorf("delete bucket error %s\n", e)
    }
}