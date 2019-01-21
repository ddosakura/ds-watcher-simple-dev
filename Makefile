build:
	./dssdc/build.sh & ./dssds/build.sh
docker:
	docker build -t ddosakura/simple-dev:v0.0.1-alpha.1 .