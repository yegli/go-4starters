.PHONY: greeter quiz calculator chatroom

greeter:
	cd ./greeter && $(MAKE) all

quiz:
	cd ./quiz && $(MAKE) all

calculator:
	cd ./calculator && $(MAKE) all

chatroom:
	cd ./chatroom && $(MAKE) all