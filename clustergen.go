package clustergen

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

// cluster represents data about a record cluster.
type cluster struct {
	ID                       string `json:"id"`
	ClusterName              string `json:"clusterName"`
	Provider                 string `json:"provider"` // AWS, GCP, Azure, vSphere...
	ProviderVersion          string `json:"providerVersion"`
	KubernetesVersion        string `json:"kubernetesVersion"`
	ControlPlaneMachineCount int    `json:"controlPlaneMachineCount"`
	WorkerMachineCount       int    `json:"workerMachineCount"`
	ClusterTopology          bool   `json:"clusterTopology"`
	Flavor                   string `json:"flavor"`
	Owner                    string `json:"owner"` // meta-data
	GitHubToken              string `json:"gitHubToken"`
}

type clusterAWS struct {
	ID                         string `json:"id"`
	AWSRegion                  string `json:"awsRegion"` // reguired
	AWSAccessKeyID             string `json:"awsAccessKeyID"`
	AWSSecretAccessKey         string `json:"awsSecretAccessKey"`
	AWSSSHKeyName              string `json:"awsSSHKeyName"`              // required
	AWSControlPlaneMachineType string `json:"awsControlPlaneMachineType"` // required
	AWSNodeMachineType         string `json:"awsNodeMachineType"`         // required
	AWSSessionToken            string `json:"awsSessionToken"`
	AWSB64EncodedCredentials   string `json:"awsB64EncodedCredentials"`
}

type clusterGCP struct {
	ID                         string `json:"id"`
	GCPRegion                  string `json:"gcpRegion"`
	GCPProject                 string `json:"gcpProject"`
	GCPControlPlaneMachineType string `json:"gcpControlPlaneMachineType"`
	GCPNodeMachineType         string `json:"gcpNodeMachineType"`
	GCPImageID                 string `json:"gcpImageID"`
	GCPNetworkName             string `json:"gcpNetworkName"`
	// GCPB64EncodedCredentials   string `json:"gcpB64EncodedCredentials"`
}

type clusterAzure struct {
	ID                                  string `json:"id"`
	AzureSubscriptionID                 string `json:"azureSubscriptionID"`
	AzureTenantID                       string `json:"azureTenantID"`
	AzureClientID                       string `json:"azureClientID"`
	AzureClientSecret                   string `json:"azureClientSecret"`
	AzureLocation                       string `json:"azureLocation"`
	AzureResourceGroup                  string `json:"azureResourceGroup"`
	AzureControlPlaneMachineType        string `json:"azureControlPlaneMachineType"`
	AzureNodeMachineType                string `json:"azureNodeMachineType"`
	AzureClusterIdentitySecretName      string `json:"azureClusterIdentitySecretName"`
	AzureClusterIdentitySecretNamespace string `json:"azureClusterIdentitySecretNamespace"`
	ClusterIdentityName                 string `json:"clusterIdentityName"`
}

type clusterOCI struct {
	ID                                 string `json:"id"`
	OCICompartmentID                   string `json:"ociCompartmentID"`
	OCIControlPlaneMachineType         string `json:"ociControlPlaneMachineType"`
	OCIControlPlaneMachineTypeOCPUs    string `json:"ociControlPlaneMachineTypeOCPUs"`
	OCINodeMachineType                 string `json:"ociNodeMachineType"`
	OCINodeMachineTypeOCPUs            string `json:"ociNodeMachineTypeOCPUs"`
	OCIImageID                         string `json:"ociImageID"`
	OCISSHKey                          string `json:"ociSSHKey"`
	OCIControlPlanePVTransitEncryption string `json:"ociControlPlanePVTransitEncryption"`
	OCINodePVTransitEncryption         string `json:"ociNodePVTransitEncryption"`
	NodeMachineCount                   string `json:"nodeMachineCount"`
}

// type clusterTanzu struct {
// 	SomeTanzuParam 	string `json:"someTanzuParam"`
// }

// type clusterAzure struct {
// 	SomeAzureParam string `json:"someAzureParam"`
// }

// getclusters responds with the list of all clusters as JSON.
func GetClusters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clusters)
}

func GetClustersAWS(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clustersAWS)
}

func GetClustersGCP(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clustersGCP)
}

func GetClustersAzure(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clustersAzure)
}

func GetClustersOCI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clustersOCI)
}

// postclusters adds an cluster from JSON received in the request body.
func PostClusters(c *gin.Context) {
	var newcluster cluster
	// Call BindJSON to bind the received JSON to
	// newcluster.
	if err := c.BindJSON(&newcluster); err != nil {
		return
	}
	// Add the new cluster to the slice.
	clusters = append(clusters, newcluster)
	c.IndentedJSON(http.StatusCreated, newcluster)
}

