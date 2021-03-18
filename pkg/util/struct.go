package util

import (
	"encoding/json"
)

// StructAdaptiveCopyByJSON uses JSON to copy the elements in the source struct to the corresponding elements in the object struct
func StructAdaptiveCopyByJSON(source interface{}, object interface{}) error {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(sourceJSON, object); err != nil {
		return err
	}
	return nil
}
