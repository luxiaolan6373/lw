package lw

//同GetX,获取上次找图，找字，找色调用完成后，设置的目标的左上角的X坐标值。
// 返回上次设置的X坐标。
func (com *LwSoft) X() int {
	var ret, _ = com.lw.CallMethod("x")
	return int(ret.Val)
}

//同GetY()获取上次找图，找字，找色调用完成后，设置的目标的左上角的y坐标值。
func (com *LwSoft) Y() int {
	var ret, _ = com.lw.CallMethod("y")
	return int(ret.Val)
}

//同GetIdx ,获取上次找图，找字，找色调用完成后，设置的序号。注意,本插件内的所有序号都是从1开始数的
func (com *LwSoft) Idx() int {
	var ret, _ = com.lw.CallMethod("idx")
	return int(ret.Val)
}

// @title    查看乐玩版本号
func (com *LwSoft) Ver() string {
	var ret, _ = com.lw.CallMethod("ver")
	return ret.ToString()
}

//设置全局路径,设置了此路径后,所有接口调用中,相关的文件都相对于此路径. 比如图片,字库等.
func (com *LwSoft) SetPath(path string) int {
	var ret, _ = com.lw.CallMethod("SetPath", path)
	return int(ret.Val)
}

//获取全局路径.(可用于调试)
func (com *LwSoft) GetPath() string {
	var ret, _ = com.lw.CallMethod("GetPath")
	return ret.ToString()
}

//设置是否弹出错误信息,默认是打开.
//show 整形数: 0表示不打开(有错误时不弹窗提示),1表示打开(有错误时弹窗提示)
//返回值: 0 : 失败  1 : 成功
func (com *LwSoft) SetShowErrorMsg(show int) int {
	var ret, _ = com.lw.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}
