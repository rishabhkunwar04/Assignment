package repo

import "targeting-engine/model"

type CampaignRepository interface {
	GetActiveCampaigns() []model.Campaign
	GetTargetingRules() []model.TargetingRule
}
