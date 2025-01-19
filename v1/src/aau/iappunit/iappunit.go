// build package provides structure to represents the build information
//
// This package includes below functions:
//   - New
//   - Initialize
//   - IsInitialized
//   - GetID
//   - Deinitialize
//   - Start
//   - IsStarted
//   - Stop
//   - Status
//   - Info
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        : D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 26/01/2024
//     Copyright     :	Open source MIT License
//     Class/module  : IKAppUnit - AgniOne Application Framework
//     Objective     : Define the interface for AgniOne Application Unit
//     ---------------------------------------------------------------------------------------------------------------------
//     This interface will be used to implement of Application Units for KAndy Application Framework.
//     Main objective is to allow KAF to load the application units and call it'e methods and
//     to controll the execution flow of application unit
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     --------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		01/01/2024	Created 	Created the initial version
//     Ajith de Silva		01/01/2024	Updated 	Defined functions with parameters & return values
//     Ajith de Silva		01/01/2024	Updated 	Add the application framework interface as parameter
package iappunit

import (
	iappfm "agnione/v1/src/appfm/iappfw"
	atypes "agnione/v1/src/appfm/types"
	build "agnione/v1/src/lib"
)

// IAppUnit the interface for the AgniOne Application Unit
type IAppUnit interface {

	//New creates a new instance and return the IZAppUnit interface
	New() interface{}

	// Initialize initializes the application unit
	//
	// Parameter frm_instance - instance of application framework (ztypes.IZApp)
	// Parameter instance_id - instance id
	// Parameter appunit_name - unit name
	// Parameter appunit_file - unit file name
	// Parameter config_fike - unit configuration file name
	// Returns true if suceess, unless false
	Initialize(frm_instance iappfm.IAgniApp, instance_id int,
		appunit_name string, appunit_path string,config_file string) (bool, error)

	// IsInitialized returns the initialize status of the application unit
	IsInitialized() bool

	// GetID retuns the pre-set id of the current instance
	GetID() (instance_id int)

	// Deinitialize clear all the related objects
	Deinitialize()

	// Start starts the process of the application unit
	Start() (bool, error)

	// IsStarted returns the start status of the application unit
	IsStarted() bool

	// Stop stops the process of the application unit
	Stop() (bool, error)

	// Status return the status of the application unit
	Status() *atypes.AppUnitInfo

	// Info returns the information of the library
	Info() build.BuildInfo
}
