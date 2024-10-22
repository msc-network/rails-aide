# rails-aide [![Build Status](https://travis-ci.org/msc-network/rails-aide.svg?branch=master)](https://travis-ci.org/msc-network/rails-aide)

A small command line utility for automating file creation and other helpers in a Rails app with Vue on the front end.

Currently specific to the mscmgmt webapp.

For a more detailed look go read the [docs](docs/index.md)

## Install

First download the binary from [here](https://github.com/msc-network-rails-aide)

Then copy to your usr/local ```sudo cp rails-aide /usr/local/bin```

```rails-aide``` is now available globally

## Usage

For information about flags and what they do see [usage.md](docs/usage.md)

### Quick Start

On first run a global config file will be created in ~/.rails-aide

Run: ```./rails-aide -help```

## Development

The latest binary for Linux is included in the repo, to make your own:

Clone repo: ```git clone github.com/msc-network/rails-aide.git```

Build binary: ```go build```

Run tool: ```./rails-aide -help```

## TODO

* ~~Write generated files to directory rather than where the binary is run~~
* Piggy back running rails commands - involves storing model attributes and generate the rails command
* ~~Better config using a local file to overwrite the default if present~~
* ~~Allow use of a global config too~~
* More involved templates including imports, mixins etc.
* ~~Rename to 'rails-aide'~~
* ~~Make generated filenames customisable via config~~
* Create log and write to basedir

## Credits

 * [Eugene Rasini](https://github.com/cromonms)

## License

The MIT License (MIT) - see [`LICENSE.md`](LICENSE) for more details
