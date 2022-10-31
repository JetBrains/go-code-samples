package main

import "database/sql"

func connect() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	create table if not exists articles (id integer not null primary key autoincrement, title text, content text);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dbCreateArticle(article *Article) error {
	query, err := db.Prepare("insert into articles(title,content) values (?,?)")
	defer query.Close()

	if err != nil {
		return err
	}
	_, err = query.Exec(article.Title, article.Content)

	if err != nil {
		return err
	}

	return nil
}

func dbGetAllArticles() ([]*Article, error) {
	query, err := db.Prepare("select id, title, content from articles")
	defer query.Close()

	if err != nil {
		return nil, err
	}
	result, err := query.Query()

	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)
	for result.Next() {
		data := new(Article)
		err := result.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, data)
	}

	return articles, nil
}

func dbGetArticle(articleID string) (*Article, error) {
	query, err := db.Prepare("select id, title, content from articles where id = ?")
	defer query.Close()

	if err != nil {
		return nil, err
	}
	result := query.QueryRow(articleID)
	data := new(Article)
	err = result.Scan(&data.ID, &data.Title, &data.Content)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func dbUpdateArticle(id string, article *Article) error {
	query, err := db.Prepare("update articles set (title, content) = (?,?) where id=?")
	defer query.Close()

	if err != nil {
		return err
	}
	_, err = query.Exec(article.Title, article.Content, id)

	if err != nil {
		return err
	}

	return nil
}

func dbDeleteArticle(id string) error {
	query, err := db.Prepare("delete from articles where id=?")
	defer query.Close()

	if err != nil {
		return err
	}
	_, err = query.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
