abigen:
	@echo "  >  generate go code for abi"
	@ls internal/pkg/abi/*.json |\
 	 xargs -I {} basename {} .json |\
 	 xargs -I {} abigen --pkg abi --out ./internal/pkg/abi/{}.go  --type {} --abi ./internal/pkg/abi/{}.json

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
