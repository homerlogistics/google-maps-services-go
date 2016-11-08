package maps

import (
	"fmt"
	"googlemaps.github.io/maps/internal"
	"net/url"
)

var staticMapAPI = &apiConfig{
	host:            "https://maps.googleapis.com",
	path:            "/maps/api/staticmap",
	acceptsClientID: true,
}

func (c *Client) StaticMapSign(r *StaticMapRequest) (string, error) {

	config := staticMapAPI

	signed, err := internal.SignURL(config.path, c.clientID, c.signature, r.params())

	if err != nil {
		return "", err
	}

	return config.host + config.path + "?" + signed, nil
}

func (r *StaticMapRequest) params() url.Values {
	q := make(url.Values)

	if r.Center != "" {
		q.Set("center", r.Center)
	}
	if r.Zoom != 0 {
		q.Set("zoom", fmt.Sprint(r.Zoom))
	}
	if r.Size != nil {
		q.Set("size", r.Size)
	}
	if r.Scale != 0 {
		q.Set("scale", fmt.Sprint(r.Scale))
	}
	if r.Format != "" {
		q.Set("format", r.Format)
	}
	if r.MapType != "" {
		q.Set("maptype", r.MapType)
	}
	if r.Language != "" {
		q.Set("language", r.Language)
	}
	if r.Region != "" {
		q.Set("region", r.Region)
	}
	if r.Markers != nil {

		s := ""

		ms := r.Markers.MarkerStyles

		if ms != nil {
			if ms.Color != "" {
				s += "color:" + ms.Color
			}
			if ms.Size != "" {
				s += "color:" + ms.Size
			}
		}

		q.Set(key, value)

	}

	return q
}

type StaticMapFormat string

const (
	StaticMapFormatPNG          = StaticMapFormat("PNG")
	StaticMapFormatPNG8         = StaticMapFormat("PNG8")
	StaticMapFormatPNG32        = StaticMapFormat("PNG32")
	StaticMapFormatGIF          = StaticMapFormat("GIF")
	StaticMapFormatJPG          = StaticMapFormat("JPG")
	StaticMapFormatJPG_BASELINE = StaticMapFormat("JPG-BASELINE")
)

type StaticMapMapType string

const (
	StaticMapMapTypeRoadmap   = StaticMapMapType("ROADMAP")
	StaticMapMapTypeSatellite = StaticMapMapType("SATELLITE")
	StaticMapMapTypeTerrain   = StaticMapMapType("TERRAIN")
	StaticMapMapTypeHybrid    = StaticMapMapType("HYBRID")
)

// StaticMapRequest is the request structure for StaticMap API
type StaticMapRequest struct {
	// StaticMap fields

	// Location Parameters

	// Defines the center of the map, equidistant from all edges of the map. This parameter takes a location as either a comma-separated {latitude,longitude} pair (e.g. "40.714728,-73.998672") or a string address (e.g. "city hall, new york, ny") identifying a unique location on the face of the earth.
	Center string
	// Defines the zoom level of the map, which determines the magnification level of the map.
	Zoom int

	// Map Parameters

	// Defines the rectangular dimensions of the map image
	Size string //is a string (600x600)
	//  Affects the number of pixels that are returned
	Scale int
	// Defines the format of the resulting image. By default, the Google Static Maps API creates PNG images. There are several possible formats including GIF, JPEG and PNG types.
	Format string
	// Defines the type of map to construct.
	MapType string
	// Defines the language to use for display of labels on map tiles.
	Language string
	// Defines the appropriate borders to display, based on geo-political sensitivities.
	Region string

	//Below here is definitely going to take a little more planning

	//Feature parameters

	// Define one or more markers to attach to the image at specified locations. This parameter takes a single marker definition with parameters separated by the pipe character (|).
	Markers *Markers
	// Defines a single path of two or more connected points to overlay on the image at specified locations. This parameter takes a string of point definitions separated by the pipe character (|).
	Path *Path //(possible array)
	// Specifies one or more locations that should remain visible on the map, though no markers or other indicators will be displayed.
	Visible string //(possible array) ??????????
	// Defines a custom style to alter the presentation of a specific feature (roads, parks, and other features) of the map.
	Style string //(possible array)

}

//We will make all the models down here for now

//Markers
type Markers struct {
	MarkerStyles    *MarkerStyles
	MarkerLocations []*LatLng
}

type MarkerStyles struct {
	Size  string //can be tiny, mid, small
	Color string //can be hex or text ie blue
	Label string
}

//Path
type Path struct {
	PathStyles *PathStyles
	PathPoints []*LatLng
}

type PathStyles struct {
	Weight    int
	Color     string
	FillColor string
	Geodesic  bool
}
