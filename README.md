# Goaxis

Goaxis is a component-oriented transformation hub.

### Describe

Goaxis follows the idea of high cohesion and low coupling, which can ensure that components are completely independent, the interaction between components is controlled by the scheduling layer, and the scheduling layer is composed of actual concrete business logic. Goaxis only provides basic data acquisition and native data acquisition. Supports synchronous asynchronous mode data request and broadcast, concurrent security.

### Installation

    go get github.com/jacky2478/goaxis

### Example

```
// fake code
package main
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
    
    // inside libdb:
    // pull datas from other module: tdb.SyncPull, tdb.ASyncPull
    // push datas by boradcast: tdb.SyncPush, tdb.ASyncPush
    libdb.Init(tdb)

    // inside libweb:
    // pull datas from other module: tweb.SyncPull, tweb.ASyncPull
    // push datas by boradcast: tweb.SyncPush, tweb.ASyncPush
    libweb.Init(tweb)

    libweb.ServeAndListen()
}
```


Notes:
