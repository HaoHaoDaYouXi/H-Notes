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

## <a id="scyfz">删除远端提交</a>

要删除远端仓库的提交记录，需要执行以下步骤：

- 使用`git log`查找你想要删除的提交的哈希值。
- 使用`git revert`创建一个新的提交来撤销指定的提交变化。
- 使用`git push`将这个撤销提交推送到远端。

例：假设你想要删除的提交哈希值是`bad_commit_hash`。

- `git log`：查找提交哈希值
- `git revert bad_commit_hash`：创建撤销提交
- `git push origin main`：将更改推送到远端，分支名为`main`

如果想要彻底删除提交（不保留提交内容），可以使用以下步骤：

- 使用`git rebase`交互式地编辑历史记录。
- 在编辑界面中，使用`drop `命令来删除特定的提交。
- 使用`git push --force`将更改强制推送到远端。

例：
```
git rebase -i HEAD~3 # 交互式地编辑最近3个提交，将会打开编辑器
# 在编辑器中，找到你想要删除的提交前面，将其对应的 `pick` 改为 `drop`
# 保存并关闭编辑器
git push --force origin main # 强制推送到远端
```

**注意：强制推送会重写历史记录，这可能会影响其他协作者的工作，所以在执行这些操作之前，请确保已经和团队内的其他成员进行了沟通**

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


----
