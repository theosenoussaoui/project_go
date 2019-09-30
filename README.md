## VueJS - Golang - PostgreSQL

- [x] Valid Dockerfiles / Docker-Compose and Makefile.
- [x] User login, creation with JWT Middleware. Blocked at 3 tries with a custom blacklist model.
- [x] Fully commented functions.
- [ ] Tests on all functions.
- [ ] Voting system.
- [ ] Deleting and modifying users.


There are the commands to setup your environment efficiently :

```bash
make install
```
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