package cipg

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

//If -UseConfigFile flag is set, then generate options from the yaml file.
//If -GenerateOptionFile flag is set, then Marshal and write options to the yaml file.
func GenerateWithYAML(opt interface{}, logger func(i ...interface{})) (exit bool) {
	return generateWithFile(opt, logger, yaml.Marshal, yaml.Unmarshal)
}

//If -UseConfigFile flag is set, then generate options from the json file.
//If -GenerateOptionFile flag is set, then Marshal and write options to the json file.
func GenerateWithJSON(opt interface{}, logger func(i ...interface{})) (exit bool) {
	return generateWithFile(opt, logger, json.Marshal, json.Unmarshal)

}

func generateWithFile(opt interface{}, logger func(i ...interface{}),
	Marshal func(in interface{}) (out []byte, err error),
	Unmarshal func(in []byte, out interface{}) (err error)) (exit bool) {

	exit = false
	file := flag.String("UseOptionFile", "", "Set this value to load option from a config file.")
	gfile := flag.String("GenerateOptionFile", "", "Set this value to generate a option file.")
	Generate(opt, logger)
	flag.Parse()
	if path := *file; path != "" {
		logger(fmt.Sprintf("-UseConfigFile flag detected, load option from %s.", path)) //文件读取模式
		data, err := ioutil.ReadFile(path)                                              //读取文件
		if err != nil {
			panic(err)
		}
		err = Unmarshal(data, opt) //反序列化
		if err != nil {
			panic(err)
		}
	}
	if path := *gfile; path != "" {
		logger(fmt.Sprintf("-GenerateConfigFile flag detected, option file will write to %s.", path))
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm) //打开文件流
		if err != nil {
			panic(err)
		}
		data, err := Marshal(opt) //序列化
		if err != nil {
			panic(err)
		}
		n, err := f.Write(data) //写入
		if err != nil {
			panic(err)
		}
		logger(fmt.Sprintf("%d bytes written to %s", n, path))
		exit = true
		return
	}
	return
}
