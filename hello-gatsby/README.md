# hello-gatsby

## Tips

### Resolve gatsby develop error

```
ERROR #98123  WEBPACK

Generating SSR bundle failed

[BABEL] /Users/ryutah/projects/github.com/ryutah/sandbox/hello-gatsby/.cache/develop-static-entry.js: No "exports" main resolved in /Users/ryutah/projects/github.com/ryutah/sandbox/hello-gatsby/node_modules/@babel/helper-compilation-targets/package.json

File: .cache/develop-static-entry.js
```

1. remove package-lock.json

    ```console
    rm package-lock.json
    ```

1. exec gatsby clean

    ```console
    gatsby clean
    ```

1. remove node_modules

    ```console
    rm -rf node_modules
    ```

1. exec npm install

    ```console
    npm install
    ```
