# git

## <a id="scyfz">删除原分支并重新创建一个</a>
~~~
1.切换到一个脱离主分支的另外一条全新主分支，不用太在意叫什么，因为后面还会修改分支名称
	git checkout --orphan latest_branch
2.暂存所有改动过的文件，内容为当前旧分支的最新版本所有文件
 	git add -A
3.提交更改
 	git commit -am "commit message"
4.删除原始主分支
	git branch -D main
5.将当前分支重命名为 main
	git branch -m main
6.强制更新您的存储库
	git push -f origin main
7.本地main关联远端main
	git branch --set-upstream-to=origin/main
~~~

## <a id="tjgxml">git 提交和更新脚本命令</a>

### <a id="win_bat">windows脚本</a>

**git pull**
```shell
echo "git pull orgin main...."

git pull origin main
```

**git push**
```shell
echo "Start submitting code to the local repository"
echo "The current directory is：%cd%"
git add *
echo;

echo "Commit the changes to the local repository"
echo "please enter the commit info...."
set now=%date% %time%
echo %now%
set /p message=
git commit -m "%now% %message%"
echo;

echo "Commit the changes to the remote git server"
git push --set-upstream origin main
echo;

echo "Batch execution complete!"
echo;
pause;
```
