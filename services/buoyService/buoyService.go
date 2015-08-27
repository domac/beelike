package buoyService
import (
	"beelike/models"
	"beelike/models/buoyModels"
	"beelike/services"
	log "github.com/goinggo/tracelog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"beelike/util/mongo"
)

//** TYPES

type (
// buoyConfiguration contains settings for running the buoy service.
	buoyConfiguration struct {
		Database string
	}
)

//** PACKAGE VARIABLES

// Config provides buoy configuration.
var Config buoyConfiguration

//** INIT

func init() {
	// Pull in the configuration.
	Config = buoyConfiguration{
		Database:models.DBConfig.Mongodb.Db,
	}
}

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(service *services.Service, stationID string) (*buoyModels.BuoyStation, error) {
	log.Startedf(service.UserID, "FindStation", "stationID[%s]", stationID)
	var buoyStation buoyModels.BuoyStation
	//DBCall 方法
	f := func(collection *mgo.Collection) error {
		queryMap := bson.M{"station_id": stationID}

		log.Trace(service.UserID, "FindStation", "MGO : db.buoy_stations.find(%s).limit(1)", mongo.ToString(queryMap))
		return collection.Find(queryMap).One(&buoyStation)
	}

	if err := service.DBAction(Config.Database, "buoy_stations", f); err != nil {
		if err != mgo.ErrNotFound {
			log.CompletedError(err, service.UserID, "FindStation")
			return nil, err
		}
	}
	log.Completedf(service.UserID, "FindStation", "buoyStation%+v", &buoyStation)
	return &buoyStation, nil
}

// FindRegion retrieves the stations for the specified region
func FindRegion(service *services.Service, region string) ([]buoyModels.BuoyStation, error) {
	log.Startedf(service.UserID, "FindRegion", "region[%s]", region)

	var buoyStations []buoyModels.BuoyStation
	f := func(collection *mgo.Collection) error {
		queryMap := bson.M{"region": region}

		log.Trace(service.UserID, "FindRegion", "Query : db.buoy_stations.find(%s)", mongo.ToString(queryMap))
		return collection.Find(queryMap).All(&buoyStations)
	}

	if err := service.DBAction(Config.Database, "buoy_stations", f); err != nil {
		log.CompletedError(err, service.UserID, "FindRegion")
		return nil, err
	}

	log.Completedf(service.UserID, "FindRegion", "buoyStations%+v", buoyStations)
	return buoyStations, nil
}