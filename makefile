# https://github.com/ogen-go/ogen
gen-by-oapi-codegen:
	docker compose run oapi-codegen-service \
		oapi-codegen \
			--config=./oapi_codegen/config.yaml \
			./oas.yaml
