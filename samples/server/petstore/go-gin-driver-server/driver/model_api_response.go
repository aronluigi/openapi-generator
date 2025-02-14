/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

// ApiResponse - Describes the result of uploading an image resource
type ApiResponse struct {
	Code int32 `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
