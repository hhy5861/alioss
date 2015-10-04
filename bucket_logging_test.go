package alioss
import "testing"

func tTestPutBucketLogging(t *testing.T) {
    e := NewTestClient().PutBucketLogging("acrazing", "hello-world-")
    if e != nil {
        t.Error("put bucket logging error:", e)
    }
}

func tTestGetBucketLogging(t *testing.T) {
    r, e := NewTestClient().GetBucketLogging()
    if e != nil {
        t.Error("get bucket logging error: ", e)
    } else {
        t.Log(r.LoggingEnabled)
    }
}

func tTestDeleteBucketLogging(t *testing.T) {
    e := NewTestClient().DeleteBucketLogging()
    if e != nil {
        t.Error("delete bucket logging error:", e)
    }
}