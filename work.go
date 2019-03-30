package goaxis

import (
    "reflect"
)

/*

// at the scheduling layer, implement goaxis.ICallback
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
       data := ds.Input().Get().(string)
       libdb.IsUserExist(data)
    }
    return nil
}

// Initialize the IHUB interface for each component
func main() {
    tdb = goaxis.Create("libdb", &workRsp{})
    tweb = goaxis.Create("libhttp", &workRsp{})

    libdb.Init(tdb)
    libweb.Init(tweb)
    libweb.ServeAndListen()
}

// Synchronization request using the IHub interface
func ServeAndListen()
{
    pullData := goaxis.DataSet()
    if err := tdb.SyncPull(tdb.Caller("ServeAndListen", "GetPort"), pullData); err == nil {
        port := pullData.Output().Get()
    }
}

// Asynchronous request using the IHub interface
func HandleLogin()
{
    // 1. verify user login

    // 2. get offline messages to push
    pullData := goaxis.DataSet()
    if err := tdb.ASyncPull(tdb.Caller("HandleLogin", "GetOfflineMessages"), pullData); err == nil {
        messages := pullData.Output().GetValueByPush()
    }
}

// Synchronization broadcast using the IHub interface
func InitDB()
{
    pushData := goaxis.DataSet("MySql")
    tdb.SyncPush(tdb.Caller("InitDB"), pushData)
    mysqlAddr := fmt.Sprintf("%v", pushData.Output().Get())
}

// Asynchronous broadcast using the IHub interface
func WaitReceive()
{
    tweb.ASyncPush(tdb.Caller("WaitReceive"), goaxis.DataSet("data from net"))
}
*/
func Create(module string, notify func(mode string, caller ICaller, ds IDataSet) error) IHub {
    return newHub(module, notify)
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

// create for hook ---> dataset := DataSetByHook(setHook, getHook)
func DataSetByHook(setHook, getHook func(IValue) error) IDataSet {
    p := new(dataSet)
    p.setHook = setHook
    p.getHook = getHook
    return p
}
