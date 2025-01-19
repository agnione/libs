// iaappfw package provides interface to the IKApp application interface of AgniOne Application Framework
//
// This package includes below functions:
//
//   - AppPath
//   - ReloadConfig
//   - StartWSMonitor
//   - StopWSMonitor
//   - Is_Interrupted
//   - Write2Console
//   - Write2Log
//   - AddRoutine
//   - RemoveRoutine
//   - Version
//   - Name
//   - Memory_Usage
//   - Routine_Count
//   - RequestHandleCount_Add
//   - RequestFailedCount_Add
//   - HandledRequestCount
//   - FailedRequestCount
//   - Started
//   - SendMonitorMessage
//   - GetContext
//   - GetAppStatus
//   - GetAppInfo
//   - Save_App_Config
//   - GetFileInfo
//   - GetFileContent
//   - GetFileContetLines
//   - Get_Mailer
//   - Get_WSClient
//   - Get_RESTClient
//   - Logconfig
//   - WriteFileContent
//   - Units_List
//   - Unit_Stop
//   - Unit_start
//   - Unit_Restart
//   - Unit_Status
//
// ---------------------------------------------------------------------------------------------------------------------
// Author		:	D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 26/01/2024
// Copyright		:	Open source MIT License
// Class/module		:	IAApp - AgniOne Application Framework
// Objective		:	Define the interface for AgniOne Application Framework
// ---------------------------------------------------------------------------------------------------------------------
// This interface instance will be passed to application units so that the functionalities
// of framework can be called using the instance.
// Main requirements is to allow application units to access the framework features with
// level of control and also minimize the application unit implementation
// ---------------------------------------------------------------------------------------------------------------------
// Author			Date		Action		Description
// ---------------------------------------------------------------------------------------------------------------------
// Ajith de Silva		01/01/2024	Created 	Created the initial version
// Ajith de Silva		01/01/2024	Updated 	Defined functions with parameters & return values
// Ajith de Silva		01/01/2024	Updated 	separate appinfo and status
// Ajith de Silva		03/04/2024	Added	 	ExecuteAndFetchResult method to execute the OS command
// Ajith de Silva		14/05/2024	Added	 	Added a method ID to return application ID
// ---------------------------------------------------------------------------------------------------------------------
package iappfw

import (
	ihttp "agnione/v1/src/afplugins/http/iahttpclient" /// import the httplcient interface
	iwl "agnione/v1/src/afplugins/mailer/iamailer"     /// import the wsclient interface
	iws "agnione/v1/src/afplugins/websocket/iawsclient"
	atypes "agnione/v1/src/appfm/types"
	"context"
	"time"
)

