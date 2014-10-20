Study notes
========

#### CodeWars 战斗笔记 
========

- 字符串重复

```javascript
function stringRepeat(c, n) {
  return new Array(n + 1).join(c);
}
// example:
stringRepeat("*", 5); // result: *****
```

- 数组去掉重复值

```javascript
function noDuplicate(arr) {
  return arr.filter(function(l, i, a) {
    return a.indexOf(l) == i;
  });
}
// example:
noDuplicate([1,2,3,4,1,2]); // result: [1,2,3,4]
```

- 截取两位小数字符

```javascript
function getFixed(num, b) {
  // toFixed方法会四舍五入
  return num.toFixed(b + 1).replace(/\d$/, "");
}
// example:
getFixed(123.956, 2); // result: 123.95
```

#### Code 水滴石穿
=======

- 固定顺序Object比较

```javascript
function compare(obj1, obj2) {
  return JSON.stringify(obj1) === JSON.stringify(obj2);
}
// example:
compare({a:1, b:2}, {a:1, b:2}); // result: true
compare({a:1, b:2}, {b:2, a:1}); // result: false
```
