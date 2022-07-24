# Crew interview test

Test instructions can be found [here](https://github.com/crewdotwork/backend-challenge).

## Application

The application has two commands: `server` and `seed`.

### `server` command

Launch the API server which exposes three endpoints:

- `GET /health`: Health check.
- `POST /v1/talents`: Add a talent.
- `GET /v1/talents`: List talents.

The command requires `--config <config file path>` flag which points to the config file. See [default config file](config/default.yaml) 
for example.

Note that for security reason, `mongo.uri` is not set via config file but via environment variable `MONGO_URI`. This 
will automatically override `mongo.uri` field in config file. The same kind of override can be applied to any other field.

### `seed` command

Seed the database with data from Crew API. The command requires four flags:

- `--mongo-uri`: URI of the mongo database to seed.
- `--database`: Name of the database to seed.
- `--collection`: Name of the collection to seed.
- `--crew-uri`: Crew API's URI to fetch data from.

## How to run

### Local server

For local run, docker-compose is used. It will launch a server instance along with a mongodb instance. 

Local API server is exposed at `localhost:8080` and mongo is exposed at `localhost:27017`.

```shell
make up   # docker compose up
make down # docker compose down
```

### Seed local database

Seed the local database with data from Crew API.

### Postman

There is a [Postman collection](postman/crew_interview.postman_collection.json) that you can import to Postman and test 
the API server.

```shell
make seed
```

## Improvements

Due to time constraint, there are multiple things that I could not do but am still aware of their necessities in a 
professional context. I would ensure all the below, if I have time and resource to do so. 

### Code

- Not enough tests are written, I would mock all the interfaces and unit test all layers/modules.
- Swagger file/endpoint is not generated. 
- Pagination of list talents endpoint. I would do key set pagination instead of offset pagination in Crew API.
- More logs and log configs.
- Integration and e2e tests.
- etc.

### Deployment

Infrastructure provisioning takes time and resource, especially in an IoC context. Here are a list of what I considered, 
but did not do:

- Deploy on a ready-to-use Kubernetes cluster: [Okteto](https://cloud.okteto.com/) is a free option. [Here](https://github.com/elliot-token/api/tree/main/deploy/kustomize) is an example of how I deploy my app on k8s.
- Deploy on AWS lambda with Terraform. [Here](https://github.com/Ventilateur/insider-acquisition-report/tree/main/aws/sec4) is another example of how I did the same in another project.
- Deploy locally with [localstack](https://localstack.cloud/) to mimic AWS EKS or Lambda. Same as above, a set of Terraform must be produced.

I am happy to discuss the options and how exactly would I proceed for each of them. However, for such a short test, 
please allow me to omit that part, since it's very time-consuming.