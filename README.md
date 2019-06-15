# go-chi-wire-template

Golang server template with Chi(router) and Wire (google's DI)

# How to use:

I have supplied a very simple script in `main.go` that will replace my go module import definitions throughout the templates. This is because of how go handles imports. In just a few steps we will be up and running:

```bash
    # 1. Open go.mod and replace github.com/roger-king/go-chi-wire with your desired go module name
    module github.com/roger-king/testing-template - replace this module name

    # 2. Run the replacement script and pass your go module name:
    go run main.go github.com/new-user/this-is-example

    # 3. Remove replacement script, commit, and push
    git rm main.go
    git commit -m "chore: remove template replacement script"
    git push origin master

    # 4. Run setup command
    make setup

    # 5. Start server
    make start
```
