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
	"time"
)

// checks if the DbEmoji type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbEmoji{}

// DbEmoji struct for DbEmoji
type DbEmoji struct {
	// Emoji identifier
	Id string `json:"id"`
	// Emoji Name
	Name string `json:"name"`
	// Emoji URL
	Url string `json:"url"`
	// Emoji display order
	DisplayOrder int32 `json:"display_order"`
	// Timestamp representing highlight reaction creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp representing highlight reaction last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDbEmoji instantiates a new DbEmoji object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbEmoji(id string, name string, url string, displayOrder int32) *DbEmoji {
	this := DbEmoji{}
	this.Id = id
	this.Name = name
	this.Url = url
	this.DisplayOrder = displayOrder
	return &this
}

// NewDbEmojiWithDefaults instantiates a new DbEmoji object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbEmojiWithDefaults() *DbEmoji {
	this := DbEmoji{}
	return &this
}

// GetId returns the Id field value
func (o *DbEmoji) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbEmoji) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DbEmoji) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DbEmoji) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *DbEmoji) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DbEmoji) SetUrl(v string) {
	o.Url = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *DbEmoji) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *DbEmoji) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DbEmoji) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DbEmoji) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DbEmoji) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DbEmoji) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbEmoji) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DbEmoji) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DbEmoji) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DbEmoji) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbEmoji) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	toSerialize["display_order"] = o.DisplayOrder
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDbEmoji struct {
	value *DbEmoji
	isSet bool
}

func (v NullableDbEmoji) Get() *DbEmoji {
	return v.value
}

func (v *NullableDbEmoji) Set(val *DbEmoji) {
	v.value = val
	v.isSet = true
}

func (v NullableDbEmoji) IsSet() bool {
	return v.isSet
}

func (v *NullableDbEmoji) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbEmoji(val *DbEmoji) *NullableDbEmoji {
	return &NullableDbEmoji{value: val, isSet: true}
}

func (v NullableDbEmoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbEmoji) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
