/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 2
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DbWorkspaceIssueStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbWorkspaceIssueStats{}

// DbWorkspaceIssueStats struct for DbWorkspaceIssueStats
type DbWorkspaceIssueStats struct {
	// Number of issues opened in a time range
	Opened int32 `json:"opened"`
	// Number of issues closed in a time range
	Closed int32 `json:"closed"`
	// Repository average open/close time for issues
	Velocity int32 `json:"velocity"`
}

// NewDbWorkspaceIssueStats instantiates a new DbWorkspaceIssueStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbWorkspaceIssueStats(opened int32, closed int32, velocity int32) *DbWorkspaceIssueStats {
	this := DbWorkspaceIssueStats{}
	this.Opened = opened
	this.Closed = closed
	this.Velocity = velocity
	return &this
}

// NewDbWorkspaceIssueStatsWithDefaults instantiates a new DbWorkspaceIssueStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbWorkspaceIssueStatsWithDefaults() *DbWorkspaceIssueStats {
	this := DbWorkspaceIssueStats{}
	var opened int32 = 0
	this.Opened = opened
	var closed int32 = 0
	this.Closed = closed
	var velocity int32 = 0
	this.Velocity = velocity
	return &this
}

// GetOpened returns the Opened field value
func (o *DbWorkspaceIssueStats) GetOpened() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Opened
}

// GetOpenedOk returns a tuple with the Opened field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceIssueStats) GetOpenedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Opened, true
}

// SetOpened sets field value
func (o *DbWorkspaceIssueStats) SetOpened(v int32) {
	o.Opened = v
}

// GetClosed returns the Closed field value
func (o *DbWorkspaceIssueStats) GetClosed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Closed
}

// GetClosedOk returns a tuple with the Closed field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceIssueStats) GetClosedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Closed, true
}

// SetClosed sets field value
func (o *DbWorkspaceIssueStats) SetClosed(v int32) {
	o.Closed = v
}

// GetVelocity returns the Velocity field value
func (o *DbWorkspaceIssueStats) GetVelocity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Velocity
}

// GetVelocityOk returns a tuple with the Velocity field value
// and a boolean to check if the value has been set.
func (o *DbWorkspaceIssueStats) GetVelocityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Velocity, true
}

// SetVelocity sets field value
func (o *DbWorkspaceIssueStats) SetVelocity(v int32) {
	o.Velocity = v
}

func (o DbWorkspaceIssueStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbWorkspaceIssueStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["opened"] = o.Opened
	toSerialize["closed"] = o.Closed
	toSerialize["velocity"] = o.Velocity
	return toSerialize, nil
}

type NullableDbWorkspaceIssueStats struct {
	value *DbWorkspaceIssueStats
	isSet bool
}

func (v NullableDbWorkspaceIssueStats) Get() *DbWorkspaceIssueStats {
	return v.value
}

func (v *NullableDbWorkspaceIssueStats) Set(val *DbWorkspaceIssueStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDbWorkspaceIssueStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDbWorkspaceIssueStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbWorkspaceIssueStats(val *DbWorkspaceIssueStats) *NullableDbWorkspaceIssueStats {
	return &NullableDbWorkspaceIssueStats{value: val, isSet: true}
}

func (v NullableDbWorkspaceIssueStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbWorkspaceIssueStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
