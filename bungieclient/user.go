package bungieclient

type UserMembershipDatatResponse struct {
	Response UserMembershipData
}

type UserMembershipData struct {
	DestinyMemberships []UserInfoCard
}

type UserInfoCard struct {
	SupplementalDisplayName string
	IconPath string
	MembershipType int32
	MembershipId string
	DisplayName string
}

// Membership types
/*
None: 0
TigerXbox: 1
TigerPsn: 2
TigerBlizzard: 4
TigerDemon: 10
BungieNext: 254
All: -1
*/

func (client *BungieClient) GetMembershipsForCurrentUser() (*UserMembershipDatatResponse, error) {
	resp := &UserMembershipDatatResponse{}
	_, err := client.sling.New().Get("User/GetMembershipsForCurrentUser/").ReceiveSuccess(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}