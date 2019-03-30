package goaxis

import (
    "testing"
)

const (
    c_Test_Int = 333
    c_Test_Int2 = 222
    c_Test_Int_Expect = 555
    c_Test_Str = "test string"
    c_Test_Str_Expect = "test from workRsp.Pull"
)

var pushData int

func Notify(mode string, caller ICaller, ds IDataSet) error {
    switch caller.Last() {
    case "TestPull":
        if str, ok := ds.Input().Get().(string); ok && str == c_Test_Str {
            ds.Output().Set(c_Test_Str_Expect)
        }
    case "TestPush":
       if idata, ok := ds.Input().Get().(int); ok {
            ds.Output().Set(idata + c_Test_Int2)
       }
    }
    return nil
}

func TestCreate(t *testing.T) {
    ttest := Create("libtest", Notify)

    // test pull
    dataPull := DataSet(c_Test_Str)
    if err := ttest.SyncPull(ttest.Caller("TestPull"), dataPull); err != nil {
        t.Errorf("test failed at SyncPull, detail: %v", err.Error())
        return
    }

    sdata := dataPull.Output().Get()
    if sdata.(string) != c_Test_Str_Expect {
        t.Errorf("test failed at SyncPull, expect: %v, but got: %v", c_Test_Str_Expect, sdata)
        return
    }

    // test push
    dataPush := DataSet(c_Test_Int)
    if err := ttest.SyncPush(ttest.Caller("TestPush"), dataPush); err != nil {
        t.Errorf("test failed at SyncPush, detail: %v", err.Error())
        return
    }

    idata := dataPush.Output().Get()
    if idata.(int) != c_Test_Int_Expect {
        t.Errorf("test failed at SyncPull, expect: %v, but got: %v", c_Test_Int_Expect, idata)
        return
    }
}