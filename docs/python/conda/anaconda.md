# Anaconda

## <a id="an_anaconda">centos 安装Anaconda</a>

### 下载Anaconda

首先，你需要下载Anaconda的安装脚本。可以使用wget命令从Anaconda的官网获取最新版本的安装脚本。

如果官网进不去，或者下载太慢，可以在清华大学开源软件镜像站下载：[清华大学开源软件镜像站](https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/)，从中选择适合自己的版本进行下载

打开终端并运行以下命令：
```shell
wget https://mirrors.tuna.tsinghua.edu.cn/anaconda/archive/Anaconda3-2024.10-1-Linux-x86_64.sh
```

### 运行安装脚本

下载完成后，给安装脚本添加执行权限，并运行它：
```shell
chmod +x Anaconda3-2024.10-1-Linux-x86_64.sh
./Anaconda3-2024.10-1-Linux-x86_64.sh
```

### 按照提示进行安装

在运行安装脚本后，会出现许可协议，按Enter键查看协议内容，最后输入yes接受协议。

你将被询问安装路径，默认路径通常是/home/username/anaconda3。你可以按Enter键使用默认路径，或者输入新的路径。

安装完成后，系统会询问是否将Anaconda添加到PATH环境变量中，选择yes。

### 初始化Anaconda
安装完成后，运行以下命令初始化Anaconda：
```shell
source ~/anaconda3/bin/activate
```

### 验证安装

验证Anaconda是否成功安装，如果安装成功，将显示conda的版本号。
```shell
conda --version
```

### 更新conda

安装完成后，建议更新conda到最新版本：
```shell
conda update conda
```

### 使用Anaconda

你可以使用conda create命令创建新的环境，使用conda install命令安装所需的包。


## <a id="sjhxz">Anaconda升级和卸载</a>

1.升级
升级Anaconda需要先升级conda
~~~
conda update conda
conda update anaconda
conda update anaconda-navigator    //update最新版本的anaconda-navigator
~~~
2.卸载Anaconda软件

由于Anaconda的安装文件都包含在一个目录中，所以直接将该目录删除即可。删除整个Anaconda目录：

计算机控制面板->程序与应用->卸载 //windows

或者

找到C:\ProgramData\Anaconda3\Uninstall-Anaconda3.exe执行卸载
~~~
rm -rf anaconda //ubuntu
~~~
最后，建议清理下.bashrc中的Anaconda路径。

conda环境使用基本命令：
~~~
conda update -n base conda #update最新版本的conda
conda create -n xxxx python=3.5 #创建python3.5的xxxx虚拟环境
conda activate xxxx #开启xxxx环境
conda deactivate #关闭环境
conda env list #显示所有的虚拟环境
conda info --envs #显示所有的虚拟环境
~~~

## <a id="azxbb">Anaconda安装最新的TensorFlow版本</a>

一般从anaconda官网下载的anaconda，查看tensorflow依然还是1.2的版本，现在用conda更新TensorFlow，解决方法：

1. 打开anaconda-prompt
2. 查看tensorflow各个版本：（查看会发现有一大堆TensorFlow源，但是不能随便选，选择可以用查找命令定位）
    ~~~
    anaconda search -t conda tensorflow
    ~~~
3. 找到自己安装环境对应的最新TensorFlow后（可以在终端搜索anaconda，定位到那一行），然后查看指定包<USER/PACKAGE>可安装版本信息命令
    ~~~
    anaconda show <USER/PACKAGE>
    ~~~

    查看tensorflow版本信息
    ~~~
    anaconda show anaconda/tensorflow
    ~~~

4. 第4步会提供一个下载地址，使用下面命令就可安装1.8.0版本tensorflow
    ~~~
    conda install --channel https://conda.anaconda.org/anaconda tensorflow=1.8.0
    ~~~
    更新，卸载安装包：
    ~~~
    conda list #查看已经安装的文件包
    conda list -n xxx #指定查看xxx虚拟环境下安装的package
    conda update xxx #更新xxx文件包
    conda uninstall xxx #卸载xxx文件包
    ~~~

5. 删除虚拟环境
    ~~~
    conda remove -n xxxx --all //创建xxxx虚拟环境
    ~~~

6. 清理（conda瘦身）
    conda clean就可以轻松搞定！
    - 第一步：通过conda clean -p来删除一些没用的包，这个命令会检查哪些包没有在包缓存中被硬依赖到其他地方，并删除它们。
    - 第二步：通过conda clean -t可以将conda保存下来的tar包。
    ~~~
    conda clean -p //删除没有用的包
    conda clean -t //tar打包
    conda clean -y -all //删除所有的安装包及cache
    ~~~

7. 重命名env
    Conda是没有重命名环境的功能的, 要实现这个基本需求, 只能通过愚蠢的克隆-删除的过程。
    切记不要直接mv移动环境的文件夹来重命名, 会导致一系列无法想象的错误的发生!
    ~~~
    conda create --name newname --clone oldname //克隆环境
    conda remove --name oldname --all //彻底删除旧环境
    ~~~

## <a id="szjxy">设置镜像源</a>

查看默认文件中的源
```
conda config --show channels
```

添加清华源
```
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/
```

添加阿里云镜像源
```
conda config --add channels https://mirrors.aliyun.com/anaconda/pkgs/free/
conda config --add channels https://mirrors.aliyun.com/anaconda/pkgs/main/
```

添加中科大源
```
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/pkgs/main/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/conda-forge/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/msys2/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/bioconda/
conda config --add channels https://mirrors.ustc.edu.cn/anaconda/cloud/menpo/
```

删除默认源
```
conda config --remove channels defaults
```

删除文件中指定的源
```
conda config --remove channels https://mirrors.aliyun.com/anaconda/pkgs/free/
```

设置搜索时显示通道地址
```
conda config --set show_channel_urls yes
```

----
