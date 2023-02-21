# Google Photos Takeout Handler

Project in go to unify and fix file dates from a google takeout of Google Photos data.~

## Usage

Update all variables to your own environments.

Update isCopy, using true to create a copy of the files and false to move them to new directory.

This will be improved in newer versions to give either an UI or a console input to provide variables.

```go
isCopy := true

dirToIterate := "<yourpath>\\Extracted"
rootToProcessTo := "<yourpath>\\Processed"
```

## Running

```bash
go run .
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## Tests

TBD

## License

[MIT](https://choosealicense.com/licenses/mit/)