# 组合（Composite）

## 问题

将对象组合成树形结构来表示`整体`/`部分`层次关系，可以用相同的方式处理单独对象和组合对象。

## 效果

组件（Component）类是组合类（Composite）和叶子类（Leaf）的父类，可以把组合类看成是树的中间节点。

组合对象拥有一个或者多个组件对象，因此组合对象的操作可以委托给组件对象去处理，而组件对象可以是另一个组合对象或者叶子对象。

## 解决方案

最简单的组合模式就是把对象组合成树形结构，然后就可以像处理树形结构一样去处理组合对象和叶子对象。

```java
public class Tree {
    private Long id;
    private String name;
    private Long pId;
    private List<Tree> child;
    
    // get、set...

    public void add(Tree tree) {
        child.add(tree);
    }

    public void remove(Tree tree) {
        child.remove(tree);
    }
}
```
有的组合模式是组件和组合对象对象分开的，比如叶子节点和组合节点都是独立的。

----
