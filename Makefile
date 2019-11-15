install: clean-front up 

down:
	docker-compose down

up:
	docker-compose up -d --build

clean-front:
	cd front && \
	rm -f yarn.lock && \
	rm -rf node_modules && \
	yarn install