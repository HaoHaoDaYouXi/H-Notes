# Mysql 编译安装步骤

1.安装依赖包
~~~
yum -y install gcc gcc-c++ make cmake ncurses ncurses-devel openssl-devel bison libaio-devel libtirpc libtirpc-devel
~~~

2.创建mysql程序用户
~~~
useradd -M -s /sbin/nologin mysql
~~~

3.创建mysql的安装路径(可直接跳过)
~~~
mkdir -p /home/mysql/mysql5.7.39
~~~

4.解压mysql安装包，配置软件模块
~~~
cd ~
tar -zxvf mysql-5.7.39.tar.gz
~~~

5.配置安装选项
~~~
cd mysql-5.7.39/
cmake . -DCMAKE_INSTALL_PREFIX=/home/mysql/mysql5.7.39 -DINSTALL_BINDIR=/home/mysql/mysql5.7.39/bin -DMYSQL_UNIX_ADDR=/home/mysql/mysql5.7.39/mysql.sock -DSYSCONFDIR=/home/mysql/mysql5.7.39/conf -DMYSQL_DATADIR=/home/mysql/mysql5.7.39/data -DDOWNLOAD_BOOST=1 -DWITH_BOOST=/home/mysql/mysql5.7.39/boost -DMYSQL_TCP_PORT=61306 -DMYSQL_USER=mysql -DEXTRA_CHARSETS=all -DDEFAULT_CHARSET=utf8 -DDEFAULT_COLLATION=utf8_general_ci -DWITH_EXTRA_CHARSETS=all -DWITH_MYISAM_STORAGE_ENGINE=1 -DWITH_INNOBASE_STORAGE_ENGINE=1 -DWITH_MEMORY_STORAGE_ENGINE=1 -DWITH_ARCHIVE_STORAGE_ENGINE=1 -DWITH_BLACKHOLE_STORAGE_ENGINE=1 -DWITH_FEDERATED_STORAGE_ENGINE=1 -DWITH_PARTITION_STORAGE_ENGINE=1 -DWITH_PERFSCHEMA_STORAGE_ENGINE=1 -DENABLED_LOCAL_INFILE=1 -DWITH_READLINE=1
~~~
### cmake 命令说明
~~~
cmake .
 # 指定mysql的安装路径,基础的文件夹,对应mysqld的--basedir参数 /usr/local/mysql
 -DCMAKE_INSTALL_PREFIX=/home/mysql/mysql5.7.39     
 # bin目录位置,MySQL 主执行文件目录  PREFIX/bin
 -DINSTALL_BINDIR=/home/mysql/mysql5.7.39/bin
 # 指定mysql进程监听套接字文件（数据库连接文件）的存储路径    /tmp/mysql.sock
 -DMYSQL_UNIX_ADDR=/home/mysql/mysql5.7.39/mysql.sock 
 # 默认配置my.cnf目录  
 -DSYSCONFDIR=/home/mysql/mysql5.7.39/conf
 # 指定数据库的家目录，数据库文件的存储路径
 -DMYSQL_DATADIR=/home/mysql/mysql5.7.39/data      
 # 表示能 BOOST 文件的下载
 -DDOWNLOAD_BOOST=1 
 # 表示BOOST 文件的下载路径
 -DWITH_BOOST=/home/mysql/mysql5.7.39/boost         
 # 指定端口号 3306
 -DMYSQL_TCP_PORT=3306            
 # 指定管理用户
 -DMYSQL_USER=mysql   
 # 安装所有扩展字符集     
 -DEXTRA_CHARSETS=all           
 # 指定默认使用的字符集编码，如 utf8
 -DDEFAULT_CHARSET=utf8                     
 # 指定默认使用的字符集校对规则
 -DDEFAULT_COLLATION=utf8_general_ci        
 # 扩展性的字符集，支持其他字符集编码  all
 -DWITH_EXTRA_CHARSETS=all                  
 # 安装MYISAM存储引擎
 -DWITH_MYISAM_STORAGE_ENGINE=1             
 # 安装INNOBASE存储引擎
 -DWITH_INNOBASE_STORAGE_ENGINE=1           
 # 安装MEMORY存储引擎
 -DWITH_MEMORY_STORAGE_ENGINE=1             
 # 安装Archive引擎
 -DWITH_ARCHIVE_STORAGE_ENGINE=1 
 # 安装BLACKHOLE引擎
 -DWITH_BLACKHOLE_STORAGE_ENGINE=1 
 # 安装FEDERATED引擎
 -DWITH_FEDERATED_STORAGE_ENGINE=1 
 # 安装PARTITION引擎
 -DWITH_PARTITION_STORAGE_ENGINE=1 
 # 安装PERFSCHEMA引擎
 -DWITH_PERFSCHEMA_STORAGE_ENGINE=1 
 # 不支持example存储引擎
 -DWITHOUT_EXAMPLE_STORAGE_ENGINE=1
 # 允许从本地导入数据
 -DENABLED_LOCAL_INFILE=1             
 # 支持readline程序平台，读取数据按行读取，一行是一个对象     OFF  
 -DWITH_READLINE=1  
 # 启用zlib库支持（zib、gzib相关）   system                        
 -DWITH_ZLIB=system
 # 启用SSL库支持（安全套接层）no
 -DWITH_SSL=no 
 # 不启用基于wrap的访问控制
 -DWITH_LIBWRAP=0
 # 启用性能分析功能
 -DENABLE_PROFILING=1
 # 不启用DEBUG功能
 -DWITH_DEBUG=0
