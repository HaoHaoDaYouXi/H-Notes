# npm to pnpm

如何将`npm`迁移到`pnpm`，需要执行以下步骤：

## 卸载`npm`包

首先，将`npm`命令下载的包（node_modules）从当前项目中卸载，`package-lock.json`文件暂时不能删除后面转换需要使用

## 安装`pnpm`

全局安装`pnpm`，以便可以在项目中使用它
```
npm install -g pnpm
```

## 创建配置文件

在项目目录下创建`.npmrc`的文件
```
# pnpm 配置
shamefully-hoist = true # 强制提升模块以避免版本冲突
auto-install-peers = true # 自动安装依赖的对等包
strict-peer-dependencies = false # 严格遵守对等依赖的版本要求
```

## 转换相关文件

将`package-lock.json`转成`pnpm-lock.yaml`文件，保证依赖版本不变
```
pnpm import
```

## 安装依赖包
   
通过`pnpm`安装依赖包
```
pnpm install
```

最后，迁移完成，项目正常运行之后，可以删除原本的`package-lock.json`文件。


----
