package lw

/*
函数简介:
默认键盘鼠标在按下弹起，移动时会在设置的延时间隔的基础加上一点随机延时，可以通过此函数设置随机的上限值，或者关闭随机延时。

参数定义:
mode  整形数: 0=关闭随机延时；其它值为随机延时的上限。

返回值: 0: 失败  1: 成功
*/
func (com *LwSoft) SetRandomDelay(delay int) int {
	var ret, _ = com.lw.CallMethod("SetRandomDelay", delay)
	return int(ret.Val)
}

/*
函数简介:
设置后台键盘鼠标采有同步或者异步消息模式。默认为异步。

参数定义:
mode  整形数: 0＝异步模式； 1＝ 开启同步

返回值:  0: 失败 1: 成功
注 : 此函数影响的接口有KeyPress
注: 此接口必须在绑定之后才能调用。
有些时候，如果是异步发送，如果发送动作太快,中间没有延时,有可能下个动作会影响前面的.
而用同步就没有这个担心.
*/
func (com *LwSoft) EnableKeyMouseSync(mode int) int {
	var ret, _ = com.lw.CallMethod("EnableKeyMouseSync", mode)
	return int(ret.Val)
}

/*
函数简介:
设置前台键鼠的模拟方式.

参数定义:
mode 整形数: 0 正常模式(默认模式) 1 硬件模拟

返回值: 0失败，1成功

如果模式为1，首次运行需要在电脑上安装两个驱动，会弹出来安装驱动的提示，要点同意。vista以上系统要管理员权限。
*/
func (com *LwSoft) SetSimMode(mode int) int {
	var ret, _ = com.lw.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

/*
函数简介:
设置按键时,键盘按下和弹起的时间间隔。高级用户使用。某些窗口可能需要调整这个参数才可以正常按键。

参数定义:
delay 整形数: 延时,单位是毫秒

返回值: 0:失败  1:成功
*/
func (com *LwSoft) SetKeypadDelay(delay int) int {
	var ret, _ = com.lw.CallMethod("SetKeypadDelay", delay)
	return int(ret.Val)
}

/*
函数简介:
设置鼠标单击或者双击时,鼠标按下和弹起的时间间隔。高级用户使用。某些窗口可能需要调整这个参数才可以正常点击。

参数定义:
delay 整形数: 延时,单位是毫秒

返回值:0:失败 1:成功

注 : 此函数影响的接口有LeftClick RightClick MiddleClick LeftDoubleClick
注意,在没有调用  binwindow 绑定窗口之前,所有的键盘鼠标命令都无效!!!
*/
func (com *LwSoft) SetMouseDelay(delay int) int {
	var ret, _ = com.lw.CallMethod("SetMouseDelay", delay)
	return int(ret.Val)
}

/*
函数简介:
鼠标动作模拟真实操作,带移动轨迹,以及点击延时随机.

参数定义:
enable 整形数: 0 关闭模拟  1 开启模拟
mousedelay 整形数: 单位是毫秒. 表示在模拟鼠标移动轨迹时,每移动一次的时间间隔.这个值越大,鼠标移动越慢.
Mousestep 整形数: 表示在模拟鼠标移动轨迹时,每移动一次的距离. 这个值越大，鼠标移动越快速.

返回值:0: 失败  1: 成功

注: 此接口同样对LeftClick RightClick MiddleClick LeftDoubleClick起作用。具体表现是鼠标按下和弹起的间隔会在
当前设定延时的基础上,上下随机浮动50%. 假如设定的鼠标延时是100,那么这个延时可能就是50-150之间的一个值.
设定延时的函数是 SetMouseDelay
注意,在没有调用  binwindow 绑定窗口之前,所有的键盘鼠标命令都无效!!!
*/
func (com *LwSoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	var ret, _ = com.lw.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

/*
函数简介:
按下指定的虚拟键码

参数定义:
vk_code 整形数:虚拟按键码
num 整形数:按键次数,默认1次.

返回值:0:失败 1:成功
注意,在没有调用  binwindow 绑定窗口之前,所有的键盘鼠标命令都无效!!!
*/
func (com *LwSoft) KeyPress(vk_code, num int) int {
	var ret, _ = com.lw.CallMethod("KeyPress", vk_code, num)
	return int(ret.Val)
}

/*
函数简介:
按下指定的虚拟键码,可以是组合键

参数定义:
vk_code 整形数:虚拟按键码
StateKey 整形数:状态键,默认是0,取值范围如下,各值可相加
0=无状态键 1=Alt 2=Ctrl 4=Shift

返回值:0:失败 1:成功
*/
func (com *LwSoft) KeyPressEx(vk_code, StateKey int) int {
	var ret, _ = com.lw.CallMethod("KeyPressEx", vk_code, StateKey)
	return int(ret.Val)
}

/*
函数简介:
按住指定的虚拟键码

参数定义:
vk_code 整形数:虚拟按键码

返回值: 0:失败 1:成功
*/
func (com *LwSoft) KeyDown(vk_code int) int {
	var ret, _ = com.lw.CallMethod("KeyDown", vk_code)
	return int(ret.Val)
}

/*
函数简介:
弹起来虚拟键vk_code

参数定义:
vk_code 整形数:虚拟按键码

返回值:0:失败 1:成功
*/
func (com *LwSoft) KeyUp(vk_code int) int {
	var ret, _ = com.lw.CallMethod("KeyUp", vk_code)
	return int(ret.Val)
}

/*
函数简介:
根据指定的字符串序列，依次按顺序按下其中的字符.

参数定义:
key_str 字符串: 需要按下的字符串序列. 比如"1234","abcd","7389,1462"等.
delay 整形数: 每按下一个按键，需要延时多久. 单位毫秒.这个值越大，按的速度越慢。

返回值:0:失败 1:成功

注: 在某些情况下，SendString和SendString2都无法输入文字时，可以考虑用这个来输入.
但这个接口只支持标英文大小写字母,数字,和半角字符.

注意,在没有调用  binwindow 绑定窗口之前,所有的键盘鼠标命令都无效!!!
*/
func (com *LwSoft) KeyPressStr(key_str string, delay int) int {
	var ret, _ = com.lw.CallMethod("KeyPressStr", key_str, delay)
	return int(ret.Val)
}

/*
函数简介:
按下鼠标左键

参数定义:
返回值:0:失败 1:成功

注意,在没有调用  binwindow 绑定窗口之前,所有的键盘鼠标命令都无效!!!
*/
func (com *LwSoft) LeftClick() int {
	var ret, _ = com.lw.CallMethod("LeftClick")
	return int(ret.Val)
}

/*
函数简介:
双击鼠标左键

参数定义:
返回值:0:失败 1:成功
*/
func (com *LwSoft) LeftDoubleClick() int {
	var ret, _ = com.lw.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

/*
函数简介:

按住鼠标左键
返回值:

整形数:
0:失败
1:成功



*/
func (com *LwSoft) LeftDown() int {
	var ret, _ = com.lw.CallMethod("LeftDown")
	return int(ret.Val)
}

/*
函数简介:

弹起鼠标左键

整形数:
0:失败
1:成功



*/
func (com *LwSoft) LeftUp() int {
	var ret, _ = com.lw.CallMethod("LeftUp")
	return int(ret.Val)
}

/*
函数简介:

鼠标相对于上次的位置移动rx,ry.
参数定义:

rx 整形数:相对于上次的X偏移
ry 整形数:相对于上次的Y偏移

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) MoveR(rx, ry int) int {
	var ret, _ = com.lw.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

/*
函数简介:

把鼠标移动到目的范围内的任意一点

参数定义:

x 整形数:X坐标
y 整形数:Y坐标
w 整形数:宽度(从x计算起)
h 整形数:高度(从y计算起)

返回值:

整形数:
0:失败
1:成功

注: 此函数的意思是移动鼠标到指定的范围(x,y,x+w,y+h)内的任意随机一点.
*/
func (com *LwSoft) MoveToEx(x, y, w, h int) string {
	var ret, _ = com.lw.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

/*
函数简介:

把鼠标移动到目的点(x,y)
参数定义:

x 整形数:X坐标
y 整形数:Y坐标

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) MoveTo(x, y int) int {
	var ret, _ = com.lw.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

/*
函数简介:

按下鼠标右键

返回值:

整形数:
0:失败
1:成功



*/
func (com *LwSoft) RightClick() int {
	var ret, _ = com.lw.CallMethod("RightClick")
	return int(ret.Val)
}

/*
函数简介:

按住鼠标右键

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) RightDown() int {
	var ret, _ = com.lw.CallMethod("RightDown")
	return int(ret.Val)
}

/*
函数简介:

弹起鼠标右键

返回值:

整形数:
0:失败
1:成功

*/
func (com *LwSoft) RightUp() int {
	var ret, _ = com.lw.CallMethod("RightUp")
	return int(ret.Val)
}

/*
函数简介:

按下鼠标中键

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) MiddleClick() int {
	var ret, _ = com.lw.CallMethod("MiddleClick")
	return int(ret.Val)
}

/*
函数简介:

滚轮向下滚

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) WheelDown() int {
	var ret, _ = com.lw.CallMethod("WheelDown")
	return int(ret.Val)
}

/*
函数简介:

滚轮向上滚

返回值:

整形数:
0:失败
1:成功


*/
func (com *LwSoft) WheelUp() int {
	var ret, _ = com.lw.CallMethod("WheelUp")
	return int(ret.Val)
}

/*
函数简介:

获取鼠标特征码. 当绑定时,附加信息里有"需要取鼠标特征码"时获取到的是后台鼠标特征，否则是前台鼠标特征.

参数定义:

mod:整数型,获取方式,取值有0和1两种,哪种能取到用哪种,默认值为0.

返回值:

整数型:
成功时，返回鼠标特征码.
失败时，返回0
*/
func (com *LwSoft) GetCursorShape(mod int) int {
	var ret, _ = com.lw.CallMethod("GetCursorShape", mod)
	return int(ret.Val)
}

/*
函数简介:

获取指定的按键状态.(前台信息,不是后台)

参数定义:

key 整形数:虚拟按键码

返回值:

整形数:
0:弹起
1:按下
*/
func (com *LwSoft) GetKeyState(key int) int {
	var ret, _ = com.lw.CallMethod("GetKeyState", key)
	return int(ret.Val)
}

/*
函数简介:

等待指定的按键按下 (前台,不是后台)

参数定义:

vk_code 整形数:虚拟按键码,当此值为0，表示等待任意按键(包含鼠标);当此值为－1，表示等待任意按键，但不包含鼠标键，其它值则为真实键代码。默认为0.
time_out 整形数:等待多久,单位毫秒. 如果是0，表示一直等待,默认为0.


返回值:

整形数:
超时返回-1,否则返回被按下的键的键代码。
鼠标左键是1,鼠标右键时2,鼠标中键是4.
*/
func (com *LwSoft) WaitKey(vk_code, time_out int) int {
	var ret, _ = com.lw.CallMethod("WaitKey", vk_code, time_out)
	return int(ret.Val)
}

/*
函数简介:

向指定窗口发送文本数据

参数定义:


str 字符串: 发送的文本数据

mod 整形数: 方式　　默认值为0,取值方式有以下几种

　　　　　　0:控件模式

　　　　　　1:模式2

　　　　　　2:输入法模式1

　　　　　　3:输入法模式2　(需要在绑定的时候设置附加信息32)

                  4:输入法模式3　(需要在绑定的时候设置附加信息32)

hwnd 整形数: 指定的窗口句柄,默认为绑定的窗口.


返回值:

整形数:
0: 失败
1: 成功

示例:

lw.SendString "我是来测试的"

注： 有时候发送中文，可能会大部分机器正常，少部分会乱码。这种情况一般有两个可能
1. 系统编码没有设置为GBK
2. 目标程序里可能安装了改变当前编码的软件，比如常见的是输入法. （尝试卸载）

*/
func (com *LwSoft) SendString(str string, mod, hwnd int) int {
	var ret, _ = com.lw.CallMethod("SendString", str, mod, hwnd)
	return int(ret.Val)
}
