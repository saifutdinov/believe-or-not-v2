generate-proto-backend:
	cd proto && protoc --go_out=../backend/grpc --go_opt=paths=source_relative \
    --go-grpc_out=../backend/grpc --go-grpc_opt=paths=source_relative \
    game.proto

generate-proto-frontend:
	cd proto && buf generate

run-backend:
	cd backend && go run main.go

run-frontend:
	cd frontend && npm run dev

lint-backend:
	cd backend && golangci-lint run --fix

lint-frontend:
	cd frontend && npm run lint:fix

configure-frontend:
	cd frontend && npm i

configure-back:
	cd backend && go mod tidy

run-docker:
	docker-compose -f config/docker/dev/docker-compose.yml up

rebuild-docker:
	docker-compose -f config/docker/dev/docker-compose.yml build
	docker-compose -f config/docker/dev/docker-compose.yml up --force-recreate --remove-orphans