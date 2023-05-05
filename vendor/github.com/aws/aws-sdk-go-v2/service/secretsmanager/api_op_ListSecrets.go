// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the secrets that are stored by Secrets Manager in the Amazon Web Services
// account, not including secrets that are marked for deletion. To see secrets
// marked for deletion, use the Secrets Manager console. ListSecrets is eventually
// consistent, however it might not reflect changes from the last five minutes. To
// get the latest information for a specific secret, use DescribeSecret . To list
// the versions of a secret, use ListSecretVersionIds . To get the secret value
// from SecretString or SecretBinary , call GetSecretValue . For information about
// finding secrets in the console, see Find secrets in Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_search-secret.html)
// . Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see Logging Secrets Manager events with CloudTrail (https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html)
// . Required permissions: secretsmanager:ListSecrets . For more information, see
// IAM policy actions for Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
// .
func (c *Client) ListSecrets(ctx context.Context, params *ListSecretsInput, optFns ...func(*Options)) (*ListSecretsOutput, error) {
	if params == nil {
		params = &ListSecretsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecrets", params, optFns, c.addOperationListSecretsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecretsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecretsInput struct {

	// The filters to apply to the list of secrets.
	Filters []types.Filter

	// Specifies whether to include secrets scheduled for deletion. By default,
	// secrets scheduled for deletion aren't included.
	IncludePlannedDeletion *bool

	// The number of results to include in the response. If there are more results
	// available, in the response, Secrets Manager includes NextToken . To get the next
	// results, call ListSecrets again with the value from NextToken .
	MaxResults *int32

	// A token that indicates where the output should continue from, if a previous
	// call did not show all results. To get the next results, call ListSecrets again
	// with this value.
	NextToken *string

	// Secrets are listed by CreatedDate .
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListSecretsOutput struct {

	// Secrets Manager includes this value if there's more output available than what
	// is included in the current response. This can occur even when the response
	// includes no values at all, such as when you ask for a filtered view of a long
	// list. To get the next results, call ListSecrets again with this value.
	NextToken *string

	// A list of the secrets in the account.
	SecretList []types.SecretListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecretsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSecrets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecrets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecrets(options.Region), middleware.Before); err != nil {
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

// ListSecretsAPIClient is a client that implements the ListSecrets operation.
type ListSecretsAPIClient interface {
	ListSecrets(context.Context, *ListSecretsInput, ...func(*Options)) (*ListSecretsOutput, error)
}

var _ ListSecretsAPIClient = (*Client)(nil)

// ListSecretsPaginatorOptions is the paginator options for ListSecrets
type ListSecretsPaginatorOptions struct {
	// The number of results to include in the response. If there are more results
	// available, in the response, Secrets Manager includes NextToken . To get the next
	// results, call ListSecrets again with the value from NextToken .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecretsPaginator is a paginator for ListSecrets
type ListSecretsPaginator struct {
	options   ListSecretsPaginatorOptions
	client    ListSecretsAPIClient
	params    *ListSecretsInput
	nextToken *string
	firstPage bool
}

// NewListSecretsPaginator returns a new ListSecretsPaginator
func NewListSecretsPaginator(client ListSecretsAPIClient, params *ListSecretsInput, optFns ...func(*ListSecretsPaginatorOptions)) *ListSecretsPaginator {
	if params == nil {
		params = &ListSecretsInput{}
	}

	options := ListSecretsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecretsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecretsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecrets page.
func (p *ListSecretsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecretsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListSecrets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecrets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "ListSecrets",
	}
}
