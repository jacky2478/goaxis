package goaxis

import (
    "reflect"
)

/*
import (
    "/path/libdb"
    "/path/libhttp"
    "github.com/jackychen/goaxis"
)

var tdb, tweb goaxis.IHub
type workRsp struct{}

func (p *workRsp) Pull(mode string, caller goaxis.ICaller, ds goaxis.IDataSet) error {
    switch caller.Format(".") {
    case "libweb.Listen":
       port := libdb.GetPort()
       dv.Output().Set(port)
    }
    return nil
}

func (p *workRsp) Push(mode string, caller goaxis.ICaller, ds goaxis.IDataSet) error {
    switch caller.Format(".") {
    case "libweb.RegistVerify":
       data := dv.Input().GetMust(reflect.String) 
       libdb.IsUserExist(data)
    }
    return nil
}

func main() {
    tdb = goaxis.Create("libdb", &workRsp{})
    tweb = goaxis.Create("libhttp", &workRsp{})

    libdb.Init(tdb)
    libweb.Init(tweb)
    libweb.ServeAndListen()
}
*/
func Create(module string, rsp ICallback) IHub {
    return newHub(module, rsp)
}

// 1. create for push or pull ---> dataset := DataSet(333)
// 2. set result for pull -------> dataset.Output.Set("test") 
// 3. get result for pull -------> output := dataset.Output().Get()
//                                
func DataSet(datas ...interface{}) IDataSet {
    p := new(dataSet)
    pv := p.Input()

    // set data for slice
    if len(datas) > 1 {
        pv.Push(datas...)
        return p
    }

    // set for simple datatype or struct, map...
    switch reflect.TypeOf(datas[0]).Kind() {
    case reflect.Struct:
    case reflect.Ptr:
    case reflect.Map:
        pv.Fill(datas[0])
    default:
        pv.Set(datas[0])
    }
    return p
}
