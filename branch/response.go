package branch

type Formatter struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func FormatBranch(branch Branch) Formatter {
	formatter := Formatter{
		ID:      branch.ID,
		Code:    branch.Code,
		Name:    branch.Name,
		Address: branch.Address,
	}

	return formatter
}