type IAgniApp interface {

	// Reload_Config reloads the configuration of the application framework
	// 	Returns true and nil if configuration loaded successfully.
	//	Unless returns false and error
	Reload_Config() (bool, error)

	// Start_WSMonitor starts the web socket monitoring with the pre-set configuration in config file
	// 	Returns true and nil if it started successfully.
	// 	Unless returns false and error
	Start_WSMonitor() (bool, error)

	// StopWSMonitor stops the web socket monitoring
	// 	Returns true if it started successfully. Unless returns false
	Stop_WSMonitor() bool

	// Is_Interrupted returns the application interrupt channel to the caller.
	// 	External routines should check this channel as a terminator of routines.
	Is_Interrupted() chan bool

	// Write2Console writes a given message to the console
	// 	Parameter message string - a valid string message to write to string
	Write2Console(pMessage string)

	// Write2Log writes the given entry to the log file
	//	Parameter entry string - valid string to write to the log
	//	Parameter log_level ztypes.LogLevel - log level to use when writing the log entry
	Write2Log(pEntry string, pLog_Level atypes.LogLevel)


	// Set_LogLevel set/override the curren runtime log level with given level
	// 	Parameter pLogLevel  atypes.LogLevel - valid enum of  atypes.LogLevel
	Set_LogLevel(pLogLevel atypes.LogLevel)



	// Add_Routine increments the routine count of the Application Framework and
	//	adds 1 to the framework waitgroup
	//	increments the internal routine count by 1
	Add_Routine()

	// Remove_Routine decrements the routine count and
	//	remove 1 from the framework waitgroup
	//	decrements the internal routine count by 1
	Remove_Routine()

	// Version returns the current framework/application  version
	Version() string

	// Name returns the application name
	Name() string

	// ID returns the application ID
	ID() string

	// PID returns the application process ssID
	PID() int

	// Memory_Usage returns the current application current memory usage
	Memory_Usage() string

	// App_Path returns the current application base path
	App_Path() *string

	// Routine_Count returns the number of routines currently running
	Routine_Count() uint16

	// Execute_Command executes the given OS command and returns result
	//	command string parameter - the command to execute with arguments
	//	Returns result of the command execution as string.
	//	If failed then return the error
	Execute_Command(pCommand *string) (string, error)

	// Add_Request_HandleCount adds 1 to the request handle count.
	// 	This function is useful to external modules to update his request handle count
	// 	Will be used in sending status messages over REST and websockets
	Add_Request_HandleCount()

	// Save_App_Config save/overwrite the given application configuration into the app.config file.
	// 	This function is useful to modify the app units and configuration while application is running.
	// 	Also, it should be used with CAUTION.
	// 	Returns the true if given content is valid app.config content and saved successfully.
	// 	Unless returns false and error
	Save_App_Config(pAppConfigData *[]byte)(bool, error)
	
	// Add_Request_Failed_Count adds 1 to the request failed handle count.
	// 	This function is useful to external modules to update his request handle count
	// 	Will be used in sending status messages over REST and websockets
	Add_Request_Failed_Count()

	// Handled_Request_Count returns the number of requests handled set by RequestHandleCount_Add()
	Handled_Request_Count() uint64

	// Failed_Request_Count returns the number of failed requests count set by RequestFailedCount_Add()
	Failed_Request_Count() uint64

	// Started returns the application started time
	Started() time.Time

	// Send_Monitor_Message broadcasts given message via wbe socket monitoring.
	// 	If the web socket monitoring is not started then this message will be discarded.
	// 	If web socket monitoring has been started then the message will be broadcasted among
	// 	connected monitoring web socket clients
	Send_Monitor_Message(pMessage []byte)

	// Get_App_Status returns the current application status as [ztypes.AppStatus] [http://example.com]
	Get_App_Status() atypes.AppStatus

	// Get_App_Info returns the current application information as [ztypes.AppInfo] http://example.com
	Get_App_Info() atypes.AppInfo

	// Get_Context returns the current application context object
	Get_Context() *context.Context
	
	// Get_FileInfo returns the information of the given file in format of FileInfo struct
	Get_FileInfo(pFileName *string) (*atypes.FileInfo, error)

	// Get_File_Content returns the file content of the given file name.
	// 	Returns file content []bytes,nil if successful.
	// 	Unless returns nil and error
	Get_File_Content(pFileName *string) (*[]byte, error)
		
	// Logfile_Basepath returns the string of the log base path.
	Logfile_Basepath() *string

	// Logfile_Name returns the name of the application log file
	Logfile_Name() string

	// Get_FileContent_Lines returns the file []content string of the given file name.
	// 	Returns file content []bstring,nil if successful.
	// 	Unless returns nil and error
	Get_FileContent_Lines(pFileName *string) (*[]string, error)

// Write_FileContent writes the given content []byte to the given filename.
	// 	Returns true and nil if file exists. Unless returns nil and error
	Write_FileContent(pFileName *string, pData *[]byte) (bool, error)
	
	// Units_List returns the units given in the app.config file.
	// 	Returns true and nil if list successfully loaded. Unless returns nil and error
	Units_List() ([]atypes.Appunit, error)
	
	// Unit_Stop stops the given unit (if unit is loaded & running).
	//	pForce parameter determine that the unit should load in force or wait until all the current execution stops.
	// 	Returns true and nil if the given uint is successfully stopped. Unless returns nil and error
	Unit_Stop(pUnitName *string,pForce bool)(bool,error)
	
	// Unit_Start starts the given unit (if unit is nit loaded & not running).
	// 	Returns true and nil if the given uint is successfully loaded and running. Unless returns nil and error
	Unit_Start(pUnitName *string)(bool,error)
	
	// Unit_Restart re-starts the given unit.
	//	pForce parameter determine that the unit should perform the stop and start in force or wait until all the current execution stops.
	// 	Returns true and nil if the given uint is successfully restarted. Unless returns nil and error
	Unit_Restart(pUnitName *string,pForce bool)(bool,error)
		
	// Unit_Status returns the status of given unit at that time.
	// 	Returns unit info and nil if the given uint's status is successfully fetched. Unless returns nil and error
	Unit_Status(pUnitName *string)(*atypes.AppUnitInfo,error)
	
	// Get_Mailer returns the instance of mailer that can be used to send emails
	//	pType parameter will determine which mail plugin in should load.
	// A new instance will be created and return.
	// If failed then returns nil and error
	Get_Mailer(pType *string) (iwl.IAMailMessage, error)

	// Get_WSClient returns the instance of the Web Socket client defined in the config file
	// A new instance will be created and return.
	// If failed then returns nil and error
	Get_WSClient(pType *string) (iws.IAWSClient, error)

	// Get_RESTClient returns the instance of the REST client defined in the config file
	// A new instance will be created and return.
	// If failed then returns nil and error
	Get_RESTClient(pType *string) (ihttp.IAHTTPClient, error)
}