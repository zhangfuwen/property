# property
load java property file content and reflect it into a structure

# Usage
Write a property file with content:

    auto=true
    color=green
    number=1500
    Number1=1500
    FloatNumber=5.5
    
Code:

	type T1 struct {
		Auto        bool   `propName:"auto"`
		Color       string `propName:"color"`
		Number      int    `propName:"number"`
		Number1     int16
		FloatNumber float32
	}
	var t1 = T1{}

	if err := property.LoadProperties(&t1, "./test.properties"); err != nil {
		t.Errorf("failed, " + err.Error())
	}
	if !t1.Auto || t1.Color != "green" || t1.Number != 1500 || t1.Number1 != 1500 || t1.FloatNumber != 5.5 {
		t.Errorf("failed")
	}
