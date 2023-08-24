TTY_SERVER_ASSETS=$(wildcard frontend/public/*) frontend/public/index.html

.PHONY: all frontend clean cleanfront rebuild

all:
	cd frontend && npm install && npm run build && cd -

clean:
	rm -fr frontend/public
	@echo "Cleaned"
