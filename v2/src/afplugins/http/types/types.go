// build package provides structure to represents http request & response types
//
// This package includes below functions:
//
//   - ATTPRequest
//
//   - AHTTPResponse
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :		D. Ajith Nilantha de Silva  | 02/01/2024
//     Copyright     :		Open source MIT License
//     Class/module  :		IAHTTPTypes
//     Objective     :		Define the common types for http client plugins
//     ---------------------------------------------------------------------------------------------------------------------
//     This types will be used to implement the http client library.
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		01/02/2024	Created 	Created the initial version
//     Ajith de Silva		01/02/2024	Updated 	Defined main types
//     ---------------------------------------------------------------------------------------------------------------------
package types

// HTTPRequest contains the request parameters.
type AHTTPRequest struct {

	// URL a valid HTTP URL string
	URL string

	// Body data (byte array) to submit as body of HTTP request. Will not be used in GET request
	Body []byte

	// Headers collection of headers that needs to pass to the HTTP request
	Headers map[string]string

	// Timeout time to wait for result in seconds
	Timeout int
}

// HTTPResponse contains the response/result of the performed HTTP request.
type AHTTPResponse struct {

	// Headers response headers collection
	Headers map[string]string
	
	// Body response body
	Body []byte
	
	// StatusCode HTTP status code after performing the request
	StatusCode int
	
}
