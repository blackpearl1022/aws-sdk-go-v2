// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is part of Amazon GameLift FleetIQ with game server groups, which is
// in preview release and is subject to change. Retrieves information for a game
// server resource. Information includes the game server statuses, health check
// info, and the instance the game server is running on. To retrieve game server
// information, specify the game server ID. If successful, the requested game
// server object is returned. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
// Related operations
//
//     * RegisterGameServer
//
//     * ListGameServers
//
//     *
// ClaimGameServer
//
//     * DescribeGameServer
//
//     * UpdateGameServer
//
//     *
// DeregisterGameServer
func (c *Client) DescribeGameServer(ctx context.Context, params *DescribeGameServerInput, optFns ...func(*Options)) (*DescribeGameServerOutput, error) {
	if params == nil {
		params = &DescribeGameServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameServer", params, optFns, addOperationDescribeGameServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameServerInput struct {

	// An identifier for the game server group where the game server is running. Use
	// either the GameServerGroup name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The identifier for the game server to be retrieved.
	//
	// This member is required.
	GameServerId *string
}

type DescribeGameServerOutput struct {

	// Object that describes the requested game server resource.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameServer{}, middleware.After)
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
	addOpDescribeGameServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameServer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeGameServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameServer",
	}
}