/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// A workflow execution represents an instantiated workflow, including all inputs and additional metadata as well as computed results included state, outputs, and duration-based attributes. Used as a response object used in Get and List execution requests.
type AdminExecution struct {
	// Unique identifier of the workflow execution.
	Id *CoreWorkflowExecutionIdentifier `json:"id,omitempty"`
	// User-provided configuration and inputs for launching the execution.
	Spec *AdminExecutionSpec `json:"spec,omitempty"`
	// Execution results.
	Closure *AdminExecutionClosure `json:"closure,omitempty"`
}