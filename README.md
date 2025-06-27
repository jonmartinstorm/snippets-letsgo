# snippets-letsgo

Følger boka til Alex Edwards: [https://lets-go.alexedwards.net/](https://lets-go.alexedwards.net/)


# Mysql

```
podman run --name mysql-test \
  -e MYSQL_ROOT_PASSWORD=my-secret-pw \
  -e MYSQL_DATABASE=testdb \
  -e MYSQL_USER=testuser \
  -e MYSQL_PASSWORD=testpass \
  -p 3306:3306 \
  -d docker.io/library/mysql:9.3.0
```

```
podman exec -it mysql-test mysql -u root -p
podman exec -it mysql-test mysql -D snippetbox -u web -p
```