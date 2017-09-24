package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Data Gets the different data recieved
type Data struct {
	Project        string   `json:"repository"`
	Owner          string   `json:"owner"`
	TopContributor string   `json:"topContributor"`
	Contributions  int      `json:"commits"`
	Languages      []string `json:"languages"`
}

//Repository has structure of some data
type Repository struct {
	FullName string `json:"full_name"`
	Owner    struct {
		Username string `json:"login"`
	} `json:"owner"`
}

//Contributors has structure of some data
type Contributors struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

func getAPIURL(url string) (link *http.Response) {
	link, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return
}

func readBody() {

}

func catchURL(number int) string {
	switch number {
	case 1:
		return "https://api.github.com/repos/Stektpotet/Amazeking"
	case 2:
		return "https://api.github.com/repos/Stektpotet/Amazeking/contributors"
	case 3:
		return "https://api.github.com/repos/Stektpotet/Amazeking/languages"
	}
	return ""
}

//serveRest is the handler in this program
func serveRest(w http.ResponseWriter, r *http.Request) {

	//Get response from url
	repo := getAPIURL(catchURL(1))
	cont := getAPIURL(catchURL(2))
	lang := getAPIURL(catchURL(3))

	defer repo.Body.Close()
	defer cont.Body.Close()
	defer lang.Body.Close()

	//Grab the body from the response
	body1, err1 := ioutil.ReadAll(repo.Body)
	if err1 != nil {
		panic(err1)
	}

	body2, err2 := ioutil.ReadAll(cont.Body)
	if err2 != nil {
		panic(err2)
	}

	body3, err3 := ioutil.ReadAll(lang.Body)
	if err3 != nil {
		panic(err3)
	}

	//Unmarshalling the body into a payload
	var p Repository
	var c []Contributors
	var v = make(map[string]int)

	err1 = json.Unmarshal(body1, &p)
	if err1 != nil {
		panic(err1)
	}

	err2 = json.Unmarshal(body2, &c)
	if err2 != nil {
		panic(err2)
	}

	err3 = json.Unmarshal(body3, &v)
	if err3 != nil {
		panic(err3)
	}

	//Working on end result
	repos := string("github.com/" + p.FullName)
	//Remember: var v = make(map[string]int)
	keys := []string{}
	for key := range v {
		keys = append(keys, key)
	}

	//Printing end result
	fmt.Fprintln(w, "Owner:         ", p.Owner.Username, "\nRepo:          ", repos)
	fmt.Fprintln(w, "Top committer: ", c[0].Login, "\nContributions: ", c[0].Contributions)
	fmt.Fprintln(w, "Languages:     ", keys)
}

/*Main retrieves data from a github api and displays
some of this on (randomly generated name)
https:secret-wave-88527.herokuapp.com/ */
func main() {
	http.HandleFunc("/", serveRest)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
