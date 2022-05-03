#!/bin/bash

pip3 install pre-commit==2.18.1

pre-commit install
pre-commit install --hook-type commit-msg

yarn global add @graphql-inspector/graphql-loader@3.1.2 @graphql-inspector/git-loader@3.1.2 @graphql-inspector/diff-command@3.1.2 @graphql-inspector/cli@3.1.2 graphql@16.4.0
