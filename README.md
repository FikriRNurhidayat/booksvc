# Book Service API

Book Service API is another implementation of DDD written in Golang.
This project was inspired by Marko Milojevic's Medium post series about
["Practical DDD in Golang"](https://levelup.gitconnected.com/practical-ddd-in-golang-domain-service-4418a1650274)

## How to run it?

I will only tell you how to run it inside Docker. Hence the following software is required:

- docker
- docker-compose

To start running the service, all you need to do is running the `docker-compose up -d` command.
But, **YOU'RE NOT DONE YET**. As you can see on the log, it fails to connect to the database.
That's because you don't have necessary user and privileges to do so.

To solve this issue, you can log into `mongosh` inside the docker by running:

```sh
docker-compose exec -it database mongosh
```

After you logged in, create new user and setup necessary privileges.

```javascript
use admin

db.createUser({
    user: "bookkeeper",
    pwd: "123456",
    roles: [
        { role: 'userAdminAnyDatabase', db: 'admin' },
        { role: 'readWriteAnyDatabase', db: 'admin' },
        { role: 'dbAdminAnyDatabase', db: 'admin' }
    ]
})
```

Then, go ahead restart the application by running

```sh
docker-compose restart app
docker-compose logs -f app
```

# License

This code is licensed under the [MIT License](./LICENSE).
