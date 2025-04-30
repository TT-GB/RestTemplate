package main

type Resource struct {
	IP           string `json:"ip"`
	ResourceName string `json:"resourceName"`
	MAC          string `json:"mac"`
	OS           string `json:"os"`
	ResourceType string `json:"resourceType"`
	RiskScore    int    `json:"riskScore"`
	Tags         string `json:"tags"`
}

// Resources is the shared list of mock data
var Resources = []Resource{
	{
		IP:           "192.168.1.10",
		ResourceName: "Server-01",
		MAC:          "00:1A:2B:3C:4D:5E",
		OS:           "Ubuntu 22.04",
		ResourceType: "Server",
		RiskScore:    75,
		Tags:         "Production, Linux",
	},
	{
		IP:           "192.168.1.20",
		ResourceName: "Workstation-Alpha",
		MAC:          "00:1B:2C:3D:4E:5F",
		OS:           "Windows 10",
		ResourceType: "Workstation",
		RiskScore:    40,
		Tags:         "Employee, Windows",
	},
	{
		IP:           "10.0.0.5",
		ResourceName: "IoT-Camera",
		MAC:          "00:FF:AA:BB:CC:DD",
		OS:           "Embedded Linux",
		ResourceType: "IoT Device",
		RiskScore:    90,
		Tags:         "IoT, Surveillance",
	},
}
