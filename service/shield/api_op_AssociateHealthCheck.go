// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds health-based detection to the Shield Advanced protection for a resource.
// Shield Advanced health-based detection uses the health of your AWS resource to
// improve responsiveness and accuracy in attack detection and mitigation. You
// define the health check in Route 53 and then associate it with your Shield
// Advanced protection. For more information, see Shield Advanced Health-Based
// Detection
// (https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option)
// in the AWS WAF and AWS Shield Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) AssociateHealthCheck(ctx context.Context, params *AssociateHealthCheckInput, optFns ...func(*Options)) (*AssociateHealthCheckOutput, error) {
	if params == nil {
		params = &AssociateHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateHealthCheck", params, optFns, addOperationAssociateHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateHealthCheckInput struct {

	// The Amazon Resource Name (ARN) of the health check to associate with the
	// protection.
	//
	// This member is required.
	HealthCheckArn *string

	// The unique identifier (ID) for the Protection object to add the health check
	// association to.
	//
	// This member is required.
	ProtectionId *string
}

type AssociateHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpAssociateHealthCheckValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateHealthCheck(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateHealthCheck(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateHealthCheck",
	}
}