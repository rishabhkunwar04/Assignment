package model

type TargetingRule struct {
	CampaignID       string
	IncludeAppIDs    []string
	ExcludeAppIDs    []string
	IncludeCountries []string
	ExcludeCountries []string
	IncludeOS        []string
	ExcludeOS        []string
}
