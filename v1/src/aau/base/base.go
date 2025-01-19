// build package provides base type and methods to use when developing plugin for AgniOne Application Framework
//
// This package includes following types & methods/functions:
//   - AUBase
//   - Initialize
//   - Deinitialize
//   - IsInitialized
//   - Start
//   - Stop
//   - IsStarted
//   - Add_Routine
//   - Remove_Routine
//   - Get_ID
//   - Get_PID
//   - Get_KAUUID
//   - Status
//   - Generate_Monitoring_Message
//   - ConvertToFloat32
//   - ConvertToInt32
//   - Get_RESTClient
//   - Get_WSClient
//   - Get_Mailer
//   - ExecuteandFetch
//   - Send_Monitor_Message
//   - Write2Log
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :   D. Ajith Nilantha de Silva  ajithdesilva@gmail.com | 02/01/2024
//     Copyright     :   Open source MIT License
//     Class/module  :   AUBase - AgniOne Application Framework
//     Objective     :   Implement the KAU base library
//     ---------------------------------------------------------------------------------------------------------------------
//     Uses to encapsulte all the the common methods onf ZAU
//     Helps to buid the custom/business specific ZAU with more focus
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     --------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		05/04/2024	Created 	Created the initial version
//     Ajith de Silva		05/04/2024	Updated 	Added common functions
//     Ajith de Silva		05/04/2004	Updated 	Added plugin functions
//     Ajith de Silva		09/04/2004	Updated 	Added Get_MQClusterClient plugin functions
//     Ajith de Silva		09/04/2004	Updated 	Added ExecuteandFetch function
//     Ajith de Silva		09/04/2004	Added 		Added ConvertToFloat32 function
//
// ------------------------------------------------------------------------------------------------------------------------
package AUBase

