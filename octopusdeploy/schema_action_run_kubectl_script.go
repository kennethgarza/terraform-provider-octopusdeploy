package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getRunKubectlScriptSchema() *schema.Schema {
	actionSchema, element := getActionSchema()
	addExecutionLocationSchema(element)
	addScriptFromPackageSchema(element)
	addPackagesSchema(element, false)
	addWorkerPoolSchema(element)
	addWorkerPoolVariableSchema(element)

	element.Schema["script_body"] = &schema.Schema{
		Optional: true,
		Type:     schema.TypeString,
	}

	element.Schema["script_syntax"] = &schema.Schema{
		Computed: true,
		Optional: true,
		Type:     schema.TypeString,
	}

	element.Schema["variable_substitution_in_files"] = &schema.Schema{
		Description: "A newline-separated list of file names to transform, relative to the package contents. Extended wildcard syntax is supported.",
		Optional:    true,
		Type:        schema.TypeString,
	}
	return actionSchema
}

func expandRunKubectlScriptAction(flattenedAction map[string]interface{}) *deployments.DeploymentAction {
	action := expandRunScriptAction(flattenedAction)
	action.ActionType = "Octopus.KubernetesRunScript"
	return action
}

func flattenKubernetesRunScriptAction(action *deployments.DeploymentAction) map[string]interface{} {
	return flattenRunScriptAction(action)
}
