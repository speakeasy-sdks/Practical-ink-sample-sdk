# TechPort SDK


## Overview

TechPort: TechPort RESTful API

### Available Operations

* [GetAPI](#getapi) - Returns the swagger specification for the API.
* [GetAPIProjectsIDFormat](#getapiprojectsidformat) - Returns information about a specific technology project.
* [GetAPIProjectsFormat](#getapiprojectsformat) - Returns a list of available technology project IDs.

## GetAPI

Returns the swagger specification for the API.

### Example Usage

```go
package main

import(
	practicalinksamplesdk "github.com/speakeasy-sdks/Practical-ink-sample-sdk"
	"context"
	"log"
	"net/http"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAPIResponse](../../models/operations/getapiresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIProjectsIDFormat

Returns information about a specific technology project.

### Example Usage

```go
package main

import(
	practicalinksamplesdk "github.com/speakeasy-sdks/Practical-ink-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/Practical-ink-sample-sdk/models/operations"
	"log"
)

func main() {
    s := practicalinksamplesdk.New()

    ctx := context.Background()
    res, err := s.GetAPIProjectsIDFormat(ctx, operations.GetAPIProjectsIDFormatRequest{
        DotFormat: "string",
        Format: operations.FormatJSON,
        ID: 355611,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Project != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAPIProjectsIDFormatRequest](../../models/operations/getapiprojectsidformatrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetAPIProjectsIDFormatResponse](../../models/operations/getapiprojectsidformatresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIProjectsFormat

Returns a list of available technology project IDs.

### Example Usage

```go
package main

import(
	practicalinksamplesdk "github.com/speakeasy-sdks/Practical-ink-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/Practical-ink-sample-sdk/models/operations"
	"github.com/speakeasy-sdks/Practical-ink-sample-sdk/types"
	"log"
)

func main() {
    s := practicalinksamplesdk.New()

    ctx := context.Background()
    res, err := s.GetAPIProjectsFormat(ctx, operations.GetAPIProjectsFormatRequest{
        DotFormat: "string",
        Format: operations.QueryParamFormatXML,
        UpdatedSince: types.MustDateFromString("2021-05-09"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAPIProjectsFormatRequest](../../models/operations/getapiprojectsformatrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAPIProjectsFormatResponse](../../models/operations/getapiprojectsformatresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
