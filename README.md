Personal
========

#### CodeWars 战斗笔记 
========

+ 字符串重复

```javascript
function stringRepeat(c, n) {
  return new Array(n + 1).join(c);
}
```

+ 数组去掉重复值

```javascript
function noDuplicate(arr) {
  return arr.filter(function(l, i, a) {
    return a.indexOf(l) == i;
  });
}
```
