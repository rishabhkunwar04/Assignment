package repo

import "targeting-engine/model"

type MemoryRepository struct{}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (r *MemoryRepository) GetActiveCampaigns() []model.Campaign {
	return []model.Campaign{
		{"spotify", "https://somelink", "Download", "ACTIVE"},
		{"duolingo", "https://somelink2", "Install", "ACTIVE"},
		{"subwaysurfer", "https://somelink3", "Play", "ACTIVE"},
	}
}

func (r *MemoryRepository) GetTargetingRules() []model.TargetingRule {
	return []model.TargetingRule{
		{"spotify", nil, nil, []string{"US", "Canada"}, nil, nil, nil},
		{"duolingo", nil, nil, nil, []string{"US"}, []string{"Android", "iOS"}, nil},
		{"subwaysurfer", []string{"com.gametion.ludokinggame"}, nil, nil, nil, []string{"Android"}, nil},
	}
}
