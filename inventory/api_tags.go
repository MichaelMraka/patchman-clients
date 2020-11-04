/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// TagsApiService TagsApi service
type TagsApiService service

// ApiTagGetTagsOpts Optional parameters for the method 'ApiTagGetTags'
type ApiTagGetTagsOpts struct {
    Tags optional.Interface
    OrderBy optional.String
    OrderHow optional.String
    PerPage optional.Int32
    Page optional.Int32
    Staleness optional.Interface
    Search optional.String
    RegisteredWith optional.String
    Filter optional.Interface
}

/*
ApiTagGetTags Get the active host tags for a given account
Required permissions: inventory:hosts:read
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApiTagGetTagsOpts - Optional Parameters:
 * @param "Tags" (optional.Interface of []string) -  filters out hosts not tagged by the given tags
 * @param "OrderBy" (optional.String) -  Ordering field name
 * @param "OrderHow" (optional.String) -  Direction of the ordering. Default to ASC
 * @param "PerPage" (optional.Int32) -  A number of items to return per page.
 * @param "Page" (optional.Int32) -  A page number of the items to return.
 * @param "Staleness" (optional.Interface of []string) -  Culling states of the hosts. Default: fresh,stale,unknown
 * @param "Search" (optional.String) -  Only include tags that match the given search string. The value is matched against namespace, key and value.
 * @param "RegisteredWith" (optional.String) -  Filters out any host not registered with the specified service
 * @param "Filter" (optional.Interface of map[string]interface{}) -  Filters hosts based on system_profile fields
@return ActiveTags
*/
func (a *TagsApiService) ApiTagGetTags(ctx _context.Context, localVarOptionals *ApiTagGetTagsOpts) (ActiveTags, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ActiveTags
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tags"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		t:=localVarOptionals.Tags.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tags", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.OrderBy.IsSet() {
		localVarQueryParams.Add("order_by", parameterToString(localVarOptionals.OrderBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderHow.IsSet() {
		localVarQueryParams.Add("order_how", parameterToString(localVarOptionals.OrderHow.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("per_page", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Staleness.IsSet() {
		t:=localVarOptionals.Staleness.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("staleness", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("staleness", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RegisteredWith.IsSet() {
		localVarQueryParams.Add("registered_with", parameterToString(localVarOptionals.RegisteredWith.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-rh-identity"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
