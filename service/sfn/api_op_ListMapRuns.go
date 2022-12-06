// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Map Runs that were started by a given state machine execution. Use
// this API action to obtain Map Run ARNs, and then call DescribeMapRun to obtain
// more information, if needed.
func (c *Client) ListMapRuns(ctx context.Context, params *ListMapRunsInput, optFns ...func(*Options)) (*ListMapRunsOutput, error) {
	if params == nil {
		params = &ListMapRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMapRuns", params, optFns, c.addOperationListMapRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMapRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMapRunsInput struct {

	// The Amazon Resource Name (ARN) of the execution for which the Map Runs must be
	// listed.
	//
	// This member is required.
	ExecutionArn *string

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMapRunsOutput struct {

	// An array that lists information related to a Map Run, such as the Amazon
	// Resource Name (ARN) of the Map Run and the ARN of the state machine that started
	// the Map Run.
	//
	// This member is required.
	MapRuns []types.MapRunListItem

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMapRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListMapRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListMapRuns{}, middleware.After)
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
	if err = addOpListMapRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMapRuns(options.Region), middleware.Before); err != nil {
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

// ListMapRunsAPIClient is a client that implements the ListMapRuns operation.
type ListMapRunsAPIClient interface {
	ListMapRuns(context.Context, *ListMapRunsInput, ...func(*Options)) (*ListMapRunsOutput, error)
}

var _ ListMapRunsAPIClient = (*Client)(nil)

// ListMapRunsPaginatorOptions is the paginator options for ListMapRuns
type ListMapRunsPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMapRunsPaginator is a paginator for ListMapRuns
type ListMapRunsPaginator struct {
	options   ListMapRunsPaginatorOptions
	client    ListMapRunsAPIClient
	params    *ListMapRunsInput
	nextToken *string
	firstPage bool
}

// NewListMapRunsPaginator returns a new ListMapRunsPaginator
func NewListMapRunsPaginator(client ListMapRunsAPIClient, params *ListMapRunsInput, optFns ...func(*ListMapRunsPaginatorOptions)) *ListMapRunsPaginator {
	if params == nil {
		params = &ListMapRunsInput{}
	}

	options := ListMapRunsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMapRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMapRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMapRuns page.
func (p *ListMapRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMapRunsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListMapRuns(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListMapRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "ListMapRuns",
	}
}