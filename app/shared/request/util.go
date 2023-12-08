package request

import (
	"fmt"
	"net/url"
	"reflect"
)

func BindForm(form url.Values, dest interface{}) error {
	destValue := reflect.ValueOf(dest)

	// Ensure dest is a pointer to a struct
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a pointer to a struct")
	}

	destElem := destValue.Elem()
	destType := destElem.Type()

	// Convert url.Values to map
	formMap := make(map[string]string)
	for key, values := range form {
		if len(values) > 0 {
			formMap[key] = values[0]
		}
	}

	// Iterate through struct fields and set values
	for i := 0; i < destElem.NumField(); i++ {
		field := destElem.Field(i)
		fieldType := destType.Field(i)

		// Use form tag if present, otherwise use field name
		formName := fieldType.Tag.Get("form")
		if formName == "" {
			formName = fieldType.Name
		}

		// Look up form value by name
		value, ok := formMap[formName]
		if ok {
			// Convert string value to the field's type and set it
			if field.CanSet() {
				err := setField(field, value)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func setField(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	default:
		// Handle unsupported field types
		return fmt.Errorf("unsupported field type: %v", field.Kind())
	}

	return nil
}
