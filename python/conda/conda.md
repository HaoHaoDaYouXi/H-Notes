# conda

## conda命令

### 创建环境
```shell
conda create --name myenv
```
创建一个名为`myenv`的新环境。同时也可以可以添加包名到此命令之后，例如：
```shell
conda create --name myenv numpy scipy
```
将会在创建环境的同时安装`numpy`和`scipy`

### 激活和退出环境
```shell
conda activate myenv
```
激活名为`myenv`的环境。激活环境后，任何安装的包都将被安装到这个特定的环境中。

切换环境也是直接输入要使能的环境即可。
```shell
conda deactivate
```
退出当前激活的环境。

### 安装和卸载包
```shell
conda install package_name
```
在当前活动的环境中安装指定的包。
```shell
conda uninstall package_name
```
卸载当前环境中的指定包。

### 更新包
```shell
conda update package_name
```
更新当前环境中的指定包到最新版本。
```shell
conda update --all
```
更新当前环境中所有可以更新的包。

### 列出环境中的包
```shell
conda list
```
显示当前活动环境中已安装的所有包及其版本。

### 删除环境
```shell
conda remove --name myenv --all
```
删除名为`myenv`的环境及其中的所有包。

### 查找包
```shell
conda search package_name
```
查找可用的包版本。

### 信息和帮助
```shell
conda info
```
显示有关当前安装的`conda`的详细信息，包括环境位置、活动环境等。
```shell
conda help
```
显示帮助菜单或为特定命令提供帮助，如`conda help install`。

### 更改下载源
列出当前通道：
```shell
conda config --show channels
```
输出如下：
```shell
channels:
  - defaults
```
输入`conda info`可以看到默认的下载通道的URLs

由于下载源服务器在国外，下载速度慢，建议更换为国内下载源：
```shell
conda config --remove channels defaults
```
添加清华大学的镜像：
```shell
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud//pytorch/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge/
```
