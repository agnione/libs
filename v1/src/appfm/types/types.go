// atypes package provides types to use for AgniOne Application Framework
//
// This package includes below type:
//   - App
//   - Appunit
//   - AppConfig
//   - AppInfo
//   - Info
//   - FileInfo
//   - ZAppUnitInfo
//   - AppStatus
//   - ConvertStoI
//   - LogLevel
//   - MainConfig
//   - Httpmonitor
//   - Wsmonitor
//   - Mqengine
//   - Websocket
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :   D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 26/01/2024
//     Copyright     :   Open source MIT License
//     Class/module  :   atypes - AgniOne Application Framework
//     Objective     :   Define the common types for AgniOne Application Framework
//     ---------------------------------------------------------------------------------------------------------------------
//     Types defined here will be used by every package in the
//     Application Framework, Application PlugIns and Application Units
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		01/01/2024	Created 	Created the initial version
//     Ajith de Silva		01/01/2024	Updated 	separate appinfo and status into to structs
//     Ajith de Silva		02/26/2024	Added 		Included the Log level constants
//     Vindhya Bandara		03/04/2024	Added		Included the MainConfig constants
//     Ajith de Silva		10/06/2024	Added		added the profiler port and changed all ports type to uint16
//     ---------------------------------------------------------------------------------------------------------------------
package atypes

import "time"

type Info struct {
	Name    string
	Version string
}

// AppInfo holds the complete Application information
type AppInfo struct {
	Info
	Started             string	// started time
	UpTime              string	// uptime
	AppUnits            []AppUnitInfo	// Loaded app Unit(s) information
	PID int						// OS process ID of the Application
	WSMonitor_Started   bool	// flag to indicate the Web Socket monitoring is started
	HTTPMonitor_Started bool	// flag to indicate the REST monitoring is started
}

// structure to hold the file information
type FileInfo struct {
	Name    string    // base name of the file
	Size    int64     // length in bytes for regular files; system-dependent for others
	ModTime time.Time // modification time
	IsDir   bool      // abbreviation for Mode().IsDir()
}

type MemUsage struct{
	Heap uint64
	HeapAlloc uint64
	Total uint64
	
}

// KAppUnit information
type AppUnitInfo struct {
	Info        Info	// holds the name & version
	Mem_Usage   MemUsage	// holds the memory usage
	Req_Handled uint32	// number of request handled successfully
	Req_Failed  uint32	// number of request failed to handle
	Routines    uint16	// count of running number of routines/threads
	Active 		uint16	// count of current active number executions
}

// / application status
type AppStatus struct {
	Mem_Usage      MemUsage	// holds the memory usage	// memory usage	
	Req_Handled    uint64	// total request handled
	Req_Failed     uint64	// total request failed
	Routines       uint16		// count of running routines/thread
	MonitorClients uint8	// number of web socket monitor clients
	StatusClients  uint8	// number of REST monitor client
}

// ConvertStoI converts given structure to given interface
// Returns the converted interface.
func ConvertStoI[T any](c any) interface{} {
	return c.(T)
}

// LogLevel type
type LogLevel int8

// constants for log levels flags
const (
	LOG_FATAL LogLevel = 1
	LOG_PANIC LogLevel = 2
	LOG_ERROR LogLevel = 3
	LOG_WARN  LogLevel = 4
	LOG_INFO  LogLevel = 5
	LOG_DEBUG LogLevel = 6
)

type Logconfig struct {
	LogLevel        string `json:"log_level"`
	LogFileBasePath string `json:"log_file_base_path"`
	LogFileMaxSize  int    `json:"log_file_max_size"`
}

// Struture to hold application configuration
type AppConfig struct {
	App      App       `json:"app"`
	Log      Logconfig `json:"log"`
	Appunits []Appunit `json:"appunits"`
}

type App struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Version string `json:"version"`
}

type Appunit struct {
	Uname    string `json:"uname"`
	Path     string `json:"path"`
	ConfigFile   string `json:"config"`
	Enable   int8   `json:"enable"`
	PoolSize int8   `json:"pool_size"`
}

type PlugIn struct {
	Type   string `json:"type"`
	Ifname string `json:"ifname"`
	Path   string `json:"path"`
	Name   string `json:"name"`
	Enable int8    `json:"enable"`
}

type FMConfig struct {
	Core struct {
		Log struct {
			Level        string `json:"leg_level"`
			File_MaxSize  int    `json:"log_file_max_size"`
			File_Base_Path string `json:"log_file_base_path"`
		} `json:"log"`
		HTTPMonitor struct {
			Host         string `json:"host"`
			Port         *int    `json:"port"`
			Enable       int8    `json:"enable"`
			AutoStart int8    `json:"auto_start"`
		} `json:"http_monitor"`
		WSMonitor struct {
			Host      string `json:"host"`
			Port      *int    `json:"port"`
			Enable    int8    `json:"enable"`
		} `json:"ws_monitor"`
	} `json:"core"`
	Plugins struct {
		HTTP []PlugIn  `json:"http"`
		Websocket []PlugIn `json:"websocket"`
		Mailer []PlugIn `json:"mailer"`
	} `json:"plugins"`
}
