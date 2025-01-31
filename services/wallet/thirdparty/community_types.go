package thirdparty

// Community-related info used by the wallet, cached in the wallet db.
type CommunityInfo struct {
	CommunityName  string `json:"community_name"`
	CommunityColor string `json:"community_color"`
	CommunityImage string `json:"community_image"`
}

type CommunityInfoProvider interface {
	GetCommunityID(tokenURI string) string
	FetchCommunityInfo(communityID string) (*CommunityInfo, error)
}
