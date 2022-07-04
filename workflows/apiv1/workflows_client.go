// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package workflows

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	workflowspb "google.golang.org/genproto/googleapis/cloud/workflows/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListWorkflows  []gax.CallOption
	GetWorkflow    []gax.CallOption
	CreateWorkflow []gax.CallOption
	DeleteWorkflow []gax.CallOption
	UpdateWorkflow []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("workflows.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("workflows.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://workflows.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListWorkflows:  []gax.CallOption{},
		GetWorkflow:    []gax.CallOption{},
		CreateWorkflow: []gax.CallOption{},
		DeleteWorkflow: []gax.CallOption{},
		UpdateWorkflow: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Workflows API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListWorkflows(context.Context, *workflowspb.ListWorkflowsRequest, ...gax.CallOption) *WorkflowIterator
	GetWorkflow(context.Context, *workflowspb.GetWorkflowRequest, ...gax.CallOption) (*workflowspb.Workflow, error)
	CreateWorkflow(context.Context, *workflowspb.CreateWorkflowRequest, ...gax.CallOption) (*CreateWorkflowOperation, error)
	CreateWorkflowOperation(name string) *CreateWorkflowOperation
	DeleteWorkflow(context.Context, *workflowspb.DeleteWorkflowRequest, ...gax.CallOption) (*DeleteWorkflowOperation, error)
	DeleteWorkflowOperation(name string) *DeleteWorkflowOperation
	UpdateWorkflow(context.Context, *workflowspb.UpdateWorkflowRequest, ...gax.CallOption) (*UpdateWorkflowOperation, error)
	UpdateWorkflowOperation(name string) *UpdateWorkflowOperation
}

// Client is a client for interacting with Workflows API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Workflows is used to deploy and execute workflow programs.
// Workflows makes sure the program executes reliably, despite hardware and
// networking interruptions.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListWorkflows lists Workflows in a given project and location.
// The default order is not specified.
func (c *Client) ListWorkflows(ctx context.Context, req *workflowspb.ListWorkflowsRequest, opts ...gax.CallOption) *WorkflowIterator {
	return c.internalClient.ListWorkflows(ctx, req, opts...)
}

// GetWorkflow gets details of a single Workflow.
func (c *Client) GetWorkflow(ctx context.Context, req *workflowspb.GetWorkflowRequest, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	return c.internalClient.GetWorkflow(ctx, req, opts...)
}

// CreateWorkflow creates a new workflow. If a workflow with the specified name already
// exists in the specified project and location, the long running operation
// will return ALREADY_EXISTS error.
func (c *Client) CreateWorkflow(ctx context.Context, req *workflowspb.CreateWorkflowRequest, opts ...gax.CallOption) (*CreateWorkflowOperation, error) {
	return c.internalClient.CreateWorkflow(ctx, req, opts...)
}

// CreateWorkflowOperation returns a new CreateWorkflowOperation from a given name.
// The name must be that of a previously created CreateWorkflowOperation, possibly from a different process.
func (c *Client) CreateWorkflowOperation(name string) *CreateWorkflowOperation {
	return c.internalClient.CreateWorkflowOperation(name)
}

// DeleteWorkflow deletes a workflow with the specified name.
// This method also cancels and deletes all running executions of the
// workflow.
func (c *Client) DeleteWorkflow(ctx context.Context, req *workflowspb.DeleteWorkflowRequest, opts ...gax.CallOption) (*DeleteWorkflowOperation, error) {
	return c.internalClient.DeleteWorkflow(ctx, req, opts...)
}

// DeleteWorkflowOperation returns a new DeleteWorkflowOperation from a given name.
// The name must be that of a previously created DeleteWorkflowOperation, possibly from a different process.
func (c *Client) DeleteWorkflowOperation(name string) *DeleteWorkflowOperation {
	return c.internalClient.DeleteWorkflowOperation(name)
}

// UpdateWorkflow updates an existing workflow.
// Running this method has no impact on already running executions of the
// workflow. A new revision of the workflow may be created as a result of a
// successful update operation. In that case, such revision will be used
// in new workflow executions.
func (c *Client) UpdateWorkflow(ctx context.Context, req *workflowspb.UpdateWorkflowRequest, opts ...gax.CallOption) (*UpdateWorkflowOperation, error) {
	return c.internalClient.UpdateWorkflow(ctx, req, opts...)
}

// UpdateWorkflowOperation returns a new UpdateWorkflowOperation from a given name.
// The name must be that of a previously created UpdateWorkflowOperation, possibly from a different process.
func (c *Client) UpdateWorkflowOperation(name string) *UpdateWorkflowOperation {
	return c.internalClient.UpdateWorkflowOperation(name)
}

