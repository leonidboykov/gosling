---
title: Gosling Documentation
---
Gosling is a lightweight HTTP redirects generator.

# Installing Gosling

There are multiple ways to get Gosling.

#### macOS

For macOS you may use [homebrew](https://brew.sh) to download `gosling` from the command line 

    brew install leonidboykov/tap/gosling

#### Windows

For Windows you may use [scoop](https://scoop.sh) to download `gosling` from the command line

    scoop bucket add gosling https://github.com/leonidboykov/scoop-bucket.git
    scoop install gosling

#### Development version

You may also get the latest version from `master` branch with

    go get -u github.com/leonidboykov/cmd/gosling

# Getting Started

Create `redirects.json` with the following JSON structure

```json
{
  "/": "https://root_redirect",
  "project_1": "https://project_1",
  "project_2": "https://project_2"
}
```

Then simply run

    gosling

This command will produce a `public` folder with HTML redirects files

    .
    ├── public
    │   ├── index.html
    │   ├── project_1
    │   │   └── index.html
    │   └── project_2
    │       └── index.html
    └── redirects.json

# Available flags

You may also override default setting with the following flags

    gosling -o build

will write HTML files to `build` folder instead of `public`

    gosling -r my_redirects.json

will use `my_redirects.json` file instead of `redirects.json`

    gosling -v

will print current verson and exit

# Using with GitLab CI

There is also Docker image available on [Docker Hub](https://hub.docker.com/r/leonidboykov/gosling). You can use the following `.gitlab-ci.yml` config file to build redirects with `gosling`

```yaml
image: leonidboykov/gosling:latest
pages:
  script:
    - gosling
  artifacts:
    paths:
      - public
  only:
    - master
```

# License

`gosling` is free software licensed under the [MIT](LICENSE) license.
