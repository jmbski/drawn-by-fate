// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pD, err := UnmarshalPD(bytes)
//    bytes, err = pD.Marshal()

package main

import "encoding/json"

func UnmarshalPD(data []byte) (PD, error) {
	var r PD
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PD) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Meta data for a content package
type PD struct {
	// Version of the game the pack was designed for. Uses 3-part semantic versioning                            
	// (<major>.<minor>.<patch>)                                                                                 
	AppVersion                                                                                  string           `json:"app_version" yaml:"app_version"`
	// List of packs that this pack requires in order to work correctly. This pack's load order                  
	// will be automatically incremented to ensure it's loaded after its dependencies                            
	Dependencies                                                                                []PackDependency `json:"dependencies" yaml:"dependencies"`
	// Unique identifier for the mod (matches pack folder name)                                                  
	ID                                                                                          string           `json:"id" yaml:"id"`
	// Indicates when the pack should be loaded. Lower = loaded earlier (with core pack at 0).                   
	// Higher load order will overwrite lower on conflict. Load order will be automatically                      
	// adjusted based on dependencies                                                                            
	LoadOrder                                                                                   int64            `json:"load_order" yaml:"load_order"`
	// Displayed name of the pack (non-unique)                                                                   
	Name                                                                                        string           `json:"name" yaml:"name"`
	// Internal version of the pack. Use 3-part semantic versioning (<major>.<minor>.<patch>)                    
	PackVersion                                                                                 string           `json:"pack_version" yaml:"pack_version"`
	// Version of the config schemas the pack was designed with. Uses 3-part semantic versioning                 
	// (<major>.<minor>.<patch>)                                                                                 
	SchemaVersion                                                                               string           `json:"schema_version" yaml:"schema_version"`
}

// ID and version of the required pack
type PackDependency struct {
	PackID      string `json:"pack_id" yaml:"pack_id"`
	PackVersion string `json:"pack_version" yaml:"pack_version"`
}
