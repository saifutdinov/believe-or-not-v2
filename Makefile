PROTO_DIR=proto
GO_OUT=backend/proto
TS_OUT=frontend/src/grpc

proto:
	buf generate

dev-backend:
	cd backend && go run main.go

dev-frontend:
	cd frontend && npm run dev

lint-backend:
	cd backend && golangci-lint run

lint-backend-fix:
	cd backend && golangci-lint run --fix

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down