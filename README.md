# Personal finance

## Running the app locally

From project's root run:

```bash
docker-compose up
```

App should now be running on [http://localhost:5000/](http://localhost:5000/).

## Building production image

From project's root run:

```bash
docker build ./server --build-arg app_env=production
```
