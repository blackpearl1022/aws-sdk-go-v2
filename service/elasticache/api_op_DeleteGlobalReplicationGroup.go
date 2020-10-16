// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deleting a Global Datastore is a two-step process:
//
//     * First, you must
// DisassociateGlobalReplicationGroup to remove the secondary clusters in the
// Global Datastore.
//
//     * Once the Global Datastore contains only the primary
// cluster, you can use DeleteGlobalReplicationGroup API to delete the Global
// Datastore while retainining the primary cluster using Retain…= true.
//
// Since the
// Global Datastore has only a primary cluster, you can delete the Global Datastore
// while retaining the primary by setting RetainPrimaryCluster=true. When you
// receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or revert
// this operation.
func (c *Client) DeleteGlobalReplicationGroup(ctx context.Context, params *DeleteGlobalReplicationGroupInput, optFns ...func(*Options)) (*DeleteGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &DeleteGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGlobalReplicationGroup", params, optFns, addOperationDeleteGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGlobalReplicationGroupInput struct {

	// The name of the Global Datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The primary replication group is retained as a standalone replication group.
	//
	// This member is required.
	RetainPrimaryReplicationGroup *bool
}

type DeleteGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//
	//     * The GlobalReplicationGroupIdSuffix represents the name
	// of the Global Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteGlobalReplicationGroup{}, middleware.After)
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
	addOpDeleteGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGlobalReplicationGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteGlobalReplicationGroup",
	}
}