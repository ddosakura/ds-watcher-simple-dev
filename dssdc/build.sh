rm -rf ./dist
xgo --targets=linux/amd64,darwin/amd64,windows/amd64    \
    -v -x -dest dist -out dssdc ./dssdc