import (
	autypes "agnione/v1/src/aau/types"
	ihttp "agnione/v1/src/afplugins/http/iahttpclient" /// import the httplcient interface
	amailer "agnione/v1/src/afplugins/mailer/iamailer"
	"agnione/v1/src/afplugins/websocket/iawsclient"
	iappfm "agnione/v1/src/appfm/iappfw"
	atypes "agnione/v1/src/appfm/types"
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// AUBase base struct to hold the propeties of the Application unit
type AUBase struct {
	Info_Lock      *sync.Mutex
	Unit_Info      *atypes.AppUnitInfo
	ID             int /// ID of the instance
	AppFramework   iappfm.IAgniApp	/// instance of the running framework
	Unit_Name      string	
	Config_File      string	
	Unit_Path      string
	App_UID string
	Stopper        chan bool
	Is_Started     bool
	Is_Initialized bool
}

// Initialize initializes the properties of the base struct.
//
// Reterns true and nil when success
// If failed then returns false and error message
func (appu *AUBase) Initialize(pFM_Instance iappfm.IAgniApp, pInstance_ID int,
	 pUnit_Name string, pUnit_Path string,pConfig_File string,) (bool, error) {

	if pFM_Instance == nil {
		appu.Is_Initialized = false
		appu.AppFramework.Send_Monitor_Message([]byte(fmt.Sprintf("%d fm_instance iappfm.IAgniApp is NIL", pInstance_ID)))
		return false, fmt.Errorf("%d fm_instance iappfm.IKApp is NIL", pInstance_ID)
	}
	
	appu.Unit_Info=new(atypes.AppUnitInfo)
	
	appu.AppFramework = pFM_Instance
	appu.ID = pInstance_ID
	appu.Unit_Name = pUnit_Name
	appu.Config_File = pConfig_File
	appu.Unit_Path = pUnit_Path
	appu.Info_Lock = &sync.Mutex{}
	appu.Unit_Info.Info.Name = appu.Unit_Name
	appu.Is_Initialized = true
	appu.Stopper = nil

	appu.Read_Memory_Usage()
	
	
	return true, nil
}

// Deinitialize clear all the related objects.
func (appu *AUBase) Deinitialize() {

	/// clear objects here
	appu.AppFramework = nil
	appu.Info_Lock = nil
	appu.Unit_Info = nil
	appu.Unit_Name = ""
	appu.Config_File = ""
	appu.Unit_Path = ""
}

func (appu *AUBase) Start() (bool, error) {
	if !appu.Is_Initialized {
		return false, fmt.Errorf("%d Failed to Start. %s instance is not initialized", appu.ID, appu.Unit_Name)
	}
	appu.Stopper = make(chan bool)
	return true, nil
}

func (appu *AUBase) Stop() (bool, error) {
	if appu.AppFramework == nil {
		return false, fmt.Errorf("instance is not initialized")
	}
	//fmt.Printf("%d In the  BASE STOP %s..... 2 \n", appu.ID, appu.Unit_name)
	appu.AppFramework.Write2Log(appu.App_UID + " - Stopping the AUBase.....", atypes.LOG_INFO)
	appu.AppFramework.Send_Monitor_Message([]byte(appu.App_UID + " - Stopping the AUBase....."))

	if !appu.Is_Started {		
		appu.AppFramework.Write2Log(appu.App_UID + " - Failed to Stopping the AUBase. Instance process is not started", atypes.LOG_INFO)
		appu.AppFramework.Send_Monitor_Message([]byte(appu.App_UID + " - Stopping the AUBase. Instance process is not started"))
		return false, fmt.Errorf("instance process is not started")
	}

	if appu.Stopper != nil {
		appu.Write2Log(appu.App_UID + " - Closing the AUBase Stopper chan .....", atypes.LOG_INFO)
		close(appu.Stopper)
		appu.Write2Log(appu.App_UID + " - Closing the AUBase Stopper chan ........DONE", atypes.LOG_INFO)
	}

	appu.AppFramework.Write2Log(appu.App_UID + " - Stopping the AUBase..... DONE", atypes.LOG_INFO)
	appu.AppFramework.Send_Monitor_Message([]byte(appu.App_UID + " - Stopping the AUBase.....DONE"))
	appu.Is_Started = false
	
	return true, nil
}

// Add_Routine increment the routine total routine count of the AppFramework
// and the current application unit
func (appu *AUBase) Add_Routine() {
	

	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Unit_Info.Routines++
	
	go func ()  {
		defer recover()
		appu.AppFramework.Add_Routine()
		///appu.AppFramework.Set_Unitinfo(&appu.ID,appu.Unit_Info)
	
	}()
	
}

// Remove_Routine decrement the total routine count of the AppFramework
// and the current application unit
func (appu *AUBase) Remove_Routine() {
	
	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Unit_Info.Routines--
	
	go func(){
		defer recover()
		appu.AppFramework.Remove_Routine()
	}()
	

}

// Add_Request_Handled_Count increment the total requests handled of the AppFramework
// and the current application unit
func (appu *AUBase) Add_Request_Handled_Count() {
	
	
	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Unit_Info.Req_Handled++
	
	go func(){
		defer recover()
		appu.AppFramework.Add_Request_HandleCount()
	}()

}

// Add_Request_Failed_Count increment the total failed requests of the AppFramework
// and the current application unit
func (appu *AUBase) Add_Request_Failed_Count() {
	
	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Unit_Info.Req_Failed++
	
	
	go func ()  {
		defer recover()
		appu.AppFramework.Add_Request_Failed_Count()
	}()
}

// IsInitialized returns the initialize status of the application unit
func (appu *AUBase) IsInitialized() bool {
	return appu.Is_Initialized
}

/// Returns the ZAU instance ID
func (appu *AUBase) Get_ID() (instance_ID int) {
	if !appu.Is_Initialized {
		fmt.Println("Not Initialized.")
		return 0
	}
	return appu.ID
}


/// Returns the Application PID
func (appu *AUBase) Get_PID() ( int) {
	if !appu.Is_Initialized {
		fmt.Println("Not Initialized.")
		return 0
	}
	return appu.AppFramework.PID()
}


/// Returns the Application PID
func (appu *AUBase) Get_AUID() ( string) {
	if !appu.Is_Initialized {
		fmt.Println("Not Initialized.")
		return ""
	}
	
	return strconv.Itoa(appu.ID) + "-" + strconv.Itoa(appu.AppFramework.PID())
}


// IsStarted returns the start status of the application unit
func (appu *AUBase) IsStarted() bool {
	return appu.Is_Started
}

func (appu *AUBase) Status() *atypes.AppUnitInfo {
	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Read_Memory_Usage()
	return appu.Unit_Info
}


// Increase_Active_Count +1 total active processing message count of 
func (appu *AUBase) Increse_Active_Count() {
	
		appu.Info_Lock.Lock()
		defer appu.Info_Lock.Unlock()
		appu.Unit_Info.Active++
}

// Decrease_Active_Count -1 total active processing message count of 
func (appu *AUBase) Decrease_Active_Count() {
	
	appu.Info_Lock.Lock()
	defer appu.Info_Lock.Unlock()
	appu.Unit_Info.Active--
}


func (appu *AUBase) Generate_Monitoring_Message(pAppID string, pID string, pStatus string, pInfo map[string]string) []byte {
	
	defer recover()
	
	if appu.AppFramework == nil {
		return nil
	}

	if _msg, _err := json.Marshal(&autypes.Status{AppID: pAppID, ID: pID, Status: pStatus, Info: pInfo}); _err != nil {
		return nil
	} else {
		return _msg
	}
}


// ConvertToFloat32 converts the given value to float32 type.
//
// Returns float32 value when success
// If failed then returns -1
func (appu *AUBase) ConvertToFloat32(pValue string) float32 {

	if pValue == "" {
		return -1
	}

	if _val, _err := strconv.ParseFloat(pValue, 32); _err != nil {
		return -1
	} else {
		return float32(_val)
	}
}


// ConvertToInt32 converts the given value to Int32 type.
//
// Returns int32 value 
// If failed then returns -1 
func (appu *AUBase) ConvertToInt32(pValue string) int32 {

	if pValue == "" {
		return -1
	}

	if _val, _err := strconv.ParseInt(pValue, 10,32); _err != nil {
		return -1
	} else {
		return int32(_val)
	}
}


// ******** PLUGIN functions *********************/
// Get_RESTClient returns the REST client plugin instance
func (appu *AUBase) Get_RESTClient(pType *string) (ihttp.IAHTTPClient, error) {
	if appu.AppFramework == nil {
		return nil, errors.New("app instance is not initialized")
	} else {
		return appu.AppFramework.Get_RESTClient(pType)
	}
}

// Get_RESTClient returns the Web Socket client plugin instance
func (appu *AUBase) Get_WSClient(pType *string) (iawsclient.IAWSClient, error) {
	if appu.AppFramework == nil {
		return nil, errors.New("app instance is not initialized")
	} else {
		return appu.AppFramework.Get_WSClient(pType)
	}
}


// Get_Mailer returns the EMail client library plugin instance
func (appu *AUBase) Get_Mailer(pType *string) (amailer.IAMailMessage, error) {
	if appu.AppFramework == nil {
		return nil, errors.New("app instance is not initialized")
	} else {
		return appu.AppFramework.Get_Mailer(pType)
	}
}
	

// Get_RESTClient returns the Logger plugin instance
func (appu *AUBase) ExecuteandFetch(os_command *string) (string, error) {
	if appu.AppFramework == nil {
		return "", errors.New("app instance is not initialized")
	} else {
		return appu.AppFramework.Execute_Command(os_command)
	}
}

func (appu *AUBase) Send_Monitor_Message(pMessage []byte) {
	
	go func ()  {
		defer recover()
		appu.AppFramework.Send_Monitor_Message(pMessage)
	}()
}

func (appu *AUBase) Write2Log(log_entry string, log_level atypes.LogLevel) {
	go func ()  {
		defer recover()
		
		appu.AppFramework.Write2Log(log_entry, log_level)
	}()
}


func (appu *AUBase) Read_Memory_Usage() {
	var _currentMem runtime.MemStats
	runtime.ReadMemStats(&_currentMem)
	
	appu.Unit_Info.Mem_Usage.Heap=_currentMem.Alloc-appu.Unit_Info.Mem_Usage.Heap
	appu.Unit_Info.Mem_Usage.HeapAlloc=_currentMem.HeapAlloc-appu.Unit_Info.Mem_Usage.HeapAlloc
	appu.Unit_Info.Mem_Usage.Total=_currentMem.TotalAlloc-appu.Unit_Info.Mem_Usage.Total
}
