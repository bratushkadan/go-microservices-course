
# services binary executables names
FRONT_END_BINARY=frontend_app
BROKER_BINARY=broker_app

CGO_ENABLED ?= 0

up:
	@echo "Starting Docker Containers..."
	docker-compose up -d &
	@echo "Docker Containers Started."

up_build: build_broker
	@echo "Stopping docker containers (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker containers..."
	docker-compose up --build -d
	@echo "Docker containers are built and running."

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Stopped docker compose."

build_broker:
	@echo "Building broker binary..."
	cd ./broker && go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	@echo "Building frontend binary..."
	@go build -o ./front-end/${FRONT_END_BINARY} ./front-end/cmd/web
	@echo "Built frontend binary."

start: build_front
	@echo "Starting frontend..."
	@cd ./front-end && ./${FRONT_END_BINARY} &
	@echo "Started frontend."

stop:
	@echo "Stopping front end..."
	@pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"