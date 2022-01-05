# lenslocked
Learning Go web development

Course: https://www.usegolang.com/


# Run the app
Start the database:
```
> docker-compose up -d
```

Run the webapp:
```
> go run .
```

(**Optional**) Connect to database with psql:
```
> docker container exec -it lenslocked_db_1 /usr/bin/psql -U user -d lenslocked
```

(**Optional**) Connect to the database in browser: 
```
http://localhost:8081/?pgsql=db&username=user&db=lenslocked&ns=public
```





# Notes
- Using the **[modd](https://github.com/cortesi/modd)** developer tool for dynamic reloading
- Using **[go-chi](https://github.com/go-chi/chi)** for HTTP routing and middleware