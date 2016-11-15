#! /bin/bash
generatedDirs=("bin/blackbox")
projectName=$(basename $(pwd))

goversion=`go version`
IFS='  ' read -ra array <<< $goversion
pkgSystemName="${array[@]: -1:1}"
pkgSystemName=`echo "$pkgSystemName" | tr / _`

pkgDirectory=$GOPATH/pkg/$pkgSystemName/$projectName
rm -Rf $pkgDirectory
cp -r $generatedDirs $pkgDirectory

echo "Binary package copied in $pkgDirectory"