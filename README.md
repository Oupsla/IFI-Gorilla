# IFI-Gorilla
Petit hello world des familles pour devenir un gopher !

## Auteurs
Benjamin Coenen - Nicolas Delperdange - Denis Hamann

## Installation
Téléchargement :
```
curl -O https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
tar xvf go1.7.3.linux-amd64.tar.gz
```

Déplacement :
```
mv go /usr/local
```

Ajout au profile (ou .bashrc)
```
vim ~/.profile
```
```
export GOPATH=$HOME/golang
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
```
vim ~/.profile
```

Création du workspace golang (**modifier user par votre nom d'utilisateur github**)
```
mkdir -p $HOME/golang/src/github.com/user/hello
cd $HOME/golang/src/github.com/user/hello
```

Vous pouvez maintenant écrire votre premier programme go avec votre éditeur de texte préféré :D
```
vim hello.go
```


Pour lancer votre programme
```
go run hello.go
```

Pour le compiler
```
go install hello.go
```

## Exercice

Produire un hello world avec les caractéristiques suivantes :
- La string 'Hello world' est définie dans une constante
- Une boucle allant de 0 jusqu'à 9 avec un for
- Quand la boucle atteint la 5ième itération, imprimer la constante Hello world

## Ressources 
Pleins d'exemple en go disponible ici : https://gobyexample.com/
