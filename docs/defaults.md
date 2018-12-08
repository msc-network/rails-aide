# Defaults

The shipped defaults assume a rigid folder structure to your rails app, and are somewhat tailored to a pattern that I have come to be familiar and productive with.

## App metadata

```json
"metadata": {
    "name": "Rails Aide",
    "stub": "rails-aide",
"basedir": "/.rails-aide",
"version": "0.0.1",
"author": "CromonMS <http://github.com/CromonMS>",
"description": "A companion for building Rails assets",
"url": "https://github.com/msc-network/rails-aide - TODO: change name",
"repository": "https://github.com/msc-network/rails-aide",
"example": "Example: ./rails-aide -m Test -a=false -v=false -u=false -r=true"
},
```

## Rails folder structure

From the config file you can change the directory structure where the generated files are located by altering the following:

```json
"BaseDir": "/app",
"FrontendPath": "/javascript/frontend",
"AdminPagesPath": "/pages/Admin/",
"UserPagesPath": "/pages/User/",
"ComponentsPath": "/components/",
```

The defaults look like this:
<small>(MODELs represents generated folders, the passed in MODEL attribute pluralised):</small>

```shell
app
│   ...
│
└───javascript
    │
    └── frontend
        │
        └── components
        │   └── MODELs
        │
        └── pages
            ├── Admin
            │   └── MODELs
            │       └── MODEL
            │       └── MODEL
            │       └── MODEL
            └── User
                └── MODELs
```

## Vue Filename conventions

At the moment these are hard-coded into the src, they will be added to the config file at some point.
<small>(MODEL represents the passed in MODEL attribute):</small>

* **Admin**
    * Create AdminRecordFile - <small>MODEL</small>Admin (admin)
    * Create AdminCollectionFile - <small>MODEL</small>sAdmin (admin)
    * Create AdminNewRecordFile - New<small>MODEL</small> (admin)
    * Create AdminEditFile - Edit<small>MODEL</small>Admin (admin)

* **Components**
    * Create ComponentFormFile - <small>MODEL</small>Form (components)
    * Create ComponentRecordDetailFile - <small>MODEL</small>Detail (components)
    * Create ComponentListFile - <small>MODEL</small>sList (components)

* **User**
    * Create UserRecordFile - User<small>MODEL</small> (user)
    * Create UserCollectionFile - User<small>MODEL</small>s (user)
    * Create UserEditFile - EditUser<small>MODEL</small> (admin)
