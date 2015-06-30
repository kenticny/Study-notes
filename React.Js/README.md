React.Js 学习笔记
===================

## 目录

#### 1. 概念

##### 1.1 什么是 React.Js

##### 1.2 React.Js 适合做什么

##### 1.3 相关学习资源

#### 2. 相关操作

##### 2.1 Hello World!

```javascript
<script src="lib/react.min.js"></script>

<div class="container"></div>
<script>
    var el = React.createElement("p", null, "Hello World!");
    React.render(el, document.querySelector("#container"));
</script>

```

##### 2.2 React 组件

定义一个组件，在定义React组件时，必须要实现一个`render`方法

```javascript
React.createClass({
    render: function() {
        return React.createElement("p", null, "Hello World!");
    }
});
```

##### 2.3 JSX

JSX是对Javascript的语法扩展（**是一种JS**），可以在JS中以类似于HTML元素的方式定义React元素（不是完全等同于HTML，有部分熟悉和特性有变化）

在使用JSX时需要引入`JSXTransformer.js`来完成转换工作。使用JSX方法如下：

```javascript
<script type="text/jsx">
    //...
</script>

<script type="text/jsx"></script>
```

在JSX中使用Javascript表达式需要使用一对花括号`{}`引用：

```javascript
var name = "kenticny"
React.render(
    <CustomComponent name={name}></CustomComponent>
);
```

##### 2.4 props 属性

props 为React中实例元素的属性

```javascript
var MyHome = React.createClass({
    render: function() {

        // 通过自身属性
        var address = this.props.address;
        return <span>{address}</span>
    }
});

React.render(
    <MyHome address="my_address" />,
    document.querySelector("body")
);
```

##### 2.5 内联样式

内联样式即直接HTML元素上通过style属性添加CSS样式，在React中可以通过给style属性赋值一个CSS的Object对象实现内联样式

```javascript
var style = {
    width: "200px",
    height: "200px"
};

React.createClass({
    render: function() {
        return <span style={style}></span>
    }
});
```

在使用React内联样式时，所有CSS属性遵循驼峰式命名

```javascript
var style = {
    borderRadius: "50%",

    // webkit moz 样式首字母大写
    WebkitTransition: "color .5s",

    // ms 样式首字母小写
    msTransition: "color .5s"
};
```

##### 2.6 状态 State

props是无状态的，仅取决于传入的值，所以需要state来记录组件的某些状态，React中状态机详细如下

- **state**
组件的状态变量，即保存组件状态的变量，任何时候都可以使用this.state读取当前状态

- **getInitialState()**
在创建组件时设置组件初始状态的方法，必须返回一个Object对象或者null

    ```javascript
    React.createClass({
        getInitalState: function() {
            return {
                name: "kenticny"
            };
        },
        render: function() {
            return <div>{this.props.name}</div>
        }
    });
    ```

- **setState(currentState)**
设置组件的状态，并且可以重新渲染组件（给this.state赋值不能重新渲染组件），currentState为当前状态，必须为object对象类型。在调用setState设置状态时，**会和当前的this.state进行合并**

##### 2.7 组件生命周期

- **componentWillMount()**
组件`初次渲染前`调用，在组件生命周期内只被调用一次

- **componentDidMount()**
组件`初次渲染后`调用，在组件生命周期内只被调用一次

- **componentWillReceiveProps(props)**
组件`即将设置新属性`时被调用，这个方法在初次渲染时不会被调用，在此方法内调用setState方法不会引起组件重新渲染
参数`props`为新的属性值

- **shouldComponentUpdate(props, state)**
组件`即将重新渲染`时调用，该方法的返回值为boolean，当为true时表示可以重新渲染，为false时不可以重新渲染，**这个方法在组件初次渲染或者调用forceUpdate方法强制渲染时不被调用**
参数`props`为新的属性值，`state`为新的状态值

- **componentWillUpdate(props, state)**
组件`即将重新渲染`时调用，这个方法在组件初次渲染时不被调用，**在此方法内不可以调用setState方法**

- **componentDidUpdate()**
组件`重新渲染后`被调用

- **componentWillUnmount()**
组件实例`即将从DOM树移除`时被调用，这个方法在整个生命周期中只会被调用一次

##### 2.8 在React中访问DOM元素
在React中访问DOM元素需要在DOM元素中添加ref属性，然后通过`findDOMNode`方法访问，一般会采用第三方类库访问DOM元素（`jQuery`, `Zepto`）

```javascript
<input name="a" ref="name" value="abc" />

var el = React.findDOMNode(this.ref.name)
el.value // abc
```