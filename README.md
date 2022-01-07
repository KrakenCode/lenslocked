# lenslocked
Learning Go web development

Course: https://www.usegolang.com/

2022-01-07 - Waiting on author's completion of v2 of course before completing continuing 

# Run the app
Start the database:
```bash
docker-compose up -d
```

Run the webapp:
```bash
go run .
```

(**Optional**) Connect to database with psql:
```bash
docker container exec -it lenslocked_db_1 /usr/bin/psql -U postgres -d lenslocked
```

(**Optional**) Connect to the database in browser: 
```
http://localhost:8081/?pgsql=db&username=postgres&db=lenslocked&ns=public

username: postgres
password: dev
```


# Notes
- Using the **[modd](https://github.com/cortesi/modd)** developer tool for dynamic reloading
- Using **[go-chi](https://github.com/go-chi/chi)** for HTTP routing and middleware