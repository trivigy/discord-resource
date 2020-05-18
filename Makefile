
dev:
	docker build -t discord-resource:dev ./

publish:
	docker build -t taylorsilva/discord-resource:latest ./ \
		&& docker push taylorsilva/discord-resource:latest
