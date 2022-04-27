# SMARTPARKING BACKEND

* Compile

```sh
make build
```

* Generate configuration files

```sh
./smartparking config --path=config.yml
```

* Initialize the database

```sh
./smartparking migrate --config=config.yml
```

* Start the service

```sh
./smartparking start --config=config.yml
```

#### Global help

```sh
smartparking applications

Usage:
  smartparking [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      initial config generation
  help        Help about any command
  migrate     migrate models to database
  start       start server

Flags:
      --config string   path to config file (default "/home/yeldar/.smartparking/config.yml")
  -h, --help            help for smartparking

Use "smartparking [command] --help" for more information about a command.
```

#### Initial config generation

```sh
Usage:
  ./smartparking config [flags]

Flags:
      --cache_host string                cache host (default "localhost")
      --cache_port string                cache port (default "11211")
      --cache_ttl duration               cache host (default 10m0s)
      --db_host string                   database host (default "localhost")
      --db_name string                   database name (default "smartparking")
      --db_password string               database password (default "000730")
      --db_port string                   database port (default "5432")
      --db_sslmode string                database sslmode (default "disable")
      --db_username string               database username (default "yeldar")
      --email_password string            email password (default "GrKrsLTy5rzGqkFC8ZhX")
      --email_sender string              email sender (default "kuanyshev_eldar@mail.ru")
      --email_smtp_host string           email smtp host (default "smtp.mail.ru")
      --email_smtp_port string           email smtp port (default "587")
      --hash_salt string                 hash salt (default "my_hash_salt")
  -h, --help                             help for config
      --jwt_access_token_ttl duration    jwt access token ttl (default 15m0s)
      --jwt_refresh_token_ttl duration   jwt refresh token ttl (default 24h0m0s)
      --jwt_secret string                jwt secret (default "my_secret_key")
      --path string                      generate config file to (default .smartparking/config.yml) (default "/home/yeldar/.smartparking/config.yml")
      --recognizer_token string          recognizer token (default "f548497ecd62d281f5ee97e497dd67426bb8a586")
      --recognizer_url string            recognizer url (default "https://api.platerecognizer.com")
      --web_certfile string              web cert file
      --web_file_storage string          web file storage (default "/home/yeldar/.smartparking/images")
      --web_host string                  web host (default "localhost")
      --web_keyfile string               web key file
      --web_port string                  web port (default "8080")
      --web_tlsenable                    web tls enable

Global Flags:
      --config string   path to config file (default "/home/yeldar/.smartparking/config.yml")
```

#### Initial migration

```sh
migrate models to database

Usage:
  smartparking migrate [flags]

Flags:
  -h, --help   help for migrate

Global Flags:
      --config string   path to config file (default "/home/yeldar/.smartparking/config.yml")
```