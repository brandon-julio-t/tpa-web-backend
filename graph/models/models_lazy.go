package models

type DiscoveryQueue struct {
	DiscoveryQueueNewReleased []*Game
}

type Community struct {
	CommunityDiscussion      CommunityDiscussion
	CommunityDiscussions     []*CommunityDiscussion
	CommunityImageAndVideo   CommunityImageAndVideo
	CommunityImagesAndVideos []*CommunityImageAndVideo
	CommunityReview          CommunityReview
	CommunityReviews         []*CommunityReview
}
