package userModel

type ConfigurationHeader struct {
	FullName   []string `json:"FullName"`
	FirstName  []string `json:"FirstName"`
	MiddleName []string `json:"MiddleName"`
	LastName   []string `json:"LastName"`
	Email      []string `json:"Email"`
	Salary     []string `json:"Salary"`
	Identifier []string `json:"Identifier"`
	Phone      []string `json:"Phone"`
	Mobile     []string `json:"Mobile"`
}

type ConfigurationHeaderExport struct {
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Salary     string `json:"Salary"`
	Identifier string `json:"Identifier"`
	Phone      string `json:"Phone"`
	Mobile     string `json:"Mobile"`
}
