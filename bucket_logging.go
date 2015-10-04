package alioss
import "encoding/xml"

type PutBucketLoggingReq struct {
    XMLName         xml.Name                `xml:"BucketLoggingStatus"`
    LoggingEnabled  putBucketLoggingEnabled
}

type putBucketLoggingEnabled struct {
    TargetBucket    string
    TargetPrefix    string
}