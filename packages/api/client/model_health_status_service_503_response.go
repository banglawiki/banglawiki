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

// checks if the HealthStatusService503Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthStatusService503Response{}

// HealthStatusService503Response struct for HealthStatusService503Response
type HealthStatusService503Response struct {
	Status  *string                                             `json:"status,omitempty"`
	Info    map[string]HealthStatusService200ResponseInfoValue  `json:"info,omitempty"`
	Error   map[string]HealthStatusService200ResponseInfoValue  `json:"error,omitempty"`
	Details *map[string]HealthStatusService200ResponseInfoValue `json:"details,omitempty"`
}

// NewHealthStatusService503Response instantiates a new HealthStatusService503Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthStatusService503Response() *HealthStatusService503Response {
	this := HealthStatusService503Response{}
	return &this
}

// NewHealthStatusService503ResponseWithDefaults instantiates a new HealthStatusService503Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthStatusService503ResponseWithDefaults() *HealthStatusService503Response {
	this := HealthStatusService503Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthStatusService503Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStatusService503Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthStatusService503Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthStatusService503Response) SetStatus(v string) {
	o.Status = &v
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthStatusService503Response) GetInfo() map[string]HealthStatusService200ResponseInfoValue {
	if o == nil {
		var ret map[string]HealthStatusService200ResponseInfoValue
		return ret
	}
	return o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthStatusService503Response) GetInfoOk() (*map[string]HealthStatusService200ResponseInfoValue, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return &o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *HealthStatusService503Response) HasInfo() bool {
	if o != nil && IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]HealthStatusService200ResponseInfoValue and assigns it to the Info field.
func (o *HealthStatusService503Response) SetInfo(v map[string]HealthStatusService200ResponseInfoValue) {
	o.Info = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthStatusService503Response) GetError() map[string]HealthStatusService200ResponseInfoValue {
	if o == nil {
		var ret map[string]HealthStatusService200ResponseInfoValue
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthStatusService503Response) GetErrorOk() (*map[string]HealthStatusService200ResponseInfoValue, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HealthStatusService503Response) HasError() bool {
	if o != nil && IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]HealthStatusService200ResponseInfoValue and assigns it to the Error field.
func (o *HealthStatusService503Response) SetError(v map[string]HealthStatusService200ResponseInfoValue) {
	o.Error = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *HealthStatusService503Response) GetDetails() map[string]HealthStatusService200ResponseInfoValue {
	if o == nil || IsNil(o.Details) {
		var ret map[string]HealthStatusService200ResponseInfoValue
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStatusService503Response) GetDetailsOk() (*map[string]HealthStatusService200ResponseInfoValue, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *HealthStatusService503Response) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]HealthStatusService200ResponseInfoValue and assigns it to the Details field.
func (o *HealthStatusService503Response) SetDetails(v map[string]HealthStatusService200ResponseInfoValue) {
	o.Details = &v
}

func (o HealthStatusService503Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthStatusService503Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableHealthStatusService503Response struct {
	value *HealthStatusService503Response
	isSet bool
}

func (v NullableHealthStatusService503Response) Get() *HealthStatusService503Response {
	return v.value
}

func (v *NullableHealthStatusService503Response) Set(val *HealthStatusService503Response) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStatusService503Response) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStatusService503Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStatusService503Response(val *HealthStatusService503Response) *NullableHealthStatusService503Response {
	return &NullableHealthStatusService503Response{value: val, isSet: true}
}

func (v NullableHealthStatusService503Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStatusService503Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
