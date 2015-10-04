package alioss
import "testing"

func tTestMarshalXML(t *testing.T) {
    c := PutBucketReq{LocationConstraint: "hello world"}
    MarshalXmlReqBody(c)
}

func tTestPutBucket(t *testing.T) {
    c := NewTestClient()
    _, err := c.PutBucket("acrazing-test", "public-read", "oss-cn-qingdao")
    if err != nil {
        t.Errorf("put bucket error: %s\n", err.Error())
    }
}