// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the definition of an existing DataBrew recipe job.
func (c *Client) UpdateRecipeJob(ctx context.Context, params *UpdateRecipeJobInput, optFns ...func(*Options)) (*UpdateRecipeJobOutput, error) {
	if params == nil {
		params = &UpdateRecipeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRecipeJob", params, optFns, c.addOperationUpdateRecipeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRecipeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecipeJobInput struct {

	// The name of the job to update.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// to be assumed when DataBrew runs the job.
	//
	// This member is required.
	RoleArn *string

	// One or more artifacts that represent the AWS Glue Data Catalog output from
	// running the job.
	DataCatalogOutputs []types.DataCatalogOutput

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job.
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//
	// * SSE-KMS -
	// Server-side encryption with keys managed by KMS.
	//
	// * SSE-S3 - Server-side
	// encryption with keys managed by Amazon S3.
	EncryptionMode types.EncryptionMode

	// Enables or disables Amazon CloudWatch logging for the job. If logging is
	// enabled, CloudWatch writes one log stream for each job run.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that DataBrew can consume when the job processes
	// data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// One or more artifacts that represent the output from running the job.
	Outputs []types.Output

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32
}

type UpdateRecipeJobOutput struct {

	// The name of the job that you updated.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateRecipeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecipeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecipeJob{}, middleware.After)
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
	if err = addOpUpdateRecipeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecipeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRecipeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "UpdateRecipeJob",
	}
}
