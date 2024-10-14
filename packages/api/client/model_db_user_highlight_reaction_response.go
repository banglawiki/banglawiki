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

// checks if the DbUserHighlightReactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbUserHighlightReactionResponse{}

// DbUserHighlightReactionResponse struct for DbUserHighlightReactionResponse
type DbUserHighlightReactionResponse struct {
	// Emoji identifier
	EmojiId       string `json:"emoji_id"`
	ReactionCount *int32 `json:"reaction_count,omitempty"`
	// Usernames of users who reacted with this emoji
	ReactionUsers []string `json:"reaction_users,omitempty"`
}

// NewDbUserHighlightReactionResponse instantiates a new DbUserHighlightReactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbUserHighlightReactionResponse(emojiId string) *DbUserHighlightReactionResponse {
	this := DbUserHighlightReactionResponse{}
	this.EmojiId = emojiId
	return &this
}

// NewDbUserHighlightReactionResponseWithDefaults instantiates a new DbUserHighlightReactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbUserHighlightReactionResponseWithDefaults() *DbUserHighlightReactionResponse {
	this := DbUserHighlightReactionResponse{}
	return &this
}

// GetEmojiId returns the EmojiId field value
func (o *DbUserHighlightReactionResponse) GetEmojiId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmojiId
}

// GetEmojiIdOk returns a tuple with the EmojiId field value
// and a boolean to check if the value has been set.
func (o *DbUserHighlightReactionResponse) GetEmojiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmojiId, true
}

// SetEmojiId sets field value
func (o *DbUserHighlightReactionResponse) SetEmojiId(v string) {
	o.EmojiId = v
}

// GetReactionCount returns the ReactionCount field value if set, zero value otherwise.
func (o *DbUserHighlightReactionResponse) GetReactionCount() int32 {
	if o == nil || IsNil(o.ReactionCount) {
		var ret int32
		return ret
	}
	return *o.ReactionCount
}

// GetReactionCountOk returns a tuple with the ReactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserHighlightReactionResponse) GetReactionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReactionCount) {
		return nil, false
	}
	return o.ReactionCount, true
}

// HasReactionCount returns a boolean if a field has been set.
func (o *DbUserHighlightReactionResponse) HasReactionCount() bool {
	if o != nil && !IsNil(o.ReactionCount) {
		return true
	}

	return false
}

// SetReactionCount gets a reference to the given int32 and assigns it to the ReactionCount field.
func (o *DbUserHighlightReactionResponse) SetReactionCount(v int32) {
	o.ReactionCount = &v
}

// GetReactionUsers returns the ReactionUsers field value if set, zero value otherwise.
func (o *DbUserHighlightReactionResponse) GetReactionUsers() []string {
	if o == nil || IsNil(o.ReactionUsers) {
		var ret []string
		return ret
	}
	return o.ReactionUsers
}

// GetReactionUsersOk returns a tuple with the ReactionUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserHighlightReactionResponse) GetReactionUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.ReactionUsers) {
		return nil, false
	}
	return o.ReactionUsers, true
}

// HasReactionUsers returns a boolean if a field has been set.
func (o *DbUserHighlightReactionResponse) HasReactionUsers() bool {
	if o != nil && !IsNil(o.ReactionUsers) {
		return true
	}

	return false
}

// SetReactionUsers gets a reference to the given []string and assigns it to the ReactionUsers field.
func (o *DbUserHighlightReactionResponse) SetReactionUsers(v []string) {
	o.ReactionUsers = v
}

func (o DbUserHighlightReactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbUserHighlightReactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emoji_id"] = o.EmojiId
	if !IsNil(o.ReactionCount) {
		toSerialize["reaction_count"] = o.ReactionCount
	}
	if !IsNil(o.ReactionUsers) {
		toSerialize["reaction_users"] = o.ReactionUsers
	}
	return toSerialize, nil
}

type NullableDbUserHighlightReactionResponse struct {
	value *DbUserHighlightReactionResponse
	isSet bool
}

func (v NullableDbUserHighlightReactionResponse) Get() *DbUserHighlightReactionResponse {
	return v.value
}

func (v *NullableDbUserHighlightReactionResponse) Set(val *DbUserHighlightReactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDbUserHighlightReactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDbUserHighlightReactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbUserHighlightReactionResponse(val *DbUserHighlightReactionResponse) *NullableDbUserHighlightReactionResponse {
	return &NullableDbUserHighlightReactionResponse{value: val, isSet: true}
}

func (v NullableDbUserHighlightReactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbUserHighlightReactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
