PROJECT_NAME=smartway-task
APP_LOCAL_NAME=web-backend

DOCKER_LOCAL_IMAGE_NAME=$(PROJECT_NAME)/$(APP_LOCAL_NAME)

WORK_DIR_LINUX=./cmd/smartway-task
CONFIG_DIR_LINUX=./cmd/smartway-task/config

docker.run: docker.build
	docker compose -f cmd/smartway-task/docker-compose.yaml up -d

docker.build: build.linux
	docker build -t $(DOCKER_LOCAL_IMAGE_NAME) -f cmd/smartway-task/Dockerfile .

run.linux: build.linux
	go run $(WORK_DIR_LINUX)/*.go \
		-config.files $(CONFIG_DIR_LINUX)/application.yaml \
		-env.vars.file $(CONFIG_DIR_LINUX)/application.env \

build.linux: build.linux.clean
	mkdir -p $(WORK_DIR_LINUX)/build
	go build -o $(WORK_DIR_LINUX)/build/main $(WORK_DIR_LINUX)/*.go
	cp -R $(CONFIG_DIR_LINUX)/* $(WORK_DIR_LINUX)/build

build.linux.local: build.linux.clean
	mkdir -p $(WORK_DIR_LINUX)/build
	go build -o $(WORK_DIR_LINUX)/build/main $(WORK_DIR_LINUX)/*.go
	cp -R $(CONFIG_DIR_LINUX)/* $(WORK_DIR_LINUX)/build
	@echo "build.local: OK"

build.linux.clean:
	rm -rf $(WORK_DIR_LINUX)/build

tests.run:
	go test ./internal/domain/employee/tests/...

tests.run.verbose:
	go test -v \
		./internal/domain/employee/tests/...

swagger.gen:
	swag init --parseDependency --parseInternal -g ./cmd/smartway-task/main.go -o ./cmd/smartway-task/docs

mock.gen: mock.employee_storage.gen

mock.employee_storage.gen:
	mockgen -source=internal/domain/employee/employee_storage.go \
	-destination=internal/domain/employee/mocks/mock_employee_storage.go
