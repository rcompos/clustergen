package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rcompos/clustergen"
)

func main() {

	// //Load the .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("error: failed to load the env file")
	// }

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	clustergen.LogEnvVars()

	router := gin.Default()
	router.GET("/", clustergen.IndexHandler)
	router.LoadHTMLGlob("views/*")

	// Get cluster config
	router.GET("/clusters", clustergen.GetClusters) // list cluster global configs
	router.GET("/clusters/:id", clustergen.GetClusterByID)
	router.POST("/clusters", clustergen.PostClusters) // create new cluster config

	// Get cluster config by provider
	router.GET("/clusters/aws", clustergen.GetClustersAWS) // list AWS cluster settings
	router.GET("/clusters/aws/:id", clustergen.GetClusterAWSByID)
	router.GET("/clusters/gcp", clustergen.GetClustersGCP) // list GCP cluster settings
	router.GET("/clusters/gcp/:id", clustergen.GetClusterGCPByID)
	router.GET("/clusters/oci", clustergen.GetClustersOCI) // list OCI cluster settings
	router.GET("/clusters/oci/:id", clustergen.GetClusterOCIByID)
	// router.GET("/azure", getclustersAzure) // list Azure cluster settings
	// router.GET("/azure/:id", getclusterAzureByID)

	// Generate Cluster-API workload cluster manifest
	router.GET("/generate/:id", clustergen.GenerateClusterByID)

	// Form to post cluster configs
	router.GET("/config/aws", clustergen.GetConfigAWS)
	router.GET("/config/gcp", clustergen.GetConfigGCP)
	router.GET("/config/oci", clustergen.GetConfigOCI)
	router.POST("/config/aws", clustergen.PostConfigAWS)
	router.POST("/config/gcp", clustergen.PostConfigGCP)
	router.POST("/config/oci", clustergen.PostConfigOCI)

	router.Run("localhost:8888")
}
