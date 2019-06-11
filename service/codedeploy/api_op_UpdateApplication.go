// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of an UpdateApplication operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/UpdateApplicationInput
type UpdateApplicationInput struct {
	_ struct{} `type:"structure"`

	// The current name of the application you want to change.
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string"`

	// The new name to give the application.
	NewApplicationName *string `locationName:"newApplicationName" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.NewApplicationName != nil && len(*s.NewApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/UpdateApplicationOutput
type UpdateApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApplication = "UpdateApplication"

// UpdateApplicationRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Changes the name of an application.
//
//    // Example sending a request using UpdateApplicationRequest.
//    req := client.UpdateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/UpdateApplication
func (c *Client) UpdateApplicationRequest(input *UpdateApplicationInput) UpdateApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateApplicationOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateApplicationRequest{Request: req, Input: input, Copy: c.UpdateApplicationRequest}
}

// UpdateApplicationRequest is the request type for the
// UpdateApplication API operation.
type UpdateApplicationRequest struct {
	*aws.Request
	Input *UpdateApplicationInput
	Copy  func(*UpdateApplicationInput) UpdateApplicationRequest
}

// Send marshals and sends the UpdateApplication API request.
func (r UpdateApplicationRequest) Send(ctx context.Context) (*UpdateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationResponse{
		UpdateApplicationOutput: r.Request.Data.(*UpdateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationResponse is the response type for the
// UpdateApplication API operation.
type UpdateApplicationResponse struct {
	*UpdateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplication request.
func (r *UpdateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}