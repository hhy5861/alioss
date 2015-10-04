package alioss
import "testing"

func tTestPutBucketWebsite(t *testing.T) {
    e := NewTestClient().PutBucketWebsite("", "")
    if e != nil {
        t.Errorf("put bucket website error: %s\n", e.Error())
    }
}

func tTestGetBucketWebsite(t *testing.T) {
    r, e := NewTestClient().GetBucketWebsite()
    if e != nil {
        t.Errorf("get bucket website error: %s\n", e.Error())
    } else {
        t.Log(r.ErrorDocument, r.IndexDocument, "\n")
    }
}

func tTestDeleteBucketWebsite(t *testing.T) {
    e := NewTestClient().DeleteBucket()
    if e != nil {
        t.Errorf("delete bucket website error: %s\n", e.Error())
    }
}