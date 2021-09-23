// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// UpdateDirectoryObject Note: Deprecated. Use the new resource model APIs instead.
// Directory object details, used to define either import or export directory objects in Data Pump Settings.
// Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error.
// Export directory will error if there are database link details specified.
// If an empty object is specified, the stored Directory Object details will be removed.
type UpdateDirectoryObject struct {

	// Name of directory object in database
	Name *string `mandatory:"false" json:"name"`

	// Absolute path of directory on database server
	Path *string `mandatory:"false" json:"path"`
}

func (m UpdateDirectoryObject) String() string {
	return common.PointerString(m)
}
