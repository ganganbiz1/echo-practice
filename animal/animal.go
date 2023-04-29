package animal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

type Animal struct {
	Name string `json:"name"`
}

type RequestBody struct {
	Name string `json:"name"`
}

type Response struct {
	Message     string `json:"message"`
	Explanation string `json:"explanation"`
}

func New(name string) *Animal {
	return &Animal{Name: name}
}

func (a *Animal) Get(c echo.Context) error {
	var res Response

	animalRes, err := http.Get("http://" + os.Getenv("ANIMAL_API_HOST") + ":" + os.Getenv("ANIMAL_API_PORT") + "/animal/" + a.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, res)
	}

	defer animalRes.Body.Close()

	body, _ := io.ReadAll(animalRes.Body)

	log.Printf("Response:%s", string(body))

	return c.String(http.StatusOK, string(body))
}

func (a *Animal) Post(c echo.Context) error {
	rbody := &RequestBody{
		Name: a.Name,
	}

	animalData, err := json.Marshal(rbody)
	if err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}

	log.Printf("Animal:%s", string(animalData))

	base, _ := url.Parse("http://" + os.Getenv("ANIMAL_API_HOST") + ":" + os.Getenv("ANIMAL_API_PORT"))

	reference, _ := url.Parse("animal/register")

	endpoint := base.ResolveReference(reference).String()

	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(animalData))

	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	log.Printf("Response:%s", string(body))

	jsonData, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	})
	return c.String(http.StatusCreated, string(jsonData))
}
