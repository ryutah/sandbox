# logging-and-alerting

Examples to output operations logging and operations error reporting.

## Logging sample and Reported Error

### Logging

![Logging](https://user-images.githubusercontent.com/6662577/81158305-ccc34c80-8fc2-11ea-8483-019652fe69f4.png)

### Error Details

![error details](https://user-images.githubusercontent.com/6662577/81157741-40188e80-8fc2-11ea-8eb1-bc8a2d62ca7d.png)

### Error Log Details

![error log details](https://user-images.githubusercontent.com/6662577/81158990-7acef680-8fc3-11ea-8727-03d55fdc1ff9.png)

## Deploy

```console
gcloud app deploy .
```

## Usage

### Exec with no error

```console
curl 'https://laa-dot-[PROJECT_ID].appspot.com' \
  -X POST \
  -d '{"message": "[SOME_MESSAGE]", "raise_error": false}'
```

### Exec with error

```console
curl 'https://laa-dot-[PROJECT_ID].appspot.com' \
  -X POST \
  -d '{"message": "[SOME_MESSAGE]", "raise_error": true}'
```
