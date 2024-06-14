# Installing Redis
1. Using a package manager:

**On Ubuntu**:
```sh
sudo apt update
sudo apt install redis-server
```

**On macOS**:
```sh
brew install redis
```

2. Using Docker:
If you prefer using Docker, you can run the following command to start a Redis container:
```sh
docker run --name redis -d -p 6379:6379 redis
```

### Configuring Redis
1. Modify the Redis configuration file if needed:

The default configuration file is usually located at /etc/redis/redis.conf on Ubuntu or /usr/local/etc/redis.conf on macOS.

You might want to modify the configuration to set a password or change other settings. For example, to set a password, you can uncomment and set the requirepass directive:

```sh
requirepass yourpassword
```

2. Restart the Redis service:
**On Ubuntu**:
```sh
sudo systemctl restart redis-server
```

**On macOS**:
```sh
brew services restart redis
```

### Testing Redis
1. Connect to Redis CLI:
```sh
redis-cli
```

2. Run some basic commands to test:
```sh
ping
set mykey myvalue
get mykey
```
You should see responses like PONG, OK, and myvalue.

### Configuring the Go Application to Use Redis
Ensure your `.env` file has the correct Redis configuration:

```makefile
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=yourpassword   # Only if you set a password in the configuration
```

### Example .env file
```plaintext
DB_DSN=postgres://username:password@localhost:5432/yourdb
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=yourpassword
```