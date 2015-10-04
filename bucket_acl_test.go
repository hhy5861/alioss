package alioss
import "testing"

func tTestPutBucketAcl(t *testing.T) {
    c := NewTestClient()
    _, err := c.PutBucketAcl("acrazing-test", "private", "oss-cn-qingdao")
    t.Log(c.Request.URL.String())
    t.Log(c.CanonicalizedResource())
    if err != nil {
        t.Errorf("put bucket acl error: %s\n", err.Error())
    }
}