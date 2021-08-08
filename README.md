# go-graphql
This is an implementation of GraphQL in GoLang

## Database
This project was setup using a postgres docker container
1) install Docker
2) run `docker run --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=somepassword -d postgres`
3) you can access the db directly with `docker exec -it some-postgres bash`
4) update the 'config.yml' file with the db values. The db values can be retrieved with running these commands will in the container
    `su postgres` this switches the current user to 'postgres'
    `psql` enables querying the db
    `\conninfo` this outputs the current connection info
5) you can exit the container with command `\q`
6) the container can be stopped with command `docker stop postgres`