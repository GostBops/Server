/*
 * Swagger Blog
 *
 * A Simple Blog
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/binary"
	"encoding/json"
	//"errors"
	//"github.com/boltdb/bolt"
	//"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	//"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
    SecretKey = "gostbops"
)


func fatal(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type Response struct {
    Data string `json:"data"`
}

type Token struct {
    Token string `json:"token"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}


func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}

/*func CreateComment(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	articleId  := strings.Split(r.URL.Path, "/")[3]
	Id, err:= strconv.Atoi(articleId)
	if err != nil {
		response := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			v := b.Get(itob(Id))
			if v == nil {
				return errors.New("Article Not Exists")
			} else {
				return nil
			}
		}
		return errors.New("Article Not Exists")
	})
	
	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	comment := &Comment{
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Content: "",
		Author: "",
		ArticleId: Id,
	}
	err = json.NewDecoder(r.Body).Decode(&comment)

	if err != nil  || comment.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			response := ErrorResponse{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
		} else {
			response := ErrorResponse{"There is no content in your article"}
			JsonResponse(response, w, http.StatusBadRequest)
		} 
		return
	}

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(comment.Author), nil
        })

    if err == nil {
        if token.Valid {


			err = db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte("Comment"))
				if err != nil {
					return err
				}
				id, _ := b.NextSequence()
				encoded, err := json.Marshal(comment)
				return b.Put(itob(int(id)), encoded)
			})
		
			if err != nil {
				response := ErrorResponse{err.Error()}
				JsonResponse(response, w, http.StatusBadRequest)
				return
			}


			JsonResponse(comment, w, http.StatusOK)
        } else {
			response := ErrorResponse{"Token is not valid"}
			JsonResponse(response, w, http.StatusUnauthorized)
        }
    } else {
		response := ErrorResponse{"Unauthorized access to this resource"}
		JsonResponse(response, w, http.StatusUnauthorized)
    }
}*/

func CreateComment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(172.18.0.2:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	articleId  := strings.Split(r.URL.Path, "/")[3]
	Id, err:= strconv.Atoi(articleId)
	if err != nil {
		response := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	query, err := db.Query("select * from test.Article where id=" + articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	if !query.Next() {
		response := ErrorResponse{"Article Not Exists"}
		JsonResponse(response, w, http.StatusBadRequest)  
		return
	}

	comment := &Comment{
		Date:  time.Now().Format("2006-01-02 15:04:05"),
		Content: "",
		Author: "",
		ArticleId: Id,
	}
	err = json.NewDecoder(r.Body).Decode(&comment)

	if err != nil  || comment.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			response := ErrorResponse{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
		} else {
			response := ErrorResponse{"There is no content in your article"}
			JsonResponse(response, w, http.StatusBadRequest)
		} 
		return
	}

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(comment.Author), nil
        })

    if err == nil {
        if token.Valid {
			query, err = db.Query("INSERT INTO `test`.`Comment` (`date`, `author`, `articleId`, `content`) VALUES ('" + comment.Date + "', '" + comment.Author + "', " + articleId + ", '" + comment.Content + "')")
			if err != nil {
				log.Fatal(err)
			}
			defer query.Close()


			JsonResponse(comment, w, http.StatusOK)
        } else {
			response := ErrorResponse{"Token is not valid"}
			JsonResponse(response, w, http.StatusUnauthorized)
        }
    } else {
		response := ErrorResponse{"Unauthorized access to this resource"}
		JsonResponse(response, w, http.StatusUnauthorized)
    }
}

/*func SignIn(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	var user User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
			response := ErrorResponse{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
			return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b != nil {
			v := b.Get([]byte(user.Username))
			if ByteSliceEqual(v, []byte(user.Password)) {
				return nil
			} else {
				return errors.New("Wrong Username or Password")
			}
		} else {
			return errors.New("Wrong Username or Password")
		}
	})

	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
			fatal(err)
	}

	tokenString, err := token.SignedString([]byte(user.Username))
	if err != nil {
			fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w, http.StatusOK)
}*/

func SignIn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(172.18.0.2:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	var user User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
			response := ErrorResponse{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
			return
	}

	query, err := db.Query("select * from test.User where username='" + user.Username + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Wrong Username or Password"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	var userQuery User
	v = v[1:len(v)-1]
	_ = json.Unmarshal(v, &userQuery)

	if userQuery.Password != user.Password {
		response := ErrorResponse{"Wrong Username or Password"}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
			fatal(err)
	}

	tokenString, err := token.SignedString([]byte(user.Username))
	if err != nil {
			fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w, http.StatusOK)
}

func ByteSliceEqual(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    if (a == nil) != (b == nil) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func JsonResponse(response interface{}, w http.ResponseWriter, code int) {
    json, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
        return
    }


    w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
    w.Write(json)
}

/*func SignUp(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil || user.Password == "" || user.Username == "" {
			response := ErrorResponse{"Wrong Username or Password"}
			JsonResponse(response, w, http.StatusBadRequest)
			return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b != nil {
			v := b.Get([]byte(user.Username))
			if v != nil {
				return errors.New("User Exists")
			}
		}
		return nil
	})

	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("User"))
		if err != nil {
			return err
		}
		return b.Put([]byte(user.Username), []byte(user.Password))
	})

	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}*/

func SignUp(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(172.18.0.2:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil || user.Password == "" || user.Username == "" {
			response := ErrorResponse{"Wrong Username or Password"}
			JsonResponse(response, w, http.StatusBadRequest)
			return
	}

	query, err := db.Query("select * from test.User where username='" + user.Username + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	if query.Next() {
		response := ErrorResponse{"User Exists"}
		JsonResponse(response, w, http.StatusBadRequest)  
		return
	}

	query, err = db.Query("INSERT INTO `test`.`User` (`username`, `password`) VALUES ('" + user.Username + "', '" + user.Password + "')")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

/*func CreateArticle(w http.ResponseWriter, r *http.Request) {



	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()



	var articleInfo ArticleCreate
	err = json.NewDecoder(r.Body).Decode(&articleInfo)

	if err != nil  || articleInfo.Content == "" || articleInfo.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Fprint(w, err)
			fmt.Print(err)
		} else if articleInfo.Content == "" {
			fmt.Fprint(w, "There is no content in your article")
			fmt.Print("There is no content in your article")
		} else {
			fmt.Fprint(w, "There is no name of your article")
			fmt.Print("There is no name of your article")
		}
		return
	}
	
	var tags []Tag

	for i := 0; i < len(articleInfo.Tags); i++ {
		tags = append(tags, Tag{
			Name: articleInfo.Tags[i],
		})
	}
	
	str, err := ioutil.ReadFile("./a.md")

	if err != nil {

		fmt.Println("read file error")

		return

	}
	articleInfo.Content = string(str)
	article := &Article {
		Id: 1,
		Name: articleInfo.Name,
		Tags: tags,
		Date: time.Now().Format("2006-01-02 15:04:05"),
		Content: articleInfo.Content,
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Article"))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()
		article.Id = int(id)
		encoded, err := json.Marshal(article)
		byte_id := itob(article.Id)
		return b.Put(byte_id, encoded)
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(err)
		fmt.Fprint(w, err)
		return
	}


	for i := 0; i < len(tags); i++ {
		err = db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("Tag"))
			if err != nil {
				return err
			}
			var n []byte
			return b.Put([]byte(tags[i].Name), n)
		})
	
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print(err)
			fmt.Fprint(w, err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	JsonResponse(article, w, http.StatusOK)

}*/