package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pterm/pterm"
)

type userRes struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		ZipCode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchphrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var results []userRes

	err = json.Unmarshal(bs, &results)

	if err != nil {
		log.Fatalf("failed to unmarshal json response:%v", err)
	}

	tableData := pterm.TableData{
		{"Name", "Username", "Email", "Phone", "Address"},
	}

	for _, r := range results {
		tableData = append(tableData, []string{
			r.Name, r.UserName, r.Email, r.Phone, r.Address.Street,
		})
	}

	err = pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData(tableData).
		Render()

	if err != nil {
		log.Fatal(err)

	}
}
