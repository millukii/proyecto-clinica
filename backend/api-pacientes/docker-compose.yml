version: '3'
services:
    mongo:
        image: mongo:3.2
    example:
        build: .
        volumes:
            - .:/go/src/api-pacientes
        links:
            - mongo