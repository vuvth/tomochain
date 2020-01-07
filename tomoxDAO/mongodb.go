package tomoxDAO

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	lru "github.com/hashicorp/golang-lru"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/ethdb"
	"github.com/tomochain/tomochain/log"
	"github.com/tomochain/tomochain/tomox/tradingstate"
	"github.com/tomochain/tomochain/tomoxlending/lendingstate"
	"strings"
	"time"
)

const (
	ordersCollection        = "orders"
	tradesCollection        = "trades"
	lendingItemsCollection  = "lending_items"
	lendingTradesCollection = "lending_trades"
)

type MongoDatabase struct {
	Session          *mgo.Session
	dbName           string
	emptyKey         []byte
	cacheItems       *lru.Cache // Cache for reading
	orderBulk        *mgo.Bulk
	tradeBulk        *mgo.Bulk
	lendingItemBulk  *mgo.Bulk
	lendingTradeBulk *mgo.Bulk
}

// InitSession initializes a new session with mongodb
func NewMongoDatabase(session *mgo.Session, dbName string, mongoURL string, replicaSetName string, cacheLimit int) (*MongoDatabase, error) {
	if session == nil {
		// in case of multiple database instances
		hosts := strings.Split(mongoURL, ",")
		dbInfo := &mgo.DialInfo{
			Addrs:          hosts,
			Database:       dbName,
			ReplicaSetName: replicaSetName,
			Timeout:        30 * time.Second,
		}
		ns, err := mgo.DialWithInfo(dbInfo)
		if err != nil {
			return nil, err
		}
		session = ns
	}
	itemCacheLimit := defaultCacheLimit
	if cacheLimit > 0 {
		itemCacheLimit = cacheLimit
	}
	cacheItems, _ := lru.New(itemCacheLimit)

	db := &MongoDatabase{
		Session:    session,
		dbName:     dbName,
		cacheItems: cacheItems,
	}
	if err := db.EnsureIndexes(); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *MongoDatabase) IsEmptyKey(key []byte) bool {
	return key == nil || len(key) == 0 || bytes.Equal(key, db.emptyKey)
}

func (db *MongoDatabase) getCacheKey(key []byte) string {
	return hex.EncodeToString(key)
}

func (db *MongoDatabase) HasObject(hash common.Hash, val interface{}) (bool, error) {
	if db.IsEmptyKey(hash.Bytes()) {
		return false, nil
	}
	cacheKey := db.getCacheKey(hash.Bytes())
	if db.cacheItems.Contains(cacheKey) {
		return true, nil
	}

	sc := db.Session.Copy()
	defer sc.Close()
	var (
		count int
		err   error
	)
	query := bson.M{"hash": hash.Hex()}
	switch val.(type) {
	case *tradingstate.OrderItem:
		// Find key in ordersCollection collection
		count, err = sc.DB(db.dbName).C(ordersCollection).Find(query).Limit(1).Count()

		if err != nil {
			return false, err
		}

		if count == 1 {
			return true, nil
		}
	case *tradingstate.Trade:
		// Find key in tradesCollection collection
		count, err = sc.DB(db.dbName).C(tradesCollection).Find(query).Limit(1).Count()

		if err != nil {
			return false, err
		}

		if count == 1 {
			return true, nil
		}
	case *lendingstate.LendingItem:
		// Find key in lendingItemsCollection collection
		count, err = sc.DB(db.dbName).C(lendingItemsCollection).Find(query).Limit(1).Count()

		if err != nil {
			return false, err
		}

		if count == 1 {
			return true, nil
		}
	case *lendingstate.LendingTrade:
		// Find key in lendingTradesCollection collection
		count, err = sc.DB(db.dbName).C(lendingTradesCollection).Find(query).Limit(1).Count()

		if err != nil {
			return false, err
		}

		if count == 1 {
			return true, nil
		}

	}
	return false, nil
}

