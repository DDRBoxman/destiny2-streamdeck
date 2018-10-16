package bungieclient

import (
	"fmt"
	"log"
)

type Destiny2ManifestResponse struct {
	Response Destiny2Manifest
}

type Destiny2Manifest struct {
	Version string
	MobileAssetContentPath string
	MobileGearAssetDataBases []GearAssetDataBaseDefinition
	MobileWorldContentPaths map[string]string
	MobileClanBannerDatabasePath string
	MobileGearCDN map[string]string
}

type GearAssetDataBaseDefinition struct {
	Version int32
	Path string
}

func (client *BungieClient) GetDestiny2Manifest() (*Destiny2ManifestResponse, error) {
	manifest := &Destiny2ManifestResponse{}
	_, err := client.sling.New().Get("Destiny2/Manifest/").ReceiveSuccess(manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}

type Destiny2ProfileResponse struct {
	Response Destiny2Profile
}

type Destiny2Profile struct {
	Characters DestinyCharacterComponent
	CharacterProgressions DestinyCharacterProgressionComponent
}

type DestinyCharacterComponent struct {
	Data map[string]DestinyCharacterComponentData
}

type DestinyCharacterComponentData struct {
	CharacterId string
	Light int
	EmblemPath string
	EmblemBackgroundPath string
}

type DestinyCharacterProgressionComponent struct {
	Data map[string]DestinyCharacterProgressionComponentData
}

type DestinyCharacterProgressionComponentData struct {
	Factions map[uint32]DestinyFactionProgression
}

type DestinyFactionProgression struct {
	FactionHash uint32
	Level int
	ProgressionHash int
	DailyProgress int
	DailyLimit int
	WeeklyProgress int
	WeeklyLimit int
	CurrentProgress int
	LevelCap int
	StepIndex int
	ProgressToNextLevel int
	NextLevelAt int
}

type ComponentsParams struct {
	Components []int32 `url:"components,omitempty"`
}

func (client *BungieClient) GetDestiny2Profile(membershipType int32, destinyMembershipId string, components []int32) (*Destiny2ProfileResponse, error) {
	profileResp := &Destiny2ProfileResponse{}

	params := ComponentsParams{
		Components: components,
	}

	url := fmt.Sprintf("Destiny2/%d/Profile/%s", membershipType, destinyMembershipId)

	resp, err := client.sling.New().Get(url).QueryStruct(params).Receive(profileResp, nil)
	if err != nil {
		return profileResp, err
	}

	log.Println(resp)

	return profileResp, nil
}