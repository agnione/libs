// package provides Interface to implement web socket client plugin for AgniOne Application Framework
//
// This package includes below functions :
//
//   - New
//
//   - Initialize
//
//   - GetID
//
//   - DeInitialize
//
//   - IsConnected
//
//   - Connect
//
//   - Disconnect
//
//   - Read
//
//   - ReadJSON
//
//   - Write
//
//   - WriteJSON
//
//   - Info
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :   D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 02/02/2024
//     Copyright     :   Open source MIT License
//     Class/module  :   iawsclient - AgniOne Application Framework
//     Objective     :   Define the interface to build web socket client library
//
//     This interface will be used to implement the web socket client library.
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		06/02/2004	Created 	Created the initial version
//     Ajith de Silva		06/02/2004	Updated 	Defined main functions
//     ---------------------------------------------------------------------------------------------------------------------
package iawsclient

import build "agnione/v1/src/lib"

type IAWSClient interface {

	// New creates a new instance of IWSClient and return the interface
	New() interface{}

	// Initialize initializes the given id to the instance. Used to identify the instance
	//
	// Returns true if success. Unless false
	Initialize(pInstance_ID int) bool

	/// GetID returns the pre-set id of the current instance
	GetID() (pInstance_ID int)

	// DeInitialize removes the matching web socket connection of the pool.
	DeInitialize()

	// IsConnected checks the fetched web socket connection is connected by writing PING message.
	//
	// Returns true if the connection is live. unless false with error message
	IsConnected() (bool, error)

	// Connect establishes the fetched connection from the pool.
	//
	// Given wsurl will be used with the request headers for connection.
	//
	// If success then returns the true,HTTP status code and nil.
	//
	// If failed then returns false,-1 and the error message
	Connect(pWS_URL string, pRequest_Headers *map[string][]string, pSub_protocols *[]string,pCompression bool) (bool, int, error)

	// Disconnect disconnects the fetched connection from the pool.
	//
	// Returns true and nil if disconnection is success.
	//
	// Unless returns false and error message
	Disconnect() (bool, error)

	// Read reads the data from the connection.
	//
	// If success returns message 1|0 (1=Text Message ,2=Binary Message), message bytes and nil for error
	//
	// Unless returns 0 as message type, nil and error
	Read() (pMessage_Type int, pMessage *[]byte, err error)

	// Write writes the binary message to the fetched web socket connection.
	//
	// Returns true and nil if write is success.
	//
	// Unless returns false and error message
	Write(pMessage_Type int, pMessage *[]byte) (bool, error)

	// Info returns the build information of the library
	Info() build.BuildInfo
}