func (db *MongoDatabase) GetObject(hash common.Hash, val interface{}) (interface{}, error) {

	if db.IsEmptyKey(hash.Bytes()) {
		return nil, nil
	}

	cacheKey := db.getCacheKey(hash.Bytes())
	if cached, ok := db.cacheItems.Get(cacheKey); ok {
		return cached, nil
	} else {
		sc := db.Session.Copy()
		defer sc.Close()

		query := bson.M{"hash": hash.Hex()}

		switch val.(type) {
		case *tradingstate.OrderItem:
			var oi *tradingstate.OrderItem
			err := sc.DB(db.dbName).C(ordersCollection).Find(query).One(&oi)
			if err != nil {
				return nil, err
			}
			db.cacheItems.Add(cacheKey, oi)
			return oi, nil
		case *tradingstate.Trade:
			var t *tradingstate.Trade
			err := sc.DB(db.dbName).C(tradesCollection).Find(query).One(&t)
			if err != nil {
				return nil, err
			}
			db.cacheItems.Add(cacheKey, t)
			return t, nil
		case *lendingstate.LendingItem:
			var li *lendingstate.LendingItem
			err := sc.DB(db.dbName).C(ordersCollection).Find(query).One(&li)
			if err != nil {
				return nil, err
			}
			db.cacheItems.Add(cacheKey, li)
			return li, nil
		case *lendingstate.LendingTrade:
			var t *lendingstate.LendingTrade
			err := sc.DB(db.dbName).C(tradesCollection).Find(query).One(&t)
			if err != nil {
				return nil, err
			}
			db.cacheItems.Add(cacheKey, t)
			return t, nil
		default:
			return nil, nil
		}
	}
}

func (db *MongoDatabase) PutObject(hash common.Hash, val interface{}) error {
	cacheKey := db.getCacheKey(hash.Bytes())
	db.cacheItems.Add(cacheKey, val)

	switch val.(type) {
	case *tradingstate.Trade:
		// PutObject trade into tradesCollection collection
		db.tradeBulk.Insert(val.(*tradingstate.Trade))
	case *tradingstate.OrderItem:
		// PutObject order into ordersCollection collection
		o := val.(*tradingstate.OrderItem)
		if o.Status == tradingstate.OrderStatusOpen {
			db.orderBulk.Insert(o)
		} else {
			query := bson.M{"hash": o.Hash.Hex()}
			db.orderBulk.Upsert(query, o)
		}
		return nil
	case *lendingstate.LendingTrade:
		lt := val.(*lendingstate.LendingTrade)
		// PutObject LendingTrade into tradesCollection collection
		if existed, err := db.HasObject(hash, val); err == nil && existed {
			query := bson.M{"hash": lt.Hash.Hex()}
			db.lendingTradeBulk.Upsert(query, lt)
		} else {
			db.lendingTradeBulk.Insert(lt)
		}
	case *lendingstate.LendingItem:
		// PutObject order into ordersCollection collection
		li := val.(*lendingstate.LendingItem)
		if li.Status == lendingstate.LendingStatusOpen {
			db.lendingItemBulk.Insert(li)
		} else {
			query := bson.M{"hash": li.Hash.Hex()}
			db.lendingItemBulk.Upsert(query, li)
		}
		return nil
	default:
		log.Error("PutObject: unknown type of object", "val", val)
	}

	return nil
}

func (db *MongoDatabase) DeleteObject(hash common.Hash, val interface{}) error {
	cacheKey := db.getCacheKey(hash.Bytes())
	db.cacheItems.Remove(cacheKey)

	sc := db.Session.Copy()
	defer sc.Close()

	query := bson.M{"hash": hash.Hex()}

	found, err := db.HasObject(hash, val)
	if err != nil {
		return err
	}

	if found {
		var err error
		switch val.(type) {
		case *tradingstate.OrderItem:
			err = sc.DB(db.dbName).C(ordersCollection).Remove(query)
			if err != nil && err != mgo.ErrNotFound {
				return fmt.Errorf("failed to delete orderItem. Err: %v", err)
			}
		case *tradingstate.Trade:
			err = sc.DB(db.dbName).C(tradesCollection).Remove(query)
			if err != nil && err != mgo.ErrNotFound {
				return fmt.Errorf("failed to delete tomox trade. Err: %v", err)
			}
		case *lendingstate.LendingItem:
			err = sc.DB(db.dbName).C(lendingItemsCollection).Remove(query)
			if err != nil && err != mgo.ErrNotFound {
				return fmt.Errorf("failed to delete lendingItem. Err: %v", err)
			}
		case *lendingstate.LendingTrade:
			err = sc.DB(db.dbName).C(lendingTradesCollection).Remove(query)
			if err != nil && err != mgo.ErrNotFound {
				return fmt.Errorf("failed to delete lendingTrade. Err: %v", err)
			}

		}
	}

	return nil
}

func (db *MongoDatabase) InitBulk() *mgo.Session {
	sc := db.Session.Copy()
	db.orderBulk = sc.DB(db.dbName).C(ordersCollection).Bulk()
	db.tradeBulk = sc.DB(db.dbName).C(tradesCollection).Bulk()
	return sc
}

