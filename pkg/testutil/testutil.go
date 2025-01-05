package testutil

import "reflect"

func IgnoreFields[T any](v T, ignoreFields ...string) T {
	// Convert ignoreFields to map for O(1) lookup
	ignoreMap := make(map[string]bool)
	for _, field := range ignoreFields {
		ignoreMap[field] = true
	}

	// Get reflect.Value of the input
	val := reflect.ValueOf(v)
	isPtr := val.Kind() == reflect.Ptr

	// If pointer, get the element it points to
	if isPtr {
		val = val.Elem()
	}

	// Only proceed if we have a struct
	if val.Kind() != reflect.Struct {
		return v
	}

	// Create a new instance of the same type
	result := reflect.New(val.Type()).Elem()

	// Copy all fields from the original value
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if ignoreMap[field.Name] {
			// Set zero value for ignored fields
			result.Field(i).Set(reflect.Zero(field.Type))
		} else {
			// Copy the value for non-ignored fields
			result.Field(i).Set(val.Field(i))
		}
	}

	// If original value was a pointer, return a pointer
	if isPtr {
		ptr := reflect.New(val.Type())
		ptr.Elem().Set(result)
		return ptr.Interface().(T)
	}

	// Convert back to original type
	return result.Interface().(T)
}
