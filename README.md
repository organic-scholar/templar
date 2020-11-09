# Templar

![templar logo](./images/logo-white.png)


Templar clones project from the remote git repository and substitute placeholders in source files.

## Installtion

Download templar binary for your platform from the release page.

## Usage

In your project template repository commit a template.json file e.g

```
{
    "parameters": {
        "name": "myApp",
        "description": "This is a sample app"
    },
    "files": [
        "package.json", "src/main.js"
    ]
}

```

In the above file, parameters are rendered into the mentioned source files using the mustache template engine. It will also prompt the user to override parameters.


You can use templar as follow

```
templar github:user/repo 
templar bitbucket:user/repo
```

Templar clones your repository in current working directory but can also specify another one.
```
templar github:user/repo my-app
```

By default templar uses ssh to clone if you wan to use https
```
templar github:user/repo --use=https
```


## License

[MIT](LICENSE.md)
