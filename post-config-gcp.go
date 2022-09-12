package clustergen

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// postConfig adds an cluster config from JSON received in the request body.
func PostConfigGCP(c *gin.Context) {
	id := strconv.Itoa(GenerateID())
	// Generate new ID
	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not create new ID"})
		return
	}
	provider := "gcp"

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

	c.Bind(&clustersGCP)
	gcpRegion := c.PostForm("gcpRegion")
	gcpProject := c.PostForm("gcpProject")
	gcpImageID := c.PostForm("gcpImageID")
	gcpControlPlaneMachineType := c.PostForm("gcpControlPlaneMachineType")
	gcpNodeMachineType := c.PostForm("gcpNodeMachineType")
	gcpNetworkName := c.PostForm("gcpNetworkName")

	// c.JSON(http.StatusOK, gin.H{
	controlPlaneMachineCountString, _ := strconv.Atoi(controlPlaneMachineCount)
	workerMachineCountString, _ := strconv.Atoi(workerMachineCount)
	clusterTopologyBool, _ := strconv.ParseBool(clusterTopology)
	// idString := strconv.Itoa(id)

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

	newclusterGCP := clusterGCP{
		ID:                         id,
		GCPRegion:                  gcpRegion,
		GCPProject:                 gcpProject,
		GCPControlPlaneMachineType: gcpControlPlaneMachineType,
		GCPNodeMachineType:         gcpNodeMachineType,
		GCPImageID:                 gcpImageID,
		GCPNetworkName:             gcpNetworkName,
		// GCPB64EncodedCredentials:   gcpB64EncodedCredentials,
	}

	// Add the new cluster to the slice
	clusters = append(clusters, newcluster)
	clustersGCP = append(clustersGCP, newclusterGCP)

	// Create config file from structs
	configString := fmt.Sprint(
		"\nCLUSTER_NAME: ", newcluster.ClusterName,
		"\nCONTROL_PLANE_MACHINE_COUNT: ", newcluster.ControlPlaneMachineCount,
		"\nWORKER_MACHINE_COUNT: ", newcluster.WorkerMachineCount,
		"\nKUBERNETES_VERSION: ", newcluster.KubernetesVersion,
		"\nCLUSTER_TOPOLOGY: ", newcluster.ClusterTopology,
		"\nFLAVOR: ", newcluster.Flavor,
		"\nGITHUB_TOKEN: ", newcluster.GitHubToken,
		"\nGCP_REGION: ", newclusterGCP.GCPRegion,
		"\nIMAGE_ID: ", newclusterGCP.GCPImageID,
		"\nGCP_PROJECT: ", newclusterGCP.GCPProject,
		"\nGCP_CONTROL_PLANE_MACHINE_TYPE: ", newclusterGCP.GCPControlPlaneMachineType,
		"\nGCP_NODE_MACHINE_TYPE: ", newclusterGCP.GCPNodeMachineType,
		"\nGCP_NETWORK_NAME: ", newclusterGCP.GCPNetworkName,
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
