#!/bin/bash

cd log
rm *
touch _.json
cd ..
cd group
rm *
touch group.json
echo "[]" >> group.json

echo "init completed. to deploy, do not forget to change the DOCKER ENVIRONMENT VARIABLE into false."