# Default arguments
ARGS ?= 

# Build and run the bot configuration
bot:
	@echo "Starting air with bot configuration..."
	@air -c .air-bot.toml -- $(ARGS)

# Build and run the api configuration
api:
	@echo "Starting air with api configuration..."
	@air -c .air-api.toml -- $(ARGS)