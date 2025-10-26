# n8n Developer Mode

```
apt-get install build-essential -y
```

```
cd /var/www/vhosts
```

```
mkdir codyclouds.com
```

```
cd codyclouds.com
```

```
git clone --depth 1 https://github.com/n8n-io/n8n
```

```
cd n8n
```

```
npm install -g pnpm
```

```
pnpm install
```

shuraiajenin@gmail.com

memory swap:
```
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

```
pnpm run build
```

```
 Tasks:    41 successful, 41 total
Cached:    0 cached, 41 total
  Time:    2m4.663s
  ```

```
pnpm run start
```

## Option-2

curl -fsSL https://deb.nodesource.com/setup_22.x | sudo -E bash -

> `npm install -g n8n`

> `mkdir -p /home/mostain/n8n-data`

> `chmod 700 /home/mostain/n8n-data`

> `chown mostain:mostain /home/mostain/n8n-data`


> nano `/etc/systemd/system/n8n.service`

```
[Unit]
Description=n8n workflow automation
After=network.target

[Service]
Type=simple
User=mostain
Environment=DB_TYPE=sqlite
Environment=DB_SQLITE_DATABASE=/home/mostain/n8n-data/database.sqlite
Environment=N8N_BASIC_AUTH_ACTIVE=true
Environment=N8N_BASIC_AUTH_USER=admin
Environment=N8N_BASIC_AUTH_PASSWORD=securepassword
Environment=N8N_ENFORCE_SETTINGS_FILE_PERMISSIONS=true
#Environment=DB_TYPE=postgresdb
#Environment=DB_POSTGRESDB_HOST=localhost
#Environment=DB_POSTGRESDB_PORT=5432
#Environment=DB_POSTGRESDB_DATABASE=n8ndb
#Environment=DB_POSTGRESDB_USER=n8nuser
#Environment=DB_POSTGRESDB_PASSWORD=n8npass
#Environment=N8N_BASIC_AUTH_ACTIVE=true
#Environment=N8N_BASIC_AUTH_USER=admin
#Environment=N8N_BASIC_AUTH_PASSWORD=securepassword
ExecStart=/usr/bin/n8n
Restart=always
RestartSec=10
WorkingDirectory=/home/mostain

[Install]
WantedBy=multi-user.target

```

```
systemctl daemon-reexec
systemctl daemon-reload
systemctl enable n8n
systemctl start n8n
```

```
chmod 600 /home/mostain/.n8n/config
```



linux systemd config:
> `/etc/systemd/system/workflow.local.service`

```

[Unit]
Description=workflow.local
After=network.target
[Service]
Type=simple
User=workflowlocal
Group=mateors
Restart=on-failure
RestartSec=5s
WorkingDirectory=/var/www/vhosts/workflow.local/public_html/n8n
ExecStartPre=/usr/bin/pnpm install
ExecStartPre=/usr/bin/pnpm build
ExecStart=/usr/bin/pnpm start
Environment="N8N_PORT=8107"
Environment="N8N_SECURE_COOKIE=false"
Environment="N8N_SSL_CERT=/etc/ssl/workflow.local/workflow.local.pem"
Environment="N8N_SSL_KEY=/etc/ssl/workflow.local/workflow.local.key.pem"
#Environment="WEBHOOK_URL=https://workflow.local"
[Install]
WantedBy=multi-user.target
```


nginx config:
> `/etc/nginx/sites-available/workflow.local`

```
server {
        listen 443 ssl;
        server_name workflow.local;
        access_log /var/www/vhosts/workflow.local/log/access.log;
        error_log /var/www/vhosts/workflow.local/log/error.log;
        add_header X-Frame-Options "SAMEORIGIN";
        add_header X-Content-Type-Options "nosniff";
        add_header X-XSS-Protection "1; mode=block";
        location / {
                #proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                #proxy_set_header X-Forwarded-Proto $scheme;
                #proxy_set_header X-Real-IP $remote_addr;
                #proxy_set_header Host $host;
                #proxy_pass  http://127.0.0.1:8107;
                #add_header Content-Type text/html;
                proxy_pass http://localhost:8107;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "upgrade";

        }
        location = /favicon.ico { access_log off; log_not_found off; }
        location = /robots.txt  { access_log off; log_not_found off; }

        # managed by lxroot
        ssl_certificate /etc/ssl/workflow.local/workflow.local.pem;
        ssl_certificate_key /etc/ssl/workflow.local/workflow.local.key.pem;
        #include /etc/letsencrypt/options-ssl-nginx.conf;
        #ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;
        add_header Strict-Transport-Security "max-age=31536000" always;
        #ssl_trusted_certificate /etc/letsencrypt/live/mostain.org/chain.pem;
        ssl_stapling off;
        ssl_stapling_verify off;
        add_header X-Powered-By LXROOT;
}

server {
        listen 80;
        # managed by lxroot
        if ($host = workflow.local) {
                return 301 https://$host$request_uri;
        }

        server_name workflow.local;
        add_header X-Frame-Options "SAMEORIGIN";
        add_header X-Content-Type-Options "nosniff";
        add_header X-XSS-Protection "1; mode=block";
        location / {
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header X-Forwarded-Proto $scheme;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header Host $host;
                proxy_pass  http://127.0.0.1:8107;
                add_header Content-Type text/html;
        }
        location = /favicon.ico { access_log off; log_not_found off; }
        location = /robots.txt  { access_log off; log_not_found off; }
        add_header X-Powered-By LXROOT;
}
```