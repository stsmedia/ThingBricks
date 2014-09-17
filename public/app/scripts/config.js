"use strict";

angular.module("config", [])

    .constant("ENV", {
        "name": "dev",
        "apiEndpoint": "http://localhost:9000"
    })

;