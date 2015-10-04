package alioss
import (
    "testing"
)

func tTestGetService(t *testing.T) {
    r, e := NewTestClient().GetService()
    if e != nil {
        t.Errorf("request error with %s", e.Error())
    } else {
        t.Log(r.Buckets)
    }
}