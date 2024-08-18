# Default arguments
ARGS ?= 

# Build and run the bot configuration
bot:
	@killall -q air || true
	@echo "Starting air with bot configuration..."
	@air -c .air-bot.toml -- $(ARGS)

# Run ngrok
ngrok:
	@killall -q ngrok || true
	@ngrok http --domain weekly-proper-alien.ngrok-free.app 80

# Build and run the api configuration
api:
	@echo "Starting air with api configuration..."
	@air -c .air-api.toml -- $(ARGS)