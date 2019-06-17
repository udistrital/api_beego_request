# api_beego_request
Este proyecto es una API generada desde cero con el framework beego, contiene los cambios pertinente y recomendados por los lineamientos de la OAS que se encuentran en el repositorio [introduccion_oas](https://github.com/udistrital/introduccion_oas)

Tiene como proposito documentar el pipeline en drone para ejecutrar el marco de pruebas con Sonarquebe.

## .drone.yml

```bash

workspace:
  base: /go
  path: src/github.com/udistrital/${DRONE_REPO##udistrital/}
  when:
    branch: [master, develop]
    event: push

pipeline:
  # build go program
  go:
    image: golang:1.9
    commands:
     - go get -t
     - GOOS=linux GOARCH=amd64 go build -o main
    when:
      branch: [master, develop]
      event: push

  # test the go program
  go-test:
    image: golang:1.9
    commands:
     - go get -u gopkg.in/alecthomas/gometalinter.v1
     - gometalinter.v1 --install
     - go get github.com/axw/gocov/...
     - go get github.com/AlekSi/gocov-xml
     - go get -u github.com/jstemmer/go-junit-report
     - gometalinter.v1 ./... --checkstyle | tee report.xml
     - gocov test ./... | gocov-xml > coverage.xml
     - go test -v ./... | go-junit-report > test.xml
    when:
      branch: [develop]
      event: push

  # run sonar-scanner
  sonar-scanner-test:
    image: openjdk:8-alpine
    commands:
     - export RELEASE=3.3.0.1492
     - apk add --no-cache  curl grep sed unzip nodejs npm
     - curl --insecure -o ./sonarscanner.zip -L https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-$RELEASE-linux.zip
     - unzip sonarscanner.zip
     - rm sonarscanner.zip
     - rm -rf sonar-scanner-$RELEASE-linux/jre
     - sed -i 's/use_embedded_jre=true/use_embedded_jre=false/g' ./sonar-scanner-$RELEASE-linux/bin/sonar-scanner
     - export PATH=$PATH:/go/src/github.com/udistrital/${DRONE_REPO##udistrital/}/sonar-scanner-$RELEASE-linux/bin
     - cp sonar-project.properties ./sonar-scanner-$RELEASE-linux/conf/sonar-scanner.properties
     - sonar-scanner
    when:
      branch: [develop]
      event: push

```

## Licencia

This file is part of api_beego_request.

api_beego_request is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Foobar is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar.  If not, see <https://www.gnu.org/licenses/>.
