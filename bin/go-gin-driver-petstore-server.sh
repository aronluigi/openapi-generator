
#!/bin/sh

SCRIPT="$0"
echo "# START SCRIPT: $SCRIPT"

while [ -h "$SCRIPT" ] ; do
  ls=`ls -ld "$SCRIPT"`
  link=`expr "$ls" : '.*-> \(.*\)$'`
  if expr "$link" : '/.*' > /dev/null; then
    SCRIPT="$link"
  else
    SCRIPT=`dirname "$SCRIPT"`/"$link"
  fi
done

if [ ! -d "${APP_DIR}" ]; then
  APP_DIR=`dirname "$SCRIPT"`/..
  APP_DIR=`cd "${APP_DIR}"; pwd`
fi

executable="./modules/openapi-generator-cli/target/openapi-generator-cli.jar"

if [ ! -f "$executable" ]
then
  mvn -B clean package
fi

SPEC="modules/openapi-generator/src/test/resources/2_0/petstore.yaml"
GENERATOR="go-gin-driver"
STUB_DIR="samples/server/petstore/go-gin-driver-server"

echo "Removing files and folders under $STUB_DIR"
rm -rf $STUB_DIR

# if you've executed sbt assembly previously it will use that instead.
export JAVA_OPTS="${JAVA_OPTS} -Xmx1024M -DloggerPath=conf/log4j.properties"

ags="generate -t modules/openapi-generator/src/main/resources/go-gin-driver-server -i $SPEC -g $GENERATOR -o $STUB_DIR --additional-properties packageName=petstoreserver --additional-properties hideGenerationTimestamp=true $@"

java $JAVA_OPTS -jar $executable $ags
#!/usr/bin/env bash