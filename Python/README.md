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
- 元组: 不可更改的数组
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
		
##### 2.5 元组(Tuple)

定义一个元组

```python
tuple = (1, 2, 3)
```
如果定义只有一个元素的元组

```python
tuple = (1,)
```
元组不可变是指元组中元素`指向的位置不可变`，所以当元组中元素是一个引用的时候，引用的内容是可更改的

```python
tuple = (1, 2, ['a', 'b'])
tuple[2][0] = 'x'
print tuple    # (1, 2, ['x', 'b'])
```

##### 2.6 流程控制（判断和循环）

- 条件判断

	条件判断使用`if`语句，程序块不需要使用`{}`包括，使用`缩进`控制，和其他语言不同的是，`if条件`后和`else`要加`:`，多重条件判断使用`elif`关键字
	
	```python
	age = 20
	if age > 18:
		print 'The age more than 18'
		print 'adult'
	elif age < 10:
		print 'You are too young'
	else:
		print 'The age less than 18'
	```

	条件判断可以使用简写，类似于Javascript，`非零数值`、`非空字符串`、`非空list`都为`True`，其余为`False`
	
	```python
	x = 'abc'
	if x:
		print 'Not Null'
	```

- 循环

	和大多数语言一直，循环有两种，`for循环`和`while循环`，`for循环`格式为`for x in ...`，可以遍历list或者tuple中的元素
	
	```python
	names = ['James', 'Davis', 'Kurie']

	# names也可以是元组
	# names = ('James', 'Davis', 'Kurie')
	
	for name in names:
		print name
	
	# James
	# Davis
	# Kurie
	```
	
	`range(n)`可以生成0-n的数组，比如计算0-100的和，可以遍历一个range生成的数组
	
	```python
	sum = 0
	for i in range(101):
		sum += i
	print sum   # 5050
	```
	
	`while循环`和其他语言实现基本相同，使用`while`实现上面的问题
	
	```python
	sum = 0
	i = 1
	while i <= 100:
		sum += i
		i += 1
	print sum 
	```
	
##### 2.7 字典

`Dictionary`和其他语言中的`Map`或者`Object`类似，都是键值对，定义一个字典可以使用`{key: value}`形式定义

```python
dic = {'name': 'kenticny', 'age': 25 }
print dic['name']     # 'kenticny'
```

可以通过key在字典中添加值

```python
dic['gender'] = 'male'
print dic    # {'name': 'kenticny', 'age': 25, 'gender': 'male' }
```

如果key不在字典中，使用`dic[key]`获取值会报错，可以通过`in`判断key是否在字典中

```python
'gender' in dic   # True
```

也可以通过`get`方法来获取值，如果不存在则返回`None`

```python
dic.get('ok')   # None
```

`get`方法如果key不存在还可以指定返回值

```python
dic.get('name', 'oo')   # kenticny
dic.get('xxx', 'oo')    # oo
```

`pop`方法可以删除字典中对应key的值并返回

```python
gender = dic.pop('gender')    # male
print dic     # {'name': 'kenticny', 'age': 25 }
```

遍历字典，可以使用`iteritems()`方法

```python
dic = {'a': 'aaa', 'b': 'bbb', 'c': 'ccc'}
for k, v in dic.iteritems():
	print k, v

# a  aaa
# b  bbb
# c  ccc
```

##### 2.8 集合（Set）

`set集合`是一组key的集合，没有值，在集合中，key没有重复的

```python
set1 = set([1, 2, 3])
print set1   # set([1, 2, 3])
```

`add`方法可以在集合中添加元素，如果添加了重复的元素则没有体现

```python
set1.add(4)
print set1   # set([1,2,3,4])
set1.add(2)
print set1   # set([1,2,3,4])
```

`remove`方法可以删除集合中的元素，如果元素不存在则会报错

```python
set1.remove(2)
print set1    # set([1,3,4])
```

通过for循环可以遍历集合

```python
for i in set1:
	print i

# 1
# 3
# 4
```

可以通过`&`和`|`来获取交集或者并集的操作

```python
set1 = set([1,2,3])
set2 = set([2,3,4,5])
set3 = set1 & set2
set4 = set1 | set2
print set1, set2
# set([2,3]), set([1,2,3,4,5])
```

##### 2.9 函数

