#!/bin/bash
basePath=$(cd `dirname $0`; pwd)
templatesPath=${basePath}/templates
datasetsPath=${basePath}/datasets
cd $GOPATH/pkg/mod/github.com/gobuffalo/packr@v1.30.1/packr
go run main.go -i ${templatesPath}
go run main.go -i ${datasetsPath}

templatesPackrFile=$(cd ${templatesPath}; ls | grep "\-packr.go")
datasetsPackrFile=$(cd ${datasetsPath}; ls | grep "\-packr.go")
sed -i '' '/.go\"/d' ${templatesPath}/${templatesPackrFile}
sed -i '' '/.go\"/d' ${datasetsPath}/${datasetsPackrFile}
