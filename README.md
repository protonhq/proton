# Proton [Work in progress]

## Usage

```sh
proton serve
```

## Configuration

### By toml

```toml
[database]
host="localhost"
port=6879
user="user"
password="password"
database="proton"
sslmode="disable"

[server]
port=8080

[keys]
jwt_signing_key="jwt_signing_key"
```

### By Environment Variables

| Name                 | Description       |
| -------------------- | ----------------- |
| PT_SERVER_PORT       | Server Port       |
| PT_DATABASE_HOST     | Database Host     |
| PT_DATABASE_PORT     | Database Port     |
| PT_DATABASE_USER     | Database User     |
| PT_DATABASE_PASSWORD | Database Password |
| PT_DATABASE_DATABASE | Database Name     |
| PT_DATABASE_SSLMODE  | Database SSL Mode |
