// hnApiTest plays with the Algolia Search API for HackerNews and the gopencils package
//
// structs for unmarshall were created with https://github.com/ChimeraCoder/gojson
// References:
// 	https://hn.algolia.com/api
//
//	BaseUrl: https://hn.algolia.com/api/v1
// 	Items: /items/:id
// 	Users: /users/:username
//	Search: /search?query=foo&tags=story
//
package main

import (
	"fmt"

	"github.com/bndr/gopencils"
	"github.com/kr/pretty"
)

func main() {
	api := gopencils.Api("https://hn.algolia.com/api/v1/")

	usernames := []string{"r0nin", "whiteacid", "gdubs"}

	searchQry := map[string]string{
		"tags": "",
	}

	for _, username := range usernames {
		// Create a new pointer to response Struct
		r := new(User)
		// Get user with id i into the newly created response struct
		_, err := api.Res("users").Id(username, r).Get()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%# v\n", pretty.Formatter(r))
		searchQry["tags"] = "author_" + username

		search := new(Search)
		_, err = api.Res("search", search).Get(searchQry)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(search.NbHits, "for user:", username)
		for _, hit := range search.Hits {
			if hit.ObjectID == "" {
				continue
			}

			i := new(Item)
			_, err = api.Res("items", i).Id(hit.ObjectID).Get()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if i.Type != "story" {
				continue
			}
			pretty.Println(i)

		}
	}
}
