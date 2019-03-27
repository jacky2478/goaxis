# Goaxis

Goaxis is a component-oriented transformation hub.

### Describe

Goaxis follows the idea of high cohesion and low coupling, which can ensure that components are completely independent, the interaction between components is controlled by the scheduling layer, and the scheduling layer is composed of actual concrete business logic. Goaxis only provides basic data acquisition and native data acquisition. Supports synchronous asynchronous mode data request and broadcast, concurrent security.

### Installation

    go get github.com/jacky2478/goaxis

### Example

```
// fake code
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
```


Notes:
