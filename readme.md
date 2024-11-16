
# Go Rest Server with MySql
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
a sample go rest server for learning Go and MySql


## Authors

- [@JeraldVictor](https://www.github.com/JeraldVictor)


## Run Locally

Clone the project

```bash
  git clone https://github.com/JeraldVictor/go-rest.git
```

Go to the project directory

```bash
  cd go-rest
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  make run
```


## Running Tests

To run tests, run the following command

```bash
  make test
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`SERVER_HOST`
`SERVER_PORT`
`DB_HOST`
`DB_PORT`
`DB_USER`
`DB_PASSWORD`
`DB_NAME`
