# mscmgmt-builder [![Build Status](https://travis-ci.org/msc-network/mscmgmt-builder.svg?branch=master)](https://travis-ci.org/msc-network/mscmgmt-builder)

A small command line utility for automating file creation and other helpers in a Rails app with Vue on the front end.

Currently specific to the mscmgmt webapp.

## Usage

For information about flags and what they do see the [docs](docs/index.md)

### Quick Start

Run: ```./mscmgmt-builder -help```

## Development

The latest binary for Linux is included in the repo, to make your own:

Clone repo: ```git clone github.com/msc-network/mscmgmt-builder.git```

Build binary: ```go build```

Run tool: ```./mscmgmt-builder -help```

## TODO

* Write generated files to directory rather than where the binary is run.
* Piggy back running rails commands - involves storing model attributes and generate the rails command
* ~~Better config using a local file to overwrite the default if present~~
* More involved templates including imports, mixins etc.
* Make it more generic for other Rails apps. Possibly rename to 'railsgo'

## Credits

 * [Eugene Rasini](https://github.com/cromonms)

## License

The MIT License (MIT) - see [`LICENSE.md`](LICENSE) for more details
