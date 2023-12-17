#!/usr/bin/env bash

docker run -ti -e LOG_LEVEL=debug -v $(PWD):/usr/src/app docker.io/renovate/renovate:full renovate --platform=local