~~~

6.编译安装
~~~
make -j 2 && make install
~~~

如出现报错，需要将文件夹中的CMakeCache文件删除，然后重新编译即可，如果不确定文件位置，可以使用find / -name CMakeCache.txt -type f 进行文件查找。

7.设置系统配置
~~~
cd /home/mysql/mysql5.7.39
mkdir conf
# 配置my.cnf
cp support-files/mysql.server /etc/init.d/mysqld   #添加mysqld系统服务，将mysql添加进系统服务管理中
chmod 755 /etc/init.d/mysqld   #赋予mysqld文件755权限
chkconfig --add /etc/init.d/mysqld   #将mysqld加入系统管理
chkconfig --level 35 mysqld on   #设置mysqld在init3和init5级别开启
echo "export PATH=$PATH:/home/mysql/mysql5.7.39/bin" >> /etc/profile   #将/usr/local/mysql/bin目录加入PATH
source /etc/profile   #刷新profile文件，重载系统环境变量PATH
echo $PATH   #查看PATH
chown -R mysql:mysql /home/mysql/   #将mysql目录的所有文件的属主和属组改为mysql用户
~~~

8.初始化数据库
~~~
/home/mysql/mysql5.7.39/bin/mysqld
 --defaults-file=/home/mysql/mysql5.7.39/conf/my.cnf                     #指定数据库文件的存储路径
 --basedir=/home/mysql/mysql5.7.39           #指定数据库的安装目录
 --datadir=/home/mysql/mysql5.7.39/data                     #指定数据库文件的存储路径
 --user=mysql                                 #指定管理用户
 --initialize-insecure
~~~

9.修改配置文件
~~~
vim /etc/init.d/mysqld   #修改mysqld文件
 basedir=/home/mysql/mysql5.7.39   #找到basedir参数，输入/home/mysql/mysql5.7.39
 datadir=/home/mysql/mysql5.7.39/data   #batadir参数输入/home/mysql/mysql5.7.39/data 
~~~

10.启动并查看mysql服务
~~~
service mysqld start
ps -ef | grep mysqld
~~~

11.登录及退出mysql
~~~
mysql   #登录，密码为空，直接回车即可
use mysql;
update user set password=password("myroot") where user="root";
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'myroot' WITH GRANT OPTION; 
flush privileges;
quit/exit   #退出，需进入数据库
mysql -uroot -pmyroot   #使用root登录，密码为myroot，-p后面不要加空格
~~~

12.配置nginx访问
~~~
stream { 
    server {
        listen 61307;
        proxy_pass 172.16.0.141:61306;
    }
}
~~~


# MySQL系列—MySQL编译安装常见问题(或缺少依赖)及解决方法
=========================
MySQL 编译安装时需要安装的依赖(全)：
yum install -y cmake

yum install ncurses ncurses-devel -y

yum install -y libarchive

yum install -y gcc gcc-c++

yum install -y openssl openssl-devel

yum install -y libtirpc libtirpc-devel

指定boost（下载、在编译项指定即可,见文章尾部： MySQL编译安装常用选项）

安装（安装方法见文章尾部）

以下问题经由（腾讯云服务器CentOS Linux release 8.0.1905 (Core)） 华为云服务器（CentOS Linux release 7.6.1810 (Core) ） 和 mysql-boost-5.7.30.tar.gz 测试而来

华为云服务器（CentOS Linux release 7.6.1810 (Core) ）预装的软件比较多，出现提示需要安装的依赖比较少

