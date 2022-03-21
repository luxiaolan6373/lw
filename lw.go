package lw

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os"
	"os/exec"
)

//lw对象
type LwSoft struct {
	lw       *ole.IDispatch
	IUnknown *ole.IUnknown
}

//注册乐玩com组件
func Reg() (unknown *ole.IUnknown, err error) {
	var com LwSoft
	dir, _ := os.Getwd()
	ret := exec.Command("regsvr32", "/s", dir+"\\lw.dll")
	_, err = ret.Output()
	if err != nil {
		panic(err)
	}
	com.IUnknown, err = oleutil.CreateObject("lw.lwsoft3")
	if err != nil {
		panic("注册失败,请注意lw.dll是否存在以及是否是管理员权限运行:" + err.Error())
	}

	return com.IUnknown, err
}
func UnReg() {
	dir, _ := os.Getwd()
	cmd := exec.Command("regsvr32", "/s", "/u", dir+"\\lw.dll")
	_, err := cmd.Output()
	if err != nil {
		panic("取消注册失败,请注意lw.dll是否存在以及是否是管理员权限运行:" + err.Error())
	}
}
func New() (lw *LwSoft, err error) {
	var com LwSoft
	// 告诉window使用什么方式调用com组件
	err = ole.CoInitialize(0)
	if err != nil {
		return nil, err
	}
	//释放对象
	defer ole.CoUninitialize()

	//创建com对象
	com.IUnknown, err = oleutil.CreateObject("lw.lwsoft3")
	if err != nil {
		com.IUnknown, err = Reg()
	}

	// 查询接口是否正常
	com.lw, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {

		return nil, err
	}
	return &com, nil
}
