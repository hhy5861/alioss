package alioss
import "testing"

func tTestPutBucketAcl(t *testing.T) {
    c := NewTestClient()
    err := c.PutBucketAcl("public-read")
    if err != nil {
        t.Errorf("put bucket acl error: %s\n", err.Error())
    }
}

func tTestGetBucketAcl(t *testing.T) {
    c := NewTestClient()
    r, e := c.GetBucketAcl()
    if e != nil {
        t.Errorf("get bucket acl error %s\n", e.Error())
    } else {
        t.Log(r.AccessControlList)
    }
}