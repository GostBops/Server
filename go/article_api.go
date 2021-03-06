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
	"encoding/json"
	//"encoding/binary"
  //"fmt"
    "log"
	"net/http"
	"net/url"
    "strings"
	//"errors"
	"strconv"
    //"github.com/codegangsta/negroni"
	//"github.com/boltdb/bolt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	//"reflect"
)


/*func GetArticleById(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	articleId := strings.Split(r.URL.Path, "/")[3]
	Id, err:= strconv.Atoi(articleId)
	if err != nil {
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	var article Article
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			v := b.Get(itob(Id))
			if v == nil {
				return errors.New("Article Not Exists")
			} else {
				_ = json.Unmarshal(v, &article)
				return nil
			}
		} else {
			return errors.New("Article Not Exists")
		}
	})

	if err != nil {
		reponse := ErrorResponse{err.Error()}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	JsonResponse(article, w, http.StatusOK)
}*/

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	articleId := strings.Split(r.URL.Path, "/")[3]
	_, err = strconv.Atoi(articleId)
	if err != nil {
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	
	query, err := db.Query("select * from test.Article where id=" + articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Article Not Exists"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	v = v[1:len(v)-1]
	str := strings.Replace(string(v), "id\":\"", "id\":", -1)
	str = strings.Replace(str, "\",\"name", ",\"name", -1)
	v = []byte(str)

	var article Article
	
	_ = json.Unmarshal(v, &article)
	file, err := ioutil.ReadFile(article.Content)
	if err != nil {
		log.Fatal(err)
	}
	article.Content = string(file)
	JsonResponse(article, w, http.StatusOK)

}
func getJSON(rows *sql.Rows) ([]byte, error) {
  columns, err := rows.Columns()
  if err != nil {
      return []byte(""), err
  }
	count := len(columns)

  tableData := make([]map[string]interface{}, 0)
  values := make([]interface{}, count)
  valuePtrs := make([]interface{}, count)
  for rows.Next() {
      for i := 0; i < count; i++ {
          valuePtrs[i] = &values[i]
      }
      rows.Scan(valuePtrs...)
      entry := make(map[string]interface{})
      for i, col := range columns {
          var v interface{}
          val := values[i]
          b, ok := val.([]byte)
          if ok {
              v = string(b)
          } else {
              v = val
          }
          entry[col] = v
      }
      tableData = append(tableData, entry)
	}
  jsonData, err := json.Marshal(tableData)
  if err != nil {
      return []byte(""), err
	}
  return jsonData, nil 
}

/*func GetArticles(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	IdIndex, err:= strconv.Atoi(page)
	IdIndex = (IdIndex - 1)* 10 + 1
	var articles ArticlesResponse
	var article ArticleResponse
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			 k, v := c.Seek(itob(IdIndex))
			 if k == nil {
				 return errors.New("Page is out of index")
			 }
			 key := binary.BigEndian.Uint64(k) 
			 fmt.Print(key)
			 if int(key) != IdIndex {
				 return errors.New("Page is out of index")
			 }
			 count := 0
			for ; k != nil && count < 10; k, v = c.Next() {
				err = json.Unmarshal(v, &article)
				if err != nil {
					return err
				}
				articles.Articles = append(articles.Articles, article)
				count = count + 1
			}
			return nil
		} else {
			return errors.New("Article Not Exists")
		}
	})
	if err != nil {
		reponse := ErrorResponse{err.Error()}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	json, err := json.Marshal(articles)
	fmt.Println(string(json))
	JsonResponse(articles, w, http.StatusOK)

}*/

func GetArticles(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	IdIndex, err:= strconv.Atoi(page)
	IdIndex = (IdIndex - 1)* 10
	Id := strconv.Itoa(IdIndex)

	query, err := db.Query("select * from test.Article limit " + Id + ",10")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	if string(v) == "[]" {
		reponse := ErrorResponse{"Page is out of index"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	var article ArticlesResponse
	v = []byte("{\"articles\":" + string(v) + "}")
	str := strings.Replace(string(v), "id\":\"", "id\":", -1)
	str = strings.Replace(str, "\",\"name", ",\"name", -1)
	v = []byte(str)

	_ = json.Unmarshal(v, &article)
	JsonResponse(article, w, http.StatusOK)
}

/*func GetCommentsOfArticle(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	articleId := strings.Split(r.URL.Path, "/")[3]
	Id, err:= strconv.Atoi(articleId)
	if err != nil {
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	var article []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			v := b.Get(itob(Id))
			if v == nil {
				return errors.New("Article Not Exists1")
			} else {
				article = v
				return nil
			}
		} else {
			return errors.New("Article Not Exists")
		}
	})

	if err != nil {
		reponse := ErrorResponse{err.Error()}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	var comments Comments
	var comment Comment
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Comment"))
		if b != nil {
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				err = json.Unmarshal(v, &comment)
				if err != nil {
					return err
				}
				if comment.ArticleId == Id {
					comments.Content = append(comments.Content, comment)
				}
			}

			return nil
		} else {
			return errors.New("Comment Not Exists")
		}
	})

	if err != nil {
		reponse := ErrorResponse{err.Error()}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}

	JsonResponse(comments, w, http.StatusOK)
}*/

func GetCommentsOfArticle(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "testuser:123@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()

	articleId := strings.Split(r.URL.Path, "/")[3]
	_, err = strconv.Atoi(articleId)
	if err != nil {
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
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

	query, err = db.Query("select * from test.Comment where articleId=" + articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}

	var comments Comments
	v = []byte("{\"content\":" + string(v) + "}")
	str := strings.Replace(string(v), "Id\":\"", "Id\":", -1)
	str = strings.Replace(str, "\",\"author", ",\"author", -1)
	v = []byte(str)

	_ = json.Unmarshal(v, &comments)
	JsonResponse(comments, w, http.StatusOK)
}
