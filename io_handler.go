package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type IoHandler struct{}

func NewIoHandler() *IoHandler {
	return &IoHandler{}
}

func (h IoHandler) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (h IoHandler) DecodeToml(data string, v *TomlSetting) (interface{}, error) {
	return toml.Decode(data, &v)
}

func (h IoHandler) MakeDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, 0777); err != nil {
			return err
		}
	}
	return nil
}

func (h IoHandler) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (h IoHandler) Write(f *os.File, b []byte) (n int, err error) {
	return f.Write(b)
}

func (h IoHandler) NotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func (h IoHandler) GetHomeDirPath() (string, error) {
	return os.UserHomeDir()
}
