package AbsLog

import (
	"github.com/jimblizz/logger"
	"github.com/bsphere/le_go"
	"os"
)

func NewAbsLog (environment string, key string) AbsLog  {
	if environment != "dev" && environment != "development" {
		// Open log file
		f, err := os.OpenFile("filterclient.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("Could not open filterclient.log for writing")
		} else {
			logger.SetOutput(f)
		}
		defer f.Close()
	}

	return AbsLog{
		Log: logger.New(key),
	}
}

type AbsLog struct {
	Log 			*logger.Logger
	Le 				*le_go.Logger
	UseLogentries 	bool
}

func (self *AbsLog) SetUseLogentries(token string) {
	var err error

	self.Le, err = le_go.Connect(token)
	if err != nil {
		panic(err)
	}
	self.UseLogentries = true
	//defer self.Le.Close()

	return
}

func (self AbsLog) Info (msg string) {
	self.addLog("info", msg)
	return
}

func (self AbsLog) Error (msg string) {
	self.addLog("error", msg)
	return
}


func (self AbsLog) addLog (level string, msg string) {
	if level == "error" {
		self.Log.Error(msg)
		if self.UseLogentries {
			self.Le.Println("{\"level\": error, \"msg\": " + msg + "}")
		}

	} else {
		self.Log.Info(msg)
		if self.UseLogentries {
			self.Le.Println("{\"level\": \"info\", \"msg\": \"" + msg + "\"}")
		}
	}
}
