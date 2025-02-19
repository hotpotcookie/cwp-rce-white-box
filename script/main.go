package main
//----------
import (
    json     "encoding/json"
    ioutil   "io/ioutil"
    log      "log"
    http     "net/http"
    mux      "github.com/gorilla/mux"
    time     "time"
)

//----------
type DATA struct {
    PASSWD  string `json:"PASSWD"`
    SHADOW  string `json:"SHADOW"`
} 
type property struct {
    IP_LISTENER     string `json:"IP_LISTENER"`
    IP_TARGET       string `json:"IP_TARGET"`
    PORT_LISTENER   string `json:"PORT_LISTENER"`
    LAST_UPDATE     string `json:"LAST_UPDATE"`
    DATA            DATA   `json:"DATA"`
}
//----------
type allProperties []property;
var properties = allProperties{
    {
        IP_LISTENER: "",
        IP_TARGET: "",        
        PORT_LISTENER: "",
        LAST_UPDATE: "",
    },
}

//----------
func getProperties(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(properties);
}
func updateProperties(w http.ResponseWriter, r *http.Request) {
    var updatedProperty property;
    reqBody,_ := ioutil.ReadAll(r.Body);
    json.Unmarshal(reqBody, &updatedProperty);
    currentTime := time.Now();    
    //----------
    for i, key := range properties {
        if len(updatedProperty.PORT_LISTENER) > 0 { key.PORT_LISTENER = updatedProperty.PORT_LISTENER; } else { key.PORT_LISTENER = key.PORT_LISTENER; }
        if len(updatedProperty.IP_LISTENER) > 0 { key.IP_LISTENER = updatedProperty.IP_LISTENER; } else { key.IP_LISTENER = key.IP_LISTENER; }        
        if len(updatedProperty.IP_TARGET) > 0 { key.IP_TARGET = updatedProperty.IP_TARGET; } else { key.IP_TARGET = key.IP_TARGET; }                
        if len(updatedProperty.DATA.PASSWD) > 0 { key.DATA.PASSWD = updatedProperty.DATA.PASSWD; } else { key.DATA.PASSWD = key.DATA.PASSWD; }        
        if len(updatedProperty.DATA.SHADOW) > 0 { key.DATA.SHADOW = updatedProperty.DATA.SHADOW; } else { key.DATA.SHADOW = key.DATA.SHADOW; }                
        key.LAST_UPDATE = currentTime.String();
        properties = append(properties[:i], key);
        json.NewEncoder(w).Encode(key);
    }
}

//----------
func main() {
    router := mux.NewRouter().StrictSlash(true);
    router.HandleFunc("/", getProperties).Methods("GET");
    router.HandleFunc("/", updateProperties).Methods("PATCH");
    log.Fatal(http.ListenAndServe(":2080", router));
}