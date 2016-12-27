package property

import "testing"

func TestLoadProperties(t *testing.T) {
	type T1 struct {
		Auto        bool   `propName:"auto"`
		Color       string `propName:"color"`
		Number      int    `propName:"number"`
		Number1     int16
		FloatNumber float32
	}
	var t1 = T1{}

	if err := LoadProperties(&t1, "./test.properties"); err != nil {
		t.Errorf("failed, " + err.Error())
	}
	if !t1.Auto || t1.Color != "green" || t1.Number != 1500 || t1.Number1 != 1500 || t1.FloatNumber != 5.5 {
		t.Errorf("failed")
	}
	t.Log("TestLoadProperties")
}