func (db *MongoDatabase) InitLendingBulk() *mgo.Session {
	sc := db.Session.Copy()
	db.lendingItemBulk = sc.DB(db.dbName).C(lendingItemsCollection).Bulk()
	db.lendingTradeBulk = sc.DB(db.dbName).C(lendingTradesCollection).Bulk()
	return sc
}

func (db *MongoDatabase) CommitBulk() error {
	if _, err := db.orderBulk.Run(); err != nil && !mgo.IsDup(err) {
		return err
	}
	if _, err := db.tradeBulk.Run(); err != nil && !mgo.IsDup(err) {
		return err
	}
	return nil
}

func (db *MongoDatabase) CommitLendingBulk() error {
	if _, err := db.lendingItemBulk.Run(); err != nil && !mgo.IsDup(err) {
		return err
	}
	if _, err := db.lendingTradeBulk.Run(); err != nil && !mgo.IsDup(err) {
		return err
	}
	return nil
}

func (db *MongoDatabase) Put(key []byte, val []byte) error {
	// for levelDB only
	return nil
}

func (db *MongoDatabase) Delete(key []byte) error {
	// for levelDB only
	return nil
}

func (db *MongoDatabase) Has(key []byte) (bool, error) {
	// for levelDB only
	return false, nil
}

func (db *MongoDatabase) Get(key []byte) ([]byte, error) {
	// for levelDB only
	return nil, nil
}

func (db *MongoDatabase) DeleteItemByTxHash(txhash common.Hash, val interface{}) {
	sc := db.Session.Copy()
	defer sc.Close()

	query := bson.M{"txHash": txhash.Hex()}
	switch val.(type) {
	case *tradingstate.OrderItem:
		if err := sc.DB(db.dbName).C(ordersCollection).Remove(query); err != nil && err != mgo.ErrNotFound {
			log.Error("DeleteItemByTxHash: failed to delete order", "txhash", txhash, "err", err)
		}
	case *tradingstate.Trade:
		if err := sc.DB(db.dbName).C(tradesCollection).Remove(query); err != nil && err != mgo.ErrNotFound {
			log.Error("DeleteItemByTxHash: failed to delete trade", "txhash", txhash, "err", err)
		}
	case *lendingstate.LendingItem:
		if err := sc.DB(db.dbName).C(lendingItemsCollection).Remove(query); err != nil && err != mgo.ErrNotFound {
			log.Error("DeleteItemByTxHash: failed to delete lendingItem", "txhash", txhash, "err", err)
		}
	case *lendingstate.LendingTrade:
		if err := sc.DB(db.dbName).C(lendingTradesCollection).Remove(query); err != nil && err != mgo.ErrNotFound {
			log.Error("DeleteItemByTxHash: failed to delete lendingTrade", "txhash", txhash, "err", err)
		}
	default:
		log.Error("DeleteItemByTxHash: Unknown object type", "txhash", txhash, "object", val)
	}

}

func (db *MongoDatabase) GetListItemByTxHash(txhash common.Hash, val interface{}) interface{} {
	sc := db.Session.Copy()
	defer sc.Close()

	query := bson.M{"txHash": txhash.Hex()}
	switch val.(type) {
	case *tradingstate.OrderItem:
		result := []*tradingstate.OrderItem{}
		if err := sc.DB(db.dbName).C(ordersCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByTxHash (orders)", "err", err, "Txhash", txhash)
		}
		return result
	case *tradingstate.Trade:
		result := []*tradingstate.Trade{}
		if err := sc.DB(db.dbName).C(tradesCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByTxHash (trades)", "err", err, "Txhash", txhash)
		}
		return result
	case *lendingstate.LendingItem:
		result := []*lendingstate.LendingItem{}
		if err := sc.DB(db.dbName).C(lendingItemsCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByTxHash (lendingItems)", "err", err, "Txhash", txhash)
		}
		return result
	case *lendingstate.LendingTrade:
		result := []*lendingstate.LendingTrade{}
		if err := sc.DB(db.dbName).C(lendingTradesCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByTxHash (lendingTrades)", "err", err, "Txhash", txhash)
		}
		return result
	default:
		log.Error("GetListItemByTxHash: Unknown object type", "txhash", txhash, "object", val)
	}
	return nil
}

