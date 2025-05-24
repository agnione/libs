// build package provides type to represents the build information of plugin/application
//
// This package includes below type:
//
//   - BuildInfo
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author		:   D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 16/01/2024
//     Copyright   :	MIT License
//     package		:   BuildInfo - AgniOne Application Framework
//     Objective	:   Define common structure to represents the build information for AgniOne Application Framework
//     ---------------------------------------------------------------------------------------------------------------------
//     This structure will be used to return build information of every module, plugin,library
//     and application.
//
//     ** During the build process, all modules will be fed with details using ldflags
//     ---------------------------------------------------------------------------------------------------------------------
//     Author                        	Date        	Action      	Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva			28/01/2024	Created 	Created the initial version
//     ---------------------------------------------------------------------------------------------------------------------
package build

// BuildInfo structure to hold the build information
type BuildInfo struct {

	// Version version information stored during the build process
	Version string
	// Time built time stored during the build process
	Time string

	// User built user name stored during the build process
	User string

	// BuildGoVersion go version stored during the build process
	BuildGoVersion string
}


