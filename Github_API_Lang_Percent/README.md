## Github account language percentage calculator

This program allow to display the percentage of most used languages in an account.
Program only calculates repos that are not forked ( Can be changed by removed this condition `!repolist[i].IsFork` )

### Usage

- Build the excecutable ( below command is for linux )

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go
```

- zip the excecutable

```sh
zip deployment.zip main
```

- Create a lambda function and provide API gateway as trigger
- Upload the zip
- `<URL><function name>?username=<github username>`
