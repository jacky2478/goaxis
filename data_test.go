package goaxis

import (
    "fmt"
    "testing"
)

func TestInput(t *testing.T) {
    ds := new(dataSet)
    ads1 := ds.Input()
    ads2 := ds.Input()

    if fmt.Sprintf("%v", ads1) != fmt.Sprintf("%v", ads2) {
        t.Errorf("TestInput failed, expect: ads1 == ads2, but got &ads1: %v, &ads2: %v", &ads1, &ads2)
    }
}

func TestOutput(t *testing.T) {
    ds := new(dataSet)
    ads1 := ds.Output()
    ads2 := ds.Output()

    if fmt.Sprintf("%v", ads1) != fmt.Sprintf("%v", ads2) {
        t.Errorf("TestOutput failed, expect: ads1 == ads2, but got &ads1: %v, &ads2: %v", &ads1, &ads2)
    }
}

func TestSetValue(t *testing.T) {
    data := "TestSetValue"
    expect := "TestSetValue"

    ds := new(dataSetV).Set(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.setV) != expect {
        t.Errorf("TestSetValue failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.Set(data).Get()))
    }
}

func TestFillValue(t *testing.T) {
    data := struct {
        Test string
    }{
        Test: "TestFillValue",
    }
    expect := "{TestFillValue}"

    ds := new(dataSetV).Fill(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.fillV) != expect {
        t.Errorf("TestFillValue failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.Fill(data).GetValueByFill()))
    }
}

func TestPushValue(t *testing.T) {
    data := "TestPushValue"
    expect := "[TestPushValue]"

    ds := new(dataSetV).Push(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.pushV) != expect {
        t.Errorf("TestPushValue failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.pushV))
    }
}

func TestGet(t *testing.T) {
    data := "TestGet"
    expect := "TestGet"

    ds := new(dataSetV)
    if fmt.Sprintf("%v", ds.Set(data).Get()) != expect {
        t.Errorf("TestGet failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.Set(data).Get()))
    }
}

func TestGetValyeByFill(t *testing.T) {
    data := "TestGetValyeByFill"
    expect := "TestGetValyeByFill"

    ds := new(dataSetV).Fill(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.GetValueByFill()) != expect {
        t.Errorf("TestGetValyeByFill failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.Set(data).Get()))
    }
}

func TestGetValyeByPush(t *testing.T) {
    data := "TestGetValyeByPush"
    expect := "[TestGetValyeByPush]"

    ds := new(dataSetV).Push(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.pushV) != expect {
        t.Errorf("TestGetValyeByPush failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.pushV))
    }
}

func TestGetByIndex(t *testing.T) {
    data := "TestGetByIndex"
    expect := "TestGetByIndex"

    ds := new(dataSetV).Push(data).(*dataSetV)
    if fmt.Sprintf("%v", ds.GetByIndex(0)) != expect {
        t.Errorf("TestGetByIndex failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds.GetByIndex(0)))
    }
}

func TestGetByName(t *testing.T) {
    data := struct {
        Test string
    }{
        Test: "TestGetByName",
    }
    expect := "TestGetByName"

    ds1 := new(dataSetV).Fill(&data).(*dataSetV)
    if fmt.Sprintf("%v", ds1.GetByName("Test")) != expect {
        t.Errorf("TestFill failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds1.GetByName("Test")))
    }

    ds2 := new(dataSetV).Fill(data).(*dataSetV)
    if fmt.Sprintf("%v", ds2.GetByName("Test")) != expect {
        t.Errorf("TestFill failed, expect: %v, but got: %v", expect, fmt.Sprintf("%v", ds2.GetByName("Test")))
    }
}