package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/icrowley/fake"
	"github.com/nylo-andry/search-service/config"
	"github.com/olivere/elastic"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RealName string `json:"real_name"`
}

var elasticClient *elastic.Client

func InitClient(conf config.Configurations) {
	var err error

	url := fmt.Sprintf("http://%s:%s", conf.ElasticHost, conf.ElasticPort)
	elasticClient, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
}

func Populate(number int) error {
	idxExists, err := elasticClient.IndexExists("users").Do(context.Background())
	if err != nil {
		return err
	}
	if !idxExists {
		elasticClient.CreateIndex("users").Do(context.Background())
	}

	for i := 0; i < number; i++ {
		user := User{
			Username: fake.UserName(),
			Email:    fake.EmailAddress(),
			RealName: fake.FullName(),
		}
		_, err = elasticClient.Index().
			Index("users").
			Type("doc").
			BodyJson(user).
			Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func Search(term string, from, size int) ([]*User, error) {
	q := elastic.NewMultiMatchQuery(term, "username", "email", "real_name").Fuzziness("AUTO:2,5")
	res, err := elasticClient.Search().
		Index("users").
		Query(q).
		From(from).
		Size(size).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	users := make([]*User, 0)

	for _, hit := range res.Hits.Hits {
		var user User
		err := json.Unmarshal(*hit.Source, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