// gRPCClient is a client for interacting with Workflows API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client workflowspb.WorkflowsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new workflows client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Workflows is used to deploy and execute workflow programs.
// Workflows makes sure the program executes reliably, despite hardware and
// networking interruptions.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           workflowspb.NewWorkflowsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListWorkflows(ctx context.Context, req *workflowspb.ListWorkflowsRequest, opts ...gax.CallOption) *WorkflowIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListWorkflows[0:len((*c.CallOptions).ListWorkflows):len((*c.CallOptions).ListWorkflows)], opts...)
	it := &WorkflowIterator{}
	req = proto.Clone(req).(*workflowspb.ListWorkflowsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*workflowspb.Workflow, string, error) {
		resp := &workflowspb.ListWorkflowsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListWorkflows(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetWorkflows(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetWorkflow(ctx context.Context, req *workflowspb.GetWorkflowRequest, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetWorkflow[0:len((*c.CallOptions).GetWorkflow):len((*c.CallOptions).GetWorkflow)], opts...)
	var resp *workflowspb.Workflow
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateWorkflow(ctx context.Context, req *workflowspb.CreateWorkflowRequest, opts ...gax.CallOption) (*CreateWorkflowOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateWorkflow[0:len((*c.CallOptions).CreateWorkflow):len((*c.CallOptions).CreateWorkflow)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteWorkflow(ctx context.Context, req *workflowspb.DeleteWorkflowRequest, opts ...gax.CallOption) (*DeleteWorkflowOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteWorkflow[0:len((*c.CallOptions).DeleteWorkflow):len((*c.CallOptions).DeleteWorkflow)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DeleteWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UpdateWorkflow(ctx context.Context, req *workflowspb.UpdateWorkflowRequest, opts ...gax.CallOption) (*UpdateWorkflowOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "workflow.name", url.QueryEscape(req.GetWorkflow().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateWorkflow[0:len((*c.CallOptions).UpdateWorkflow):len((*c.CallOptions).UpdateWorkflow)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateWorkflowOperation manages a long-running operation from CreateWorkflow.
type CreateWorkflowOperation struct {
	lro *longrunning.Operation
}

// CreateWorkflowOperation returns a new CreateWorkflowOperation from a given name.
// The name must be that of a previously created CreateWorkflowOperation, possibly from a different process.
func (c *gRPCClient) CreateWorkflowOperation(name string) *CreateWorkflowOperation {
	return &CreateWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateWorkflowOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	var resp workflowspb.Workflow
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateWorkflowOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	var resp workflowspb.Workflow
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateWorkflowOperation) Metadata() (*workflowspb.OperationMetadata, error) {
	var meta workflowspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateWorkflowOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateWorkflowOperation) Name() string {
	return op.lro.Name()
}

// DeleteWorkflowOperation manages a long-running operation from DeleteWorkflow.
type DeleteWorkflowOperation struct {
	lro *longrunning.Operation
}

// DeleteWorkflowOperation returns a new DeleteWorkflowOperation from a given name.
// The name must be that of a previously created DeleteWorkflowOperation, possibly from a different process.
func (c *gRPCClient) DeleteWorkflowOperation(name string) *DeleteWorkflowOperation {
	return &DeleteWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteWorkflowOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteWorkflowOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteWorkflowOperation) Metadata() (*workflowspb.OperationMetadata, error) {
	var meta workflowspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteWorkflowOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteWorkflowOperation) Name() string {
	return op.lro.Name()
}

// UpdateWorkflowOperation manages a long-running operation from UpdateWorkflow.
type UpdateWorkflowOperation struct {
	lro *longrunning.Operation
}

// UpdateWorkflowOperation returns a new UpdateWorkflowOperation from a given name.
// The name must be that of a previously created UpdateWorkflowOperation, possibly from a different process.
func (c *gRPCClient) UpdateWorkflowOperation(name string) *UpdateWorkflowOperation {
	return &UpdateWorkflowOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateWorkflowOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	var resp workflowspb.Workflow
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateWorkflowOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*workflowspb.Workflow, error) {
	var resp workflowspb.Workflow
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateWorkflowOperation) Metadata() (*workflowspb.OperationMetadata, error) {
	var meta workflowspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateWorkflowOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateWorkflowOperation) Name() string {
	return op.lro.Name()
}

// WorkflowIterator manages a stream of *workflowspb.Workflow.
type WorkflowIterator struct {
	items    []*workflowspb.Workflow
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*workflowspb.Workflow, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkflowIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkflowIterator) Next() (*workflowspb.Workflow, error) {
	var item *workflowspb.Workflow
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkflowIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkflowIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}