// getclusterByID locates the cluster whose ID value matches the id
// parameter sent by the client, then returns that cluster as a response.
func GetClusterByID(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	for _, a := range clusters {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cluster not found"})
}

// getclusterByID locates the cluster whose ID value matches the id
// parameter sent by the client, then returns that cluster as a response.
func GetClusterAWSByID(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	for _, a := range clustersAWS {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "clusterAWS not found"})
}

// getclusterByID locates the cluster whose ID value matches the id
// parameter sent by the client, then returns that cluster as a response.
func GetClusterGCPByID(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	for _, a := range clustersGCP {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "clusterGCP not found"})
}

// getclusterByID locates the cluster whose ID value matches the id
// parameter sent by the client, then returns that cluster as a response.
func GetClusterAzureByID(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	for _, a := range clustersAzure {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "clusterAzure not found"})
}

// getclusterByID locates the cluster whose ID value matches the id
// parameter sent by the client, then returns that cluster as a response.
func GetClusterOCIByID(c *gin.Context) {
	id := c.Param("id")
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	for _, a := range clustersOCI {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "clusterOCI not found"})
}

func queryParams(c *gin.Context) {
	c.Request.URL.Query()
}

func GetConfigAWS(c *gin.Context) {
	c.HTML(200, "cluster-form-aws.html", nil)
}

func GetConfigGCP(c *gin.Context) {
	c.HTML(200, "cluster-form-gcp.html", nil)
}

func GetConfigAzure(c *gin.Context) {
	c.HTML(200, "cluster-form-azure.html", nil)
}

func GetConfigOCI(c *gin.Context) {
	c.HTML(200, "cluster-form-oci.html", nil)
}

// generateclusterByID generates Cluster-API manifests for the cluster
// whose ID value matches the id parameter sent by the client,
// then returns that cluster as a response.
func GenerateClusterByID(c *gin.Context) {
	id := c.Param("id")
	found := ClusterExists(id)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cluster not found"})
	}

	// godotenv.Load("/Users/composr/keys/capi-dev-aws-env")
	// godotenv.Load("/Users/composr/work/capi-work/capi-dev-aws-env")

	cname := ""
	provider := ""
	providerVersion := ""
	// var vals cluster
	for _, a := range clusters {
		if a.ID == id {
			cname = a.ClusterName
			provider = a.Provider
			providerVersion = a.ProviderVersion
		}
	}
	if cname == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not retrieve cluster name."})
	}

	clusterctlYAML := fmt.Sprintf("./cluster-api/%v/clusterctl.yaml", id)
	// TODO:  Add other options such as flavor...
	cmd := fmt.Sprintf("clusterctl generate cluster %v --infrastructure=%v:%v --config %v", cname, provider, providerVersion, clusterctlYAML)
	log.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Printf("Failed to execute command: %s", cmd)
		log.Printf("Error: %v", err)
	}

	//
	// Return generated CAPI cluster YAML
	//
	c.String(http.StatusOK, string(out))
}

func IndexHandler(c *gin.Context) {
	// Handle /index
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func ClusterExists(id string) bool {
	// Loop through the list of clusters, looking for
	// an cluster whose ID value matches the parameter.
	found := false
	for _, a := range clusters {
		if a.ID == id {
			found = true
		}
	}
	return found
}

func GenerateID() int {
	// Generate new ID
	id := 0
	for i := 1; i < 1000; i++ {
		match := false
		for _, a := range clusters {
			tmp, _ := strconv.Atoi(a.ID)
			if tmp == i {
				match = true
			}
		}
		if match == false {
			id = i
			break
		}
	}
	if id == 0 {
		log.Println("Could not create new ID")
	}
	return id
}

func createClusterctlYAML(p string, d string) {
	// p path to file
	// d filename
	configData := []byte(d)
	file := fmt.Sprintf("%v/clusterctl.yaml", p)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777) // Create your file
	}
	// add error return if fail
	writeToFile(file, configData)
}

func writeToFile(f string, d []byte) {
	// d1 := []byte("hello\ngo\n")
	// err := os.WriteFile("./tmp/dat1", d, 0644)
	err := os.Remove(f)
	check(err)
	err2 := os.WriteFile(f, d, 0644)
	check(err2)
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func LogEnvVars() {
	cmd := exec.Command("env")
	stdout0, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(stdout0))
}
