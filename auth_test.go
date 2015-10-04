package alioss
import "testing"

func tTestHeaderAuth(t *testing.T) {
    c := NewTestClient()
    c.SetAuthPair("44CF9590006BF252F707", "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV")
    c.Request.Header.Set("Content-Md5", "ODBGOERFMDMzQTczRUY3NUE3NzA5QzdFNUYzMDQxNEM=")
    c.Request.Header.Set("Content-Type", "text/html")
    c.BucketName = "oss-example"
    c.ObjectName = "nelson"
    c.Request.Method = "PUT"
    c.Request.Header.Set("X-OSS-Meta-Author", "foo@bar.com")
    c.Request.Header.Set("X-OSS-Magic", "abracadabra")
    c.Request.Header.Set("Date", "Thu, 17 Nov 2005 18:49:58 GMT")
    c.InitAuth()
    sig := c.Request.Header.Get("Authorization")
    if sig != "OSS 44CF9590006BF252F707:26NBxoKdsyly4EDv6inkoDft/yA=" {
        t.Errorf("auth signature error as %s", sig)
    }
}

func tTestUrlAuth(t *testing.T) {
    c := NewTestClient()
    c.SetAuthPair("44CF9590006BF252F707", "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV")
    c.Request.Method = "GET"
    c.SetQuery("Expires", "1141889120")
    c.BucketName = "oss-example"
    c.ObjectName = "oss-api.pdf"
    c.InitAuth()
    sig := c.Query["Signature"]
    if sig != "vjbyPxybdZaNmGa%2ByT272YEAiv4%3D" {
        t.Errorf("auth signature error as %s", sig)
    }
}