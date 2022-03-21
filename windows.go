package lw

/*
函数简介:

查找符合条件的窗口

参数定义:

title 字符串: 窗口标题,如果为空，则匹配所有.这里的匹配是模糊匹配.默认为空

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是默认是模糊匹配.默认为空

processname 字符串:  进程名,如果为空，则匹配所有.这里的匹配是精确匹配.默认为空

pid 整形数:进程ID,如果为0,则匹配所有,默认为0

fhwnd 整形数:父窗口,如果为0则查找顶级窗口,默认为0
isvisble 整形数:是否不判断可视状态,0=只查找具有可视属性的窗口,1=忽略可视属性.默认为0
fpid 整形数:父进程ID,如果为0,则匹配所有,.默认为0

exacttitle 整形数:是否精确标题,0=模糊匹配标题,1=精确匹配标题.默认为0

w 整形数:窗口宽度,如果为0则忽略此参数,默认为0

h 整形数:窗口高度,如果为0则忽略此参数,默认为0

lose 整形数:跳过次数,匹配成功多少次才返回,从1开始数,0=首次找到就返回.默认为0

comd 字符串:启动命令,是否精确匹配由下一个参数决定,为空则匹配所有,默认为空

exactcomd　整形数:是否精确匹配启动启动命令,0=模糊匹配,1=精确匹配.默认为0



返回值:

整形数:
整形数表示的窗口句柄，没找到返回0

*/
func (com *LwSoft) FindWindow(title, class, processname string, pid, fhwnd, isvisble, fpid, exacttitle, w, h, lose int, comd string, exactcomd int) int {
	var ret, _ = com.lw.CallMethod("FindWindow", title, class, processname, pid, fhwnd, isvisble, fpid, exacttitle, w, h, lose, comd, exactcomd)
	return int(ret.Val)
}

/*
函数简介:

查找符合条件的下层窗口。

参数定义:

fhwnd 整形数: 父窗口句柄

title 字符串: 窗口标题,如果提供一个空文本，则匹配所有.这里的匹配是模糊匹配.默认为一个空文本

class 字符串: 窗口类名，如果提供一个空文本，则匹配所有. 这里的匹配是默认是模糊匹配.默认为一个空文本

id   整形数: 控件ID,为0匹配所有,默认为0



返回值:

整形数:
整形数表示的窗口句柄，没找到返回0

注意,这里的查找,并不是只匹配子级窗口,而是匹配在fhwnd下层的所有窗口.
*/
func (com *LwSoft) FindSonWindow(fhwnd int, title, class string) int {
	var ret, _ = com.lw.CallMethod("FindSonWindow", fhwnd, title, class)
	return int(ret.Val)
}

/*
函数简介:

获取源窗口的第指定个数的子窗口.

参数定义:

fhwnd 整形数: 父窗口句柄

var_ 整形数: 要取第几个子窗口,从1开始数,默认为1.

返回值:



整形数:
整形数表示的窗口句柄，没找到返回0
*/
func (com *LwSoft) GetSonWindow(fhwnd, var_ int) int {
	var ret, _ = com.lw.CallMethod("GetSonWindow", fhwnd, var_)
	return int(ret.Val)
}

/*
函数简介:

根据指定条件,枚举系统中符合条件的窗口,可以枚举到按键自带的无法枚举到的窗口

参数定义:

title 字符串: 窗口标题,如果为空，则匹配所有.这里的匹配是模糊匹配.默认为空

class 字符串: 窗口类名，如果为空，则匹配所有. 这里的匹配是默认是模糊匹配.默认为空

processname 字符串:  进程名,如果为空，则匹配所有.这里的匹配是精确匹配,但不区分大小写.默认为空

pid 整形数:进程ID,如果为0,则匹配所有,默认为0

fhwnd 整形数:父窗口,如果为0则查找顶级窗口,默认为0

isvisble 整形数:是否不判断可视状态,0=只查找具有可视属性的窗口,1=忽略可视属性.默认为0


exacttitle 整形数:是否精确标题,0=模糊匹配标题,1=精确匹配标题.默认为0



返回值:

字符串 :
返回所有匹配的窗口句柄字符串,格式"hwnd1,hwnd2,hwnd3"


*/
func (com *LwSoft) EnumWindow(title, class, processname string, pid, fhwnd, isvisble, exacttitle int) string {
	var ret, _ = com.lw.CallMethod("EnumWindow", title, class, processname, pid, fhwnd, isvisble, exacttitle)
	return ret.ToString()
}

/*
函数简介:

根据控件ID,获取子窗口.

参数定义:

fhwnd 整形数:父窗口句柄.
id 整形数:控件ID,如果某个窗口的控件ID不为0,那么相对于同级窗口,这个ID值一定是唯一的,且不会变的,控件ID可以用乐玩编程助手查看.


返回值:

整形数:
整形数表示的窗口句柄，没找到返回0



*/
func (com *LwSoft) GetWindowById(fhwnd, id int) int {
	var ret, _ = com.lw.CallMethod("GetWindowById", fhwnd, id)
	return int(ret.Val)
}

