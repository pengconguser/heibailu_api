package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func ReadJsonFile(jsonPath string) (map[string]string, error) {
	var json_result = map[string]string{}

	dw, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	if err = json.Unmarshal(dw, &json_result); err == nil {
		return json_result, nil
	}

	return json_result, nil
}

//转换时间格式的函数,适应前端格式的需要
func ToDateTimeString(inter ...interface{}) {

	for _, v := range inter {
		value := reflect.ValueOf(v)

		value = value.Elem()

		f := value.FieldByName("created_at")

		if f.Kind() == reflect.String {
			timeP, _ := time.Parse("2006-01-02 15:04:05", f.String())
			f.SetString(timeP.Format("2006-01-02 15:04:05"))
		}
	}
}

//返回当前项目的绝对路径

func Object_path() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		fmt.Println(err)
	}

	path = strings.Replace(path, "\\", "/", -1)

	return path
}
