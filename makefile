# https://github.com/ogen-go/ogen
gen-by-oapi-codegen:
	docker compose run oapi-codegen-service \
		oapi-codegen \
			--config=./config/oapi-codegen.yaml \
			./oas.yaml