/*
函数简介:

获取给定窗口相关的窗口句柄

参数定义:

hwnd 整形数: 窗口句柄

flag 整形数: 取值定义如下

0 : 获取源窗口的父窗口　如果该窗口没有父窗口,则返回0

1 : 获取源窗口的获取第一个儿子窗口

2 : 获取源窗口的第一个兄弟窗口

3 : 预留

4 : 获取源窗口的下一个兄弟窗口

5 : 获取源窗口的上一个兄弟窗口

6 : 获取源窗口的拥有者窗口

7 : 获取源窗口的顶层窗口,如果该窗口本身就是顶级窗口,则返回其自身.

返回值:

整形数:
返回整型表示的窗口句柄


*/
func (com *LwSoft) GetWindow(hwnd, flag int) int {
	var ret, _ = com.lw.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

获取给定坐标的可见窗口句柄,可以获取到按键自带的插件无法获取到的句柄

参数定义:

X 整形数: 屏幕X坐标

Y 整形数: 屏幕Y坐标

返回值:

整形数:
返回整型表示的窗口句柄
*/
func (com *LwSoft) GetPointWindow(x, y int) int {
	var ret, _ = com.lw.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

/*
函数简介:

获取当前鼠标指向的可见窗口句柄,可以获取到按键自带的插件无法获取到的句柄

返回值:

整形数:
返回整型表示的窗口句柄


*/
func (com *LwSoft) GetMousePointWindow() int {
	var ret, _ = com.lw.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

/*
函数简介:

获取指定窗口所在的进程ID.

参数定义:

hwnd 整形数: 窗口句柄

返回值:

整形数:
返回整型表示的是进程ID


*/
func (com *LwSoft) GetWindowProcessId(hwnd int) int {
	var ret, _ = com.lw.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

/*
函数简介:

获取窗口包含外边框的尺寸，会设置X,Y

整形数:
0: 失败
1: 成功

注意，获取的宽高要用lw.x(),lw.y()获取。
*/
func (com *LwSoft) GetWindowSize(hwnd int) int {
	var ret, _ = com.lw.CallMethod("GetWindowSize", hwnd)
	return int(ret.Val)
}

/*
函数简介:

获取窗口客户区域的宽度和高度,会设置X,Y。

参数定义:

hwnd 整形数: 指定的窗口句柄


返回值:

整形数:
0: 失败
1: 成功

注意，获取的宽高要用lw.x(),lw.y()获取。
*/
func (com *LwSoft) GetClientSize(hwnd int) int {
	var ret, _ = com.lw.CallMethod("GetClientSize", hwnd)
	return int(ret.Val)
}

/*
函数简介:

获取窗口包含外边框的左上角在桌面的位置。

参数定义:

hwnd 整形数: 指定的窗口句柄


返回值:

整形数:
0: 失败
1: 成功
注意，获取的位置要用lw.x(),lw.y()获取。
*/
func (com *LwSoft) GetWindowVertex(hwnd int) int {
	var ret, _ = com.lw.CallMethod("GetWindowVertex", hwnd)
	return int(ret.Val)
}

/*
函数简介:

获取指定窗口的一些属性


参数定义:

hwnd 整形数: 指定的窗口句柄

flag 整形数: 取值定义如下

0 : 判断窗口是否存在

1 : 判断窗口是否处于激活

2 : 判断窗口是否可见

3 : 判断窗口是否最小化

4 : 判断窗口是否最大化

5 : 判断窗口是否置顶

6 : 判断窗口是否无响应

7 : 判断窗口是否可用(灰色为不可用)

返回值:

整形数:
0: 不满足条件
1: 满足条件


*/
func (com *LwSoft) GetWindowState(hwnd, flag int) int {
	var ret, _ = com.lw.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

获取窗口的标题或者内容

参数定义:

hwnd 整形数: 指定的窗口句柄

flag  整形数:获取方式，默认为0，取值范围如下。

         1，要取的是窗口的标题

         2，要取的是窗口里的内容

 返回值:

字符串:
窗口的标题或者内容

大部分情况flag=0和flag=1取值的返回值都是一样的，但少数控件标题并不等于内容，具体可以用工具查看。
*/
func (com *LwSoft) GetWindowTitle(hwnd, flag int) string {
	var ret, _ = com.lw.CallMethod("GetWindowTitle", hwnd, flag)
	return ret.ToString()
}

/*
函数简介:

获取窗口的类名

参数定义:

hwnd 整形数: 指定的窗口句柄

返回值:

字符串:
窗口的类名


*/
func (com *LwSoft) GetWindowClass(hwnd int) string {
	var ret, _ = com.lw.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

/*
函数简介:

获取指定窗口所在的进程的exe文件全路径.

参数定义:

hwnd 整形数: 窗口句柄

返回值:

字符串:
返回字符串表示的是exe全路径名


*/
func (com *LwSoft) GetWindowProcessPath(hwnd int) string {
	var ret, _ = com.lw.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

/*
函数简介:

把屏幕坐标转换为窗口坐标

参数定义:

hwnd 整形数: 指定的窗口句柄

x 整形数: 屏幕X坐标

y 整形数: 屏幕Y坐标

返回值:

整形数:
0: 失败
1: 成功

转换后的坐标要用lw.x(),lw.y()获取。
*/
func (com *LwSoft) ScreenToClient(hwnd, x, y int) int {
	var ret, _ = com.lw.CallMethod("ScreenToClient", hwnd, x, y)
	return int(ret.Val)
}

/*
函数简介:

把窗口坐标转换为屏幕坐标

参数定义:

hwnd 整形数: 指定的窗口句柄

x 整形数: 窗口X坐标

y 整形数: 窗口Y坐标

返回值:

整形数:
0: 失败
1: 成功


转换后的坐标要用lw.x(),lw.y()获取。
*/
func (com *LwSoft) ClientToScreen(hwnd, x, y int) int {
	var ret, _ = com.lw.CallMethod("ClientToScreen", hwnd, x, y)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的标题

参数定义:

hwnd 整形数: 指定的窗口句柄

titie 字符串: 标题

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetWindowText(hwnd int, title string) int {
	var ret, _ = com.lw.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

/*
函数简介:

移动指定窗口到指定位置

参数定义:

hwnd 整形数: 指定的窗口句柄

x 整形数: X坐标

y 整形数: Y坐标

width  整形数:新的窗口宽度，为0表示不改变宽度，默认为0

height 整形数:新的窗口高度，为0表示不改变高度，默认为0

返回值:

整形数:
0: 失败
1: 成功

*/
func (com *LwSoft) MoveWindow(hwnd, x, y, width, height int) int {
	var ret, _ = com.lw.CallMethod("MoveWindow", hwnd, x, y, width, height)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的大小

参数定义:

hwnd 整形数: 指定的窗口句柄
width 整形数: 宽度

height 整形数: 高度

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetWindowSize(hwnd, width, height int) int {
	var ret, _ = com.lw.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的状态

参数定义:

hwnd 整形数: 指定的窗口句柄

flag 整形数: 取值定义如下

0 : 关闭指定窗口(发送销毁消息)

1 : 激活指定窗口

2 : 最小化指定窗口,但不激活

3 : 最小化指定窗口,并释放内存,但同时也会激活窗口

4 : 最大化指定窗口,同时激活窗口.

5 : 恢复指定窗口 ,但不激活

6 : 隐藏指定窗口

7 : 显示指定窗口

8 : 置顶指定窗口

9 : 取消置顶指定窗口

10 : 禁止指定窗口

11 : 取消禁止指定窗口

12 : 恢复并激活指定窗口

13 : 强制结束窗口所在进程.

14 : 闪烁指定的窗口

15 : 使指定的窗口获取输入焦点

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetWindowState(hwnd, flag int) int {
	var ret, _ = com.lw.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

设置窗口的透明度

参数定义:

hwnd 整形数: 指定的窗口句柄

trans 整形数: 透明度取值(0-255) 越小透明度越大 0为完全透明(不可见) 255为完全显示(不透明)

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetWindowTransparent(hwnd, trans int) int {
	var ret, _ = com.lw.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}

/*
函数简介:

更改窗口的边框属性。

参数定义:

hwnd 整形数: 指定的窗口句柄

flag 整形数: 有以下几种取值方式

                   0 移除窗口的边框

                   1 移除窗口的关闭按钮

                   2 移除窗口的最小化按钮

                   3 移除窗口的最大化按钮

                   4 禁止通过拉动窗口边框调整窗口尺寸（移除窗口的可调边框属性）

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetWindowBorder(hwnd, flag int) int {
	var ret, _ = com.lw.CallMethod("SetWindowBorder", hwnd, flag)
	return int(ret.Val)
}

/*
函数简介:

为指定窗口指定一个新的父窗口

参数定义:

shwnd 整形数: 要指定父窗口的窗口

fhwnd 整形数: 指定的父窗口，0表示桌面窗口。

返回值:

整形数:
0: 失败
1: 成功


*/
func (com *LwSoft) SetParent(shwnd, fhwnd int) int {
	var ret, _ = com.lw.CallMethod("SetParent", shwnd, fhwnd)
	return int(ret.Val)
}

/*
函数简介:

给目标窗口设置一个自定义的标志值

参数定义:

hwnd 整形数: 要设置标志的窗口

flag 整形数: 要设置的标志

返回值:

整形数:
返回该窗口原来的标志。

注意：这个标志是全局的，任何进程的乐玩对象都可以访问。
*/
func (com *LwSoft) SetWindowFlag(hwnd, flag int) int {
	var ret, _ = com.lw.CallMethod("SetWindowFlag", hwnd, flag)
	return int(ret.Val)
}

/*


函数简介:

取窗口标记

参数定义:

hwnd 整形数: 要取标志的窗口

返回值:

整形数:
返回该窗口的标志。

*/
func (com *LwSoft) GetWindowFlag(hwnd int) int {
	var ret, _ = com.lw.CallMethod("SetWindowFlag", hwnd)
	return int(ret.Val)
}
