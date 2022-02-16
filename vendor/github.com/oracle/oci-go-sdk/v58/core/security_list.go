// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SecurityList A set of virtual firewall rules for your VCN. Security lists are configured at the subnet
// level, but the rules are applied to the ingress and egress traffic for the individual instances
// in the subnet. The rules can be stateful or stateless. For more information, see
// Security Lists (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm).
// **Note:** Compare security lists to NetworkSecurityGroups,
// which let you apply a set of security rules to a *specific set of VNICs* instead of an entire
// subnet. Oracle recommends using network security groups instead of security lists, although you
// can use either or both together.
// **Important:** Oracle Cloud Infrastructure Compute service images automatically include firewall rules (for example,
// Linux iptables, Windows firewall). If there are issues with some type of access to an instance,
// make sure both the security lists associated with the instance's subnet and the instance's
// firewall rules are set correctly.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type SecurityList struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the security list.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules []EgressSecurityRule `mandatory:"true" json:"egressSecurityRules"`

	// The security list's Oracle Cloud ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules []IngressSecurityRule `mandatory:"true" json:"ingressSecurityRules"`

	// The security list's current state.
	LifecycleState SecurityListLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the security list was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the security list belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m SecurityList) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityList) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityListLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityListLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityListLifecycleStateEnum Enum with underlying type: string
type SecurityListLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityListLifecycleStateEnum
const (
	SecurityListLifecycleStateProvisioning SecurityListLifecycleStateEnum = "PROVISIONING"
	SecurityListLifecycleStateAvailable    SecurityListLifecycleStateEnum = "AVAILABLE"
	SecurityListLifecycleStateTerminating  SecurityListLifecycleStateEnum = "TERMINATING"
	SecurityListLifecycleStateTerminated   SecurityListLifecycleStateEnum = "TERMINATED"
)

var mappingSecurityListLifecycleStateEnum = map[string]SecurityListLifecycleStateEnum{
	"PROVISIONING": SecurityListLifecycleStateProvisioning,
	"AVAILABLE":    SecurityListLifecycleStateAvailable,
	"TERMINATING":  SecurityListLifecycleStateTerminating,
	"TERMINATED":   SecurityListLifecycleStateTerminated,
}

// GetSecurityListLifecycleStateEnumValues Enumerates the set of values for SecurityListLifecycleStateEnum
func GetSecurityListLifecycleStateEnumValues() []SecurityListLifecycleStateEnum {
	values := make([]SecurityListLifecycleStateEnum, 0)
	for _, v := range mappingSecurityListLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityListLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityListLifecycleStateEnum
func GetSecurityListLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingSecurityListLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityListLifecycleStateEnum(val string) (SecurityListLifecycleStateEnum, bool) {
	mappingSecurityListLifecycleStateEnumIgnoreCase := make(map[string]SecurityListLifecycleStateEnum)
	for k, v := range mappingSecurityListLifecycleStateEnum {
		mappingSecurityListLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSecurityListLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}