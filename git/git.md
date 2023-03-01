# 删除github中的提交历史记录
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