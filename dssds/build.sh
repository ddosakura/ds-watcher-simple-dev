for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64 arm; do
        echo $GOOS-$GOARCH
        CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -v -o ./build/dssds-$GOOS-$GOARCH ./dssds
    done
done
