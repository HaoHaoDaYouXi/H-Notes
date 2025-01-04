# GitHub

## 远程推送问题

当出现以下错误时，一般是防火墙导致的
```shell
ssh: connect to host github.com port 22: Connection timed out
fatal: Could not read from remote repository.
```

解决办法：

可以换个端口访问，先测试下 443
```shell
ssh -T -p 443 git@ssh.github.com
```
如果是第一次的，会出现一个确定信息，输入yes即可，若出现以下信息，表示可以正常访问
```shell
Hi xxxx! You've successfully authenticated, but GitHub does not provide shell access.
```

现在，我们需要在`~/.ssh/config`（windows目录为用户目录下的ssh）文件中覆盖 SSH 设置
```
HOST github.com
	Hostname ssh.github.com
	Port 443
	User git
```

保存后，再次尝试使用以下命令进行 SSH 连接：
```shell
ssh -T git@github.com
```
若出现一下信息，表示可以正常访问
```shell
Hi xxxx! You've successfully authenticated, but GitHub does not provide shell access.
```
后续正常使用即可


----
