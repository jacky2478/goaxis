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

    ptrWaitSetCh *chan byte
}

// set for simple data type, example: int, bool, string...
func (p *dataSetV) Set(data interface{}) IValue {
    p.setV = data
    p.endWaitSet()
    return p
}

// push for array or slice, example: []int, []string...
func (p *dataSetV) Push(datas ...interface{}) IValue {
    p.pushV = append(p.pushV, datas...)
    p.endWaitSet()
    return p
}

// fill for map or struct, example: map[string]int, TData...
func (p *dataSetV) Fill(data interface{}) IValue {
    p.fillV = data
    p.endWaitSet()
    return p
}

func (p *dataSetV) ReadyWaitSet() {
    wsch := make(chan byte, 1)
    p.ptrWaitSetCh = &wsch
}

func (p *dataSetV) endWaitSet() {
    if p.ptrWaitSetCh != nil && *p.ptrWaitSetCh != nil {
        *p.ptrWaitSetCh <- byte(0)
        close(*p.ptrWaitSetCh)
        *p.ptrWaitSetCh = nil
    }
}

func (p *dataSetV) startWaitSet() {
    if p.ptrWaitSetCh != nil && *p.ptrWaitSetCh != nil {
        <- *p.ptrWaitSetCh
    }
}

// get original value by Fill
func (p *dataSetV) GetValueByFill() interface{} {
    p.startWaitSet()
    return p.fillV
}

// get original value by Push
func (p *dataSetV) GetValueByPush() []interface{} {
    p.startWaitSet()
    return p.pushV
}

// for simple data type, example: int, bool, string...
func (p *dataSetV) Get() interface{} {
    p.startWaitSet()
    if p.setV == nil {
        return nil
    }
    return p.setV
}

// for array or slice, example: []int, []string...
func (p *dataSetV) GetByIndex(index int) interface{} {
    p.startWaitSet()
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
    p.startWaitSet()
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