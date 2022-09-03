// when we will write our controller file we will need data that is unmarshalled 
// to unmarshall the data we will need a utility function and a corresponding file for the same


package utils


import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{} ){
	if body, err := ioutil.ReadAll(r.Body); err == nil{

		if err := json.Unmarshal([] byte(body), x ); err != nil{
			return 
		}
	}


}