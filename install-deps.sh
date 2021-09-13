#!/bin/bash

pip3 install pre-commit==2.15.0

pre-commit install
pre-commit install --hook-type commit-msg

yarn global add @graphql-inspector/graphql-loader@2.9.0 @graphql-inspector/git-loader@2.9.0 @graphql-inspector/diff-command@2.9.0 @graphql-inspector/cli@2.9.0 graphql@15.5.3
