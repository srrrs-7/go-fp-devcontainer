PHONY: help
help:
	cat ./Makefile

#################
# Dev Container #
#################
.PHONY: dev-cp

dev-cp:
	cp .devcontainer/compose.override.yaml.example .devcontainer/compose.override.yaml

#############
# Container #
#############
.PHONY: cp fmt vet test tidy graph env

API_MOD = ./apps/api
PKGS_MOD = ./apps/pkgs
MODS = $(API_MOD) $(PKGS_MOD)

cp:
	cp compose.override.yaml.example compose.override.yaml

test:
	for mod in $(MODS); do \
		(cd $$mod && go test ./...); \
	done

tidy:
	for mod in $(MODS); do \
		(cd $$mod && go mod tidy); \
	done

fmt:
	for mod in $(MODS); do \
		(cd $$mod && go fmt ./...); \
	done	

vet:
	for mod in $(MODS); do \
		(cd $$mod && go vet ./...); \
	done

graph:
	for mod in $(MODS); do \
		(cd $$mod && go mod graph); \
	done

env:
	for mod in $(MODS); do \
		(cd $$mod && go env); \
	done

########
# sqlc #
########
DB_MOD = ./apps/pkgs/db

.PHONY: sqlc-gen sqlc-gen sqlc-compile sqlc-verify sqlc-help

sqlc-gen:
	cd ${DB_MOD} && sqlc generate

sqlc-compile:
	cd ${DB_MOD} && sqlc compile

sqlc-verify:
	cd ${DB_MOD} && sqlc verify

sqlc-help:
	cd ${DB_MOD} && sqlc help