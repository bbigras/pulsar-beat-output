// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateExpirationForHITRequest
type UpdateExpirationForHITInput struct {
	_ struct{} `type:"structure"`

	// The date and time at which you want the HIT to expire
	//
	// ExpireAt is a required field
	ExpireAt *time.Time `type:"timestamp" required:"true"`

	// The HIT to update.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateExpirationForHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateExpirationForHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateExpirationForHITInput"}

	if s.ExpireAt == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExpireAt"))
	}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateExpirationForHITResponse
type UpdateExpirationForHITOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateExpirationForHITOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateExpirationForHIT = "UpdateExpirationForHIT"

// UpdateExpirationForHITRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The UpdateExpirationForHIT operation allows you update the expiration time
// of a HIT. If you update it to a time in the past, the HIT will be immediately
// expired.
//
//    // Example sending a request using UpdateExpirationForHITRequest.
//    req := client.UpdateExpirationForHITRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateExpirationForHIT
func (c *Client) UpdateExpirationForHITRequest(input *UpdateExpirationForHITInput) UpdateExpirationForHITRequest {
	op := &aws.Operation{
		Name:       opUpdateExpirationForHIT,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateExpirationForHITInput{}
	}

	req := c.newRequest(op, input, &UpdateExpirationForHITOutput{})
	return UpdateExpirationForHITRequest{Request: req, Input: input, Copy: c.UpdateExpirationForHITRequest}
}

// UpdateExpirationForHITRequest is the request type for the
// UpdateExpirationForHIT API operation.
type UpdateExpirationForHITRequest struct {
	*aws.Request
	Input *UpdateExpirationForHITInput
	Copy  func(*UpdateExpirationForHITInput) UpdateExpirationForHITRequest
}

// Send marshals and sends the UpdateExpirationForHIT API request.
func (r UpdateExpirationForHITRequest) Send(ctx context.Context) (*UpdateExpirationForHITResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateExpirationForHITResponse{
		UpdateExpirationForHITOutput: r.Request.Data.(*UpdateExpirationForHITOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateExpirationForHITResponse is the response type for the
// UpdateExpirationForHIT API operation.
type UpdateExpirationForHITResponse struct {
	*UpdateExpirationForHITOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateExpirationForHIT request.
func (r *UpdateExpirationForHITResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}