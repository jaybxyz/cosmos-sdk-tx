#!/usr/bin/make -f

###############################################################################
###                                Build                                    ###
###############################################################################


###############################################################################
###                               DevOps                                    ###
###############################################################################

local-node: 
	@echo "Running single local network..."
	./scripts/localnet.sh

.PHONY: local-node