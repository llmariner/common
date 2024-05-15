package assert

import (
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Equal asserts that two Gorm objects are equal except their Model field.
func Equal(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if err := clearModel(expected); err != nil {
		t.Errorf("failed to clear model from the expected object: %s", err)
		return false
	}
	if err := clearModel(actual); err != nil {
		t.Errorf("failed to clear model from the actual object: %s", err)
		return false
	}
	return assert.Equal(t, expected, actual, msgAndArgs)
}

// ElementsMatch asserts that the specified slices are equal ignoring the order of the elements.
func ElementsMatch(t assert.TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	if err := clearModelFromSlice(listA); err != nil {
		t.Errorf("failed to clear model from the expected object: %s", err)
		return false
	}
	if err := clearModelFromSlice(listB); err != nil {
		t.Errorf("failed to clear model from the actual object: %s", err)
		return false
	}
	return assert.ElementsMatch(t, listA, listB, msgAndArgs)
}

func clearModelFromSlice(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Type().Kind() != reflect.Slice {
		return fmt.Errorf("%+v is not a slice", v)
	}
	for i := 0; i < v.Len(); i++ {
		if err := clearModel(v.Index(i).Interface()); err != nil {
			return err
		}
	}
	return nil
}

func clearModel(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("%+v is not a pointer", v)
	}

	fv := v.Elem().FieldByName("Model")
	if !fv.IsValid() {
		return fmt.Errorf("%+v does not have the Model field", v)
	}
	fv.Set(reflect.ValueOf(gorm.Model{}))
	return nil
}
