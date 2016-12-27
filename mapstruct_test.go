package property

import (
	"testing"
)


func TestUnmarshal(t *testing.T) {
	m := map[string]string {
		"auto":"true",
		"color":"green",
		"number":"1500",
		"Number1":"1500",
		"FloatNumber":"5.5",
	}
	type T1 struct {
		Auto bool `propName:"auto"`
		Color string `propName:"color"`
		Number int `propName:"number"`
		Number1 int16
		FloatNumber float32
	}
	var t1 = T1{}
	if err:=Unmarshal(&t1, m); err!=nil {
		t.Errorf("failed, "+err.Error())
	}
	if(!t1.Auto || t1.Color!="green" || t1.Number!=1500 || t1.Number1 != 1500|| t1.FloatNumber!=5.5) {
		t.Errorf("failed")
	}
}