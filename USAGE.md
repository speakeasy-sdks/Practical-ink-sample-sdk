<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	practicalinksamplesdk "github.com/speakeasy-sdks/Practical-ink-sample-sdk"
	"log"
)

func main() {
	s := practicalinksamplesdk.New()

	ctx := context.Background()
	res, err := s.GetAPI(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->