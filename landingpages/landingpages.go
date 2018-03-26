package landingpages

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/gorilla/mux"
)

// Catalog  NOTE the so slice does nothing..  the template is just static content.
// This function is here only in the thought this could be a dynamic list at some point.
func Catalog(w http.ResponseWriter, r *http.Request) {
	log.Println("Generating dataset catalog")

	so := []string{"p418graph", "p418prov", "p418querylogs", "p418spatial"}

	templateFile := "./templates/datacatalog.html"

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", so)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}

}

func LandingPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Landing page for ID for resource: %s\n", vars["id"])

	so := schemaFromPackage(vars["id"]) // TODO..   injection danger?   strings.Trim this...
	templateFile := "./templates/lptemplate.html"

	ht, err := template.New("Template").ParseFiles(templateFile) //open and parse a template text file
	if err != nil {
		log.Printf("template parse failed: %s", err)
	}

	err = ht.ExecuteTemplate(w, "Q", so)
	if err != nil {
		log.Printf("Template execution failed: %s", err)
	}
}

func schemaFromPackage(datapkg string) string {
	pkg, err := datapackage.Load(fmt.Sprintf("./static/datapackages/%s/datapackage.json", datapkg))
	// pkg, err := datapackage.Load("https://geodex.org/id/datafiles/REFID/datapackage.json")
	if err != nil {
		log.Printf("Error getting package: %v\n", err)
		return "Error getting package"
	}

	resource := pkg.GetResource("schemaorg") // could I also read the "schemaorg" JSON file in and use it?
	rc, err := resource.RawRead()
	defer rc.Close()
	contents, _ := ioutil.ReadAll(rc)
	if err != nil {
		log.Printf("Error getting raw read of entry: %v\n", err)
		return "Error getting schema.org entry"
	}

	return string(contents) // just grab the first "thing" and return it..  we are just testing pipelines...
}