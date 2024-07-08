### How to use
1. Search for `INSERT` to find replacements that need to be made
2. Import types-template as [custom repository](https://github.com/Infinitare/types-template)
   1. replace `"github.com/Infinitare/types-template` with your own repository
   2. run `go get -u github.com/Infinitare/types-template` to remove the dependency
3. When you use, only import "github.com/pkg/errors"
   1. that's important for stack trace
   2. DO NOT USE "errors" package