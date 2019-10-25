# Go & VueJS voting application

Note exclusively for the Golang project (will be removed in the future) :

-   [x] Valid Dockerfiles / Docker-Compose and Makefile.
-   [x] User login, creation with JWT Middleware. Blocked at 3 tries with a custom blacklist model.
-   [x] Fully commented functions.
-   [x] Voting system.
-   [x] DELETE and PUT methods for users.
-   [ ] Tests on all functions.

There are the commands to setup your environment efficiently :

```bash
make install
```

After that, you can see and test your fully functional API on `localhost:8080`.

You can also get our `Postman` Collection [here](https://github.com/theosenoussaoui/project_go/blob/dev/back/postman/Golang Soutenance.postman_collection.json). Be careful as you have to initialize all your variables in Postman.