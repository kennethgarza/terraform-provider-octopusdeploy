package octopusdeploy

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type DeploymentProcessService struct {
	sling *sling.Sling
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	return &DeploymentProcessService{
		sling: sling,
	}
}

type DeploymentProcesses struct {
	Items []DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	ID             string           `json:"Id,omitempty"`
	LastModifiedBy string           `json:"LastModifiedBy,omitempty"`
	LastModifiedOn string           `json:"LastModifiedOn,omitempty"`
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	Links          Links            `json:"Links,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	Steps          []DeploymentStep `json:"Steps"`
	Version        *int32           `json:"Version"`
}

type DeploymentStep struct {
	ID                 string             `json:"Id"`
	Name               string             `json:"Name"`
	PackageRequirement string             `json:"PackageRequirement,omitempty"`                                         // may need its own model / enum
	Properties         map[string]string  `json:"Properties"`                                                           // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Condition          string             `json:"Condition,omitempty" validate:"oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	StartTrigger       string             `json:"StartTrigger,omitempty" validate:"oneof=StartAfterPrevious StartWithPrevious"`
	Actions            []DeploymentAction `json:"Actions"`
}

type DeploymentAction struct {
	ID                            string            `json:"Id"`
	Name                          string            `json:"Name"`
	ActionType                    string            `json:"ActionType"`
	IsDisabled                    bool              `json:"IsDisabled"`
	CanBeUsedForProjectVersioning bool              `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string          `json:"Environments"`
	ExcludedEnvironments          []string          `json:"ExcludedEnvironments"`
	Channels                      []string          `json:"Channels"`
	TenantTags                    []string          `json:"TenantTags"`
	Properties                    map[string]string `json:"Properties"`     // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	LastModifiedOn                string            `json:"LastModifiedOn"` // datetime
	LastModifiedBy                string            `json:"LastModifiedBy"`
	Links                         Links             `json:"Links"` // may be wrong
}

func (d *DeploymentProcess) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		return err
	}

	return nil
}

func (d *DeploymentProcessService) Get(deploymentProcessID string) (*DeploymentProcess, error) {
	var deploymentProcess DeploymentProcess
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcessID)

	resp, err := d.sling.New().Get(path).Receive(&deploymentProcess, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &deploymentProcess, err
}

func (d *DeploymentProcessService) GetAll() (*[]DeploymentProcess, error) {
	var listOfDeploymentProcess []DeploymentProcess
	path := fmt.Sprintf("deploymentprocesses")

	for {
		var deploymentProcesses DeploymentProcesses
		octopusDeployError := new(APIError)

		resp, err := d.sling.New().Get(path).Receive(&deploymentProcesses, &octopusDeployError)

		apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

		if apiErrorCheck != nil {
			return nil, apiErrorCheck
		}

		for _, deploymentProcess := range deploymentProcesses.Items {
			listOfDeploymentProcess = append(listOfDeploymentProcess, deploymentProcess)
		}

		if deploymentProcesses.PagedResults.Links.PageNext != "" {
			path = deploymentProcesses.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return &listOfDeploymentProcess, nil // no more pages to go through
}

func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcess) (*DeploymentProcess, error) {
	var updated DeploymentProcess
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcess.ID)

	resp, err := s.sling.New().Put(path).BodyJSON(deploymentProcess).Receive(&updated, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)

	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return &updated, nil
}