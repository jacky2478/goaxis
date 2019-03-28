package goaxis

import (
    "reflect"
)

/*
desc: implement the interface of IDataSet
input: cache for the data from component to the scheduling layer
output: cache for the data from the scheduling layer to component

setHook: do something before set defined by caller
setHook: do something before get defined by caller
*/
type dataSet struct{
    input *dataSetV
    output *dataSetV

    setHook func(IValue) error
    getHook func(IValue) error
}

// get input value by the interface of IDataSet
func (p *dataSet) Input() IValue {
    if p.input == nil {
        p.input = new(dataSetV)
        p.input.setHook = p.setHook
    }
    return p.input
}

// get output value by the interface of IDataSet
func (p *dataSet) Output() IValue {
    if p.output == nil {
        p.output = new(dataSetV)
        p.output.setHook = p.setHook
    }
    return p.output
}

// reset hook for set action
func (p *dataSet) ResetHookForSet(hk func(IValue) error) {
    if p.input != nil {
        p.input.setHook = hk
    }
}

// reset hook for get action
func (p *dataSet) ResetHookForGet(hk func(IValue) error) {
    if p.output != nil {
        p.output.getHook = hk
    }
}

/*
desc: implement the interface of IValue
setV: for simple data type, example: int, bool, string...
fillV: for map or struct, example: map[string]int, TData...
pushV: for array or slice, example: []int, []string...

setHook: do something before set defined by caller
setHook: do something before get defined by caller
*/
type dataSetV struct {
    setV interface{}
    fillV interface{}
    pushV []interface{}

    setHook func(IValue) error
    getHook func(IValue) error
}

// do the hook of before operation for setting
func (p *dataSetV) beforeSet() error {
    if p.setHook == nil {
        return nil
    }
    return p.setHook(p)
}

// do the hook of before operation for getting
func (p *dataSetV) beforeGet() error {
    if p.getHook == nil {
        return nil
    }
    return p.getHook(p)
}

// set for simple data type, example: int, bool, string...
func (p *dataSetV) Set(data interface{}) IValue {
    if p.beforeSet() == nil {
        p.setV = data
    }
    return p
}

// push for array or slice, example: []int, []string...
func (p *dataSetV) Push(datas ...interface{}) IValue {
    if p.beforeSet() == nil {
        p.pushV = append(p.pushV, datas...)
    }
    return p
}

// fill for map or struct, example: map[string]int, TData...
func (p *dataSetV) Fill(data interface{}) IValue {
    if p.beforeSet() == nil {
        p.fillV = data
    }
    return p
}

// get original value by Fill
func (p *dataSetV) GetValueByFill() interface{} {
    if p.beforeGet() == nil {
        return p.fillV
    }
    return nil
}

// get original value by Push
func (p *dataSetV) GetValueByPush() []interface{} {
    if p.beforeGet() == nil {
        return p.pushV
    }
    return nil
}

// for simple data type, example: int, bool, string...
func (p *dataSetV) Get() interface{} {
    if p.beforeGet() == nil {
        return p.setV
    }
    return nil
}

// for array or slice, example: []int, []string...
func (p *dataSetV) GetByIndex(index int) interface{} {
    if p.beforeGet() == nil {
        if len(p.pushV) == 0 {
            return nil
        }

        for i, v := range p.pushV {
            if i != index {
                continue
            }
            return v
        }
    }
    return nil
}

// for map or struct, example: map[string]int, TData...
func (p *dataSetV) GetByName(name string) interface{} {
    if p.beforeGet() == nil {
        if p.fillV == nil {
            return nil
        }

        var fillVV reflect.Value
        if reflect.TypeOf(p.fillV).Kind() == reflect.Ptr {
            fillVV = reflect.ValueOf(p.fillV).Elem()
        } else {
            fillVV = reflect.ValueOf(p.fillV)
        }

        if !fillVV.IsValid() {
            return nil
        }

        fieldV := fillVV.FieldByName(name)
        if fieldV.IsValid() {
            return fieldV.Interface()

        }
    }
    return nil
}