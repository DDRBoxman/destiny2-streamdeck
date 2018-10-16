package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"sync"
	"time"
	"context"
	_ "image/png"

	"net/http"

	"golang.org/x/oauth2"
	"github.com/DDRBoxman/streamdeck-go"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"

	"github.com/DDRBoxman/destiny2-streamdeck/bungieclient"
)

var ICON_SIZE = 72

func main() {
	draw2d.SetFontFolder("./fonts")

	/*bungieClient := bungieclient.NewBungieClient(http.DefaultClient)

	result, err := bungieClient.GetDestiny2Manifest()
	if err != nil {
		log.Fatal(err)
	}

	err = bungieclient.DownloadLatestWorldContent(result.Response)
	if err != nil {
		log.Fatal(err)
	}*/

	decks := streamdeck.FindDecks()

	err := decks[0].Open()
	if err != nil {
		log.Fatal(err)
	}

	decks[0].Reset()

	decks[0].SetBrightness(99)


	wc, err := bungieclient.OpenWorldContent()
	if err != nil {
		log.Fatal(err)
	}

	factions, err := wc.GetFactions()
	if err != nil {
		log.Fatal(err)
	}

	key := 0

	for _, faction := range factions {
		if faction.Redacted || faction.ProgressionHash == 0 || !faction.DisplayProperties.HasIcon {
			continue
		}

		resp, err := http.Get("https://bungie.net" + faction.DisplayProperties.Icon)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		log.Println(resp.StatusCode)
		
		factionImage, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(faction.DisplayProperties.Name)
		log.Println(faction.DisplayProperties.Description)

		drawFactionToKey(decks[0], factionImage, key)
		key += 1
	}

}

func drawFactionToKey(deck streamdeck.StreamDeck, icon image.Image, key int) {
	dest := image.NewRGBA(image.Rect(0, 0, ICON_SIZE, ICON_SIZE))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetLineWidth(2)

	progress := 1500.0 / 2000.0
	distance := 204.0

	dashLength := progress * distance

	gc.SetLineDash([]float64{dashLength, 300}, 0)

	gc.MoveTo(72/2, 0)
	gc.LineTo(72,72/2)
	gc.LineTo(72/2,72)
	gc.LineTo(0,72/2)
	gc.Stroke()

	gc.Save()

	scaleX := float64(ICON_SIZE) / float64(icon.Bounds().Dx())
	scaleY := float64(ICON_SIZE) / float64(icon.Bounds().Dy())

	gc.Scale(scaleX, scaleY)
	gc.DrawImage(icon)

	gc.Restore()

	gc.SetFontSize(12)
	gc.SetFontData(draw2d.FontData{
		Name: "Roboto",
	})

	gc.FillStringAt("20", 55, 72)

	deck.WriteImageToKey(dest, key)
}

func client() {
	ctx := context.Background()
	conf := &oauth2.Config{
	    ClientID:     "21384",
	    ClientSecret: "lolol",
	    Endpoint: oauth2.Endpoint{
	        AuthURL:  "https://www.bungie.net/en/OAuth/Authorize",
	        TokenURL: "https://www.bungie.net/Platform/App/OAuth/token/",
	    },
	}

	token := &oauth2.Token{
		AccessToken: "lolol",
		RefreshToken: "lolol",
	}

	client := conf.Client(ctx, token)

	bungieClient := bungieclient.NewBungieClient(client)

	result, err := bungieClient.GetDestiny2Manifest()
	if err != nil {
		log.Fatal(err)
	}

	err = bungieclient.DownloadLatestWorldContent(result.Response)
	if err != nil {
		log.Fatal(err)
	}

	memberships, err := bungieClient.GetMembershipsForCurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(memberships)

	for _, membership := range memberships.Response.DestinyMemberships {
		if membership.MembershipType == 2 { // PSN
			profile, err := bungieClient.GetDestiny2Profile(membership.MembershipType, membership.MembershipId, []int32{200, 202})
			if err != nil {
				log.Fatal(err)
			}

			log.Println(profile)
		}
	}
}

func doOauth() {
	/*ctx := context.Background()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	log.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
	    log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
	    log.Fatal(err)
	}

	log.Println(tok)*/
}

func doDeck() {
	decks := streamdeck.FindDecks()

	err := decks[0].Open()
	if err != nil {
		log.Fatal(err)
	}

	decks[0].Reset()

	decks[0].SetBrightness(99)

	var wg sync.WaitGroup
	wg.Add(1)

	for {
		time.Sleep(2000)

		drawTempToKey(decks[0], "CPU", 2, 4)
	}
	
	wg.Wait()
}

func drawTempToKey(deck streamdeck.StreamDeck, label string, value float32, key int) {
	dest := image.NewRGBA(image.Rect(0, 0, ICON_SIZE, ICON_SIZE))
	gc := draw2dimg.NewGraphicContext(dest)

	gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})

	gc.SetFontSize(28)
	gc.SetFontData(draw2d.FontData{
		Name: "Roboto",
	})

	gc.FillStringAt(fmt.Sprintf("%.0f°", value), 10, 32+12)

	gc.SetFontSize(14)
	gc.FillStringAt(label, 10, 72-8)

	deck.WriteImageToKey(dest, key)
}