## VueJS - Golang - PostgreSQL

- [x] Valid Dockerfiles / Docker-Compose and Makefile.
- [x] User login, creation with JWT Middleware. Blocked at 3 tries with a custom blacklist model.
- [x] Fully commented functions.
- [ ] Tests on all functions.
- [ ] Voting system.
- [ ] DELETE and PUT methods for users.

There are the commands to setup your environment efficiently :

```bash
make install
```
And if you didn't uncomment the VueJS setup in the `docker-compose.yml`, you can execute the following command (requires yarn to be installed) to run your front on [localhost:8081](http://localhost:8081) : 

```bash
cd front && yarn serve
```

The VueJS container in the `docker-compose.yml` is commented because Docker's performance on MacOS sucks ass. Uncomment it if you want.

You can check the backend condition on [localhost:8080/ping](http://localhost:8080/ping) (it's running the Gonic main example).
 

### If you want to install a Golang extension : 

Install Golang locally first and install your extension using its own documentation.
After that, rebuild your backend using :

```bash
make up
```

### If you changed your code in the back folder (app.go) :

You have to rebuild your code (until and if I find a way to get some kind of hot reloading to work) by executing the command : 

```bash
make up
```
