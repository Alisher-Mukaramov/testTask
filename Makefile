appBuild:
	@echo " > start app building"
	docker-compose build
appDeploy:
	@echo " > start app deploy"
	docker-compose up -d

run:
	$(MAKE) appBuild
	$(MAKE) appDeploy