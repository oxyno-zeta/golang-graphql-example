#!/bin/bash

pip3 install pre-commit==2.20.0

pre-commit install
pre-commit install --hook-type commit-msg

yarn global add @graphql-inspector/graphql-loader@3.3.0 @graphql-inspector/git-loader@3.3.0 @graphql-inspector/diff-command@3.3.0 @graphql-inspector/cli@3.3.0 graphql@16.5.0
