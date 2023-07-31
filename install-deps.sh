#!/bin/bash

pipx install pre-commit==3.3.2

pre-commit install

npm install -g @graphql-inspector/graphql-loader@3.3.0 @graphql-inspector/git-loader@3.3.0 @graphql-inspector/diff-command@3.3.0 @graphql-inspector/cli@3.3.0 graphql@16.5.0
