package wiim

import (
	"encoding/xml"
	"log"
)

type DIDL struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:metadata-1-0/DIDL-Lite/ DIDL-Lite"`
	Class   string   `xml:"upnp:class"`
	Item    struct {
		Quality       int    `xml:"http://www.wiimu.com/song/ quality"`
		ActualQuality string `xml:"http://www.wiimu.com/song/ actualQuality"`
		Atmos         int    `xml:"http://www.wiimu.com/song/ atmos"`
		ID            int    `xml:"http://www.wiimu.com/song/ id"`
		AlbumID       int    `xml:"http://www.wiimu.com/song/ albumid"`
		SingerID      int    `xml:"http://www.wiimu.com/song/ singerid"`
		Title         string `xml:"http://purl.org/dc/elements/1.1/ title"`
		Artist        string `xml:"urn:schemas-upnp-org:metadata-1-0/upnp/ artist"`
		Album         string `xml:"urn:schemas-upnp-org:metadata-1-0/upnp/ album"`
		Resource      struct {
			ProtocolInfo string `xml:"protocolInfo,attr"`
			Duration     string `xml:"duration,attr"`
			URL          string `xml:",chardata"`
		} `xml:"res"`
		AlbumArtURI string `xml:"urn:schemas-upnp-org:metadata-1-0/upnp/ albumArtURI"`
		RateHz      string `xml:"rate_hz"`
		FormatS     string `xml:"format_s"`
		Bitrate     string `xml:"bitrate"`
	} `xml:"item"`
}

func (didl DIDL) MediaInfo() MediaInfo {
	log.Print(didl.Item.ActualQuality)
	return MediaInfo{
		SampleRate: didl.Item.RateHz,
		Bitrate:    didl.Item.Bitrate,
		BitDepth:   didl.Item.FormatS,
		Artist:     didl.Item.Artist,
		Album:      didl.Item.Album,
		Title:      didl.Item.Title,
		Quality:    didl.Item.ActualQuality,
		Artwork:    didl.Item.AlbumArtURI,
	}
}

func ParseMediaInfo(xmlData string) (DIDL, error) {
	var didl DIDL
	err := xml.Unmarshal([]byte(xmlData), &didl)

	return didl, err
}

type MediaInfo struct {
	SampleRate string `json:"sample_rate,omitempty"`
	Bitrate    string `json:"bitrate,omitempty"`
	BitDepth   string `json:"bit_depth,omitempty"`
	Artist     string `json:"artist,omitempty"`
	Title      string `json:"title,omitempty"`
	Album      string `json:"album,omitempty"`
	Quality    string `json:"quality,omitempty"`
	Artwork    string `json:"artwork,omitempty"`
}
