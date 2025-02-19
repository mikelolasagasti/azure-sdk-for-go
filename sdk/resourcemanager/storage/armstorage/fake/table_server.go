//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// TableServer is a fake server for instances of the armstorage.TableClient type.
type TableServer struct {
	// Create is the fake for method TableClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientCreateOptions) (resp azfake.Responder[armstorage.TableClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TableClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientDeleteOptions) (resp azfake.Responder[armstorage.TableClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TableClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientGetOptions) (resp azfake.Responder[armstorage.TableClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TableClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armstorage.TableClientListOptions) (resp azfake.PagerResponder[armstorage.TableClientListResponse])

	// Update is the fake for method TableClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *armstorage.TableClientUpdateOptions) (resp azfake.Responder[armstorage.TableClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTableServerTransport creates a new instance of TableServerTransport with the provided implementation.
// The returned TableServerTransport instance is connected to an instance of armstorage.TableClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTableServerTransport(srv *TableServer) *TableServerTransport {
	return &TableServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armstorage.TableClientListResponse]](),
	}
}

// TableServerTransport connects instances of armstorage.TableClient to instances of TableServer.
// Don't use this type directly, use NewTableServerTransport instead.
type TableServerTransport struct {
	srv          *TableServer
	newListPager *tracker[azfake.PagerResponder[armstorage.TableClientListResponse]]
}

// Do implements the policy.Transporter interface for TableServerTransport.
func (t *TableServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TableClient.Create":
		resp, err = t.dispatchCreate(req)
	case "TableClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TableClient.Get":
		resp, err = t.dispatchGet(req)
	case "TableClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	case "TableClient.Update":
		resp, err = t.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TableServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if t.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tableServices/default/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Table](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	tableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	var options *armstorage.TableClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.TableClientCreateOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := t.srv.Create(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, tableNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tableServices/default/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	tableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, tableNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tableServices/default/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	tableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, tableNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tableServices/default/tables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListPager(resourceGroupNameUnescaped, accountNameUnescaped, nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstorage.TableClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}

func (t *TableServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tableServices/default/tables/(?P<tableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorage.Table](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	tableNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tableName")])
	if err != nil {
		return nil, err
	}
	var options *armstorage.TableClientUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstorage.TableClientUpdateOptions{
			Parameters: &body,
		}
	}
	respr, errRespr := t.srv.Update(req.Context(), resourceGroupNameUnescaped, accountNameUnescaped, tableNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Table, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
