# Should install swagger first (https://github.com/go-swagger/go-swagger)
swagger generate spec /b ../ /o swagger.json
# Should install swagger-markdown from npm first (https://www.npmjs.com/package/swagger-markdown)
# swagger-markdown -i swagger.json -o swagger.md