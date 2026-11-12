package aur

type Package struct {
	Name        string `json:"Name"`
	PackageBase string `json:"PackageBase"`
	Version     string `json:"Version"`
	Description string `json:"Description"`
	Maintainer  string `json:"Maintainer"`
}

type Response struct {
	Results []Package `json:"results"`
}
