.PHONY: create-jar generate start up stop down
create-jar:
	cd backend/openapi-custom-generator && mvn package -DskipTests=true && cp target/go-custom-server-openapi-generator-1.0.0.jar ../tools/

generate:
	cd backend && java -cp tools/openapi-generator-cli-6.3.0.jar:tools/go-custom-server-openapi-generator-1.0.0.jar \
	org.openapitools.codegen.OpenAPIGenerator generate -g go-custom-server -i ../openapi.yaml && \
	cd ../frontend && npm run generate-typescript-axios

build-up:
	docker-compose up -d --build

up:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down