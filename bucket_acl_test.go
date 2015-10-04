package alioss
import "testing"

func tTestPutBucketAcl(t *testing.T) {
    c := NewTestClient()
    err := c.PutBucketAcl("acrazing-test22", "private1", "oss-cn-qingdao")
    if err != nil {
        t.Errorf("put bucket acl error: %s\n", err.Error())
    }
}