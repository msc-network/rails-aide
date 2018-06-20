# Config file

How to use a local config file.

## Install local config

Run:

```shell
./mscmgmt-builder -c
```

This will copy a ```builder.config.example.json``` file to your local filesystem in the directory where you are using the tool, rename this to ```builder.config.json```. 

The app will automatically pick up the file and use the settings in this file over the defaults.

## Schema

```json
{
  "BaseDir": "/app",
  "FrontendDir": "/javascript/frontend",
  "Admin": true,
  "Vue": true,
  "Rails": false
}
```

|Key | Value | Comment |
|----|-------|---------|
| BaseDir | "/app" | This is the base directory for your Rails app |
| FrontendDir | "/javascript/frontend" | The location of your Vue Files |
| Admin | true | Default setting is to generate Admin files |
| Vue | true | Default setting is to generate Vue files |
