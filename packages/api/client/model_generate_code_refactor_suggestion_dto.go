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

// checks if the GenerateCodeRefactorSuggestionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateCodeRefactorSuggestionDto{}

// GenerateCodeRefactorSuggestionDto struct for GenerateCodeRefactorSuggestionDto
type GenerateCodeRefactorSuggestionDto struct {
	// Suggestion Length
	DescriptionLength int32 `json:"descriptionLength"`
	// Description Temperature
	Temperature int32 `json:"temperature"`
	// Suggestion Language
	Language string `json:"language"`
	// Code
	Code string `json:"code"`
}

// NewGenerateCodeRefactorSuggestionDto instantiates a new GenerateCodeRefactorSuggestionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateCodeRefactorSuggestionDto(descriptionLength int32, temperature int32, language string, code string) *GenerateCodeRefactorSuggestionDto {
	this := GenerateCodeRefactorSuggestionDto{}
	this.DescriptionLength = descriptionLength
	this.Temperature = temperature
	this.Language = language
	this.Code = code
	return &this
}

// NewGenerateCodeRefactorSuggestionDtoWithDefaults instantiates a new GenerateCodeRefactorSuggestionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateCodeRefactorSuggestionDtoWithDefaults() *GenerateCodeRefactorSuggestionDto {
	this := GenerateCodeRefactorSuggestionDto{}
	var language string = "english"
	this.Language = language
	return &this
}

// GetDescriptionLength returns the DescriptionLength field value
func (o *GenerateCodeRefactorSuggestionDto) GetDescriptionLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DescriptionLength
}

// GetDescriptionLengthOk returns a tuple with the DescriptionLength field value
// and a boolean to check if the value has been set.
func (o *GenerateCodeRefactorSuggestionDto) GetDescriptionLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionLength, true
}

// SetDescriptionLength sets field value
func (o *GenerateCodeRefactorSuggestionDto) SetDescriptionLength(v int32) {
	o.DescriptionLength = v
}

// GetTemperature returns the Temperature field value
func (o *GenerateCodeRefactorSuggestionDto) GetTemperature() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value
// and a boolean to check if the value has been set.
func (o *GenerateCodeRefactorSuggestionDto) GetTemperatureOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Temperature, true
}

// SetTemperature sets field value
func (o *GenerateCodeRefactorSuggestionDto) SetTemperature(v int32) {
	o.Temperature = v
}

// GetLanguage returns the Language field value
func (o *GenerateCodeRefactorSuggestionDto) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *GenerateCodeRefactorSuggestionDto) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *GenerateCodeRefactorSuggestionDto) SetLanguage(v string) {
	o.Language = v
}

// GetCode returns the Code field value
func (o *GenerateCodeRefactorSuggestionDto) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GenerateCodeRefactorSuggestionDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GenerateCodeRefactorSuggestionDto) SetCode(v string) {
	o.Code = v
}

func (o GenerateCodeRefactorSuggestionDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateCodeRefactorSuggestionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptionLength"] = o.DescriptionLength
	toSerialize["temperature"] = o.Temperature
	toSerialize["language"] = o.Language
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

type NullableGenerateCodeRefactorSuggestionDto struct {
	value *GenerateCodeRefactorSuggestionDto
	isSet bool
}

func (v NullableGenerateCodeRefactorSuggestionDto) Get() *GenerateCodeRefactorSuggestionDto {
	return v.value
}

func (v *NullableGenerateCodeRefactorSuggestionDto) Set(val *GenerateCodeRefactorSuggestionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateCodeRefactorSuggestionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateCodeRefactorSuggestionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateCodeRefactorSuggestionDto(val *GenerateCodeRefactorSuggestionDto) *NullableGenerateCodeRefactorSuggestionDto {
	return &NullableGenerateCodeRefactorSuggestionDto{value: val, isSet: true}
}

func (v NullableGenerateCodeRefactorSuggestionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateCodeRefactorSuggestionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
