# Mongodb ping

mongodb ping allow to check of a mongodb is alive. it s woll be use as init container inside infrastructure to check if dependencies are respected before startup a container.

## How to use :

You can specify the following env var : 

MONGODB_AUTH ( boolean) add the authentication with the var MONGODB_USER and MONGODB_PASSWORD 
MONGODB_PORT and MONGODB_HOST represent the server which respond or not


exemple with auth : 
```
docker run -e MONGODB_AUTH="true" -e  MONGODB_USER="root" -e MONGODB_PASSWORD="rootPassXXX" -e MONGODB_PORT="27017" -e MONGODB_HOST="172.17.0.2" -it mongodb-ping
```

exemple without auth :
```
docker run  -e MONGODB_PORT="27017" -e MONGODB_HOST="172.17.0.2" -it mongodb-ping
```

If connexion is success the following message appear : "Mongodb respond OK"


