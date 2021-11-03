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
	"github.com/oracle/oci-go-sdk/v54/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v54/loganalytics"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NamespaceScheduledTaskRequiredOnlyResource = NamespaceScheduledTaskResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Required, Create, purgeTaskRepresentation)

	NamespaceScheduledTaskResourceConfig = NamespaceScheduledTaskResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Optional, Update, purgeTaskRepresentation)

	namespaceScheduledTaskSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":         Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"scheduled_task_id": Representation{RepType: Required, Create: `${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.scheduled_task_id}`},
	}

	namespaceScheduledTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"display_name":   Representation{RepType: Optional, Create: `tfPurgeTask`, Update: `tfPurgeTask2`},
		"task_type":      Representation{RepType: Optional, Create: `PURGE`},
		"filter":         RepresentationGroup{Required, namespaceScheduledTaskDataSourceFilterRepresentation}}

	namespaceScheduledTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.id}`}},
	}

	purgeTaskRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"kind":           Representation{RepType: Required, Create: `STANDARD`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"task_type":      Representation{RepType: Required, Create: `PURGE`},
		"action":         RepresentationGroup{Required, purgeActionRepresentation},
		"schedules":      RepresentationGroup{Required, schedulesRepresentation},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{RepType: Required, Create: `tfPurgeTask`, Update: `tfPurgeTask2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	purgeActionRepresentation = map[string]interface{}{
		"type":                      Representation{RepType: Required, Create: `PURGE`},
		"compartment_id_in_subtree": Representation{RepType: Required, Create: `false`},
		"data_type":                 Representation{RepType: Required, Create: `LOG`},
		"purge_compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"purge_duration":            Representation{RepType: Required, Create: `-P7D`},
		"query_string":              Representation{RepType: Required, Create: `fake_query`},
	}

	schedulesRepresentation = map[string]interface{}{
		"schedule": []RepresentationGroup{{Required, fixedSchedulesRepresentation}},
	}

	/*
		    TODO: use for acceleration tasks. for purge, only one schedule is allowed.
			schedulesUpdatedRepresentation = map[string]interface{}{
				"schedule": []RepresentationGroup{{Required, fixedSchedulesRepresentation}, {Required, cronSchedulesRepresentation}},
			}
	*/

	fixedSchedulesRepresentation = map[string]interface{}{
		"type":               Representation{RepType: Required, Create: `FIXED_FREQUENCY`, Update: `FIXED_FREQUENCY`},
		"recurring_interval": Representation{RepType: Required, Create: `P1D`, Update: `P2D`},
		"repeat_count":       Representation{RepType: Required, Create: `4`, Update: `6`},
		"misfire_policy":     Representation{RepType: Required, Create: `RETRY_ONCE`, Update: `RETRY_INDEFINITELY`},
	}
	cronSchedulesRepresentation = map[string]interface{}{
		"type":           Representation{RepType: Required, Create: `CRON`, Update: `CRON`},
		"expression":     Representation{RepType: Required, Create: `0 0 * ? * 2,3,4,5,6`, Update: `0 0 * ? * 2,3`},
		"misfire_policy": Representation{RepType: Required, Create: `RETRY_INDEFINITELY`, Update: `RETRY_ONCE`},
		"time_zone":      Representation{RepType: Required, Create: `UTC`, Update: `UTC`},
	}

	NamespaceScheduledTaskResourceDependencies = DefinedTagsDependencies +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestLogAnalyticsNamespaceScheduledTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceScheduledTaskResource_basic")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentIdU := GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task"
	datasourceName := "data.oci_log_analytics_namespace_scheduled_tasks.test_namespace_scheduled_tasks"
	singularDatasourceName := "data.oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task"
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLogAnalyticsNamespaceScheduledTaskDestroy,
		Steps: []resource.TestStep{
			// verify creation of purge task
			{
				Config: config + compartmentIdVariableStr + NamespaceScheduledTaskResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Required, Create, purgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify update to the compartment (the compartment will be switched back in the next step) of purge task
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NamespaceScheduledTaskResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Required, Create,
						RepresentationCopyWithNewProperties(purgeTaskRepresentation, map[string]interface{}{
							"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NamespaceScheduledTaskResourceDependencies,
			},
			// verify creation of purge task with optionals
			{
				Config: config + compartmentIdVariableStr + NamespaceScheduledTaskResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Optional, Create, purgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters of purge task
			{
				Config: config + compartmentIdVariableStr + NamespaceScheduledTaskResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Optional, Update, purgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask2"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					// TODO: add check for defined_tags value change
					// TODO: add check for freeform tags change
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_INDEFINITELY",
						"recurring_interval": "P2D",
						"repeat_count":       "6",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_tasks", "test_namespace_scheduled_tasks", Optional, Update, namespaceScheduledTaskDataSourceRepresentation) +
					compartmentIdVariableStr + NamespaceScheduledTaskResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Optional, Update, purgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "tfPurgeTask2"),
					resource.TestCheckResourceAttr(datasourceName, "task_type", "PURGE"),

					resource.TestCheckResourceAttr(datasourceName, "scheduled_task_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", Required, Create, namespaceScheduledTaskSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NamespaceScheduledTaskResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_task_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "num_occurrences"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tfPurgeTask2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "task_status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.schedule.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_INDEFINITELY",
						"recurring_interval": "P2D",
						"repeat_count":       "6",
					}, []string{}),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NamespaceScheduledTaskResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: namespaceScheduledTaskImportId,
				ImportStateVerifyIgnore: []string{
					"kind",
					"namespace",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckLogAnalyticsNamespaceScheduledTaskDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).logAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_scheduled_task" {
			noResourceFound = false
			request := oci_log_analytics.GetScheduledTaskRequest{}

			if namespace, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &namespace
			}

			if scheduledTaskId, ok := rs.Primary.Attributes["scheduled_task_id"]; ok {
				request.ScheduledTaskId = &scheduledTaskId
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "log_analytics")

			response, err := client.GetScheduledTask(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.ScheduledTaskLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
		InitDependencyGraph()
	}
	if !InSweeperExcludeList("LogAnalyticsNamespaceScheduledTask") {
		resource.AddTestSweepers("LogAnalyticsNamespaceScheduledTask", &resource.Sweeper{
			Name:         "LogAnalyticsNamespaceScheduledTask",
			Dependencies: DependencyGraph["namespaceScheduledTask"],
			F:            sweepLogAnalyticsNamespaceScheduledTaskResource,
		})
	}
}

func sweepLogAnalyticsNamespaceScheduledTaskResource(compartment string) error {
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()
	namespaceScheduledTaskIds, err := getNamespaceScheduledTaskIds(compartment)
	if err != nil {
		return err
	}
	for _, namespaceScheduledTaskId := range namespaceScheduledTaskIds {
		if ok := SweeperDefaultResourceId[namespaceScheduledTaskId]; !ok {
			deleteScheduledTaskRequest := oci_log_analytics.DeleteScheduledTaskRequest{}

			deleteScheduledTaskRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteScheduledTask(context.Background(), deleteScheduledTaskRequest)
			if error != nil {
				fmt.Printf("Error deleting NamespaceScheduledTask %s %s, It is possible that the resource is already deleted. Please verify manually \n", namespaceScheduledTaskId, error)
				continue
			}
			WaitTillCondition(TestAccProvider, &namespaceScheduledTaskId, namespaceScheduledTaskSweepWaitCondition, time.Duration(3*time.Minute),
				namespaceScheduledTaskSweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getNamespaceScheduledTaskIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "NamespaceScheduledTaskId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := GetTestClients(&schema.ResourceData{}).logAnalyticsClient()

	listScheduledTasksRequest := oci_log_analytics.ListScheduledTasksRequest{}
	listScheduledTasksRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for NamespaceScheduledTask resource requests \n")
	}
	for _, namespace := range namespaces {
		listScheduledTasksRequest.NamespaceName = &namespace

		listScheduledTasksResponse, err := logAnalyticsClient.ListScheduledTasks(context.Background(), listScheduledTasksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NamespaceScheduledTask list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, namespaceScheduledTask := range listScheduledTasksResponse.Items {
			id := *namespaceScheduledTask.Id
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "NamespaceScheduledTaskId", id)
		}

	}
	return resourceIds, nil
}

func namespaceScheduledTaskSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if namespaceScheduledTaskResponse, ok := response.Response.(oci_log_analytics.GetScheduledTaskResponse); ok {
		return namespaceScheduledTaskResponse.GetLifecycleState() != oci_log_analytics.ScheduledTaskLifecycleStateDeleted
	}
	return false
}

func namespaceScheduledTaskSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.logAnalyticsClient().GetScheduledTask(context.Background(), oci_log_analytics.GetScheduledTaskRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func namespaceScheduledTaskImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_scheduled_task" {
			return getNamespaceScheduledTaskCompositeId(rs.Primary.Attributes["namespace"], rs.Primary.Attributes["scheduled_task_id"]), nil
			//return rs.Primary.Attributes["ID"], nil
		}
	}

	return "", fmt.Errorf("unable to create import id as no resource of type oci_log_analytics_namespace_scheduled_task found in state")
}