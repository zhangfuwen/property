package property

import (
	"github.com/zhangfuwen/props"
	"os"
)

func LoadProperties(structPointer interface{}, filename string) error {
	prop := props.NewProperties()
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := prop.Load(f); err != nil {
		return err
	}
	return Unmarshal(structPointer, prop.GetMap())
}
