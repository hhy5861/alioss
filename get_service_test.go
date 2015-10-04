package alioss
import (
    "testing"
)

func tTestGetService(t *testing.T) {
    c := NewTestClient()
    r, e := c.GetService()
    if e != nil {
        t.Errorf("request error with %s", e.Error())
    } else {
        t.Log(r.Buckets)
    }
}