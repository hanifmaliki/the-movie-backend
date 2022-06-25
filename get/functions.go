package get

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopRated(c *gin.Context) {
	page := c.Param("page")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/movie/top_rated?api_key="+api_key+"&language=en-US&page="+page, nil)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	var responseObject TopRated
	json.Unmarshal(bodyBytes, &responseObject)
	c.IndentedJSON(http.StatusOK, responseObject)
}
