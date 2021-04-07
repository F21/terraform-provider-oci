// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v38/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v38/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OnPremConnectorRequiredOnlyResource = OnPremConnectorResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Required, Create, onPremConnectorRepresentation)

	OnPremConnectorResourceConfig = OnPremConnectorResourceDependencies +
		generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Optional, Update, onPremConnectorRepresentation)

	onPremConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"on_prem_connector_id": Representation{repType: Required, create: `${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`},
	}

	onPremConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                      Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"on_prem_connector_id":              Representation{repType: Optional, create: `${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`},
		"on_prem_connector_lifecycle_state": Representation{repType: Optional, create: `INACTIVE`},
		"filter":                            RepresentationGroup{Required, onPremConnectorDataSourceFilterRepresentation}}
	onPremConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`}},
	}

	onPremConnectorRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      RepresentationGroup{Required, ignoreDefinedTagsDS},
	}

	ignoreDefinedTagsDS = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	OnPremConnectorResourceDependencies = DefinedTagsDependencies
)

func TestDataSafeOnPremConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeOnPremConnectorResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_on_prem_connector.test_on_prem_connector"
	datasourceName := "data.oci_data_safe_on_prem_connectors.test_on_prem_connectors"
	singularDatasourceName := "data.oci_data_safe_on_prem_connector.test_on_prem_connector"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDataSafeOnPremConnectorDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Required, Create, onPremConnectorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Optional, Create, onPremConnectorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OnPremConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Optional, Create,
						representationCopyWithNewProperties(onPremConnectorRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Optional, Update, onPremConnectorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_data_safe_on_prem_connectors", "test_on_prem_connectors", Optional, Update, onPremConnectorDataSourceRepresentation) +
					compartmentIdVariableStr + OnPremConnectorResourceDependencies +
					generateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Optional, Update, onPremConnectorRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connector_id"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connector_lifecycle_state", "INACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.created_version"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", Required, Create, onPremConnectorSingularDataSourceRepresentation) +
					compartmentIdVariableStr + OnPremConnectorResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "on_prem_connector_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "available_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_version"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + OnPremConnectorResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDataSafeOnPremConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_on_prem_connector" {
			noResourceFound = false
			request := oci_data_safe.GetOnPremConnectorRequest{}

			tmp := rs.Primary.ID
			request.OnPremConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "data_safe")

			response, err := client.GetOnPremConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.OnPremConnectorLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataSafeOnPremConnector") {
		resource.AddTestSweepers("DataSafeOnPremConnector", &resource.Sweeper{
			Name:         "DataSafeOnPremConnector",
			Dependencies: DependencyGraph["onPremConnector"],
			F:            sweepDataSafeOnPremConnectorResource,
		})
	}
}

func sweepDataSafeOnPremConnectorResource(compartment string) error {
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()
	onPremConnectorIds, err := getOnPremConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, onPremConnectorId := range onPremConnectorIds {
		if ok := SweeperDefaultResourceId[onPremConnectorId]; !ok {
			deleteOnPremConnectorRequest := oci_data_safe.DeleteOnPremConnectorRequest{}

			deleteOnPremConnectorRequest.OnPremConnectorId = &onPremConnectorId

			deleteOnPremConnectorRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteOnPremConnector(context.Background(), deleteOnPremConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OnPremConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", onPremConnectorId, error)
				continue
			}
			waitTillCondition(testAccProvider, &onPremConnectorId, onPremConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				onPremConnectorSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getOnPremConnectorIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "OnPremConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := GetTestClients(&schema.ResourceData{}).dataSafeClient()

	listOnPremConnectorsRequest := oci_data_safe.ListOnPremConnectorsRequest{}
	listOnPremConnectorsRequest.CompartmentId = &compartmentId
	listOnPremConnectorsResponse, err := dataSafeClient.ListOnPremConnectors(context.Background(), listOnPremConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OnPremConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, onPremConnector := range listOnPremConnectorsResponse.Items {
		id := *onPremConnector.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "OnPremConnectorId", id)
	}
	return resourceIds, nil
}

func onPremConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if onPremConnectorResponse, ok := response.Response.(oci_data_safe.GetOnPremConnectorResponse); ok {
		return string(onPremConnectorResponse.LifecycleState) != string(oci_data_safe.OnPremConnectorLifecycleStateDeleted)
	}
	return false
}

func onPremConnectorSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataSafeClient().GetOnPremConnector(context.Background(), oci_data_safe.GetOnPremConnectorRequest{
		OnPremConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}