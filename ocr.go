package lw

/*
函数简介:
设置字库,成功返回1,失败返回0

参数定义:
index 整形数:字库的序号,取值为0-9,目前最多支持10个字库
file 字符串:这里可以是字库的文件路径(相对全SetPath的路径或者绝对路径),也可以直接用字库里的内容.
pwd 字符串:字库密码，如果对字库进行过加密，要提供字库的字库密码。可空，默认为一个空的字符串。

返回值: 0:失败  1:成功
*/
func (com *LwSoft) SetDict(index int, file, pwd string) int {
	var ret, _ = com.lw.CallMethod("SetDict", index, file, pwd)
	return int(ret.Val)
}

/*
函数简介:

表示使用哪个字库文件进行识别(index范围:0-9)
设置之后，永久生效，除非再次设定 成功返回1,失败返回0

参数定义:
index 整形数:字库编号(0-9)

返回值: 0:失败  1:成功
*/
func (com *LwSoft) UseDict(index int) int {
	var ret, _ = com.lw.CallMethod("SetDict", index)
	return int(ret.Val)
}

/*
函数简介:

识别屏幕范围(x1,y1,x2,y2)内符合color_format的字符串,并且相似度为sim,sim取值范围(0.1-1.0),

这个值越大越精确,越大速度越快,越小速度越慢,请斟酌使用!

参数定义:

x1 整形数:区域的左上X坐标
y1 整形数:区域的左上Y坐标
x2 整形数:区域的右下X坐标
y2 整形数:区域的右下Y坐标
color_format 字符串:颜色格式串,有两种格式：

                            1，直接写RGB色加-RGB色偏，如“9f2e3f-030303|2d3f2f-000000"，支持多个。

                            2，#后面加数值，如“#128”，表示采用灰度阀值二值化，值的取值范围为-6到510，只能设置1个。具体可用工具测试。

sim 小数型:相似度,取值范围0.1-1.0　　此参数的缺省值为0.95

linesign  文本型 :换行标志,表示如果碰到换行,用什么代替换行符,此参数的缺省值为一个空文本.

isbackcolor 整数型:是否采用背景色识字,1=color_format 参数所提供的颜色为背景色,其它值表示提供的是文本颜色,该参数缺省值为0.

返回值:

字符串:
返回识别到的字符串


*/
func (com *LwSoft) Ocr(x1, y1, x2, y2 int, color_format string, sim float32, linesign string, isbackcolor int) string {
	var ret, _ = com.lw.CallMethod("Ocr", x1, y1, x2, y2, color_format, sim, linesign, isbackcolor)
	return ret.ToString()
}

/*
函数简介:
在屏幕范围(x1,y1,x2,y2)内,查找string(可以是任意个字符串的组合),并返回符合color_format的坐标位置,相似度sim同Ocr接口描述. 此方法会设置　x,y,idx
(多色,差色查找类似于Ocr接口,不再重述)

参数定义:
x1 整形数:区域的左上X坐标
y1 整形数:区域的左上Y坐标
x2 整形数:区域的右下X坐标
y2 整形数:区域的右下Y坐标
string 字符串:待查找的字符串,可以是字符串组合，比如"长安|洛阳|大雁塔",中间用"|"来分割字符串
color_format 字符串:颜色格式串,有两种格式：
   1，直接写RGB色加-RGB色偏，如“9f2e3f-030303|2d3f2f-000000"，支持多个。
   2，#后面加数值，如“#128”，表示采用灰度阀值二值化，值的取值范围为-6到510，只能设置1个。具体可用工具测试。
sim 小数型:相似度,取值范围0.1-1.0　此参数的缺省值为0.95
isbackcolor 整数型:是否背景色,为1表示color_format 所提供的颜色为背景色,否则表示提供的为文字颜色.此参数的缺省值为0

返回值: 找到返回1,没找到返回0
*/
func (com *LwSoft) FindStr(x1, y1, x2, y2 int, string, color_format string, sim float32, isbackcolor int) int {
	var ret, _ = com.lw.CallMethod("FindStr", x1, y1, x2, y2, string, color_format, sim, isbackcolor)
	return int(ret.Val)
}

//还有其它的实在是用不上,自己封装吧
