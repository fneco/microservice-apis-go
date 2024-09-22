hello:
	echo "Hello, World"

# https://github.com/OpenAPITools/openapi-generator
gen-by-openapi-generator:
	docker run --rm -v "${PWD}:/local" -w /local \
		openapitools/openapi-generator-cli generate \
			-i ./oas.yaml \
			-g go \
			-o ./openapi-generator/generated
