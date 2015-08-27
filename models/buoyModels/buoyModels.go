package buoyModels
import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type BuoyCondition struct {
	WindSpeed     float64 `bson:"wind_speed_milehour" json:"wind_speed_milehour"`
	WindDirection int     `bson:"wind_direction_degnorth" json:"wind_direction_degnorth"`
	WindGust      float64 `bson:"gust_wind_speed_milehour" json:"gust_wind_speed_milehour"`
}

type BuoyLocation struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type BuoyStation struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	StationID string        `bson:"station_id" json:"station_id"`
	Name      string        `bson:"name" json:"name"`
	LocDesc   string        `bson:"location_desc" json:"location_desc"`
	Condition BuoyCondition `bson:"condition" json:"condition"`
	Location  BuoyLocation  `bson:"location" json:"location"`
}

// DisplayWindSpeed pretty prints wind speed.
func (buoyCondition *BuoyCondition) DisplayWindSpeed() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindSpeed)
}

// DisplayWindGust pretty prints wind gust.
func (buoyCondition *BuoyCondition) DisplayWindGust() string {
	return fmt.Sprintf("%.2f", buoyCondition.WindGust)
}