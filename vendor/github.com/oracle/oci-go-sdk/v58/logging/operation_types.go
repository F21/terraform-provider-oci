// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateLog           OperationTypesEnum = "CREATE_LOG"
	OperationTypesUpdateLog           OperationTypesEnum = "UPDATE_LOG"
	OperationTypesDeleteLog           OperationTypesEnum = "DELETE_LOG"
	OperationTypesMoveLog             OperationTypesEnum = "MOVE_LOG"
	OperationTypesCreateLogGroup      OperationTypesEnum = "CREATE_LOG_GROUP"
	OperationTypesUpdateLogGroup      OperationTypesEnum = "UPDATE_LOG_GROUP"
	OperationTypesDeleteLogGroup      OperationTypesEnum = "DELETE_LOG_GROUP"
	OperationTypesMoveLogGroup        OperationTypesEnum = "MOVE_LOG_GROUP"
	OperationTypesCreateConfiguration OperationTypesEnum = "CREATE_CONFIGURATION"
	OperationTypesUpdateConfiguration OperationTypesEnum = "UPDATE_CONFIGURATION"
	OperationTypesDeleteConfiguration OperationTypesEnum = "DELETE_CONFIGURATION"
	OperationTypesMoveConfiguration   OperationTypesEnum = "MOVE_CONFIGURATION"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_LOG":           OperationTypesCreateLog,
	"UPDATE_LOG":           OperationTypesUpdateLog,
	"DELETE_LOG":           OperationTypesDeleteLog,
	"MOVE_LOG":             OperationTypesMoveLog,
	"CREATE_LOG_GROUP":     OperationTypesCreateLogGroup,
	"UPDATE_LOG_GROUP":     OperationTypesUpdateLogGroup,
	"DELETE_LOG_GROUP":     OperationTypesDeleteLogGroup,
	"MOVE_LOG_GROUP":       OperationTypesMoveLogGroup,
	"CREATE_CONFIGURATION": OperationTypesCreateConfiguration,
	"UPDATE_CONFIGURATION": OperationTypesUpdateConfiguration,
	"DELETE_CONFIGURATION": OperationTypesDeleteConfiguration,
	"MOVE_CONFIGURATION":   OperationTypesMoveConfiguration,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_LOG",
		"UPDATE_LOG",
		"DELETE_LOG",
		"MOVE_LOG",
		"CREATE_LOG_GROUP",
		"UPDATE_LOG_GROUP",
		"DELETE_LOG_GROUP",
		"MOVE_LOG_GROUP",
		"CREATE_CONFIGURATION",
		"UPDATE_CONFIGURATION",
		"DELETE_CONFIGURATION",
		"MOVE_CONFIGURATION",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	mappingOperationTypesEnumIgnoreCase := make(map[string]OperationTypesEnum)
	for k, v := range mappingOperationTypesEnum {
		mappingOperationTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOperationTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}