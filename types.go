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

// More information about Google Distance Matrix API is available on
// https://developers.google.com/maps/documentation/distancematrix/

package maps

// Mode is for specifying travel mode.
type Mode string

// Avoid is for specifying routes that avoid certain features.
type Avoid string

// Units specifies which units system to return human readable results in.
type Units string

// TransitMode is for specifying a transit mode for a request
type TransitMode string

// TransitRoutingPreference biases which routes are returned
type TransitRoutingPreference string

// Travel mode preferences.
const (
	TravelModeDriving   = Mode("driving")
	TravelModeWalking   = Mode("walking")
	TravelModeBicycling = Mode("bicycling")
	TravelModeTransit   = Mode("transit")
)

// Features to avoid.
const (
	AvoidTolls    = Avoid("tolls")
	AvoidHighways = Avoid("highways")
	AvoidFerries  = Avoid("ferries")
)

// Units to use on human readable distances.
const (
	UnitsMetric   = Units("metric")
	UnitsImperial = Units("imperial")
)

// Transit mode of directions or distance matrix request.
const (
	TransitModeBus    = TransitMode("bus")
	TransitModeSubway = TransitMode("subway")
	TransitModeTrain  = TransitMode("train")
	TransitModeTram   = TransitMode("tram")
	TransitModeRail   = TransitMode("rail")
)

// Transit Routing preferences for transit mode requests
const (
	TransitRoutingPreferenceLessWalking    = TransitRoutingPreference("less_walking")
	TransitRoutingPreferenceFewerTransfers = TransitRoutingPreference("fewer_transfers")
)

// Distance is the API representation for a distance between two points.
type Distance struct {
	// HumanReadable is the human friendly distance. This is rounded and in an appropriate unit for the
	// request. The units can be overriden with a request parameter.
	HumanReadable string `json:"text"`
	// Meters is the numeric distance, always in meters. This is intended to be used only in
	// algorithmic situations, e.g. sorting results by some user specified metric.
	Meters int `json:"value"`
}
