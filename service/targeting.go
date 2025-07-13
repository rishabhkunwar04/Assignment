package service

import (
	"strings"
	"targeting-engine/model"
	"targeting-engine/repo"
)

type TargetingService struct {
	repository repo.CampaignRepository
}

func NewTargetingService(repository repo.CampaignRepository) *TargetingService {
	return &TargetingService{repository: repository}
}

func (s *TargetingService) GetEligibleCampaigns(app, country, os string) []model.Campaign {
	campaigns := s.repository.GetActiveCampaigns()
	rules := s.repository.GetTargetingRules()
	var matched []model.Campaign

	for _, c := range campaigns {
		for _, r := range rules {
			if r.CampaignID == c.ID && matchRule(r, app, country, os) {
				matched = append(matched, c)
				break
			}
		}
	}
	return matched
}

func matchRule(rule model.TargetingRule, app, country, os string) bool {
	match := func(include, exclude []string, val string) bool {
		if len(include) > 0 && !contains(include, val) {
			return false
		}
		if len(exclude) > 0 && contains(exclude, val) {
			return false
		}
		return true
	}
	return match(rule.IncludeAppIDs, rule.ExcludeAppIDs, app) &&
		match(rule.IncludeCountries, rule.ExcludeCountries, country) &&
		match(rule.IncludeOS, rule.ExcludeOS, os)
}

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, val) {
			return true
		}
	}
	return false
}
