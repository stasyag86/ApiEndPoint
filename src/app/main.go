package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "strings"
  "log"
  "sync"
)

import ej "app/easyJson"


const default_port = 9087
const endPointURL = "http://localhost:3000"
var proxyMap = make(map[string][] chan string)
var mux sync.Mutex

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
    if (r.Method == http.MethodPost){
      handlePostRequest(w, r)
    }
  }

  func handlePostRequest(w http.ResponseWriter, r *http.Request){
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
      log.Panic(err)
    }
    
    domainAndPAth := &ej.RequestPattern{}
    domainAndPAth.UnmarshalJSON(reqBody)

    fullUrl := domainAndPAth.Domain + domainAndPAth.Path
    done := make(chan string)
    if (proxyRequest(fullUrl, done)){
       //wait
      resFromProxy := <-done
      w.Write([]byte(resFromProxy))
    }else{
      resFromDb := findInDB(fullUrl)
      proxyResultForAll(fullUrl, resFromDb)
      //resFromProxy := <-done
      w.Write([]byte(resFromDb))
    }
    
  }

  func findInDB(fullUrl string) string{
    resp, err := http.Post(endPointURL, "application/x-www-form-urlencoded", strings.NewReader(fullUrl))
    if err != nil {
      log.Panic(err)
    }

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
      log.Panic(err)
    }
    bodyString := string(bodyBytes)

    var resultFromDB ej.ResponsePattern
    resultFromDB.Location = bodyString
    data, err := resultFromDB.MarshalJSON()
    if (err != nil){
      log.Panic(err)
    }
    return string(data)
  }

  func proxyRequest(url string, done chan string) bool{
    mux.Lock()
    values, ok := proxyMap[url]
    if ok {   //just add to map
      values = append(values,done)
      proxyMap[url] = values
      defer mux.Unlock()
      return true
    } else {  //create new record in map
      var values []chan string
      values = append(values,done)
      proxyMap[url] = values
      defer mux.Unlock()
      return false
    }
  }

  func proxyResultForAll(url string, resFromDb string){
    values := proxyMap[url]
    for _, doneChannel := range values {
      go triggetEvent(doneChannel, resFromDb)
    }
    mux.Lock()
    delete(proxyMap, url)
    defer mux.Unlock()
  }

  func triggetEvent(doneChannel chan string, resFromDb string){
    doneChannel <- resFromDb
    close(doneChannel)
  }


func main() {
  log.Println("Http server is ready on port: ", default_port)
  http.HandleFunc("/", handleRequestAndRedirect)
  port := fmt.Sprintf(":%d", default_port) 
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
    } 
  
}
