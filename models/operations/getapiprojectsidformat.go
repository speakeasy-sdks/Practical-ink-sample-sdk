// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/Practical-ink-sample-sdk/internal/utils"
	"github.com/speakeasy-sdks/Practical-ink-sample-sdk/models/components"
	"net/http"
)

// Format - The response type desired.
type Format string

const (
	FormatJSON Format = "json"
	FormatXML  Format = "xml"
)

func (e Format) ToPointer() *Format {
	return &e
}

func (e *Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "xml":
		*e = Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Format: %v", v)
	}
}

type GetAPIProjectsIDFormatRequest struct {
	// Automatically added
	DotFormat string `pathParam:"style=simple,explode=false,name=.format"`
	// The response type desired.
	Format Format `default:"xml" queryParam:"style=form,explode=true,name=format"`
	// ID of project to fetch
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (g GetAPIProjectsIDFormatRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAPIProjectsIDFormatRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAPIProjectsIDFormatRequest) GetDotFormat() string {
	if o == nil {
		return ""
	}
	return o.DotFormat
}

func (o *GetAPIProjectsIDFormatRequest) GetFormat() Format {
	if o == nil {
		return Format("")
	}
	return o.Format
}

func (o *GetAPIProjectsIDFormatRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAPIProjectsIDFormatResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	Project *components.Project
}

func (o *GetAPIProjectsIDFormatResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetAPIProjectsIDFormatResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIProjectsIDFormatResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIProjectsIDFormatResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPIProjectsIDFormatResponse) GetProject() *components.Project {
	if o == nil {
		return nil
	}
	return o.Project
}
