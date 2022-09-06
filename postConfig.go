package clustergen

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// postConfig adds an cluster config from JSON received in the request body.
func PostConfigAWS(c *gin.Context) {
	id := strconv.Itoa(GenerateID())
	// Generate new ID
	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not create new ID"})
		return
	}
	provider := "aws"

	// Validate provider
	c.Bind(&clusters)
	clusterName := c.PostForm("clusterName")
	// provider := c.PostForm("provider")
	providerVersion := c.PostForm("providerVersion")
	kubernetesVersion := c.PostForm("kubernetesVersion")
	controlPlaneMachineCount := c.PostForm("controlPlaneMachineCount")
	workerMachineCount := c.PostForm("workerMachineCount")
	clusterTopology := c.PostForm("clusterTopology")
	owner := c.PostForm("owner")
	gitHubToken := c.PostForm("gitHubToken")

	c.Bind(&clustersAWS)
	awsRegion := c.PostForm("awsRegion")
	awsControlPlaneMachineType := c.PostForm("awsControlPlaneMachineType")
	awsNodeMachineType := c.PostForm("awsNodeMachineType")
	awsSSHKeyName := c.PostForm("awsSSHKeyName")
	awsSecretAccessKey := c.PostForm("awsSecretAccessKey")
	awsAccessKeyID := c.PostForm("awsAccessKeyID")
	awsB64EncodedCredentials := c.PostForm("awsB64EncodedCredentials")

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
		Owner:                    owner,
		GitHubToken:              gitHubToken,
	}

	newclusterAWS := clusterAWS{
		ID:                         id,
		AWSRegion:                  awsRegion,
		AWSAccessKeyID:             awsAccessKeyID,
		AWSSecretAccessKey:         awsSecretAccessKey,
		AWSSSHKeyName:              awsSSHKeyName,
		AWSControlPlaneMachineType: awsControlPlaneMachineType,
		AWSNodeMachineType:         awsNodeMachineType,
		AWSB64EncodedCredentials:   awsB64EncodedCredentials,
	}

	// Add the new cluster to the slicem
	clusters = append(clusters, newcluster)
	clustersAWS = append(clustersAWS, newclusterAWS)

	// Create config file from structs
	configString := fmt.Sprint(
		"\nCLUSTER_NAME: ", newcluster.ClusterName,
		"\nCONTROL_PLANE_MACHINE_COUNT: ", newcluster.ControlPlaneMachineCount,
		"\nWORKER_MACHINE_COUNT: ", newcluster.WorkerMachineCount,
		"\nKUBERNETES_VERSION: ", newcluster.KubernetesVersion,
		"\nCLUSTER_TOPOLOGY: ", newcluster.ClusterTopology,
		"\nGITHUB_TOKEN: ", newcluster.GitHubToken,
		"\nAWS_REGION: ", newclusterAWS.AWSRegion,
		"\nAWS_ACCESS_KEY_ID: ", newclusterAWS.AWSAccessKeyID,
		"\nAWS_SECRET_ACCESS_KEY: ", newclusterAWS.AWSSecretAccessKey,
		"\nAWS_SESSION_TOKEN: ", newclusterAWS.AWSSessionToken,
		"\nAWS_CONTROL_PLANE_MACHINE_TYPE: ", newclusterAWS.AWSControlPlaneMachineType,
		"\nAWS_NODE_MACHINE_TYPE: ", newclusterAWS.AWSNodeMachineType,
		"\nAWS_SSH_KEY_NAME: ", newclusterAWS.AWSSSHKeyName, "\n",
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

// postConfig adds an cluster config from JSON received in the request body.
func PostConfigOCI(c *gin.Context) {
	id := strconv.Itoa(GenerateID())
	// Generate new ID
	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not create new ID"})
		return
	}
	provider := "oci"

	c.Bind(&clusters)
	clusterName := c.PostForm("clusterName")
	// provider := c.PostForm("provider")
	providerVersion := c.PostForm("providerVersion")
	kubernetesVersion := c.PostForm("kubernetesVersion")
	controlPlaneMachineCount := c.PostForm("controlPlaneMachineCount")
	workerMachineCount := c.PostForm("workerMachineCount")
	clusterTopology := c.PostForm("clusterTopology")
	owner := c.PostForm("owner")
	gitHubToken := c.PostForm("gitHubToken")

	c.Bind(&clustersOCI)
	ociCompartmentID := c.PostForm("ociCompartmentID")
	ociImageID := c.PostForm("ociImageID")
	ociControlPlaneMachineType := c.PostForm("ociControlPlaneMachineType")
	ociControlPlaneMachineTypeOCPUs := c.PostForm("ociControlPlaneMachineTypeOCPUs")
	ociNodeMachineType := c.PostForm("ociNodeMachineType")
	ociNodeMachineTypeOCPUs := c.PostForm("ociNodeMachineTypeOCPUs")
	ociSSHKey := c.PostForm("ociSSHKey")
	ociControlPlanePVTransitEncryption := c.PostForm("ociControlPlanePVTransitEncryption")
	ociNodePVTransitEncryption := c.PostForm("ociNodePVTransitEncryption")
	nodeMachineCount := c.PostForm("nodeMachineCount")

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
		Owner:                    owner,
		GitHubToken:              gitHubToken,
	}

	newclusterOCI := clusterOCI{
		ID:                                 id,
		OCICompartmentID:                   ociCompartmentID,
		OCIImageID:                         ociImageID,
		OCIControlPlaneMachineType:         ociControlPlaneMachineType,
		OCIControlPlaneMachineTypeOCPUs:    ociControlPlaneMachineTypeOCPUs,
		OCINodeMachineType:                 ociNodeMachineType,
		OCINodeMachineTypeOCPUs:            ociNodeMachineTypeOCPUs,
		OCISSHKey:                          ociSSHKey,
		OCIControlPlanePVTransitEncryption: ociControlPlanePVTransitEncryption,
		OCINodePVTransitEncryption:         ociNodePVTransitEncryption,
		NodeMachineCount:                   nodeMachineCount,
	}

	// Add the new cluster to the slicem
	clusters = append(clusters, newcluster)
	clustersOCI = append(clustersOCI, newclusterOCI)

	// Create config file from structs
	configString := fmt.Sprint(
		"\nCLUSTER_NAME: ", newcluster.ClusterName,
		"\nCONTROL_PLANE_MACHINE_COUNT: ", newcluster.ControlPlaneMachineCount,
		"\nWORKER_MACHINE_COUNT: ", newcluster.WorkerMachineCount,
		"\nKUBERNETES_VERSION: ", newcluster.KubernetesVersion,
		"\nCLUSTER_TOPOLOGY: ", newcluster.ClusterTopology,
		"\nGITHUB_TOKEN: ", newcluster.GitHubToken,
		"\nOCI_COMPARTMENT_ID: ", newclusterOCI.OCICompartmentID,
		"\nOCI_IMAGE_ID: ", newclusterOCI.OCIImageID,
		"\nOCI_CONTROL_PLANE_MACHINE_TYPE: ", newclusterOCI.OCIControlPlaneMachineType,
		"\nOCI_CONTROL_PLANE_MACHINE_TYPE_OCPUS: ", newclusterOCI.OCIControlPlaneMachineTypeOCPUs,
		"\nOCI_NODE_MACHINE_TYPE: ", newclusterOCI.OCINodeMachineType,
		"\nOCI_NODE_MACHINE_TYPE_OCPUS: ", newclusterOCI.OCINodeMachineTypeOCPUs,
		"\nOCI_SSH_KEY: ", newclusterOCI.OCISSHKey,
		"\nOCI_CONTROL_PLANE_PV_TRANSIT_ENCRYPTION: ", newclusterOCI.OCIControlPlanePVTransitEncryption,
		"\nOCI_NODE_PV_TRANSIT_ENCRYPTION: ", newclusterOCI.OCINodePVTransitEncryption,
		"\nNODE_MACHINE_COUNT: ", newclusterOCI.NodeMachineCount,
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
