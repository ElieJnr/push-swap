#!/bin/bash
rm -rf check
rm -rf push
cd checker
go build -o check
mv check ../
cd ..
cd push-swap
go build -o push
mv push ../
cd ..
