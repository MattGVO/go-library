package utils

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// ErrorResponse standardizes the error response format
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{"error": message})
}

func QueryParamInt(param string, defaultVal int) int {
	result, err := strconv.Atoi(param)
	if err != nil {
		return defaultVal
	}
	return result
  }

  func ScanRows(rows *sql.Rows, modelSlice interface{}) error {
	// Get the reflection value of the model slice
	sliceValue := reflect.ValueOf(modelSlice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("modelSlice must be a pointer to a slice")
	}

	// Get the element type (the model struct type)
	elemType := sliceValue.Elem().Type().Elem()

	// Iterate over the rows
	for rows.Next() {
		// Create a new instance of the model type
		elem := reflect.New(elemType).Elem()

		// Prepare the fields to be scanned into the model
		scanTargets := make([]interface{}, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			scanTargets[i] = elem.Field(i).Addr().Interface()
		}

		// Scan the row into the model
		if err := rows.Scan(scanTargets...); err != nil {
			return err
		}

		// Append the populated model to the slice
		sliceValue.Elem().Set(reflect.Append(sliceValue.Elem(), elem))
	}

	return rows.Err()
}

func ScanRow(row *sql.Row, model interface{}) error {
	// Get the reflection value of the model
	modelValue := reflect.ValueOf(model)

	// Ensure the model is a pointer to a struct
	if modelValue.Kind() != reflect.Ptr || modelValue.Elem().Kind() != reflect.Struct {
		return errors.New("model must be a pointer to a struct")
	}

	// Get the element type (the struct type)
	elemType := modelValue.Elem().Type()

	// Prepare the fields to be scanned into the model
	scanTargets := make([]interface{}, elemType.NumField())
	for i := 0; i < elemType.NumField(); i++ {
		scanTargets[i] = modelValue.Elem().Field(i).Addr().Interface()
	}

	// Scan the row into the model
	if err := row.Scan(scanTargets...); err != nil {
		return err
	}

	return nil
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetDueDate() string {
	return time.Now().AddDate(0, 0, 14).Format("2006-01-02")
}
