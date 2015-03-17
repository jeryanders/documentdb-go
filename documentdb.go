//
// This project start as a fork of `github.com/nerdylikeme/go-documentdb` version
// but changed a bit, and may be changed later
//
// Goal: add the full functionality of documentdb, align with the other sdks
// and make it more testable
//
package documentdb

type Config struct {
	MasterKey	string
}

type DocumentDB struct {
	client	Clienter
}

// Create DocumentDBClient
func New(url string, config Config) *DocumentDB {
	client := &Client{}
	client.Url = url
	client.Config = config
	return &DocumentDB{client}
}

// TODO: Add `requestOptions` arguments
// Read database by self link
func (c *DocumentDB) ReadDatabase(link string) (db *Database, err error) {
	err = c.client.Read(link, &db)
	if err != nil {
		return nil, err
	}
	return
}

// Read collection by self link
func (c *DocumentDB) ReadCollection(link string) (coll *Collection, err error) {
	err = c.client.Read(link, &coll)
	if err != nil {
		return nil, err
	}
	return
}

// Read document by self link
func (c *DocumentDB) ReadDocument(link string, doc interface{}) error {
	err := c.client.Read(link, &doc)
	return err
}

// Read sporc by self link
func (c *DocumentDB) ReadStoredProcedure(link string) (sporc *Sporc, err error) {
	err = c.client.Read(link, &sporc)
	if err != nil {
		return nil, err
	}
	return
}

// Read udf by self link
func (c *DocumentDB) ReadUserDefinedFunction(link string) (udf *UDF, err error) {
	err = c.client.Read(link, &udf)
	if err != nil {
		return nil, err
	}
	return
}