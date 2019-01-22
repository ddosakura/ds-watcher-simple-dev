cd $GOPATH/src/github.com/ddosakura/simple-dev
rm -rf ./dist
xgo --targets=linux/*,darwin/amd64,windows/amd64    \
    -v -x -dest dist -out dssdc ./dssdc
