for GOOS in darwin linux; do
    for GOARCH in 386 amd64 arm; do
        echo $GOOS-$GOARCH
        # TODO: sqlite need cgo, so can't use `CGO_ENABLED=0`, next step, i will use xgo to replace it
        GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/dssdc-$GOOS-$GOARCH ./dssdc
    done
done

for GOOS in windows; do
    for GOARCH in 386 amd64 arm; do
        echo $GOOS-$GOARCH
        # TODO: sqlite need cgo, so can't use `CGO_ENABLED=0`, next step, i will use xgo to replace it
        GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/dssdc-$GOOS-$GOARCH.exe ./dssdc
    done
done