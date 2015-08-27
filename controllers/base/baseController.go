package base
import (
	"github.com/astaxie/beego"
	"beelike/services"
	log "github.com/goinggo/tracelog"
	"beelike/util/mongo"
	"runtime"
	"fmt"
	"github.com/astaxie/beego/validation"
	"reflect"
)

//基本控制器结构
type BaseController struct {
	beego.Controller
	services.Service
}

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the baseController method.
func (baseController *BaseController) Prepare() {
	baseController.UserID = baseController.GetString("userID")
	if baseController.UserID == "" {
		baseController.UserID = baseController.GetString(":userID")
	}

	if baseController.UserID == "" {
		baseController.UserID = "Unknow"
	}

	if err := baseController.Service.Prepare(); err != nil {
		log.Errorf(err, baseController.UserID, "BaseController.Prepare", baseController.Ctx.Request.URL.Path)
		baseController.ServeError(err)
		return
	}
	log.Trace(baseController.UserID, "BaseController.Prepare", "UserID[%s] Path[%s]", baseController.UserID, baseController.Ctx.Request.URL.Path)
}

// Finish is called once the baseController method completes.
func (baseController *BaseController) Finish() {
	defer func() {
		if baseController.MongoSession != nil {
			mongo.CloseSession(baseController.UserID, baseController.MongoSession)
			baseController.MongoSession = nil
		}
	}()
	log.Completedf(baseController.UserID, "Finish", baseController.Ctx.Request.URL.Path)
}


//** VALIDATION

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message.
func (baseController *BaseController) ParseAndValidate(params interface{}) bool {

	if err := baseController.ParseForm(params); err != nil {
		baseController.ServeError(err)
		return false
	}

	var valid validation.Validation
	ok, err := valid.Valid(params)
	if err != nil {
		baseController.ServeError(err)
		return false
	}

	if ok == false {
		// Build a map of the Error messages for each field
		messages2 := make(map[string]string)

		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			// Look for an Error tag in the field
			typeField := val.Type().Field(i)
			tag := typeField.Tag
			tagValue := tag.Get("Error")

			// Was there an Error tag
			if tagValue != "" {
				messages2[typeField.Name] = tagValue
			}
		}

		// Build the Error response
		var errors []string
		for _, err := range valid.Errors {
			// Match an Error from the validation framework Errors
			// to a field name we have a mapping for
			message, ok := messages2[err.Field]
			if ok == true {
				// Use a localized message if one exists
				errors = append(errors, message)
				continue
			}

			// No match, so use the message as is
			errors = append(errors, err.Message)
		}

		baseController.ServeValidationErrors(errors)
		return false
	}
	return true
}

// ServeValidationErrors prepares and serves a validation exception.
func (baseController *BaseController) ServeValidationErrors(Errors []string) {
	baseController.Data["json"] = struct {
		Errors []string `json:"Errors"`
	}{Errors}
	baseController.Ctx.Output.SetStatus(409)
	baseController.ServeJson()
}


//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response.
func (baseController *BaseController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		log.Warning(baseController.Service.UserID, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))
		baseController.ServeError(fmt.Errorf("%v", r))
	}
}


//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response.
func (baseController *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}
	baseController.Data["json"] = response
	baseController.ServeJson()
}



// ServeError prepares and serves an Error exception.
func (baseController *BaseController) ServeError(err error) {
	baseController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	baseController.Ctx.Output.SetStatus(500)
	baseController.ServeJson()
}

