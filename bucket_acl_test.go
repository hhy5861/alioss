package alioss
import "testing"

func tTestPutBucketAcl(t *testing.T) {
    c := NewTestClient()
    err := c.PutBucketAcl("acrazing-test22", "private1", "oss-cn-qingdao")
    if err != nil {
        t.Errorf("put bucket acl error: %s\n", err.Error())
    }
}

func tTestGetBucketAcl(t *testing.T) {
    c := NewTestClient()
    r, e := c.GetBucketAcl("acrazing-test", "oss-cn-qingdao")
    if e != nil {
        t.Errorf("get bucket acl error %s\n", e.Error())
    } else {
        t.Log(r.AccessControlList)
    }
}