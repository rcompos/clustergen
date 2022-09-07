package clustergen

// clusters to seed record cluster data.
// Instantiate configs for DEVELOPMENT

var clusters = []cluster{
	{
		ID:                       "1",
		ClusterName:              "DE-Prod-Alpha",
		ControlPlaneMachineCount: 6,
		WorkerMachineCount:       20,
		KubernetesVersion:        "1.24.3",
		Provider:                 "aws",
		ProviderVersion:          "v1.5.0",
		Owner:                    "Data-Eng",
		ClusterTopology:          true,
	},
	{
		ID:                       "2",
		ClusterName:              "CS-Dev-Beta",
		ControlPlaneMachineCount: 3,
		WorkerMachineCount:       8,
		KubernetesVersion:        "1.24.3",
		Provider:                 "aws",
		ProviderVersion:          "v1.5.0",
		Owner:                    "Cyber-Sec",
		ClusterTopology:          true,
	},
	{
		ID:                       "3",
		ClusterName:              "AD-Dev-Alpha",
		ControlPlaneMachineCount: 9,
		WorkerMachineCount:       100,
		KubernetesVersion:        "1.23.6",
		Provider:                 "aws",
		ProviderVersion:          "v1.5.0",
		Owner:                    "AD",
		ClusterTopology:          true,
	},
	{
		ID:                       "4",
		ClusterName:              "gates-dev-0",
		ControlPlaneMachineCount: 3,
		WorkerMachineCount:       5,
		KubernetesVersion:        "1.25.1",
		Provider:                 "azure",
		ProviderVersion:          "v1.4.2",
		Owner:                    "BG",
		ClusterTopology:          true,
	},
	{
		ID:                       "20",
		ClusterName:              "page-brin-prod-1",
		ControlPlaneMachineCount: 3,
		WorkerMachineCount:       12,
		KubernetesVersion:        "1.24.1",
		Provider:                 "GCP",
		ProviderVersion:          "v1.1.1",
		Owner:                    "RC",
		ClusterTopology:          true,
	},
	{
		ID:                       "30",
		ClusterName:              "oci-dev-0",
		ControlPlaneMachineCount: 3,
		WorkerMachineCount:       12,
		KubernetesVersion:        "1.24.1",
		Provider:                 "OCI",
		ProviderVersion:          "v0.5",
		Owner:                    "DE",
		ClusterTopology:          true,
	},
}

var clustersAWS = []clusterAWS{
	{
		ID:                         "1",
		AWSRegion:                  "us-west-2",
		AWSControlPlaneMachineType: "t3.large",
		AWSNodeMachineType:         "t3.small",
	},
	{
		ID:                         "2",
		AWSRegion:                  "us-west-2",
		AWSControlPlaneMachineType: "m5a.large",
		AWSNodeMachineType:         "m5a.medium",
	},
	{
		ID:                         "3",
		AWSRegion:                  "us-west-2",
		AWSControlPlaneMachineType: "t3.large",
		AWSNodeMachineType:         "t3.medium",
	},
}

var clustersGCP = []clusterGCP{
	{
		ID:                         "20",
		GCPRegion:                  "us-west-1",
		GCPProject:                 "mygcpproject",
		GCPControlPlaneMachineType: "n1-standard-2",
		GCPNodeMachineType:         "n1-standard-2",
		GCPImageID:                 "",
		GCPNetworkName:             "n1-standard-2",
		// GCPB64EncodedCredentials:   "NmdwaTZobjUwNnB0OGGVqdXE4M2RpMzQxaHVyLxmFwcHM",
	},
}

var clustersAzure = []clusterAzure{
	{
		ID:                                  "4",
		AzureSubscriptionID:                 "1234141",
		AzureTenantID:                       "default",
		AzureClientID:                       "default",
		AzureClientSecret:                   "default",
		AzureLocation:                       "useast",
		AzureResourceGroup:                  "default",
		AzureControlPlaneMachineType:        "Standard_D2s_v3",
		AzureNodeMachineType:                "Standard_D2s_v3",
		AzureClusterIdentitySecretName:      "supersecret",
		AzureClusterIdentitySecretNamespace: "secretns",
		ClusterIdentityName:                 "somename",
	},
}

var clustersOCI = []clusterOCI{
	{
		ID:                                 "30",
		OCICompartmentID:                   "green",
		OCIControlPlaneMachineType:         "VM.Standard.E4.Flex",
		OCIImageID:                         "",
		OCIControlPlaneMachineTypeOCPUs:    "",
		OCINodeMachineType:                 "VM.Standard.E4.Flex",
		OCINodeMachineTypeOCPUs:            "",
		OCISSHKey:                          "",
		OCIControlPlanePVTransitEncryption: "",
		OCINodePVTransitEncryption:         "",
	},
}
