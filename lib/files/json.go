package files

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func ReadJSON(fileName string, configPath []string, entity interface{}) error {
	var jsonFile *os.File
	found := false
	var err error

	for _, v := range configPath {
		jsonFile, err = os.Open(fmt.Sprintf("%s/%s", v, fileName))
		defer jsonFile.Close()
		if err == nil {
			found = true
			bytes, _ := ioutil.ReadAll(jsonFile)
			err := json.Unmarshal(bytes, &entity)
			if err != nil {
				log.Println("[readJSON] failed Unmarshal, err: ", err.Error())
			}
			break
		}
	}

	if !found {
		return errors.New("file not found")
	}

	return nil
}
