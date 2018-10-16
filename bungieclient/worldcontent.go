package bungieclient

import (
	"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DownloadLatestWorldContent(manifest Destiny2Manifest) error {
	return DownloadContentDB(manifest.MobileWorldContentPaths["en"], "worldContent.sqlite")
}

type WorldContent struct {
	db *sql.DB
}

func OpenWorldContent() (*WorldContent, error) {
	wc := &WorldContent{}

	db, err := sql.Open("sqlite3", "worldContent.sqlite")
	if err != nil {
		return nil, err
	}

	wc.db = db

	return wc, nil
}

type Destiny2DisplayProperties struct {
	Description string
	Name string
	Icon string
	HasIcon bool
}

type Destiny2Faction struct {
	DisplayProperties Destiny2DisplayProperties
	ProgressionHash int
	Hash int
	Index int
	Redacted bool
}

func (wc *WorldContent) GetFactions() ([]Destiny2Faction, error) {
	factions := []Destiny2Faction{}

	rows, err := wc.db.Query("SELECT json FROM DestinyFactionDefinition")
	if err != nil {
		return factions, err
	}
	defer rows.Close()

	for rows.Next() {
		var jsonData string

		err = rows.Scan(&jsonData)
		if err != nil {
			return factions, err
		}

		faction := Destiny2Faction{}

		err = json.Unmarshal([]byte(jsonData), &faction)
		if err != nil {
			return factions,err
		}

		factions = append(factions, faction)
	}

	return factions, nil
}