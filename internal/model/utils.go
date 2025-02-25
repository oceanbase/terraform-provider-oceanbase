/*
 * Copyright 2025 OceanBase
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func parseListValue[T any](list basetypes.ListValue, parser func(basetypes.ObjectValue) (*T, error)) ([]T, error) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}

	elements := list.Elements()
	result := make([]T, 0, len(elements))

	for i, elem := range elements {
		objValue, ok := elem.(basetypes.ObjectValue)
		if !ok {
			return nil, fmt.Errorf("element at index %d is not an ObjectValue: got %T", i, elem)
		}
		parsed, err := parser(objValue)
		if err != nil {
			return nil, fmt.Errorf("failed to parse element at index %d: %w", i, err)
		}
		result = append(result, *parsed)
	}

	return result, nil
}

func parseObjectValue[T any](obj basetypes.ObjectValue) (*T, error) {
	// Check if the object is null or unknown
	if obj.IsNull() || obj.IsUnknown() {
		return nil, nil
	}

	// Create a new struct of the target type
	var result T

	// Use reflection to convert the object value to a struct
	val := reflect.ValueOf(&result).Elem()
	typ := val.Type()

	// Iterate through struct fields
	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)

		// Get the Terraform tag (if exists)
		tfTag := fieldType.Tag.Get("tfsdk")
		if tfTag == "" {
			continue // Skip fields without a Terraform tag
		}

		// Find the corresponding attribute in the object value
		attr, ok := obj.Attributes()[tfTag]
		if !ok {
			continue
		}

		// Convert the attribute based on its type
		switch field := val.Field(i); field.Kind() {
		case reflect.String:
			strVal, diags := attr.(basetypes.StringValue).ToStringValue(context.Background())
			if diags.HasError() {
				return nil, fmt.Errorf("error converting string: %v", diags)
			}
			field.SetString(strVal.ValueString())

		case reflect.Int, reflect.Int64:
			intVal, diags := attr.(basetypes.Int64Value).ToInt64Value(context.Background())
			if diags.HasError() {
				return nil, fmt.Errorf("error converting int64: %v", diags)
			}
			field.SetInt(intVal.ValueInt64())

		case reflect.Bool:
			boolVal, diags := attr.(basetypes.BoolValue).ToBoolValue(context.Background())
			if diags.HasError() {
				return nil, fmt.Errorf("error converting bool: %v", diags)
			}
			field.SetBool(boolVal.ValueBool())

		// Add more type conversions as needed
		default:
			return nil, fmt.Errorf("unsupported field type: %v", field.Kind())
		}
	}

	return &result, nil
}
