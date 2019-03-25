package goaxis

type hub struct {
    module string
    response ICallback
}

func newHub(module string, rsp ICallback) IHub {
    return &hub{module: module, response: rsp}
}

func (p *hub) Caller(names ...string) ICaller {
    return newCaller(p.module).Push(names...)
}

// request by sync, example: module A get data from module B
func (p *hub) SyncPull(caller ICaller, ds IDataSet) error {
    return p.response.Pull(C_Mode_Sync, caller, ds)
}

// request by async, example: websocket module A wait data to send from a channel generated by module B
func (p *hub) ASyncPull(caller ICaller, ds IDataSet) error {
    return p.response.Pull(C_Mode_ASync, caller, ds)
}

// broadcast data to others module, example: websocket module A broadcast data once receive something from network
func (p *hub) SyncPush(caller ICaller, ds IDataSet) error {
    return p.response.Push(C_Mode_Sync, caller, ds)
}

// broadcast data to others module, example: websocket module A broadcast data once receive something from network
func (p *hub) ASyncPush(caller ICaller, ds IDataSet) error {
    return p.response.Push(C_Mode_ASync, caller, ds)
}

