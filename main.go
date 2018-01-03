package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
	"github.com/wheatandcat/dotstamp_graphql/types"
	"github.com/wheatandcat/dotstamp_graphql/utils/auth"
	"github.com/wheatandcat/dotstamp_graphql/utils/contributions"
	"github.com/wheatandcat/dotstamp_graphql/utils/contributionsDetail"
	"github.com/wheatandcat/dotstamp_graphql/utils/follows"
	"github.com/wheatandcat/dotstamp_graphql/utils/login"
	"github.com/wheatandcat/dotstamp_graphql/utils/tags"
	"github.com/wheatandcat/dotstamp_graphql/utils/users"
)

// DB database connection
var DB *sqlx.DB

// CONF config info
var CONF ConfiInfo

var authKey string

// ConfiInfo config info type
type ConfiInfo struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	LoginKey string `yaml:"loginkey"`
}

func connectDB() {
	var buf []byte
	var err error

	if os.Getenv("ENV_CONF") == "prod" {
		buf, err = ioutil.ReadFile("./config/prod.yaml")
	} else {
		buf, err = ioutil.ReadFile("./config/devlop.yaml")
	}

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &CONF)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Connect("mysql", CONF.User+":"+CONF.Password+"@"+CONF.Host+"/"+CONF.Dbname)
	log.Println(CONF.User + ":" + CONF.Password + "@" + CONF.Host + "/" + CONF.Dbname)
	if err != nil {
		panic(err)
	}

	DB = db
}

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"me": &graphql.Field{
			Type:        types.UserType,
			Description: "find me",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, err := authJwt.Auth(authKey, CONF.LoginKey)
				if err != nil {
					u := types.UserMaster{}
					return u, nil
				}

				u := types.UserMaster{}
				err = DB.Get(&u, "SELECT * FROM user_masters WHERE id=?", id)
				if err != nil {
					u = types.UserMaster{}
					return u, nil
				}

				return u, nil
			},
		},
		"user": &graphql.Field{
			Type:        types.UserType,
			Description: "find user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "user id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				u := types.UserMaster{}
				err := DB.Get(&u, "SELECT * FROM user_masters WHERE id=?", id)

				return u, err
			},
		},
		"userList": &graphql.Field{
			Type:        graphql.NewList(types.UserType),
			Description: "user list",
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "number of item displayed",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				first, _ := p.Args["first"].(int)
				u := []types.UserMaster{}
				err := DB.Select(&u, "SELECT * FROM user_masters ORDER BY id ASC LIMIT ?", first)

				return u, err
			},
		},
		"contribution": &graphql.Field{
			Type:        types.ContributionType,
			Description: "find contribution",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "contribution id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idQuery, _ := p.Args["id"].(int)
				u := types.UserContribution{}
				err := DB.Get(&u, "SELECT * FROM user_contributions WHERE id=?", idQuery)

				return u, err
			},
		},
		"contributionList": &graphql.Field{
			Type:        graphql.NewList(types.ContributionType),
			Description: "contribution list",
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "number of item displayed",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				first, _ := p.Args["first"].(int)
				u, err := contributions.GetContributions(DB, first)
				if err != nil {
					return nil, err
				}

				u, err = tags.MapTgas(DB, u)
				if err != nil {
					return nil, err
				}

				u, err = follows.MapTgas(DB, u)
				if err != nil {
					return nil, err
				}

				return u, nil
			},
		},
		"contributionDetail": &graphql.Field{
			Type:        graphql.NewList(types.BodyType),
			Description: "find contribution detail",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "contribution id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				u, err := contributionsDetail.GetByID(DB, id)
				if err != nil {
					return nil, err
				}

				body, err := contributionsDetail.GetBody(u.Body)

				return body, err
			},
		},
		"problemList": &graphql.Field{
			Type:        graphql.NewList(types.LogProblemContributionReportType),
			Description: "proble list",
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "number of item displayed",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				first, _ := p.Args["first"].(int)
				r := []types.LogProblemContributionReport{}
				err := DB.Select(&r, "SELECT l.id, l.user_contribution_id, c.title, l.type, l.created_at, l.updated_at, l.deleted_at FROM log_problem_contribution_reports as l INNER JOIN user_contributions  as c ON  l.user_contribution_id = c.id ORDER BY l.id DESC LIMIT ?", first)
				if err != nil {
					return nil, nil
				}

				return r, nil
			},
		},
		"questionList": &graphql.Field{
			Type:        graphql.NewList(types.LogQuestionType),
			Description: "question list",
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "number of item displayed",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				first, _ := p.Args["first"].(int)
				r := []types.LogQuestion{}
				err := DB.Select(&r, "SELECT * FROM log_questions ORDER BY id ASC LIMIT ?", first)
				if err != nil {
					return nil, nil
				}

				return r, nil
			},
		},
	},
})

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"hideContribution": &graphql.Field{
			Type:        types.HideType,
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "contribution id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)

				now := time.Now().Format("2006/01/02/15:04:05")

				DB.MustExec("UPDATE user_contributions SET deleted_at = '"+now+"' WHERE id = ?", id)
				u := types.Hide{}
				u.ID = uint(id)

				return u, nil
			},
		},
		"showContribution": &graphql.Field{
			Type:        types.HideType,
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "contribution id",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				u := types.Hide{}

				DB.MustExec("UPDATE user_contributions SET deleted_at = NULL WHERE id = ?", id)
				u.ID = uint(id)

				return u, nil
			},
		},
		"createUser": &graphql.Field{
			Type:        types.UserType,
			Description: "create user",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email, _ := p.Args["email"].(string)
				password := p.Args["password"].(string)

				u, err := users.Create(DB, email, password, CONF.LoginKey)
				if err != nil {
					return nil, err
				}

				return u, nil
			},
		},
		"login": &graphql.Field{
			Type:        types.LoginType,
			Description: "login check",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "email",
				},
				"password": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "password",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email, _ := p.Args["email"].(string)
				password, _ := p.Args["password"].(string)

				u, err := login.GetLogin(DB, email, password, CONF.LoginKey)
				if err != nil {
					return nil, err
				}
				key, err := authJwt.CreateTokenString(u.ID, CONF.LoginKey)
				if err != nil {
					return nil, err
				}
				var r types.AuthKey
				r.Key = key

				return r, nil
			},
		},
	},
})

func customHandler(schema *graphql.Schema) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		authKey = r.Header.Get("Authorization")

		opts := handler.NewRequestOptions(r)

		rootValue := map[string]interface{}{
			"response": rw,
			"request":  r,
		}

		params := graphql.Params{
			Schema:         *schema,
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			RootObject:     rootValue,
		}

		result := graphql.Do(params)

		jsonStr, err := json.Marshal(result)

		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")

		rw.Write(jsonStr)
	}
}
func main() {

	connectDB()

	Schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", customHandler(&Schema))
	http.ListenAndServe(":8080", nil)

}
