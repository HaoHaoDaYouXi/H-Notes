# 设计模式和原则

开发中不管做什么都涉及到一个名词，它叫 **设计** ，比如这个功能怎么设计，这个系统架构怎么设计，等等

设计又有两种概念：模式、原则

## <a id="sjms">设计模式</a>
设计模式是指在软件工程中，针对常见的设计问题和重复出现的场景，提出的一种标准化的解决方案。

它是软件开发人员在实践中总结出来的一套成熟且被广泛接受的指导原则，用于解决特定类型的问题，以提高代码的可重用性、可读性和可维护性。

设计模式通常包括以下要素：
- 模式名称：一个简洁的名字，用来标识这个模式。
- 问题：描述在何种情况下应用该模式。
- 效果：描述模式带来的好处和可能的权衡。
- 解决方案：描述如何通过设计来解决这个问题。

设计模式可以分为三大类：
- 创建型模式：关注的是对象的创建机制，使创建过程更加灵活和可控制。
- 结构型模式：关注的是如何组合类或对象以获得新的功能。
- 行为型模式：关注的是对象间的职责分配和通信机制。

设计模式的应用有助于：
- 代码复用：利用已经过验证的解决方案。
- 提高代码可读性：遵循共同的设计思路，便于理解和维护。
- 方便沟通：提供了一种通用的语言，帮助团队成员之间交流设计思想。
- 增强软件的可维护性和可扩展性：通过清晰的结构和职责分离，使得代码更易于修改和扩展。

### <a id="cjx">创建型模式</a>
- [单例](create/singleton.md)
- [工厂](create/factory.md)
  - [简单工厂](create/factory.md#jdgc)
  - [工厂方法](create/factory.md#gcff)
  - [抽象工厂](create/factory.md#cxgc)
- [建造者](create/builder.md)
- [原型模式](create/prototype.md)

### <a id="jgx">结构型模式</a>
- [代理](structure/proxy.md)
- [适配器](structure/adapter.md)
- [桥接](structure/bridge.md)
- [组合](structure/composite.md)
- [装饰](structure/decorator.md)
- [外观](structure/facade.md)
- [享元](structure/flyweight.md)

### <a id="xwx">行为型模式</a>
- [责任链](behavior/responsibility.md)
- [命令](behavior/command.md)
- [解释器](behavior/interpreter.md)
- [迭代器](behavior/iterator.md)
- [中介者](behavior/mediator.md)
- [备忘录](behavior/memento.md)
- [观察者](behavior/observer.md)
- [状态](behavior/state.md)
- [策略](behavior/strategy.md)
- [模板方法](behavior/template.md)`
- [访问者](behavior/visitor.md)
- [空对象](behavior/null.md)

## <a id="sjyz">设计原则</a>

设计原则是在软件工程和软件架构设计中用于指导开发人员创建高质量、可维护、可扩展的软件系统的准则。

这些原则帮助确保软件的结构合理，能够适应未来的变化，并且易于理解和维护。

以下是一些软件设计原则：
- 单一职责原则（Single Responsibility Principle, SRP）
  - 一个类应该只有一个引起它变化的原因。这意味着类应该专注于单一的功能或责任。
- 开闭原则（Open-Closed Principle, OCP）
  - 软件实体（如类、模块、函数等）应该对扩展开放，对修改关闭。即容易添加新功能而不需改变现有代码。
- 里氏替换原则（Liskov Substitution Principle, LSP）
  - 子类必须能够替换其基类并保持程序的所有期望的性质不变。这有助于确保继承的正确使用。
- 依赖倒置原则（Dependency Inversion Principle, DIP）
  - 高层模块不应该依赖于低层模块，二者都应该依赖于抽象；抽象不应该依赖于细节，细节应该依赖于抽象。
- 接口隔离原则（Interface Segregation Principle, ISP）
  - 使用多个专门的接口比使用单一的总接口要好。客户端不应该被迫依赖于它不需要的方法。
- 迪米特法则（Law of Demeter, LoD）
  - 一个对象应该对其他对象有尽可能少的了解。也称为最少知识原则，减少对象之间的耦合。
- 合成复用原则（Composite/Aggregate Reuse Principle, CARP）
  - 优先使用对象组合而不是继承来达到复用的目的。

这些原则有助于创建健壮、灵活和易于维护的软件系统。

它们鼓励模块化、解耦和良好的封装，从而促进代码的重用和降低维护成本。




----
