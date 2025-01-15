APP = samsamoohooh

SWAGGO_VERSION := v1.16.4

.PHONY: swag
swag:
	@go get github.com/swaggo/swag/cmd/swag@$(SWAGGO_VERSION)
	@go install github.com/swaggo/swag/cmd/swag@$(SWAGGO_VERSION)
	@swag init -g cmd/app/app.go -o ./docs/swagger