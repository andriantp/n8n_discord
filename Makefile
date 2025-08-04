config:
	echo "🔌 ENVIRONMENT creating..."; \
	log=$$(grep '^log=' .env | cut -d '=' -f2); \
	token=$$(grep '^token=' .env | cut -d '=' -f2); \
	channel_name=$$(grep '^channel_name=' .env | cut -d '=' -f2); \
	channel_id=$$(grep '^channel_id=' .env | cut -d '=' -f2); \
	url=$$(grep '^url=' .env | cut -d '=' -f2); \
	name=$$(grep '^name=' .env | cut -d '=' -f2); \
	value=$$(grep '^value=' .env | cut -d '=' -f2); \
	host=$$(grep '^host=' .env | cut -d '=' -f2); \
	tag=$$(grep '^tag=' .env | cut -d '=' -f2); \
	que=$$(grep '^que=' .env | cut -d '=' -f2); \
	routing_key=$$(grep '^routing_key=' .env | cut -d '=' -f2); \

webhook: config
	@bash -c "set -a; source .env; set +a; mode=webhook go run ."

rabbit: config	
	@bash -c "set -a; source .env; set +a; mode=rabbit go run ."
