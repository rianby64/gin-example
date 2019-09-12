FROM gin-example-go

CMD [ "go", "test", "-v" ]

#CMD [ "dlv", "debug", "--headless", "--listen=:2345", "--log", "--api-version", "2", "--accept-multiclient", "github.com/rianby64/gin-example" ]