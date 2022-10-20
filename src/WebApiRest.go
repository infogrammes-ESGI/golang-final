package main

import ( // Les imports de bibliothèque.
	"encoding/json" // Format de données textuelle d'information structurée.
	"fmt"           // Paquet pour formater des strings, valeurs et de collecter les entrées d'utilisateur.
	"io/ioutil"
	"ioutil"
	"log"       // Paquet pour les logs en cas d'erreur par exemple.
	"math/rand" // Paquet qui permet d'utiliser la fonction rand
	"net/http"  // Paquet qui permet les connexions et la mise en place de serveur et client HTTP.
)

type Parc struct { // Création de la structure, contenant diverses données.
	id           int    `json:"id"`
	Name         string `json:"Name"`
	Parc         string `json:"inPark"`
	Place        string `json:"Place"`
	Manufacturer string `json:"Manufacturer"`
}

var Parcs []Parc     // On met les données de la structure dans la variable Parcs.
var results []string // Variable qui va contenir des informations de type string.

func homePage(w http.ResponseWriter, r *http.Request) { // Création de la page web
	fmt.Fprintf(w, "Bienvenu sur la page principale")
	fmt.Println("Endpoint Hit: homePage")
	log.Fatal(http.ListenAndServe(":8080", nil)) // On utilise le port 8080

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) { // On récupère les informations de la structure viala variable Parcs.
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Parcs)
}

func handleRequests() { // On indique l'arborescence de l'appli web.
	http.HandleFunc("/", homePage)               // http://localhost:8080/ = page principale.
	http.HandleFunc("/infos", returnAllArticles) // http://localhost:8080/infos = page affichant les données de la structure.

}

func SearchId(id int) (parc *Parc, index int) { // Fonction pour rechercher un parc par son ID
	if len(Parcs) == 0 {
		return nil, -1
	}

	for i, v := range Parcs {
		if v.id == id {
			return &v, i
		}
	}
	return nil, -1
}

func Post(w http.ResponseWriter, r *http.Request) { // Fonction qui permet d'utiliser la méthode POST qui pourras mettre à jour ou créer un parc.
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body) // On récupère la saisie et on la définie dans la variable 'body'
		if err != nil {                     // On renvoi une erreur si la saisie est erronée.
			http.Error(w, "Erreur de requête", http.StatusInternalServerError)
		}
		results = append(results, string(body)) // On envoi la saisie dans la variable 'results'

		fmt.Fprint(w, "POST éffectué")
	}

}

func delete(w http.Response, r *http.Request) { // Fonction qui utilise la méthode Delete pour supprimer une donnée.
	if r.Method == "Delete" {
		body, err := ioutil.ReadAll(r.Body) // On défini la variable 'body' avec la saisie.
		if err != nil {
			http.Error(w, "Erreur de requête", http.StatusInternalServerError)
		}
		results = append(results, string(body))

		fmt.Fprint(w, "Delete effectué")
	}

}

func main() {

	Parcs = []Parc{
		Parc{id: rand.Int(), Name: "Osiris", Parc: "Asterix", Place: "France", Manufacturer: "Vortex"},
		Parc{id: rand.Int(), Name: "Taron", Parc: "Phantasialand", Place: "Allemagne", Manufacturer: "Vortex"},
		Parc{id: rand.Int(), Name: "TheMonster", Parc: "Walygator Parc", Place: "France", Manufacturer: "Vortex"},
	}

	homePage()
	returnAllArticles()
	handleRequests()
	SearchId()
	Post()
	Delete()
}
