# Config

How to use a local config file.

## Global configuration

Upon first run use a global config file will be copied into your home directory in the ```.rails-aide``` folder. You can customise this file just like a local version.

## Install local config

Run:

```shell
./mscmgmt-builder -c
```

This will copy a ```builder.config.json``` file to your local filesystem in the directory where you are using the tool.

The utility will automatically pick up the file and use the settings in this file over the defaults.

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
    "url": "https://github.com/msc-network/rails-aide",
    "repository": "https://github.com/msc-network/rails-aide"
  },
  "filenames": {
    "admin": {
      "AdminRecordFile": "!Admin",
      "AdminCollectionFile": "=Admin",
      "AdminNewRecordFile": "New!",
      "AdminEditFile": "Edit!Admin"
    },
    "components": {
      "ComponentFormFile": "!Form",
      "ComponentRecordDetailFile": "!Detail",
      "ComponentListFile": "=List"
    },
    "user": {
      "UserRecordFile": "User!",
      "UserCollectionFile": "User=",
      "UserEditFile": "EditUser!"
    }
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

### Metadata settings

Contains information about the metadata of the application - this doesn't need changing and is included for information purposes only.

| Key                   | Type   | Comment                                                    |
|-----------------------|:------:|------------------------------------------------------------|
| name                  | string | The name of this software                                  |
| stub                  | string | A kebab-cased format of the name                           |
| basedir               | string | The directory which any config files are installed into    |
| version               | string | The current software version                               |
| author                | string | The software author(s)                                     |
| description           | string | A short description of this software                       |
| url                   | string | The hoempage URL for this software                         |
| repository            | string | The respository URL for this software                      |

### Filename settings

Contains settings specific to file creation, here you can set your own defaults for common types of files.

| Key                   | Type   | Comment                                                          |
|-----------------------|:------:|------------------------------------------------------------------|
| AdminRecordFile             | string | An Admin file to display a single record                   |
| AdminCollectionFile         | string | An Admin file to display a collection of records           |
| AdminNewRecordFile          | string | An Admin file for creating a new record                    |
| AdminEditFile               | string | An Admin file for editing an existing record               |
| ComponentFormFile           | string | A reusable form component                                  |
| ComponentRecordDetailFile   | string | A reusable record detail component                         |
| ComponentListFile           | string | A reusable list of records component                       |
| UserRecordFile              | string | A User file to display a single record                     |
| UserCollectionFile          | string | A User file to display a collection of records             |
| UserEditFile                | string | A User file for editing an existing record                 |

##### Inflection

This utility uses StringReplacement to add either a singular or plural model name to each file. To use the singular model name (e.g. Robot) use a **!** in the place you require the model name to be replaced, likewise for a plural model name (e.g. Robots) use a **=**

##### Notes 

As you can see there is no UserNewFile, this is because these can be specific to an app, and not all models require a User to be able to create records.

##### Customising your own file structure

If you add a new Key and value for a file you'd like created when you run this utility, please be sure to follow the inflection rules above.

### General settings

Contains information about the structure of your rails app, and default flag settings

| Key                   | Type   | Comment                                                    |
|-----------------------|:------:|------------------------------------------------------------|
| BaseDir               | string | The base of your Rails app directory                       |
| FrontendDir           | string | The path of your Vue Files                                 |
| AdminPagesPath        | string | The path of your Vue Admin Files                           |
| UserPagesPath         | string | The path of your Vue User Files                            |
| ComponentsPath        | string | The path of your Vue Components Files                      |
| Admin                 | bool   | Default setting is to generate Admin files                 |
| Vue                   | bool   | Default setting is to generate Vue files                   |
| Rails                 | bool   | Default setting is to **NOT** run Rails commands           |