定义函数使用`def`关键字

```python
def say_hello(name):
	print 'Hello, ' + name
	
say_hello('kenticny')   # Hello, kenticny
```

`pass`作为占位符，可以避免没有执行代码的时候报错

```python
if x > 0:
	pass
```

`isinstance`可以检查函数出参数的类型

```python
def my_abs(a):
	if not isinstance(a, (int, float)):
		raise TypeError('Error Parameters Type')
	if a > 0:
		return a
	else:
		return -a
```

函数还可以返回多个值

```python
def coordinate(x, y):
	if x > 0 and y > 0:
		return x, y
	else:
		return 0, 0

a, b = coordinate(1, 8)
print a, b     # 1, 8
```

由于调用函数传递的参数数量必须和定义函数时一致，所以当有可选参数时，函数可以使用默认参数来实现

```python
def register(name, age, city = 'Beijing', gender = 'Unknown'):
	if age < 18:
		return 'You are Young'
	else:
		return '%s, %d, %s, %s' % (name, age, city, gender)
	
register('kenticny', 20)
register('kenticny', 20, 'Shanghai')
register('kenticny', 20, gender = 'Male')
register('kenticny', 20, 'Shanghai', 'Female')
```

> 容易犯的错误
> 
> 默认参数不可以默认为引用的参数，比如数组，如果需要使用数组，可以使用None作为参数的默认值
> 
```python
def plus(L = None):
	if L is None:
		L = []
	L.append('End')
	return L
```

函数还可以传入可变参数，可变参数使用`*param`表示，可变参数的`类型`必须是一样的

```python
def addd(*numbers):
	sum = 0
	for i in numbers:
		sum += i
	return sum
```

可变参数传递可以直接传递某个数的参数，也可以在`数组`或者`元组`前加`*`表示

```python
addd(1,2,3)   # 6

list = [1,2,3,4]
addd(*list)   # 10

tuple = (1,2,3,4)
addd(*tuple)  # 10
```

函数可以传入关键字参数，关键字参数使用`**param`表示，可以以一个字典的形式传递参数

```python
def person(name, age, **option):
	print name, age, option
```

关键字参数可以直接传递字典，在字典前加`**`表示

```python
person('kenticny', 15)     # kenticny 15 {}
person('kenticny', 15, gender = 'male')   # kenticny 15 {'gender': 'male'}

dic = {'a': 'aaaa', 'b': 'bbbb'}
person('kenticny', 17, **dic)
```

> 当在一个函数中使用多种类型的参数时，必须按照 `固定参数`, `默认参数`, `可变参数`, `关键字参数`的顺序使用


##### 2.10 高阶函数

在函数也可将函数类型传入作为参数

```python
def say_hello(name):
	print 'Hello, ' + name
	
def whoareyou(id, say):
	if id == '0001':
		say('Peter')
	else:
		say('Amy')
		
whoareyou('0001', say_hello)
# Hello, Peter
```

### 3 特性

##### 3.1 切片

`切片`可以截取数组、元组、字符串中的一部分

```python
list = [1,2,3,4,5,6,7,8,9,10]
tuple = (1,2,3,4,5,6,7,8,9,10)
string = 'ABCDEFGHIJK'

print list[0:4], tuple[0:4], string[0:4]
# [1,2,3,4]  (1,2,3,4)  'ABCD'
# 以下举例中list,tuple,string都适用

print list[:4]     # [1,2,3,4]
print list[:]      # [1,2,3,4,5,6,7,8,9,10] 该数组为复制，并非引用
print list[5:]     # [6,7,8,9,10]
print list[-3:-1]  # [8,9]
print list[0:5:2]  # [1,3,5]  每隔两个取一个
```

##### 3.2 列表生成器

即生成一个数组，当只需要生成简单的数列时，可以使用`range`实现

```python
range(1,11)    # [1,2,3,4,5,6,7,8,9,10]
```

生成复杂列表时可以使用生成器，下面例子为生成1-10的平方列表

```python
[x * x for x in range(1, 11)]  # [1,4,9,16,...,100]
```

列表生成器中循环可以进行嵌套

```python
[a + b for a in 'XY' for b in 'MN']  
# ['XM', 'XN', 'YM', 'YN']
```

##### 3.3 生成器（generator）