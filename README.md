# Wit - The Weed Information Tracker

Wit aims to help cannabis patients and users to manage and monitor their cannabis consumption and inventory.
It provides a CLI with a syntax which may seem familiar to users of git. :)

## Usage examples

### Not implemented yet

- `wit init [path]`: (Example: `wit init .`) Create a new wit repository at the given `path`
- `wit checkout [name]`: (Example: `wit checkout hindu-kush`) Switch to the stash with the given `name`
- `wit add <amount as floating point number>`: (Example: `wit add 12.0`) Add the given amount to the current selected stash. Note: Similar to git, after `add`-ing something, you also need to `commit` it and give a message, read on
- `wit commit -m "<commit-message>"`: (Example: `wit commit -m "add 30 day delivery for month may") Commits the changes to the current stash

## Planned features

- Tracking of consumption
- Inventory trend
- Graphs, graphs, more graphs

## Building

Run `nx build wit` to build the project. The build artifacts will be stored in the `dist/apps` directory. Use the `--prod` flag for a production build. The built binary can be executed, e.g. via:

```shell
$ ./dist/apps/wit
Wit aims to help cannabis patients and users to manage and monitor their cannabis consumption and inventory.
It provides a TUI with a syntax which may seem familiar to users of git. :)

Usage:
  wit [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initializes a new wit repository at the given path

Flags:
      --config string   config file (default is $HOME/.wit.yaml)
  -h, --help            help for wit
  -t, --toggle          Help message for toggle
  -v, --version         version for wit

Use "wit [command] --help" for more information about a command.
```

## Running

Run `nx serve wit` to run the application. Note: You won't be able to pass program arguments to the application. For that, first build the application and afterwards run the built binary located in the `dist/apps` directory.

## Running unit tests

Run `nx test wit` to execute the unit tests.

Run `nx affected:test` to execute the unit tests affected by a change.

## Building Docker image locally and pushing to Dockerhub

Run `nx run wit:docker` to build the docker image and push it to the Dockerhub registry. Note: you must have executed a `$ docker login` in the shell you are executing the command.

## Building Docker image on GitHub and pushing to GitHub Container Registry

These steps are handled through GitHub Workflows / Actions, see: [.github/workflows/docker.yml](./github/workflows/docker.yml)

## Semantic Version

Run `nx run wit:version` to create a new semantic git version tag and update the [CHANGELOG.md](./apps/wit/CHANGELOG.md). Note: you do need to manually push the changes via `$ git push`, respectively `$ git push --tags`.

## Understand your workspace

Run `nx graph` to see a diagram of the dependencies of your projects.
