package alioss
import "testing"

func tTestGetBucketLocation(t *testing.T) {
    l, e := NewTestClient().GetBucketLocation()
    if e != nil {
        t.Errorf("get bucket location error: %s\n", e)
    } else {
        t.Logf("location: %s\n", l)
    }
}