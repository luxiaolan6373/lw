package lw

/*FindPic
函数简介:
查找指定区域内的图片,位图必须是24位色格式,支持透明色,当图像上下左右4个顶点的颜色一样时,则这个颜色将作为透明色处理.
这个函数可以查找多个图片,找到其中任何一张就返回.本函数会设置x,y,idx.

参数定义:
x1 整形数:区域的左上X坐标
y1 整形数:区域的左上Y坐标
x2 整形数:区域的右下X坐标
y2 整形数:区域的右下Y坐标
pic_name 字符串:图片名，或者图像数据缓冲区（十进制）,可以是多个图片,比如"test.bmp|test2.bmp|123456"
delta_color 字符串:颜色色偏比如"203040" 表示RGB的色偏分别是20 30 40 (这里是16进制表示).该参数的缺省值为：“000000”
sim 小数型:相似度,取值范围0.1-1.0.该参数的缺省值为0.95
dir 整形数:查找方向 0: 从左到右,从上到下 1: 从左到右,从下到上 2: 从右到左,从上到下 3: 从右到左, 从下到上；该参数的缺省值为0
timeout 整形数:超时，单位为毫秒，如果小于等于0则找一次就返回，否一直找，直到找到或者超时为止；该参数的缺省值为0
ischick 整形数:找到以后是否点击 1＝找到后点击，0＝找到后并不点击。该参数的缺省值为0
chickdex 整形数:点击偏移X，如果要点击，相对于找到的点的X坐标偏移多少点击，该参数的缺省值为0
chickdey 整形数:点击偏移Y，如果要点击，相对于找到的点的Y坐标偏移多少点击，该参数的缺省值为0
chickdely 整形数:点击延时， 如果要点击，找到后等多久单击，单位为毫秒，该参数的缺省值为0

返回值:
找到返回1，没有找到返回0

*/
func (com *LwSoft) FindPic(x1, y1, x2, y2 int, pic_name, delta_color string, sim int, dir float32, timeout, ischick, chickdex, chickdey, chickdely int) int {
	var ret, _ = com.lw.CallMethod("FindPic", x1, y1, x2, y2, pic_name,
		delta_color, sim, dir, timeout, ischick, chickdex, chickdey, chickdely)
	return int(ret.Val)
}

/*
函数简介:
取上次调用FindPic命令找到的图片的图片名字。
返回值:
如果上次找到，返回上次找到的图片的图片名字，否则返回一个空文本。
*/
func (com *LwSoft) GetFindedPicName() string {
	var ret, _ = com.lw.CallMethod("GetFindedPicName")
	return ret.ToString()
}

//其它的自己依葫芦画瓢自己补吧.
