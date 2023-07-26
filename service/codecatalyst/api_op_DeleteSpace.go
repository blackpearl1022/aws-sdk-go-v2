// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a space. Deleting a space cannot be undone. Additionally, since space
// names must be unique across Amazon CodeCatalyst, you cannot reuse names of
// deleted spaces.
func (c *Client) DeleteSpace(ctx context.Context, params *DeleteSpaceInput, optFns ...func(*Options)) (*DeleteSpaceOutput, error) {
	if params == nil {
		params = &DeleteSpaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSpace", params, optFns, c.addOperationDeleteSpaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSpaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSpaceInput struct {

	// The name of the space. To retrieve a list of space names, use ListSpaces .
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteSpaceOutput struct {

	// The name of the space.
	//
	// This member is required.
	Name *string

	// The friendly name of the space displayed to users of the space in Amazon
	// CodeCatalyst.
	DisplayName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSpaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSpace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSpace{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteSpaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSpace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSpace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteSpace",
	}
}