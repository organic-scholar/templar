# Templar

![templar logo](./images/logo.png)


Templar clones project from the remote git repository and substitute placeholders in source files.

## Installtion

Download templar binary for your platform from the release page.

## Usage

In your project template repository create a template.json file e.g

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

In the above file, parameters are rendered into the mentioned source files using the mustache template engine.

It will also prompt the user to override parameters.

```
templar github:user/repo .
templar github:user/repo my-app
templar bitbucket:user/repo
```
## License

[MIT](LICENSE.md).
