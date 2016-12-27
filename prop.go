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
	prop.Load(f)
	return Unmarshal(structPointer, prop.GetMap())
}
