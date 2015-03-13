// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maps // import "google.golang.org/maps"

import (
	"math"
)

const (
	Epsilon = 0.0001
)

// LatLng represents a location on the Earth.
type LatLng struct {
	Lat float64 `json:"lat"` // latitude
	Lng float64 `json:"lng"` // longitude
}

// AlmostEqual returns whether this LatLng is almost equal (below Epsilon) to
// the other LatLng.
func (l *LatLng) AlmostEqual(other *LatLng) bool {
	return math.Abs(l.Lat-other.Lat) < Epsilon && math.Abs(l.Lng-other.Lng) < Epsilon
}

// Bounds represents a bounded square area on the Earth.
type Bounds struct {
	NorthEast LatLng `json:"northeast"` // ne corner
	SouthWest LatLng `json:"southwest"` // sw corner
}
