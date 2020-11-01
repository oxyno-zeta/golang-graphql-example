# golang-graphql-example

For my personal use :)

The aim of this project is to provide a fully working example with a based project already working with copy/paste features.

## Convention

- 1 folder is related to one project (backend, ui, ...) in a specific language
- the git commit convention is the angular one (see [here](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines))
- Editorconfig is used to keep file content in a uniform way

## Install

This project is using the python software called `pre-commit`. This is used to install and have git pre-commit hooks.

Those ones are here to validate code, lint projects, lint and validate GraphQL, ...

Moreover, some tools are used in the backend project. These tools are using NodeJS and Yarn for package installation.

Just run the script called `./install-deps.sh` in order to install needed dependencies.

## How to use ?

The project is using VSCode Workspaces. Just open the `workspace.code-workspace` in order to have the right integration for sub projects.

## Release

In order to release, the project is using `semantic-release` in order to generate a release directly with the Changelog (git based).
