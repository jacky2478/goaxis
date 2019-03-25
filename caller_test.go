package goaxis

import (
    "testing"
)

const (
    c_Test_Master = "test master"
    c_Test_First_Add = "test add"
    c_Test_Last_Expect = "test add"
    c_Test_First_Expect = "test master"
    c_Test_Push_Data = "test data for push"
    c_Test_Push_Data_Expect = "test data for push"

    c_Test_Format = "work for format"
    c_Test_Format_Expect = "test master->work for format"
)

func TestFirst(t *testing.T) {
    caller := newCaller(c_Test_Master)
    if caller.First() !=  c_Test_First_Expect {
        t.Errorf("TestFirst failed, expect: %v, but got: %v", c_Test_First_Expect, caller.First())
    }
}

func TestLast(t *testing.T) {
    caller := newCaller(c_Test_Master)
    caller.Push(c_Test_First_Add)
    if caller.Last() !=  c_Test_Last_Expect {
        t.Errorf("TestLast failed, expect: %v, but got: %v", c_Test_Last_Expect, caller.First())
    }
}

func TestPush(t *testing.T) {
    caller := newCaller(c_Test_Master)
    caller.Push(c_Test_Push_Data)
    if caller.Last() !=  c_Test_Push_Data_Expect {
        t.Errorf("TestPush failed, expect: %v, but got: %v", c_Test_Push_Data_Expect, caller.Last())
    }
}

func TestFormat(t *testing.T) {
    caller := newCaller(c_Test_Master)
    caller.Push(c_Test_Format)
    if caller.Format("->") !=  c_Test_Format_Expect {
        t.Errorf("TestFormat failed, expect: %v, but got: %v", c_Test_Format_Expect, caller.Format("->"))
    }
}