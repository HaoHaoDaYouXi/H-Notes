@echo off

start cmd /k "mqnamesrv.cmd"
start cmd /k "mqbroker.cmd -n localhost:9876 -c ../conf/broker.conf autoCreateTopicEnable=true"

@pause
