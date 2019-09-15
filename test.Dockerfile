FROM gin-example-go

CMD [ "go", "test", "-v", "github.com/rianby64/gin-example", "github.com/rianby64/gin-example/lib/client1c" ]

#CMD [ "dlv", "debug", "--headless", "--listen=:2345", "--log", "--api-version", "2", "--accept-multiclient", "github.com/rianby64/gin-example" ]