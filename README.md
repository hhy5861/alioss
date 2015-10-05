# ALIOSS, the SDK for Aliyun OSS with golang

## How to use

```
package main

import (
    "github.com/acrazing/alioss"
    "log"
    "os"
)

func main() {
    client := alioss.NewClient("AccessKeyId", "AccessSecret", "BucketName", "Location")
    file, e := os.Open("./filename.ext")
    if e != nil {
        log.Printf("open file error: %s", e)
    } else {
        stat, e := file.Stat()
        if e != nil {
            log.Printf("get file stat error: %s", e)
        } else {
            err := client.PutObject("objectName.ext", file, stat.Size(), alioss.EmptyStringMap)
            if err != nil {
                log.Printf("put object error: %s", err)
            }
        }
    }
}
```

## TODO

- Multipart Upload
- multipart/form Upload, (useless)