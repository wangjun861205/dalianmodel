package models

import (
	"database/sql"
	"fmt"
	"github.com/wangjun861205/nbmysql"
	"strings"
	"time"
)

var BkDalian *sql.DB

func init() {
	db, err := sql.Open("mysql", "wangjun:Wt20110523@tcp(127.0.0.1:12345)/bk_dalian")
	if err != nil {
		panic(err)
	}
	BkDalian = db
}

var AuthMap = map[string]string{
	"@Id":            "`id`",
	"@Username":      "`username`",
	"@Password":      "`password`",
	"@Phone":         "`phone`",
	"@Status":        "`status`",
	"@Sessionid":     "`sessionid`",
	"@ExpireTime":    "`expire_time`",
	"@Email":         "`email`",
	"@CreateTime":    "`create_time`",
	"@LastLoginTime": "`last_login_time`",
}

type Auth struct {
	Id            *int64
	Username      *string
	Password      *string
	Phone         *string
	Status        *int64
	Sessionid     *string
	ExpireTime    *time.Time
	Email         *string
	CreateTime    *time.Time
	LastLoginTime *time.Time
}

func NewAuth(authId *nbmysql.Int, authUsername *nbmysql.String, authPassword *nbmysql.String, authPhone *nbmysql.String, authStatus *nbmysql.Int, authSessionid *nbmysql.String, authExpireTime *nbmysql.Time, authEmail *nbmysql.String, authCreateTime *nbmysql.Time, authLastLoginTime *nbmysql.Time) *Auth {
	_id := authId.ToGo()
	_username := authUsername.ToGo()
	_password := authPassword.ToGo()
	_phone := authPhone.ToGo()
	_status := authStatus.ToGo()
	_sessionid := authSessionid.ToGo()
	_expireTime := authExpireTime.ToGo()
	_email := authEmail.ToGo()
	_createTime := authCreateTime.ToGo()
	_lastLoginTime := authLastLoginTime.ToGo()
	auth := &Auth{_id, _username, _password, _phone, _status, _sessionid, _expireTime, _email, _createTime, _lastLoginTime}
	return auth
}
func AllAuth() ([]*Auth, error) {
	rows, err := BkDalian.Query("SELECT * FROM `auth`")
	if err != nil {
		return nil, err
	}
	list := make([]*Auth, 0, 256)
	for rows.Next() {
		model, err := AuthFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func QueryAuth(query string) ([]*Auth, error) {
	for k, v := range AuthMap {
		query = strings.Replace(query, k, v, -1)
	}
	rows, err := BkDalian.Query(fmt.Sprintf("SELECT * FROM `auth` WHERE %s", query))
	if err != nil {
		return nil, err
	}
	list := make([]*Auth, 0, 256)
	for rows.Next() {
		model, err := AuthFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func (m *Auth) Insert() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)
	if m.Id != nil {
		colList = append(colList, "`id`")
		valList = append(valList, fmt.Sprintf("%d", *m.Id))
	}
	if m.Username != nil {
		colList = append(colList, "`username`")
		valList = append(valList, fmt.Sprintf("%q", *m.Username))
	}
	if m.Password != nil {
		colList = append(colList, "`password`")
		valList = append(valList, fmt.Sprintf("%q", *m.Password))
	}
	if m.Phone != nil {
		colList = append(colList, "`phone`")
		valList = append(valList, fmt.Sprintf("%q", *m.Phone))
	}
	if m.Status != nil {
		colList = append(colList, "`status`")
		valList = append(valList, fmt.Sprintf("%d", *m.Status))
	}
	if m.Sessionid != nil {
		colList = append(colList, "`sessionid`")
		valList = append(valList, fmt.Sprintf("%q", *m.Sessionid))
	}
	if m.ExpireTime != nil {
		colList = append(colList, "`expire_time`")
		valList = append(valList, fmt.Sprintf("%q", m.ExpireTime.Format("2006-01-02 15:04:05")))
	}
	if m.Email != nil {
		colList = append(colList, "`email`")
		valList = append(valList, fmt.Sprintf("%q", *m.Email))
	}
	if m.CreateTime != nil {
		colList = append(colList, "`create_time`")
		valList = append(valList, fmt.Sprintf("%q", m.CreateTime.Format("2006-01-02 15:04:05")))
	}
	if m.LastLoginTime != nil {
		colList = append(colList, "`last_login_time`")
		valList = append(valList, fmt.Sprintf("%q", m.LastLoginTime.Format("2006-01-02 15:04:05")))
	}
	res, err := BkDalian.Exec(fmt.Sprintf("INSERT INTO `auth` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
	if err != nil {
		return err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = &lastInsertId
	return nil
}
func (m *Auth) Update() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)

	if m.Username != nil {
		colList = append(colList, "`username`")
		valList = append(valList, fmt.Sprintf("%q", *m.Username))
	}
	if m.Password != nil {
		colList = append(colList, "`password`")
		valList = append(valList, fmt.Sprintf("%q", *m.Password))
	}
	if m.Phone != nil {
		colList = append(colList, "`phone`")
		valList = append(valList, fmt.Sprintf("%q", *m.Phone))
	}
	if m.Status != nil {
		colList = append(colList, "`status`")
		valList = append(valList, fmt.Sprintf("%d", *m.Status))
	}
	if m.Sessionid != nil {
		colList = append(colList, "`sessionid`")
		valList = append(valList, fmt.Sprintf("%q", *m.Sessionid))
	}
	if m.ExpireTime != nil {
		colList = append(colList, "`expire_time`")
		valList = append(valList, fmt.Sprintf("%q", m.ExpireTime.Format("2006-01-02 15:04:05")))
	}
	if m.Email != nil {
		colList = append(colList, "`email`")
		valList = append(valList, fmt.Sprintf("%q", *m.Email))
	}
	if m.CreateTime != nil {
		colList = append(colList, "`create_time`")
		valList = append(valList, fmt.Sprintf("%q", m.CreateTime.Format("2006-01-02 15:04:05")))
	}
	if m.LastLoginTime != nil {
		colList = append(colList, "`last_login_time`")
		valList = append(valList, fmt.Sprintf("%q", m.LastLoginTime.Format("2006-01-02 15:04:05")))
	}
	updateList := make([]string, 0, 32)
	for i := 0; i < len(colList); i++ {
		updateList = append(updateList, fmt.Sprintf("%s=%s", colList[i], valList[i]))
	}
	_, err := BkDalian.Exec(fmt.Sprintf("UPDATE `auth` SET %s WHERE `id` = ?", strings.Join(updateList, ", ")), *m.Id)
	return err
}
func (m *Auth) Delete() error {
	tx, err := BkDalian.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM `auth` WHERE `id` = ?", *m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func AuthFromRows(rows *sql.Rows) (*Auth, error) {
	_id := new(nbmysql.Int)
	_username := new(nbmysql.String)
	_password := new(nbmysql.String)
	_phone := new(nbmysql.String)
	_status := new(nbmysql.Int)
	_sessionid := new(nbmysql.String)
	_expireTime := new(nbmysql.Time)
	_email := new(nbmysql.String)
	_createTime := new(nbmysql.Time)
	_lastLoginTime := new(nbmysql.Time)
	err := rows.Scan(_id, _username, _password, _phone, _status, _sessionid, _expireTime, _email, _createTime, _lastLoginTime)
	if err != nil {
		return nil, err
	}
	return NewAuth(_id, _username, _password, _phone, _status, _sessionid, _expireTime, _email, _createTime, _lastLoginTime), nil
}

var BookMap = map[string]string{
	"@Id":         "`id`",
	"@Isbn":       "`isbn`",
	"@Volume":     "`volume`",
	"@UniqueCode": "`unique_code`",
}

type Book struct {
	Id         *int64
	Isbn       *string
	Volume     *int64
	UniqueCode *string
}
type BookToBookInfo struct {
	All    func() ([]*BookInfo, error)
	Filter func(query string) ([]*BookInfo, error)
	Insert func(bookInfo *BookInfo) error
}

func (m *Book) BookInfoByIsbn() BookToBookInfo {
	return BookToBookInfo{
		All: func() ([]*BookInfo, error) {
			rows, err := BkDalian.Query("SELECT `book_info`.* FROM `book` JOIN `book_info` ON `book`.`isbn`=`book_info`.`isbn` where `book`.`isbn` = ?", *m.Isbn)
			if err != nil {
				return nil, err
			}
			list := make([]*BookInfo, 0, 256)
			for rows.Next() {
				model, err := BookInfoFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Filter: func(query string) ([]*BookInfo, error) {
			for k, v := range BookInfoMap {
				query = strings.Replace(query, k, v, -1)
			}
			rows, err := BkDalian.Query("SELECT `book_info`.* FROM `book` JOIN `book_info` ON `book`.`isbn`=`book_info`.`isbn` where `book`.`isbn` = ? AND ?", *m.Isbn, query)
			if err != nil {
				return nil, err
			}
			list := make([]*BookInfo, 0, 256)
			for rows.Next() {
				model, err := BookInfoFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Insert: func(bookInfo *BookInfo) error {
			tx, err := BkDalian.Begin()
			if err != nil {
				return err
			}
			colList := make([]string, 0, 32)
			valList := make([]string, 0, 32)
			if bookInfo.Id != nil {
				colList = append(colList, "`id`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Id))
			}
			if bookInfo.Title != nil {
				colList = append(colList, "`title`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Title))
			}
			if bookInfo.Price != nil {
				colList = append(colList, "`price`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Price))
			}
			if bookInfo.Author != nil {
				colList = append(colList, "`author`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Author))
			}
			if bookInfo.Publisher != nil {
				colList = append(colList, "`publisher`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Publisher))
			}
			if bookInfo.Series != nil {
				colList = append(colList, "`series`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Series))
			}
			if bookInfo.Isbn != nil {
				colList = append(colList, "`isbn`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Isbn))
			}
			if bookInfo.PublishDate != nil {
				colList = append(colList, "`publish_date`")
				valList = append(valList, fmt.Sprintf("%q", bookInfo.PublishDate.Format("2006-01-02 15:04:05")))
			}
			if bookInfo.Binding != nil {
				colList = append(colList, "`binding`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Binding))
			}
			if bookInfo.Format != nil {
				colList = append(colList, "`format`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Format))
			}
			if bookInfo.Pages != nil {
				colList = append(colList, "`pages`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Pages))
			}
			if bookInfo.WordCount != nil {
				colList = append(colList, "`word_count`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.WordCount))
			}
			if bookInfo.ContentIntro != nil {
				colList = append(colList, "`content_intro`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.ContentIntro))
			}
			if bookInfo.AuthorIntro != nil {
				colList = append(colList, "`author_intro`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.AuthorIntro))
			}
			if bookInfo.Menu != nil {
				colList = append(colList, "`menu`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Menu))
			}
			res, err := tx.Exec("INSERT INTO `book_info` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", "))
			if err != nil {
				tx.Rollback()
				return err
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				return err
			}
			bookInfo.Id = &lastInsertId
			return tx.Commit()
		},
	}
}
func NewBook(bookId *nbmysql.Int, bookIsbn *nbmysql.String, bookVolume *nbmysql.Int, bookUniqueCode *nbmysql.String) *Book {
	_id := bookId.ToGo()
	_isbn := bookIsbn.ToGo()
	_volume := bookVolume.ToGo()
	_uniqueCode := bookUniqueCode.ToGo()
	book := &Book{_id, _isbn, _volume, _uniqueCode}
	return book
}
func AllBook() ([]*Book, error) {
	rows, err := BkDalian.Query("SELECT * FROM `book`")
	if err != nil {
		return nil, err
	}
	list := make([]*Book, 0, 256)
	for rows.Next() {
		model, err := BookFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func QueryBook(query string) ([]*Book, error) {
	for k, v := range BookMap {
		query = strings.Replace(query, k, v, -1)
	}
	rows, err := BkDalian.Query(fmt.Sprintf("SELECT * FROM `book` WHERE %s", query))
	if err != nil {
		return nil, err
	}
	list := make([]*Book, 0, 256)
	for rows.Next() {
		model, err := BookFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func (m *Book) Insert() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)
	if m.Id != nil {
		colList = append(colList, "`id`")
		valList = append(valList, fmt.Sprintf("%d", *m.Id))
	}
	if m.Isbn != nil {
		colList = append(colList, "`isbn`")
		valList = append(valList, fmt.Sprintf("%q", *m.Isbn))
	}
	if m.Volume != nil {
		colList = append(colList, "`volume`")
		valList = append(valList, fmt.Sprintf("%d", *m.Volume))
	}
	if m.UniqueCode != nil {
		colList = append(colList, "`unique_code`")
		valList = append(valList, fmt.Sprintf("%q", *m.UniqueCode))
	}
	res, err := BkDalian.Exec(fmt.Sprintf("INSERT INTO `book` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
	if err != nil {
		return err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = &lastInsertId
	return nil
}
func (m *Book) Update() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)

	if m.Isbn != nil {
		colList = append(colList, "`isbn`")
		valList = append(valList, fmt.Sprintf("%q", *m.Isbn))
	}
	if m.Volume != nil {
		colList = append(colList, "`volume`")
		valList = append(valList, fmt.Sprintf("%d", *m.Volume))
	}
	if m.UniqueCode != nil {
		colList = append(colList, "`unique_code`")
		valList = append(valList, fmt.Sprintf("%q", *m.UniqueCode))
	}
	updateList := make([]string, 0, 32)
	for i := 0; i < len(colList); i++ {
		updateList = append(updateList, fmt.Sprintf("%s=%s", colList[i], valList[i]))
	}
	_, err := BkDalian.Exec(fmt.Sprintf("UPDATE `book` SET %s WHERE `id` = ?", strings.Join(updateList, ", ")), *m.Id)
	return err
}
func (m *Book) Delete() error {
	tx, err := BkDalian.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM `book` WHERE `id` = ?", *m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func BookFromRows(rows *sql.Rows) (*Book, error) {
	_id := new(nbmysql.Int)
	_isbn := new(nbmysql.String)
	_volume := new(nbmysql.Int)
	_uniqueCode := new(nbmysql.String)
	err := rows.Scan(_id, _isbn, _volume, _uniqueCode)
	if err != nil {
		return nil, err
	}
	return NewBook(_id, _isbn, _volume, _uniqueCode), nil
}

var BookInfoMap = map[string]string{
	"@Id":           "`id`",
	"@Title":        "`title`",
	"@Price":        "`price`",
	"@Author":       "`author`",
	"@Publisher":    "`publisher`",
	"@Series":       "`series`",
	"@Isbn":         "`isbn`",
	"@PublishDate":  "`publish_date`",
	"@Binding":      "`binding`",
	"@Format":       "`format`",
	"@Pages":        "`pages`",
	"@WordCount":    "`word_count`",
	"@ContentIntro": "`content_intro`",
	"@AuthorIntro":  "`author_intro`",
	"@Menu":         "`menu`",
}

type BookInfo struct {
	Id           *int64
	Title        *string
	Price        *int64
	Author       *string
	Publisher    *string
	Series       *string
	Isbn         *string
	PublishDate  *time.Time
	Binding      *string
	Format       *string
	Pages        *int64
	WordCount    *int64
	ContentIntro *string
	AuthorIntro  *string
	Menu         *string
}
type BookInfoToTag struct {
	All    func() ([]*Tag, error)
	Filter func(query string) ([]*Tag, error)
	Insert func(tag *Tag) error
}

func (m *BookInfo) TagByIsbn() BookInfoToTag {
	return BookInfoToTag{
		All: func() ([]*Tag, error) {
			rows, err := BkDalian.Query("SELECT `tag`.* FROM `book_info` JOIN `book_info__tag` ON `book_info`.`isbn`=`book_info__tag`.`book_info__isbn` JOIN `tag` on `book_info__tag`.`tag__id` = `tag`.`id` WHERE `book_info`.`isbn` = ?", *m.Isbn)
			if err != nil {
				return nil, err
			}
			list := make([]*Tag, 0, 256)
			for rows.Next() {
				model, err := TagFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Filter: func(query string) ([]*Tag, error) {
			for k, v := range TagMap {
				query = strings.Replace(query, k, v, -1)
			}
			rows, err := BkDalian.Query("SELECT `tag`.* FROM `book_info` JOIN `book_info__tag` ON `book_info`.`isbn`=`book_info__tag`.`book_info__isbn` JOIN `tag` on `book_info__tag`.`tag__id` = `tag`.`id` WHERE `book_info`.`isbn` = ? AND ?", *m.Isbn, query)
			if err != nil {
				return nil, err
			}
			list := make([]*Tag, 0, 256)
			for rows.Next() {
				model, err := TagFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Insert: func(tag *Tag) error {
			tx, err := BkDalian.Begin()
			if err != nil {
				return err
			}
			colList := make([]string, 0, 32)
			valList := make([]string, 0, 32)
			if tag.Id != nil {
				colList = append(colList, "`id`")
				valList = append(valList, fmt.Sprintf("%d", *tag.Id))
			}
			if tag.Tag != nil {
				colList = append(colList, "`tag`")
				valList = append(valList, fmt.Sprintf("%q", *tag.Tag))
			}
			res, err := tx.Exec(fmt.Sprintf("INSERT INTO `tag` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
			if err != nil {
				tx.Rollback()
				return err
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				return err
			}
			tag.Id = &lastInsertId
			_, err = tx.Exec("INSERT INTO `book_info__tag` (`book_info__isbn`, `tag__id`) VALUES (?, ?)", *m.Isbn, *tag.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
			return tx.Commit()
		},
	}
}
func NewBookInfo(bookInfoId *nbmysql.Int, bookInfoTitle *nbmysql.String, bookInfoPrice *nbmysql.Int, bookInfoAuthor *nbmysql.String, bookInfoPublisher *nbmysql.String, bookInfoSeries *nbmysql.String, bookInfoIsbn *nbmysql.String, bookInfoPublishDate *nbmysql.Time, bookInfoBinding *nbmysql.String, bookInfoFormat *nbmysql.String, bookInfoPages *nbmysql.Int, bookInfoWordCount *nbmysql.Int, bookInfoContentIntro *nbmysql.String, bookInfoAuthorIntro *nbmysql.String, bookInfoMenu *nbmysql.String) *BookInfo {
	_id := bookInfoId.ToGo()
	_title := bookInfoTitle.ToGo()
	_price := bookInfoPrice.ToGo()
	_author := bookInfoAuthor.ToGo()
	_publisher := bookInfoPublisher.ToGo()
	_series := bookInfoSeries.ToGo()
	_isbn := bookInfoIsbn.ToGo()
	_publishDate := bookInfoPublishDate.ToGo()
	_binding := bookInfoBinding.ToGo()
	_format := bookInfoFormat.ToGo()
	_pages := bookInfoPages.ToGo()
	_wordCount := bookInfoWordCount.ToGo()
	_contentIntro := bookInfoContentIntro.ToGo()
	_authorIntro := bookInfoAuthorIntro.ToGo()
	_menu := bookInfoMenu.ToGo()
	bookInfo := &BookInfo{_id, _title, _price, _author, _publisher, _series, _isbn, _publishDate, _binding, _format, _pages, _wordCount, _contentIntro, _authorIntro, _menu}
	return bookInfo
}
func AllBookInfo() ([]*BookInfo, error) {
	rows, err := BkDalian.Query("SELECT * FROM `book_info`")
	if err != nil {
		return nil, err
	}
	list := make([]*BookInfo, 0, 256)
	for rows.Next() {
		model, err := BookInfoFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func QueryBookInfo(query string) ([]*BookInfo, error) {
	for k, v := range BookInfoMap {
		query = strings.Replace(query, k, v, -1)
	}
	rows, err := BkDalian.Query(fmt.Sprintf("SELECT * FROM `book_info` WHERE %s", query))
	if err != nil {
		return nil, err
	}
	list := make([]*BookInfo, 0, 256)
	for rows.Next() {
		model, err := BookInfoFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func (m *BookInfo) Insert() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)
	if m.Id != nil {
		colList = append(colList, "`id`")
		valList = append(valList, fmt.Sprintf("%d", *m.Id))
	}
	if m.Title != nil {
		colList = append(colList, "`title`")
		valList = append(valList, fmt.Sprintf("%q", *m.Title))
	}
	if m.Price != nil {
		colList = append(colList, "`price`")
		valList = append(valList, fmt.Sprintf("%d", *m.Price))
	}
	if m.Author != nil {
		colList = append(colList, "`author`")
		valList = append(valList, fmt.Sprintf("%q", *m.Author))
	}
	if m.Publisher != nil {
		colList = append(colList, "`publisher`")
		valList = append(valList, fmt.Sprintf("%q", *m.Publisher))
	}
	if m.Series != nil {
		colList = append(colList, "`series`")
		valList = append(valList, fmt.Sprintf("%q", *m.Series))
	}
	if m.Isbn != nil {
		colList = append(colList, "`isbn`")
		valList = append(valList, fmt.Sprintf("%q", *m.Isbn))
	}
	if m.PublishDate != nil {
		colList = append(colList, "`publish_date`")
		valList = append(valList, fmt.Sprintf("%q", m.PublishDate.Format("2006-01-02 15:04:05")))
	}
	if m.Binding != nil {
		colList = append(colList, "`binding`")
		valList = append(valList, fmt.Sprintf("%q", *m.Binding))
	}
	if m.Format != nil {
		colList = append(colList, "`format`")
		valList = append(valList, fmt.Sprintf("%q", *m.Format))
	}
	if m.Pages != nil {
		colList = append(colList, "`pages`")
		valList = append(valList, fmt.Sprintf("%d", *m.Pages))
	}
	if m.WordCount != nil {
		colList = append(colList, "`word_count`")
		valList = append(valList, fmt.Sprintf("%d", *m.WordCount))
	}
	if m.ContentIntro != nil {
		colList = append(colList, "`content_intro`")
		valList = append(valList, fmt.Sprintf("%q", *m.ContentIntro))
	}
	if m.AuthorIntro != nil {
		colList = append(colList, "`author_intro`")
		valList = append(valList, fmt.Sprintf("%q", *m.AuthorIntro))
	}
	if m.Menu != nil {
		colList = append(colList, "`menu`")
		valList = append(valList, fmt.Sprintf("%q", *m.Menu))
	}
	res, err := BkDalian.Exec(fmt.Sprintf("INSERT INTO `book_info` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
	if err != nil {
		return err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = &lastInsertId
	return nil
}
func (m *BookInfo) Update() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)

	if m.Title != nil {
		colList = append(colList, "`title`")
		valList = append(valList, fmt.Sprintf("%q", *m.Title))
	}
	if m.Price != nil {
		colList = append(colList, "`price`")
		valList = append(valList, fmt.Sprintf("%d", *m.Price))
	}
	if m.Author != nil {
		colList = append(colList, "`author`")
		valList = append(valList, fmt.Sprintf("%q", *m.Author))
	}
	if m.Publisher != nil {
		colList = append(colList, "`publisher`")
		valList = append(valList, fmt.Sprintf("%q", *m.Publisher))
	}
	if m.Series != nil {
		colList = append(colList, "`series`")
		valList = append(valList, fmt.Sprintf("%q", *m.Series))
	}
	if m.Isbn != nil {
		colList = append(colList, "`isbn`")
		valList = append(valList, fmt.Sprintf("%q", *m.Isbn))
	}
	if m.PublishDate != nil {
		colList = append(colList, "`publish_date`")
		valList = append(valList, fmt.Sprintf("%q", m.PublishDate.Format("2006-01-02 15:04:05")))
	}
	if m.Binding != nil {
		colList = append(colList, "`binding`")
		valList = append(valList, fmt.Sprintf("%q", *m.Binding))
	}
	if m.Format != nil {
		colList = append(colList, "`format`")
		valList = append(valList, fmt.Sprintf("%q", *m.Format))
	}
	if m.Pages != nil {
		colList = append(colList, "`pages`")
		valList = append(valList, fmt.Sprintf("%d", *m.Pages))
	}
	if m.WordCount != nil {
		colList = append(colList, "`word_count`")
		valList = append(valList, fmt.Sprintf("%d", *m.WordCount))
	}
	if m.ContentIntro != nil {
		colList = append(colList, "`content_intro`")
		valList = append(valList, fmt.Sprintf("%q", *m.ContentIntro))
	}
	if m.AuthorIntro != nil {
		colList = append(colList, "`author_intro`")
		valList = append(valList, fmt.Sprintf("%q", *m.AuthorIntro))
	}
	if m.Menu != nil {
		colList = append(colList, "`menu`")
		valList = append(valList, fmt.Sprintf("%q", *m.Menu))
	}
	updateList := make([]string, 0, 32)
	for i := 0; i < len(colList); i++ {
		updateList = append(updateList, fmt.Sprintf("%s=%s", colList[i], valList[i]))
	}
	_, err := BkDalian.Exec(fmt.Sprintf("UPDATE `book_info` SET %s WHERE `id` = ?", strings.Join(updateList, ", ")), *m.Id)
	return err
}
func (m *BookInfo) Delete() error {
	tx, err := BkDalian.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM `book_info__tag` WHERE `book_info__isbn` = ?", *m.Isbn)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM `book_info` WHERE `id` = ?", *m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func BookInfoFromRows(rows *sql.Rows) (*BookInfo, error) {
	_id := new(nbmysql.Int)
	_title := new(nbmysql.String)
	_price := new(nbmysql.Int)
	_author := new(nbmysql.String)
	_publisher := new(nbmysql.String)
	_series := new(nbmysql.String)
	_isbn := new(nbmysql.String)
	_publishDate := new(nbmysql.Time)
	_binding := new(nbmysql.String)
	_format := new(nbmysql.String)
	_pages := new(nbmysql.Int)
	_wordCount := new(nbmysql.Int)
	_contentIntro := new(nbmysql.String)
	_authorIntro := new(nbmysql.String)
	_menu := new(nbmysql.String)
	err := rows.Scan(_id, _title, _price, _author, _publisher, _series, _isbn, _publishDate, _binding, _format, _pages, _wordCount, _contentIntro, _authorIntro, _menu)
	if err != nil {
		return nil, err
	}
	return NewBookInfo(_id, _title, _price, _author, _publisher, _series, _isbn, _publishDate, _binding, _format, _pages, _wordCount, _contentIntro, _authorIntro, _menu), nil
}

var TagMap = map[string]string{
	"@Id":  "`id`",
	"@Tag": "`tag`",
}

type Tag struct {
	Id  *int64
	Tag *string
}
type TagToBookInfo struct {
	All    func() ([]*BookInfo, error)
	Filter func(query string) ([]*BookInfo, error)
	Insert func(bookInfo *BookInfo) error
}

func (m *Tag) BookInfoById() TagToBookInfo {
	return TagToBookInfo{
		All: func() ([]*BookInfo, error) {
			rows, err := BkDalian.Query("SELECT `book_info`.* FROM `tag` JOIN `book_info__tag` ON `tag`.`id`=`book_info__tag`.`tag__id` JOIN `book_info` on `book_info__tag`.`book_info__isbn` = `book_info`.`isbn` WHERE `tag`.`id` = ?", *m.Id)
			if err != nil {
				return nil, err
			}
			list := make([]*BookInfo, 0, 256)
			for rows.Next() {
				model, err := BookInfoFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Filter: func(query string) ([]*BookInfo, error) {
			for k, v := range BookInfoMap {
				query = strings.Replace(query, k, v, -1)
			}
			rows, err := BkDalian.Query("SELECT `book_info`.* FROM `tag` JOIN `book_info__tag` ON `tag`.`id`=`book_info__tag`.`tag__id` JOIN `book_info` on `book_info__tag`.`book_info__isbn` = `book_info`.`isbn` WHERE `tag`.`id` = ? AND ?", *m.Id, query)
			if err != nil {
				return nil, err
			}
			list := make([]*BookInfo, 0, 256)
			for rows.Next() {
				model, err := BookInfoFromRows(rows)
				if err != nil {
					return nil, err
				}
				list = append(list, model)
			}
			return list, nil
		},
		Insert: func(bookInfo *BookInfo) error {
			tx, err := BkDalian.Begin()
			if err != nil {
				return err
			}
			colList := make([]string, 0, 32)
			valList := make([]string, 0, 32)
			if bookInfo.Id != nil {
				colList = append(colList, "`id`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Id))
			}
			if bookInfo.Title != nil {
				colList = append(colList, "`title`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Title))
			}
			if bookInfo.Price != nil {
				colList = append(colList, "`price`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Price))
			}
			if bookInfo.Author != nil {
				colList = append(colList, "`author`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Author))
			}
			if bookInfo.Publisher != nil {
				colList = append(colList, "`publisher`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Publisher))
			}
			if bookInfo.Series != nil {
				colList = append(colList, "`series`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Series))
			}
			if bookInfo.Isbn != nil {
				colList = append(colList, "`isbn`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Isbn))
			}
			if bookInfo.PublishDate != nil {
				colList = append(colList, "`publish_date`")
				valList = append(valList, fmt.Sprintf("%q", bookInfo.PublishDate.Format("2006-01-02 15:04:05")))
			}
			if bookInfo.Binding != nil {
				colList = append(colList, "`binding`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Binding))
			}
			if bookInfo.Format != nil {
				colList = append(colList, "`format`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Format))
			}
			if bookInfo.Pages != nil {
				colList = append(colList, "`pages`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.Pages))
			}
			if bookInfo.WordCount != nil {
				colList = append(colList, "`word_count`")
				valList = append(valList, fmt.Sprintf("%d", *bookInfo.WordCount))
			}
			if bookInfo.ContentIntro != nil {
				colList = append(colList, "`content_intro`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.ContentIntro))
			}
			if bookInfo.AuthorIntro != nil {
				colList = append(colList, "`author_intro`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.AuthorIntro))
			}
			if bookInfo.Menu != nil {
				colList = append(colList, "`menu`")
				valList = append(valList, fmt.Sprintf("%q", *bookInfo.Menu))
			}
			res, err := tx.Exec(fmt.Sprintf("INSERT INTO `book_info` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
			if err != nil {
				tx.Rollback()
				return err
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				tx.Rollback()
				return err
			}
			bookInfo.Id = &lastInsertId
			_, err = tx.Exec("INSERT INTO `book_info__tag` (`tag__id`, `book_info__isbn`) VALUES (?, ?)", *m.Id, *bookInfo.Isbn)
			if err != nil {
				tx.Rollback()
				return err
			}
			return tx.Commit()
		},
	}
}
func NewTag(tagId *nbmysql.Int, tagTag *nbmysql.String) *Tag {
	_id := tagId.ToGo()
	_tag := tagTag.ToGo()
	tag := &Tag{_id, _tag}
	return tag
}
func AllTag() ([]*Tag, error) {
	rows, err := BkDalian.Query("SELECT * FROM `tag`")
	if err != nil {
		return nil, err
	}
	list := make([]*Tag, 0, 256)
	for rows.Next() {
		model, err := TagFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func QueryTag(query string) ([]*Tag, error) {
	for k, v := range TagMap {
		query = strings.Replace(query, k, v, -1)
	}
	rows, err := BkDalian.Query(fmt.Sprintf("SELECT * FROM `tag` WHERE %s", query))
	if err != nil {
		return nil, err
	}
	list := make([]*Tag, 0, 256)
	for rows.Next() {
		model, err := TagFromRows(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, model)
	}
	return list, nil
}
func (m *Tag) Insert() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)
	if m.Id != nil {
		colList = append(colList, "`id`")
		valList = append(valList, fmt.Sprintf("%d", *m.Id))
	}
	if m.Tag != nil {
		colList = append(colList, "`tag`")
		valList = append(valList, fmt.Sprintf("%q", *m.Tag))
	}
	res, err := BkDalian.Exec(fmt.Sprintf("INSERT INTO `tag` (%s) VALUES (%s)", strings.Join(colList, ", "), strings.Join(valList, ", ")))
	if err != nil {
		return err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.Id = &lastInsertId
	return nil
}
func (m *Tag) Update() error {
	colList := make([]string, 0, 32)
	valList := make([]string, 0, 32)

	if m.Tag != nil {
		colList = append(colList, "`tag`")
		valList = append(valList, fmt.Sprintf("%q", *m.Tag))
	}
	updateList := make([]string, 0, 32)
	for i := 0; i < len(colList); i++ {
		updateList = append(updateList, fmt.Sprintf("%s=%s", colList[i], valList[i]))
	}
	_, err := BkDalian.Exec(fmt.Sprintf("UPDATE `tag` SET %s WHERE `id` = ?", strings.Join(updateList, ", ")), *m.Id)
	return err
}
func (m *Tag) Delete() error {
	tx, err := BkDalian.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM `book_info__tag` WHERE `tag__id` = ?", *m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM `tag` WHERE `id` = ?", *m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func TagFromRows(rows *sql.Rows) (*Tag, error) {
	_id := new(nbmysql.Int)
	_tag := new(nbmysql.String)
	err := rows.Scan(_id, _tag)
	if err != nil {
		return nil, err
	}
	return NewTag(_id, _tag), nil
}
