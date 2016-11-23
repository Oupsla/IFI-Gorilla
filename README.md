# IFI-Gorilla

## Avant toute chose

Pour commencer, clonez ce projet à l'adresse suivante sur votre poste :
```bash
mkdir -p $HOME/golang/src/github.com/oupsla
cd $HOME/golang/src/github.com/oupsla
git clone https://github.com/Oupsla/IFI-Gorilla.git
```

## Tutorial
Ce tp consiste à résoudre les énigmes de la boite mystère (la blackbox), qui est un fichier compilé contenant une méthode accessible *ExecuteFunction(nomDeLaFonction, method, params)*.

Votre but est d'appeler cette méthode en écrivant un serveur http avec Gorilla. Une base vous est fournie afin de faciliter le début.

La base est découpée comme suit:
- bin et blackbox : contient la boite mystère compilée (**A ne pas modifier**)
- GoDeps et vendor : contient les dépendances (**A ne pas modifier**)
- Utils : contient des méthodes de formatage (**A ne pas modifier**)
- routes : c'est ici que vous allez déclarer les routes pour Gorilla
- controllers : c'est ici que vous allez appeler la blackbox, les controlleurs traitant les routes.


## Comment

### Build l'API
```bash
$ make dependencies
$ ./deploy.sh
```

### Lancer l'API
```bash
$ make start
```

### Contacter l'API

Pour contacter votre api, nous vous conseillons d'utiliser PostMan pour Chrome : https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop

Url : localhost:8080/api/v1/welcome

Voilà à quoi devrait ressembler Postman pour la première étape :
![Postman.png](https://s16.postimg.org/wya70juad/Postman.png)

## Ajouter une route
Pour la prochaine étape, vous devez donc rajouter une route au tableau de route du fichier : routes/ifi/ifi.go  
Rendez vous dans routes/ifi/ifi.go et dans la partie
```golang
Routes = append(Routes, routes.Route{
  Name:        "Get welcome message",
  Method:      "GET",
  Path:        "/welcome",
  HandlerFunc: ifi.Welcome,
})
```
Il suffit d'ajouter un nouvel objet de ce type à la suite dans la fonction append, exemple :
```golang
Routes = append(Routes, routes.Route{
  Name:        "Get welcome message",
  Method:      "GET",
  Path:        "/welcome",
  HandlerFunc: ifi.Welcome,
}, routes.Route{
  Name:        "2nd route",
  Method:      "GET",
  Path:        "/myroute", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
  HandlerFunc: ifi.MyHandlerFunctionForRoute,
})
```
Attention, votre **HandlerFunc** doit avoir le même nom que celle que vous déclarerez dans votre fichier controllers/ifi/ifi.go dans l'étape suivante, et surtout n'oubliez pas de la faire débuter par une majuscule pour qu'elle soit publique et accessible dans un autre package.

## Ajouter un controller
Rendez vous dans **controllers/ifi/ifi.go** et ajoutez y une fonction handler qui répondra à l'appel de votre route comme pour l'exemple Welcome.
```golang
func MyHandlerFunctionForRoute(res http.ResponseWriter, req *http.Request) {
	response, err := blackbox.ExecuteFunction("fonctionBlackBox", req.Method, params.Result)

	if err != nil {
		myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
		utils.SendError(res, myError)
		return
	}

	utils.SendJSON(res, response)
}
```
Maintenant votre API répondra à la route que vous avez définie précédemment et appelera la fonction correspondante.

## Fonctions utilitaires

### Récupérer le body
Afin de vous faciliter la tâche nous avons développé des fonctions utilitaires dans **utils/requestFormatter.go**

SendJSON et SendError sont utilisées dans la fonction Welcome déjà existante dans le controller et servent respectivement à faire un retour JSON et un retour erreur de votre API.

Pour lire le body d'un POST que vous faites via Postman vous pouvez utiliser la fonction GetJSONContent, voici un exemple :
```golang
type postParamsEnigma struct {
  Result int `json:"result"`
}

func PostSomething(res http.ResponseWriter, req *http.Request) {
  params := &postParamsEnigma{}

  err := utils.GetJSONContent(params, req)

  if err != nil {
    myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
    utils.SendError(res, myError)
    return
  }

}
```
Votre structure déclarée ci-dessus (`postParamsEnigma`) est obligatoire et représente le format que JSON que vous lui envoyez dans le body de votre POST

### Récupérer le path params
La fonction `ParamAsString` sert à récupérer un query params dans une URL -> exemple `monUrl/:monParam`

Voici comment l'utiliser :
```
func RemoveSomething(res http.ResponseWriter, req *http.Request) {
  param := utils.ParamAsString("monParam", req)

  //...
}
```
De cette manière dans param vous récupèrerez votre query param.
