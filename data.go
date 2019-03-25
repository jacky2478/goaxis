package goaxis

import (
    "reflect"
)

type dataSet struct{
    input *dataSetV
    output *dataSetV
}

func (p *dataSet) Input() IValue {
    if p.input == nil {
        p.input = new(dataSetV)
    }
    return p.input
}

func (p *dataSet) Output() IValue {
    if p.output == nil {
        p.output = new(dataSetV)
    }
    return p.output
}

type dataSetV struct {
    setV interface{}
    fillV interface{}
    pushV []interface{}
}

// set for simple data type, example: int, bool, string...
func (p *dataSetV) Set(data interface{}) IValue {
    p.setV = data
    return p
}

// push for array or slice, example: []int, []string...
func (p *dataSetV) Push(datas ...interface{}) IValue {
    p.pushV = append(p.pushV, datas...)
    return p
}

// fill for map or struct, example: map[string]int, TData...
func (p *dataSetV) Fill(data interface{}) IValue {
    p.fillV = data
    return p
}

// get original value by Fill
func (p *dataSetV) GetValueByFill() interface{} {
    return p.fillV
}

// get original value by Push
func (p *dataSetV) GetValueByPush() []interface{} {
    return p.pushV
}

// for simple data type, example: int, bool, string...
func (p *dataSetV) Get() interface{} {
    if p.setV == nil {
        return nil
    }
    return p.setV
}

// for array or slice, example: []int, []string...
func (p *dataSetV) GetByIndex(index int) interface{} {
    if len(p.pushV) == 0 {
        return nil
    }

    for i, v := range p.pushV {
        if i != index {
            continue
        }
        return v
    }
    return nil
}

// for map or struct, example: map[string]int, TData...
func (p *dataSetV) GetByName(name string) interface{} {
    if p.fillV == nil {
        return nil
    }

    fillVV := reflect.ValueOf(p.fillV).Elem()
    if !fillVV.IsValid() {
        return nil
    }

    fieldV := fillVV.FieldByName(name)
    if fieldV.IsValid() {
        return fieldV.Interface()

    }
    return nil
}