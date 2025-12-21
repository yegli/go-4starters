.PHONY: greeter quiz calculator

greeter:
	cd ./greeter && $(MAKE) all

quiz:
	cd ./quiz && $(MAKE) all

calculator:
	cd ./calculator && $(MAKE) all