package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	c map[string]string
	m sync.Mutex
}

func NewCache() *Cache {
	c := make(map[string]string)
	return &Cache{c: c}
}
func (c *Cache) GetToken(token string) (string, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	value, found := c.c[token]
	if !found {
		return "", found
	}
	return value, found
}
func (c *Cache) NewToken(Token string, ID string) error {
	c.m.Lock()
	defer c.m.Unlock()
	_, found := c.c[Token]
	if found {
		return errors.New("User exists")
	}
	c.c[Token] = ID
	return nil
}
func (c *Cache) DeleteToken(token string) error {
	c.m.Lock()
	defer c.m.Unlock()

	delete(c.c, token)
	return nil
}

//func (c *Cache) UpdateTransaction(key string, status int) error {
//	value, found := c.GetTransaction(key)
//	if !found {
//		return errors.New("Transaction not found")
//	}
//	value.Status = status
//	err := c.DeleteTranscation(key)
//	if err != nil {
//		return err
//	}
//	err = c.NewTranscation(value)
//	if err != nil {
//		return err
//	}
//	return nil
//}
