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
	router.GET("/clusters", clustergen.GetClusters) // list cluster global configs
	router.GET("/clusters/:id", clustergen.GetClusterByID)
	router.POST("/clusters", clustergen.PostClusters) // create new cluster config
	router.GET("/aws", clustergen.GetClustersAWS)     // list AWS cluster settings
	router.GET("/aws/:id", clustergen.GetClusterAWSByID)
	router.GET("/gcp", clustergen.GetClustersGCP) // list GCP cluster settings
	router.GET("/gcp/:id", clustergen.GetClusterGCPByID)
	router.GET("/oci", clustergen.GetClustersOCI) // list GCP cluster settings
	router.GET("/oci/:id", clustergen.GetClusterOCIByID)
	// router.GET("/azure", getclustersAzure) // list GCP cluster settings
	// router.GET("/azure/:id", getclusterAzureByID)
	router.GET("/generate/:id", clustergen.GenerateClusterByID)
	router.GET("/config/aws", clustergen.GetConfigAWS)
	router.GET("/config/gcp", clustergen.GetConfigGCP)
	router.GET("/config/oci", clustergen.GetConfigOCI)
	router.POST("/config/aws", clustergen.PostConfigAWS)
	router.POST("/config/gcp", clustergen.PostConfigGCP)
	router.POST("/config/oci", clustergen.PostConfigOCI)
	// router.GET("/package/:id", getPackageByID)
	// router.POST("/package", postPackage)
	router.LoadHTMLGlob("views/*")
	router.GET("/", clustergen.IndexHandler)

	router.Run("localhost:8080")
}
