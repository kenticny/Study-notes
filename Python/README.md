Python学习笔记
===============

### 1. 开始

##### 1.1 几种常见的python解释器

- CPython: 官方默认的解释器
- IPython: 基于官方解释器增强
- PyPy: 编译型解释器
- Jython: 运行于Java平台
- IronPython: 运行于.Net平台

##### 1.2 Hello World (输出关键字)

```python
print 'hello world'
```
输出信息关键字`print`, 输出多条信息可以使用逗号`,`分隔

```python
print 'hello world'
print 'hello world', 'kenticny'
```

`print`也可以输出数值

```python
print 100 + 200    # 300
print '100 + 200 = ', 100 + 200   # 300
```

##### 1.3 输入关键字

输入信息python中提供了`raw_input`方法，可以将输入的信息赋值给变量，并且可以使用该变量，这个方法中可以附加提示信息

```python
a = raw_input()
a = raw_input('please input:')
```

### 2 基础内容

Python采用`缩进语法`，使用`#`注释，而且`大小写敏感`

##### 2.1 数据类型

- 整数: 可以进行`+`,`-`,`*`,`\`等数值运算
- 浮点数: 可以进行`+`,`-`,`*`,`\`等数值运算
- 字符串: 可以通过`\`对特殊字符进行转移
- 布尔值: 使用`True`和`False`表示（首字母大写），可以进行`and`,`or`,`not`运算
	- `True and True` 为 True
	- `True and False` 为 False
	- `False and False` 为 False
	- `True or True` 为 True
	- `True or False` 为 True
	- `False or False` 为 False
	- `not True` 为 Fasle
	- `not False` 为 True

- 空值: 使用`None`表示
- 列表: 类似数组, 使用`[]`定义, 通过`len`获取列表长度
- 字典:
- 自定义类型:

##### 2.2 字符串编码

Python默认采用ASCII编码，如果使用Unicode编码的字符，可以使用`u'...'`来表示，在文件中时，需要在开头指定编码

```python
# -*- coding: utf-8 -*-
a = '你好'
print len(a)  # 6 因为utf-8编码每个字符占用3个字节
a = u'你好'
print len(a)  # 2
```

##### 2.3 字符串格式化

字符串中可以使用占位符，和常见语言的占位符基本相同

- 整数: %d
- 浮点数: %f
- 字符串: %s
- 十六进制整数: %x

```python
a = '我是%s, 今年%d岁, 大家好' % ('kenticny', 25)
print a
```

占位符还可以补全位数，字符串以空格补全，整数开头以0补全，浮点数四舍五入

```python
a = '%02d - %.2f' % (3, 2.1353131)
print a   # 03 - 2.14
```

##### 2.4 列表(LIST)

类似Javascript的数组，并且可以进行数组的相关操作，列表的赋值为引用，即:

```python
a = [1, 2, 3]
print a  # [1, 2, 3]

b = a
b.append(4)
print a  # [1, 2, 3, 4]
print b  # [1, 2, 3, 4]

```

- 定义: 可以包含不同类型的元素, 包括列表嵌套

	```python
	list = [1, 'a', 2.5, True]
	print list
	
	# 正序遍历
	print list[0], list[1], list[2], list[3]
	
	# 逆序遍历
	print list[-1], list[-2], list[-3], list[-4]
	```
	
- 方法: 
	-  实例
		
		```python
		list = [1, 'a', True]
		```

	- `len`: 获取列表长度

		```python
		len(list)   # 3
		```
		
	- `append(item)`: 在列表最后追加元素
		
		```python
		list.append('abc')   
		# [1, 'a', True, 'abc']
		```
		
	- `insert(index, item)`: 在指定位置插入元素
		
		```python
		list.insert(1, False)   
		# [1, False, 'a', True, 'abc']
		```
		
	- `pop()`: 删除列表最后一个元素并返回

		```python
		list.pop()  # 'abc'
		# [1, False, 'a', True]
		```
		
	- `pop(index)`: 删除列表制定位置的元素并返回
		
		```python
		list.pop(2)  # 'a'
		# [1, False, True]
		```