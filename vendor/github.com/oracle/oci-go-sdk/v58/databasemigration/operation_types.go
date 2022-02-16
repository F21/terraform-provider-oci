// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateAgent       OperationTypesEnum = "CREATE_AGENT"
	OperationTypesDeleteAgent       OperationTypesEnum = "DELETE_AGENT"
	OperationTypesCreateMigration   OperationTypesEnum = "CREATE_MIGRATION"
	OperationTypesCloneMigration    OperationTypesEnum = "CLONE_MIGRATION"
	OperationTypesDeleteMigration   OperationTypesEnum = "DELETE_MIGRATION"
	OperationTypesUpdateMigration   OperationTypesEnum = "UPDATE_MIGRATION"
	OperationTypesStartMigration    OperationTypesEnum = "START_MIGRATION"
	OperationTypesValidateMigration OperationTypesEnum = "VALIDATE_MIGRATION"
	OperationTypesCreateConnection  OperationTypesEnum = "CREATE_CONNECTION"
	OperationTypesDeleteConnection  OperationTypesEnum = "DELETE_CONNECTION"
	OperationTypesUpdateConnection  OperationTypesEnum = "UPDATE_CONNECTION"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_AGENT":       OperationTypesCreateAgent,
	"DELETE_AGENT":       OperationTypesDeleteAgent,
	"CREATE_MIGRATION":   OperationTypesCreateMigration,
	"CLONE_MIGRATION":    OperationTypesCloneMigration,
	"DELETE_MIGRATION":   OperationTypesDeleteMigration,
	"UPDATE_MIGRATION":   OperationTypesUpdateMigration,
	"START_MIGRATION":    OperationTypesStartMigration,
	"VALIDATE_MIGRATION": OperationTypesValidateMigration,
	"CREATE_CONNECTION":  OperationTypesCreateConnection,
	"DELETE_CONNECTION":  OperationTypesDeleteConnection,
	"UPDATE_CONNECTION":  OperationTypesUpdateConnection,
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
		"CREATE_AGENT",
		"DELETE_AGENT",
		"CREATE_MIGRATION",
		"CLONE_MIGRATION",
		"DELETE_MIGRATION",
		"UPDATE_MIGRATION",
		"START_MIGRATION",
		"VALIDATE_MIGRATION",
		"CREATE_CONNECTION",
		"DELETE_CONNECTION",
		"UPDATE_CONNECTION",
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