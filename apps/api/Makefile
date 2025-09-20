include .env
export

tailwind:
	tailwindcss -i ./templ/assets/css/input.css -o ./templ/assets/css/output.css --watch

air:
	air -c .air.toml

templ:
	go tool templ generate -path templ

ent:
	go generate ent/generate.go

dev:
	make -j3 templ tailwind air

clean_temp:
	find ./templ -name "*_templ.go" -delete

.PHONY: air templ templ_dev ent tailwind dev clean_temp

