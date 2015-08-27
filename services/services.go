package services
import (
	"gopkg.in/mgo.v2"
	"beelike/util/mongo"
	"study/log"
	"beelike/util/helper"
)

//基本服务类
type Service struct {
	MongoSession *mgo.Session
	UserID       string
}

// Prepare is called before any controller.
func (service *Service) Prepare() (err error) {
	service.MongoSession, err = mongo.CopyMonotonicSession(service.UserID)
	if err != nil {
		log.Error(err, service.UserID, "Service.Prepare")
		return nil
	}
	return err
}

// Finish is called after the controller.
func (service *Service) Finish() (err error) {
	defer helper.CatchPanic(&err, service.UserID, "Service.Finish")
	if service.MongoSession != nil {
		mongo.CloseSession(service.UserID, service.MongoSession)
		service.MongoSession = nil
	}
	return err
}

// DBAction executes the MongoDB literal function
func (service *Service) DBAction(databaseName string, collectionName string, dbCall mongo.DBCall) (err error) {
	return mongo.Execute(service.UserID, service.MongoSession, databaseName, collectionName, dbCall)
}