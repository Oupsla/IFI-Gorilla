#! /bin/bash
generatedDirs=("bin/blackbox.a")
projectName=$(basename $(pwd))

goversion=`go version`
IFS='  ' read -ra array <<< $goversion
pkgSystemName="${array[@]: -1:1}"
pkgSystemName=`echo "$pkgSystemName" | tr / _`

pkgDirectory=$GOPATH/pkg/$pkgSystemName/github.com/oupsla/$projectName
rm -Rf $pkgDirectory
mkdir -p $pkgDirectory
cp -r $generatedDirs $pkgDirectory

echo "Binary package copied in $pkgDirectory"
