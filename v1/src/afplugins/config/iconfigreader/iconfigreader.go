// build package provides Interface to implement the Configuration reader plugin for AgniOne Application Framework
//
// This interface defines functions that needs to implement when building plugin
//
//   - Load
//
//   - Content
//
//   - Get
//
//   - GetInt
//
//   - GetKeyValPairs
//
//   - GetArray
//
//   - Info
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :   D. Ajith Nilantha de Silva  ajithdesilva@gmail.com | 14/03/2024
//     Copyright     :   Open source MIT License
//     Class/module  :   iconfigreader - AgniOne Application Framework
//     Objective     :   Define the config reader interface plugin
//     ---------------------------------------------------------------------------------------------------------------------
//     This interface will be used to implement the configuration reader plugin.
//     It is required to provide ihttpclient/ihttpclient.go and httpclient/httpclient.go
//     files to build the client plug-in.
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		01/02/2024	Created 	Created the initial version
//     Ajith de Silva		02/02/2024	Updated 	Defined main functions
//     ---------------------------------------------------------------------------------------------------------------------
package iconfigreader

import build "agnione/v1/src/lib"

type IAConfigReader interface {

	// Load Loads the configuration file.
	// config_file parameter is the file name to read and load
	// Returns true if load successfully. Unless false
	Load(config_file string) error

	// Content returns the loaded config file content
	// Returns string with content of the loaded configuration file
	Content() string

	// Get returns the string value of the given config section/element
	// element_name parameters define which element to read
	// Returns valid sting if successful. Unless empty string
	Get(element_name string) string

	// GetInt returns the int value of the given config section/element
	// element_name parameters define which element to read
	// Returns valid int if successful. Unless -1
	GetInt(element_name string) int

	// GetKeyValPairs returns the key value paris of the given config section/element
	// element_name parameters define which element to read
	// Returns valid key/value pairs. Unless nil
	GetKeyValPairs(element_name string) map[any]any

	// GetArray returns the array of the given config section/element
	// element_name parameters define which element to read
	// Returns valid array. Unless nil
	GetArray(element_name string) []any

	// Info returns the build information of the library
	Info() build.BuildInfo
}
