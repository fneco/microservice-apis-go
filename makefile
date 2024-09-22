hello:
	echo "Hello, World"

# https://github.com/OpenAPITools/openapi-generator
gen-by-openapi-generator:
	docker run --rm -v "${PWD}:/local" -w /local \
		openapitools/openapi-generator-cli generate \
			-i ./oas.yaml \
			-g go \
			-o ./openapi-generator/generated

# https://github.com/ogen-go/ogen
gen-by-oapi-codegen:
	docker compose run oapi-codegen-service \
		oapi-codegen \
			--config=./oapi-codegen/config.yaml \
			./oas.yaml
gen-model-by-oapi-codegen:
	docker compose run oapi-codegen-service \
		oapi-codegen \
			--config=./oapi-codegen/model.yaml \
			./oas.yaml

# https://github.com/ogen-go/ogen
gen-by-ogen:
	docker run --rm --volume ".:/workspace" \
		ghcr.io/ogen-go/ogen:latest \
			--target workspace/ogen/generated --clean workspace/oas.yaml
