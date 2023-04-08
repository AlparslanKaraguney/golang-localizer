# Localizer

This is a simple example of how to use the Localizer library.
First install the library:

```bash
$ go install golang.org/x/text/cmd/gotext@latest
```
Then edit your tanslations.go file and add your strings:

```go
package translations

//go:generate gotext -srclang=en-GB update -out=catalog.go -lang=en-GB,tr-TR ../../cmd/main.go
```

To generate the localization files, run the following command:

```bash
$ go generate ./internal/translations/translations.go
```

Some files will be created in the `internal/translations/locals` folder. You must copy them to same folder. Do not edit `out.gotext.json` files just copy them and create a new file named `messages.gotext.json` and edit your strings for example:

```bash
$ cp internal/translations/locales/tr-TR/out.gotext.json internal/translations/locales/tr-TR/messages.gotext.json
```

Each time you add a new string to the code, you must run the following command to apply the changes to file `out.gotext.json`:

```bash
$ go generate ./internal/translations/translations.go
```