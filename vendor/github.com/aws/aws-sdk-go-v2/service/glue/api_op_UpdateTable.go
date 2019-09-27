// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTableRequest
type UpdateTableInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the table resides. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// By default, UpdateTable always creates an archived version of the table before
	// updating it. However, if skipArchive is set to true, UpdateTable does not
	// create the archived version.
	SkipArchive *bool `type:"boolean"`

	// An updated TableInput object to define the metadata table in the catalog.
	//
	// TableInput is a required field
	TableInput *TableInput `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTableInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.TableInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableInput"))
	}
	if s.TableInput != nil {
		if err := s.TableInput.Validate(); err != nil {
			invalidParams.AddNested("TableInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTableResponse
type UpdateTableOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTable = "UpdateTable"

// UpdateTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Updates a metadata table in the Data Catalog.
//
//    // Example sending a request using UpdateTableRequest.
//    req := client.UpdateTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateTable
func (c *Client) UpdateTableRequest(input *UpdateTableInput) UpdateTableRequest {
	op := &aws.Operation{
		Name:       opUpdateTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTableInput{}
	}

	req := c.newRequest(op, input, &UpdateTableOutput{})
	return UpdateTableRequest{Request: req, Input: input, Copy: c.UpdateTableRequest}
}

// UpdateTableRequest is the request type for the
// UpdateTable API operation.
type UpdateTableRequest struct {
	*aws.Request
	Input *UpdateTableInput
	Copy  func(*UpdateTableInput) UpdateTableRequest
}

// Send marshals and sends the UpdateTable API request.
func (r UpdateTableRequest) Send(ctx context.Context) (*UpdateTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTableResponse{
		UpdateTableOutput: r.Request.Data.(*UpdateTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTableResponse is the response type for the
// UpdateTable API operation.
type UpdateTableResponse struct {
	*UpdateTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTable request.
func (r *UpdateTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}