func (db *MongoDatabase) GetListItemByHashes(hashes []string, val interface{}) interface{} {
	sc := db.Session.Copy()
	defer sc.Close()

	query := bson.M{"hash": bson.M{"$in": hashes}}

	switch val.(type) {
	case *tradingstate.OrderItem:
		result := []*tradingstate.OrderItem{}
		if err := sc.DB(db.dbName).C(ordersCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByHashes (orders)", "err", err, "hashes", hashes)
		}
		return result
	case *tradingstate.Trade:
		result := []*tradingstate.Trade{}
		if err := sc.DB(db.dbName).C(tradesCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByHashes (trades)", "err", err, "hashes", hashes)
		}
		return result
	case *lendingstate.LendingItem:
		result := []*lendingstate.LendingItem{}
		if err := sc.DB(db.dbName).C(lendingItemsCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByHashes (lendingItems)", "err", err, "hashes", hashes)
		}
		return result
	case *lendingstate.LendingTrade:
		result := []*lendingstate.LendingTrade{}
		if err := sc.DB(db.dbName).C(lendingTradesCollection).Find(query).All(&result); err != nil && err != mgo.ErrNotFound {
			log.Error("failed to GetListItemByHashes (lendingTrades)", "err", err, "hashes", hashes)
		}
		return result
	default:
		log.Error("GetListItemByHashes: Unknown object type", "hashes", hashes, "object", val)
	}
	return nil
}

func (db *MongoDatabase) EnsureIndexes() error {
	orderHashIndex := mgo.Index{
		Key:        []string{"hash"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_order_hash",
	}
	orderTxHashIndex := mgo.Index{
		Key:        []string{"txHash"},
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_order_tx_hash",
	}
	tradeHashIndex := mgo.Index{
		Key:        []string{"hash"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_trade_hash",
	}
	tradeTxHashIndex := mgo.Index{
		Key:        []string{"txHash"},
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_trade_tx_hash",
	}
	lendingItemHashIndex := mgo.Index{
		Key:        []string{"hash"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_lending_item_hash",
	}
	lendingItemTxHashIndex := mgo.Index{
		Key:        []string{"txHash"},
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_lending_item_tx_hash",
	}
	lendingTradeHashIndex := mgo.Index{
		Key:        []string{"hash"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_lending_trade_hash",
	}
	lendingTradeTxHashIndex := mgo.Index{
		Key:        []string{"txHash"},
		DropDups:   true,
		Background: true,
		Sparse:     true,
		Name:       "index_lending_trade_tx_hash",
	}
	sc := db.Session.Copy()
	defer sc.Close()
	if err := sc.DB(db.dbName).C(ordersCollection).EnsureIndex(orderHashIndex); err != nil {
		return fmt.Errorf("failed to index orders.hash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(ordersCollection).EnsureIndex(orderTxHashIndex); err != nil {
		return fmt.Errorf("failed to index orders.txHash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(tradesCollection).EnsureIndex(tradeHashIndex); err != nil {
		return fmt.Errorf("failed to index trades.hash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(tradesCollection).EnsureIndex(tradeTxHashIndex); err != nil {
		return fmt.Errorf("failed to index trades.txHash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(lendingItemsCollection).EnsureIndex(lendingItemHashIndex); err != nil {
		return fmt.Errorf("failed to index lending_items.hash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(lendingItemsCollection).EnsureIndex(lendingItemTxHashIndex); err != nil {
		return fmt.Errorf("failed to index lending_items.txHash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(lendingTradesCollection).EnsureIndex(lendingTradeHashIndex); err != nil {
		return fmt.Errorf("failed to index lending_trades.hash . Err: %v", err)
	}
	if err := sc.DB(db.dbName).C(lendingTradesCollection).EnsureIndex(lendingTradeTxHashIndex); err != nil {
		return fmt.Errorf("failed to index lending_trades.txHash . Err: %v", err)
	}
	return nil
}

func (db *MongoDatabase) Close() {
	db.Close()
}

func (db *MongoDatabase) NewBatch() ethdb.Batch {
	// for levelDB only
	return nil
}

type keyvalue struct {
	key   []byte
	value []byte
}
type Batch struct {
	db         *MongoDatabase
	collection string
	b          []keyvalue
	size       int
}

func (b *Batch) SetCollection(collection string) {
	// for levelDB only
}

func (b *Batch) Put(key, value []byte) error {
	// for levelDB only
	return nil
}

func (b *Batch) Write() error {
	// for levelDB only
	return nil
}

func (b *Batch) ValueSize() int {
	// for levelDB only
	return int(0)
}
func (b *Batch) Reset() {
	// for levelDB only
}