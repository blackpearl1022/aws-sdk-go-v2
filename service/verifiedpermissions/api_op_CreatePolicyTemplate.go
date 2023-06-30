// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a policy template. A template can use placeholders for the principal
// and resource. A template must be instantiated into a policy by associating it
// with specific principals and resources to use for the placeholders. That
// instantiated policy can then be considered in authorization decisions. The
// instantiated policy works identically to any other policy, except that it is
// dynamically linked to the template. If the template changes, then any policies
// that are linked to that template are immediately updated as well.
func (c *Client) CreatePolicyTemplate(ctx context.Context, params *CreatePolicyTemplateInput, optFns ...func(*Options)) (*CreatePolicyTemplateOutput, error) {
	if params == nil {
		params = &CreatePolicyTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicyTemplate", params, optFns, c.addOperationCreatePolicyTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyTemplateInput struct {

	// The ID of the policy store in which to create the policy template.
	//
	// This member is required.
	PolicyStoreId *string

	// Specifies the content that you want to use for the new policy template, written
	// in the Cedar policy language.
	//
	// This member is required.
	Statement *string

	// Specifies a unique, case-sensitive ID that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	// Specifies a description for the policy template.
	Description *string

	noSmithyDocumentSerde
}

type CreatePolicyTemplateOutput struct {

	// The date and time the policy template was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time the policy template was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The ID of the policy store that contains the policy template.
	//
	// This member is required.
	PolicyStoreId *string

	// The unique ID of the new policy template.
	//
	// This member is required.
	PolicyTemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreatePolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreatePolicyTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreatePolicyTemplateMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePolicyTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicyTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreatePolicyTemplate struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePolicyTemplate) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePolicyTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePolicyTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePolicyTemplateInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePolicyTemplateMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePolicyTemplate{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePolicyTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "verifiedpermissions",
		OperationName: "CreatePolicyTemplate",
	}
}
