package lydia

import (
    "net/http"
    "encoding/json"
    "errors"
    "io/ioutil"
    "github.com/jaskaranSM/lydia-client-go/types"
)

func constructUrl(url string, key string,params string) string{
	finalUrl := url + "?access_key="+key 
	if params != ""{
		finalUrl += "&"+params
	}
	return finalUrl
}

type LydiaAI struct {
    API_KEY string
    API_ENDPOINT string
}

func(L *LydiaAI) CreateSession() (types.SessionResponse,error) {
	url := constructUrl(L.API_ENDPOINT+"session/create",L.API_KEY,"")
	resp,err := http.Get(url)
	if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    session := types.SessionResponse{}
    bytes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return session,err
	}
    if resp.StatusCode == 200{
    	json.Unmarshal(bytes,&session)
    	return session,nil
    } else {
    	return session,errors.New(string(bytes))
    }
}

func(L *LydiaAI) ThinkThought(session_id string, text string) (types.ThoughtResponse,error) {
	params := "session_id="+session_id+"&input="+text
	url := constructUrl(L.API_ENDPOINT+"session/think",L.API_KEY,params)
	resp,err := http.Get(url)
	if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    thought := types.ThoughtResponse{}
    bytes,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return thought,err
	}
    if resp.StatusCode == 200{
    	json.Unmarshal(bytes,&thought)
    	return thought,nil
    } else {
    	return thought,errors.New(string(bytes))
    }
}

func NewClient(api_key string) *LydiaAI{
	return &LydiaAI{API_KEY:api_key, API_ENDPOINT:"https://api.intellivoid.net/coffeehouse/v1/lydia/"}
}