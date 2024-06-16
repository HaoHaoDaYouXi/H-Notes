## [Gradle](https://gradle.org/)
Gradle是一款Google推出的基于JVM、 通用灵活的 项目构建工具，支持Maven，JCenter多种第三方仓库;支持传递性依赖管理、废弃了繁杂的xml文件，转而使用 简洁的 、 支持多种语言 (例如：java、groovy等)的 build脚本文件 。

官网地址:https://gradle.org/

学习一个东西的时候，要了解其概念，下面先了解一下gradle的相关概念：

- Gradle是一个基于Apache Ant和Apache Maven概念的项目自动化构建开源工具。它使用一种基于Groovy的特定领域语言(DSL)来声明项目设置，也增加了基于Kotlin语言的kotlin-based DSL，抛弃了基于XML的各种繁琐配置，面向Java应用为主。当前其支持的语言C++、Java、Groovy、Kotlin、Scala和Swift，计划未来将支持更多的语言。----摘自百度百科

- 自我理解与总结如下：
  - gradle类似于maven，是一个集项目jar包管理、依赖管理、项目打包等操作为一体的工具。
  - gradle不同于maven的地方在于，取消maven笨重的xml配置，以独特简便的groovy语言代替大量繁琐的xml配置项。
  - 拥有自己脚本语言的gradle更加灵活，我们可以在项目构建的时候通过[groovy](../groovy/groovy.md)语言，去灵活的创建任务，依据自己的需求对项目进行构建，相比于maven来说，使用groovy进行构建的gradle，扩展性和灵活性更高。

目前已经有相当一部分公司在逐渐使用Gradle作为项目构建工具了。

作为Java开发程序员,如果想下载Spring、SpringBoot等Spring家族的源码，基本上基于Gradle构建的

## [Gradle安装](gradle_install.md)
