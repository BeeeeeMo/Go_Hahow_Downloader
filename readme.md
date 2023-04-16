
## Install dependency & Build
```bash
go tidy
go build main.go
```

## Usage
Support multiple course download, using `,` to separate course ID

e.g `main -T <token> -C <course_ID>,<course_ID>,<course_ID>`

```bash
main -T <token> -C <course_ID> -L <Concurrency Limit:Optional>
```
### or using config file
```bash
cat config.json
{
    "TOKEN": "Bearer <token>",
    "COURSE_ID": "<course_ID>",
    "ConcurrencyLimit": 2
}

main
```
