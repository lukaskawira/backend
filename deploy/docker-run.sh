#! /bin/bash
docker run -d -p 18000:8080 \
       -e HASURA_GRAPHQL_DATABASE_URL=postgres://postgres:root@host.docker.internal:5432/postgres \
       -e HASURA_GRAPHQL_ENABLE_CONSOLE=true \
       -e HASURA_GRAPHQL_DEV_MODE=true \
       hasura/graphql-engine:latest
