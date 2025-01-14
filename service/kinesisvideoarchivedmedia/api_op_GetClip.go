// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Downloads an MP4 file (clip) containing the archived, on-demand media from the
// specified video stream over the specified time range. Both the StreamName and
// the StreamARN parameters are optional, but you must specify either the
// StreamName or the StreamARN when invoking this API operation. As a prerequisite
// to using GetCLip API, you must obtain an endpoint using GetDataEndpoint ,
// specifying GET_CLIP for the APIName parameter. An Amazon Kinesis video stream
// has the following requirements for providing data through MP4:
//   - The media must contain h.264 or h.265 encoded video and, optionally, AAC or
//     G.711 encoded audio. Specifically, the codec ID of track 1 should be
//     V_MPEG/ISO/AVC (for h.264) or V_MPEGH/ISO/HEVC (for H.265). Optionally, the
//     codec ID of track 2 should be A_AAC (for AAC) or A_MS/ACM (for G.711).
//   - Data retention must be greater than 0.
//   - The video track of each fragment must contain codec private data in the
//     Advanced Video Coding (AVC) for H.264 format and HEVC for H.265 format. For more
//     information, see MPEG-4 specification ISO/IEC 14496-15 (https://www.iso.org/standard/55980.html)
//     . For information about adapting stream data to a given format, see NAL
//     Adaptation Flags (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/producer-reference-nal.html)
//     .
//   - The audio track (if present) of each fragment must contain codec private
//     data in the AAC format ( AAC specification ISO/IEC 13818-7 (https://www.iso.org/standard/43345.html)
//     ) or the MS Wave format (http://www-mmsp.ece.mcgill.ca/Documents/AudioFormats/WAVE/WAVE.html)
//     .
//
// You can monitor the amount of outgoing data by monitoring the
// GetClip.OutgoingBytes Amazon CloudWatch metric. For information about using
// CloudWatch to monitor Kinesis Video Streams, see Monitoring Kinesis Video
// Streams (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/monitoring.html)
// . For pricing information, see Amazon Kinesis Video Streams Pricing (https://aws.amazon.com/kinesis/video-streams/pricing/)
// and AWS Pricing (https://aws.amazon.com/pricing/) . Charges for outgoing AWS
// data apply.
func (c *Client) GetClip(ctx context.Context, params *GetClipInput, optFns ...func(*Options)) (*GetClipOutput, error) {
	if params == nil {
		params = &GetClipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClip", params, optFns, c.addOperationGetClipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClipInput struct {

	// The time range of the requested clip and the source of the timestamps.
	//
	// This member is required.
	ClipFragmentSelector *types.ClipFragmentSelector

	// The Amazon Resource Name (ARN) of the stream for which to retrieve the media
	// clip. You must specify either the StreamName or the StreamARN.
	StreamARN *string

	// The name of the stream for which to retrieve the media clip. You must specify
	// either the StreamName or the StreamARN.
	StreamName *string

	noSmithyDocumentSerde
}

type GetClipOutput struct {

	// The content type of the media in the requested clip.
	ContentType *string

	// Traditional MP4 file that contains the media clip from the specified video
	// stream. The output will contain the first 100 MB or the first 200 fragments from
	// the specified start timestamp. For more information, see Kinesis Video Streams
	// Limits (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/limits.html) .
	Payload io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetClip{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetClip{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addGetClipResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetClipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClip(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetClip(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetClip",
	}
}

type opGetClipResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetClipResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetClipResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "kinesisvideo"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "kinesisvideo"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("kinesisvideo")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetClipResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetClipResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
