package alioss
import "testing"

func tTestPutBucketReferrer(t *testing.T) {
    e := NewTestClient().PutBucketReferrer(true, []string{"acrazing.me"})
    if e != nil {
        t.Errorf("put bucket referrer error: %s\n", e)
    }
}

func tTestGetBucketReferrer(t *testing.T) {
    r, e := NewTestClient().GetBucketReferrer()
    if e != nil {
        t.Errorf("get bucket referrer error: %s\n", e)
    } else {
        t.Log(r.AllowEmptyReferrer, r.ReferrerList, "\n")
    }
}