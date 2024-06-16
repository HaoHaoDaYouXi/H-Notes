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