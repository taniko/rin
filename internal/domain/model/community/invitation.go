package community

type InvitationID string

type Invitation struct {
	ID          InvitationID
	CommunityID ID
}
