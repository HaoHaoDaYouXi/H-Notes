# Gradle常用命令

当我们想知道一个工具有哪些命令的时候`gradlew --help`

查看版本号`gradle -v`

检查依赖并编译打包`gradle build`

编译跳过测试`gradle build -x test`

编译并打出`Debug`包`gradlew assembleDebug`

编译打出`Debug`包并安装`gradlew installDebug`

编译并打出`Release`包`gradlew assembleRelease`

编译打出`Release`包并安装`gradlew installRelease`

`Debug`/`Release`编译并打印日志`gradlew assembleDebug --info or gradlew assembleRelease --info`

清除命令`gradlew clean`
