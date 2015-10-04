package alioss
import "encoding/xml"

type GetBucketLocation struct {
    string
    XMLName         xml.Name        `xml:"LocationConstraint"`
}

type LocationConstraint []byte

func (c *Client) GetBucketLocation() (loc string, err Error) {
    resp := &LocationConstraint{}
    err = c.SetQuery("location", "").DoAll("GET", resp, nil)
    if err == nil {
        loc = string(*resp)
    }
    return
}