**问题1**
-- Could NOT find Curses (missing: CURSES_LIBRARY CURSES_INCLUDE_PATH)

CMake Error at cmake/readline.cmake:71 (MESSAGE):

Curses library not found.  Please install appropriate package,

remove CMakeCache.txt and rerun cmake.On Debian/Ubuntu, package name is libncurses5-dev, on Redhat and derivates it is ncurses-devel.

Call Stack (most recent call first):

cmake/readline.cmake:100 (FIND_CURSES)

cmake/readline.cmake:193 (MYSQL_USE_BUNDLED_EDITLINE)

CMakeLists.txt:581 (MYSQL_CHECK_EDITLINE)



出现原因：

缺少依赖

解决方法：

yum install ncurses ncurses-devel -y

=====================================

**问题2**

CMake Error at cmake/boost.cmake:88 (MESSAGE):
  You can download it with -DDOWNLOAD_BOOST=1 -DWITH_BOOST=<directory>

This CMake script will look for boost in <directory>.  If it is not there,

it will download and unpack it (in that directory) for you.

If you are inside a firewall, you may need to use an http proxy:

export http_proxy=http://example.com:80

Call Stack (most recent call first):

cmake/boost.cmake:174 (COULD_NOT_FIND_BOOST)

CMakeLists.txt:547 (INCLUDE)



出现原因：没有指定boost位置，或boost位置错误

解决方法：

cd 指令或ls等指令验证boost路径是否正确，修改即可。

若果没有boost就需要下载，然后指定就可以了

=====================================

**问题3**

-bash: cmake: command not found

出现原因：

没有安装cmake

解决方法：

yum install -y cmake

=====================================

**问题4**

cmake: symbol lookup error: cmake: undefined symbol: archive_write_add_filter_zstd

出现原因：

缺少依赖

解决方法：

yum install -y libarchive

=====================================

**问题5**

CMake Error at CMakeLists.txt:146 (PROJECT):
  No CMAKE_CXX_COMPILER could be found.

Tell CMake where to find the compiler by setting either the environment

variable "CXX" or the CMake cache entry CMAKE_CXX_COMPILER to the full path

to the compiler, or to the compiler name if it is in the PATH.



出现原因：

缺少gcc-c++

解决方法：

yum install -y gcc gcc-c++

=====================================

**问题6**

Cannot find appropriate system libraries for WITH_SSL=system.

Make sure you have specified a supported SSL version.

Valid options are :

system (use the OS openssl library),

yes (synonym for system),

</path/to/custom/openssl/installation>

CMake Error at cmake/ssl.cmake:63 (MESSAGE):

Please install the appropriate openssl developer package.

Call Stack (most recent call first):

cmake/ssl.cmake:280 (FATAL_SSL_NOT_FOUND_ERROR)

CMakeLists.txt:579 (MYSQL_CHECK_SSL)

出现原因：

缺少依赖

解决方法：

yum install -y openssl openssl-devel

=====================================

**问题7**

-- Found PkgConfig: /usr/bin/pkg-config (found version "1.4.2")
-- Checking for module 'libtirpc'
--  Package 'libtirpc', required by 'virtual:world', not found
CMake Error at cmake/rpc.cmake:76 (MESSAGE):
  Could not find rpc/rpc.h in /usr/include or /usr/include/tirpc

Call Stack (most recent call first):

rapid/plugin/group_replication/configure.cmake:60 (MYSQL_CHECK_RPC)

rapid/plugin/group_replication/CMakeLists.txt:25 (INCLUDE)

出现原因：

缺少依赖

解决方法：

yum install -y libtirpc libtirpc-devel

=====================================

**问题8**

CMake Error at rapid/plugin/group_replication/rpcgen.cmake:100 (MESSAGE):
  Could not find rpcgen

Call Stack (most recent call first):

rapid/plugin/group_replication/CMakeLists.txt:36 (INCLUDE)

出现原因：

缺少依赖

解决方法：

**安装rpcsvc-proto**
安装rpcsvc-proto方法如下：

下载rpcsvs-proto

https://github.com/thkukuk/rpcsvc-proto/releases/download/v1.4.3/rpcsvc-proto-1.4.3.tar.xz

解压

tar -xvf rpcsvc-proto-1.4.3.tar.xz

.configure

make && make install

注意：如果下载的是tar.gz包，这个包需要手动生成configure文件后才能编译安装rpcvsc-proto,需要安装很多依赖项，很繁琐，后面单独介绍

https://github.com/thkukuk/rpcsvc-proto/archive/refs/tags/v1.4.3.tar.gz

====================================


