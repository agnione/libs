// package provides Interface to implement the HTTP client plugin for AgniOne Application Framework
//
// This interface defines functions that needs to implement when building client plugin
//
//   - New
//
//   - Initialize
//
//   - GetID
//
//   - Get
//
//   - Post
//
//   - Put
//
//   - Delete
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :  D. Ajith Nilantha de Silva  | 02/01/2024
//     Copyright     :  Open source MIT License
//     Class/module  :	IAHTTPClient - AgniOne Application Framework
//     Objective     :  Define the http client interface plugin
//     ---------------------------------------------------------------------------------------------------------------------
//     This interface will be used to implement the http client library.
//     It is required to provide ihttpclient/ihttpclient.go and httpclient/httpclient.go
//     files to build the client plug-in.
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		01/02/2024	Created 	Created the initial version
//     Ajith de Silva		02/02/2024	Updated 	Defined main functions
//     ---------------------------------------------------------------------------------------------------------------------
package iahttpclient

import (
	atypes "agnione/v1/src/afplugins/http/types"
	build "agnione/v1/src/lib"
)

// IHTTPClient interface expose the functions relates to HTTP protocol
type IAHTTPClient interface {

	//Cretes a new isntance of IZHTTPClient
	New() interface{}

	//Initialize the instance.
	//
	// Returns the ture if initialized scussessfully. Unless false
	Initialize(pInstance_ID int) bool

	// GetID retuns the pre-set id of the current instance
	GetID() (pInstance_ID int)

	// Get perfoms a HTTP GET request based on the given AHTTPRequest.
	//	Parameter http_request atypes.AHTTPRequest - valid AHTTPRequest
	//
	// If success returns AHTTPResponse with result (status code, headers and body in []bytes)
	// If failed then returns AHTTPResponse with valid status code and error with menaningful error
	Get(pHTTP_Request *atypes.AHTTPRequest) (*atypes.AHTTPResponse, error)

	// Post perfoms a HTTP POST request based on the given AHTTPRequest.
	//
	//	Parameter http_request atypes.AHTTPRequest - valid AHTTPRequest
	// If success returns AHTTPResponse with result (status code, headers and body in []bytes)
	// If failed then returns AHTTPResponse with valid status code and error with menaningful error
	Post(httppHTTP_Request_request *atypes.AHTTPRequest) (*atypes.AHTTPResponse, error)

	// Put perfoms a HTTP PUT request based on the given AHTTPRequest.
	//
	//	Parameter http_request atypes.AHTTPRequest - valid AHTTPRequest
	// If success returns AHTTPResponse with result (status code, headers and body in []bytes)
	// If failed then returns AHTTPResponse with valid status code and error with menaningful error
	Put(pHTTP_Request *atypes.AHTTPRequest) (*atypes.AHTTPResponse, error)

	// Delete perfoms a HTTP DELETE request based on the given AHTTPRequest.
	//
	//	Parameter http_request atypes.AHTTPRequest - valid AHTTPRequest
	// If success returns AHTTPResponse with result (status code, headers and body in []bytes)
	// If failed then returns AHTTPResponse with valid status code and error with menaningful error
	Delete(pHTTP_Request *atypes.AHTTPRequest) (*atypes.AHTTPResponse, error)

	// Info returns the build information of the library
	Info() build.BuildInfo
}
