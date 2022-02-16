// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BackupLocationUri PreAuthenticated object storage URI to upload or download the backup
type BackupLocationUri struct {
	Uri *string `mandatory:"true" json:"uri"`
}

func (m BackupLocationUri) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupLocationUri) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BackupLocationUri) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBackupLocationUri BackupLocationUri
	s := struct {
		DiscriminatorParam string `json:"destination"`
		MarshalTypeBackupLocationUri
	}{
		"PRE_AUTHENTICATED_REQUEST_URI",
		(MarshalTypeBackupLocationUri)(m),
	}

	return json.Marshal(&s)
}