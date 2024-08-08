package utils

import (
	"fmt"
	"reflect"
)

// GetTRNSegmentByType returns a segment from the TRN APIs based on the type property
func GetTRNSegmentByType[T any](segments []T, targetType string) (*T, error) {
	for _, segment := range segments {
		v := reflect.ValueOf(segment)
		field := v.FieldByName("Type")
		if field.IsValid() && field.String() == targetType {
			return &segment, nil
		}
	}
	return nil, fmt.Errorf("no segment found with type %s", targetType)
}
