# 安装
1、下载并上传[安装包](cronolog-1.6.2.tar.gz)
2、解压
~~~
tar zxvf cronolog-1.6.2.tar.gz
~~~
3、进入cronolog安装文件所在目录
~~~
# cd cronolog-1.6.2
~~~
4、运行安装
~~~
# ./configure
# make
# make install
~~~
5、tomcat日志按日期分割脚本修改
~~~
if [ -z "$CATALINA_OUT" ] ; then
  CATALINA_OUT="$CATALINA_BASE"/logs/catalina.%Y-%m-%d.out
shift
  #touch "$CATALINA_OUT"
  if [ "$1" = "-security" ] ; then
    if [ $have_tty -eq 1 ]; then
      echo "Using Security Manager"
    fi
    shift
    $_NOHUP "$_RUNJAVA" "$LOGGING_CONFIG" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
      -Djava.endorsed.dirs="$JAVA_ENDORSED_DIRS" -classpath "$CLASSPATH" \
      -Djava.security.manager \
      -Djava.security.policy=="$CATALINA_BASE"/conf/catalina.policy \
      -Dcatalina.base="$CATALINA_BASE" \
      -Dcatalina.home="$CATALINA_HOME" \
      -Djava.io.tmpdir="$CATALINA_TMPDIR" \
      org.apache.catalina.startup.Bootstrap "$@" start 2>&1 \
      | /usr/local/sbin/cronolog "$CATALINA_OUT" >> /dev/null&

  else
    $_NOHUP "$_RUNJAVA" "$LOGGING_CONFIG" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
      -Djava.endorsed.dirs="$JAVA_ENDORSED_DIRS" -classpath "$CLASSPATH" \
      -Dcatalina.base="$CATALINA_BASE" \
      -Dcatalina.home="$CATALINA_HOME" \
      -Djava.io.tmpdir="$CATALINA_TMPDIR" \
      org.apache.catalina.startup.Bootstrap "$@" start 2>&1 \
      | /usr/local/sbin/cronolog "$CATALINA_OUT" >> /dev/null&
~~~