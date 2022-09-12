package clustergen

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// postConfig adds an cluster config from JSON received in the request body.
func PostConfigAzure(c *gin.Context) {
	id := strconv.Itoa(GenerateID())
	// Generate new ID
	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not create new ID"})
		return
	}
	provider := "azure"

	// Validate provider
	c.Bind(&clusters)
	clusterName := c.PostForm("clusterName")
	// provider := c.PostForm("provider")
	providerVersion := c.PostForm("providerVersion")
	kubernetesVersion := c.PostForm("kubernetesVersion")
	controlPlaneMachineCount := c.PostForm("controlPlaneMachineCount")
	workerMachineCount := c.PostForm("workerMachineCount")
	clusterTopology := c.PostForm("clusterTopology")
	flavor := c.PostForm("flavor")
	owner := c.PostForm("owner")
	gitHubToken := c.PostForm("gitHubToken")

	c.Bind(&clustersAzure)
	azureSubscriptionID := c.PostForm("azureSubscriptionID")
	azureTenantID := c.PostForm("azureTenantID")
	azureClientID := c.PostForm("azureClientID")
	azureClientSecret := c.PostForm("azureClientSecret")
	azureLocation := c.PostForm("azureLocation")
	azureResourceGroup := c.PostForm("azureResourceGroup")
	azureControlPlaneMachineType := c.PostForm("azureControlPlaneMachineType")
	azureNodeMachineType := c.PostForm("azureNodeMachineType")
	azureClusterIdentitySecretName := c.PostForm("azureClusterIdentitySecretName")
	azureClusterIdentitySecretNamespace := c.PostForm("azureClusterIdentitySecretNamespace")
	clusterIdentityName := c.PostForm("clusterIdentityName")

	controlPlaneMachineCountString, _ := strconv.Atoi(controlPlaneMachineCount)
	workerMachineCountString, _ := strconv.Atoi(workerMachineCount)
	clusterTopologyBool, _ := strconv.ParseBool(clusterTopology)

	newcluster := cluster{
		ID:                       id,
		ClusterName:              clusterName,
		Provider:                 provider,
		ProviderVersion:          providerVersion,
		KubernetesVersion:        kubernetesVersion,
		ControlPlaneMachineCount: controlPlaneMachineCountString,
		WorkerMachineCount:       workerMachineCountString,
		ClusterTopology:          clusterTopologyBool,
		Flavor:                   flavor,
		Owner:                    owner,
		GitHubToken:              gitHubToken,
	}

	newclusterAzure := clusterAzure{
		ID:                                  id,
		AzureSubscriptionID:                 azureSubscriptionID,
		AzureTenantID:                       azureTenantID,
		AzureClientID:                       azureClientID,
		AzureClientSecret:                   azureClientSecret,
		AzureLocation:                       azureLocation,
		AzureResourceGroup:                  azureResourceGroup,
		AzureControlPlaneMachineType:        azureControlPlaneMachineType,
		AzureNodeMachineType:                azureNodeMachineType,
		AzureClusterIdentitySecretName:      azureClusterIdentitySecretName,
		AzureClusterIdentitySecretNamespace: azureClusterIdentitySecretNamespace,
		ClusterIdentityName:                 clusterIdentityName,
	}

	// Add the new cluster to the slicem
	clusters = append(clusters, newcluster)
	clustersAzure = append(clustersAzure, newclusterAzure)

	// Create config file from structs
	configString := fmt.Sprint(
		"\nCLUSTER_NAME: ", newcluster.ClusterName,
		"\nCONTROL_PLANE_MACHINE_COUNT: ", newcluster.ControlPlaneMachineCount,
		"\nWORKER_MACHINE_COUNT: ", newcluster.WorkerMachineCount,
		"\nKUBERNETES_VERSION: ", newcluster.KubernetesVersion,
		"\nCLUSTER_TOPOLOGY: ", newcluster.ClusterTopology,
		"\nFLAVOR: ", newcluster.Flavor,
		"\nGITHUB_TOKEN: ", newcluster.GitHubToken,
		"\nAZURE_LOCATION: ", newclusterAzure.AzureLocation,
		"\nAZURE_RESOURCE_GROUP: ", newclusterAzure.AzureResourceGroup,
		"\nAZURE_CONTROL_PLANE_MACHINE_TYPE: ", newclusterAzure.AzureLocation,
		"\nAZURE_NODE_MACHINE_TYPE: ", newclusterAzure.AzureNodeMachineType,
		"\nAZURE_SUBSCRIPTION_ID: ", newclusterAzure.AzureSubscriptionID,
		"\nAZURE_TENANT_ID: ", newclusterAzure.AzureTenantID,
		"\nAZURE_CLIENT_SECRET: ", newclusterAzure.AzureClientSecret,
		"\nAZURE_CLIENT_ID: ", newclusterAzure.AzureClientID,
		"\nAZURE_CLUSTER_IDENTITY_SECRET_NAME: ", newclusterAzure.AzureClusterIdentitySecretName,
		"\nAZURE_CLUSTER_IDENTITY_SECRET_NAMESPACE: ", newclusterAzure.AzureClusterIdentitySecretNamespace,
		"\nCLUSTER_IDENTITY_NAME: ", newclusterAzure.ClusterIdentityName,
	)

	// Cluster-API default $HOME/.cluster-api/clusterctl.yaml
	path := fmt.Sprintf("./cluster-api/%v", id)
	createClusterctlYAML(path, configString)

	// Trim leading and trailing newlines
	configTemp := strings.TrimSuffix(configString, "\n")
	configTemp = strings.TrimPrefix(configTemp, "\n")
	// Split string into slice
	configSlice := strings.Split(configTemp, "\n")

	// GenerateHandler
	c.HTML(http.StatusOK, "generate.tmpl", gin.H{"provider": strings.ToUpper(provider), "id": id, "cfg": configSlice})
}
