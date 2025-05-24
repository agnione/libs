// build package provides type to represents the status of the application units
//
// This package includes following structures:
//
//   - Status
//
//     ---------------------------------------------------------------------------------------------------------------------
//     Author        :   D. Ajith Nilantha de Silva ajithdesilva@gmail.com | 26/01/2024
//     Copyright     :   Open source MIT License
//     Class/module  :   autypes - AgniOne Application Framework
//     Objective     :   Define the common types for AgniOne Application Units
//     Types defined here will be used by KAUs
//     ---------------------------------------------------------------------------------------------------------------------
//     Author			Date		Action		Description
//     ---------------------------------------------------------------------------------------------------------------------
//     Ajith de Silva		24/04/2024	Created 	Created the initial version
//     ---------------------------------------------------------------------------------------------------------------------
package autypes

// Status type to holds the current status of the call in each state.
//
// Used to pass the monitoring message
type Status struct {
  AppID string
  ID string
  Info map[string]string
  Status string
}

