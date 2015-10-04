package alioss
import (
    "testing"
    "encoding/xml"
)

func tTestPutLifeCycleDays(t *testing.T) {
    r := BucketLifeCycleRule{
        "",
        "",
        "Enabled",
        LifeCycleExpiration{Days: 5},
    }
    req := []BucketLifeCycleRule{r}
    buf, _ := xml.Marshal(req)
    t.Log("req: ", string(buf), "\n")
    t.Log("date", r.Expiration.Date)
    e := NewTestClient().PutBucketLifeCycle(req)
    if e != nil {
        t.Errorf("put bucket life cycle days error: %s\n", e)
    }
}

func tTestGetLifeCycle(t *testing.T) {
    r, e := NewTestClient().GetBucketLifeCycle()
    if e != nil {
        t.Errorf("get life cycle error: %s\n", e)
    } else {
        t.Log(r.Rules[0].Expiration, r.Rules[0].ID, "\n")
    }
}

func tTestDeleteLifeCycle(t *testing.T) {
    c := NewTestClient()
    c.SetQuery("lifecycle", "")
    c.InitAuth()
    c.InitRequest()
    t.Log("url: ", c.Request.URL.String())
    e := c.DeleteBucketLifeCycle()
    if e != nil {
        t.Errorf("delete life cycle error: %s\n", e)
    } else {
        t.Log("delete life cycle done\n")
    }
}