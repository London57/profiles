now=$(date +%m-%d_%H-%M-%S)
wsl -e bash -c "
 cd / && \
 cd home/serzh/ && \
 sudo mkdir profiles_$now &&\
 sudo cp -r /mnt/c/Users/serzh/.vscode/tin/profiles ~/profiles_$now &&\
 cd profiles_$now/profiles &&
 exec bash \
"