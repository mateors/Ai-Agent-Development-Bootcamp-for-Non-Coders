# pgadmin4

### Step 1: Locate the pgAdmin configuration and database

```
root@MSI:~# ls -la /var/lib/pgadmin4/pgadmin4.db
-rw------- 1 root root 2281472 Oct 28 19:10 /var/lib/pgadmin4/pgadmin4.db
```

> find /opt/mateors/lxroot/pgadmin -name "pgadmin4.db"


### Step 2: Open the pgAdmin internal database

> `sqlite3 /var/lib/pgadmin4/pgadmin4.db`

sqlite>

### Step 3: Inspect the user table
> SELECT id, email, locked, login_attempts FROM user WHERE email='pgdadmin@lxroot.com';

> `update user set locked=0,login_attempts=0 where id=1;`
.quit