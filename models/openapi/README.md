# Go API Server for openapi

NRF Bootstrapping.  
© 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.2.0-alpha.1
- Build date: 2024-07-17T12:50:24.865922+05:30[Asia/Calcutta]
- Generator version: 7.7.0

### Running the server

To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once the image is built, just run
```
docker run --rm -it openapi
```

### Known Issue

Endpoints sharing a common path may result in issues. For example, `/v2/pet/findByTags` and `/v2/pet/:petId` will result in an issue with the Gin framework. For more information about this known limitation, please refer to [gin-gonic/gin#388](https://github.com/gin-gonic/gin/issues/388) for more information.

A workaround is to manually update the path and handler. Please refer to [gin-gonic/gin/issues/205#issuecomment-296155497](https://github.com/gin-gonic/gin/issues/205#issuecomment-296155497) for more information.
