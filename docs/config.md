# Config

How to use a local config file.

## Global configuration

Upon first run use a global config file will be copied into your home directory in the ```.rails-aide``` folder.

## Install local config

Run:

```shell
./mscmgmt-builder -c
```

This will copy a ```builder.config.json``` file to your local filesystem in the directory where you are using the tool.

The app will automatically pick up the file and use the settings in this file over the defaults.

## Config file precedence

The order in which this software reads config files is as follows, the internal config shouldn't normally be hit unless for some reason a global config cannot be written.

* Local config
* Global config
* Internal config

## Schema

```json
{
  "metadata": {
		"name": "Rails Aide",
		"stub": "rails-aide",
    "basedir": "/.rails-aide",
    "version": "0.0.1",
    "author": "CromonMS <http://github.com/CromonMS>",
    "description": "A companion for building Rails assets",
    "url": "https://github.com/msc-network/rails-aide - TODO"
  },
  "BaseDir": "/app",
  "FrontendPath": "/javascript/frontend",
  "AdminPagesPath": "/pages/Admin/",
  "UserPagesPath": "/pages/User/",
  "ComponentsPath": "/components/",
  "Admin": true,
  "Vue": true,
  "Rails": false
}
```

| Key                   | Value  | Comment                                                    |
|-----------------------|:------:|------------------------------------------------------------|
| **metadata**          | {}     | Contains information about the metadata of the application |
| metadata/name         | string | The name of this software                                  |
| metadata/stub         | string | A kebab-cased format of the name                           |
| metadata/basedir      | string | The directory which any config files are installed into    |
| metadata/version      | string | The current software version                               |
| metadata/author       | string | The software author(s)                                     |
| metadata/description  | string | A short description of this software                       |
| metadata/url          | string | The URL for this software                                  |
| BaseDir               | string | This is the base directory for your Rails app              |
| FrontendDir           | string | The path of your Vue Files                                 |
| AdminPagesPath        | string | The path of your Vue Admin Files                           |
| UserPagesPath         | string | The path of your Vue User Files                            |
| ComponentsPath        | string | The path of your Vue Components Files                      |
| Admin                 | bool   | Default setting is to generate Admin files                 |
| Vue                   | bool   | Default setting is to generate Vue files                   |
| Rails                 | bool   | Default setting is to not run Rails commands               |
