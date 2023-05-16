package main

import (
	"fmt"
)

func main() {
	// // initialize db
	// dal, _ := newDal("db.db")

	// // create a new page
	// p := dal.allocateEmptyPage()
	// p.num = dal.getNextPage()
	// copy(p.data[:], "hello,")

	// // commit it
	// _ = dal.writePage(p)
	// _, _ = dal.writeFreelist()

	// // Close the db
	// _ = dal.close()

	// // We expect the freelist state was saved, so we write to
	// // page number 3 and not overwrite the one at number 2
	// dal, _ = newDal("db.db")
	// p = dal.allocateEmptyPage()
	// p.num = dal.getNextPage()
	// copy(p.data[:], "world!")
	// _ = dal.writePage(p)

	// // Create a page and free it so the released pages will be updated
	// pageNum := dal.getNextPage()
	// dal.releasePage(pageNum)

	// // commit it
	// _, _ = dal.writeFreelist()
	// // dal, _ := newDal("mainTest")
	// // node, _ := dal.getNode(dal.root)
	// // node.dal = dal
	// // index, containingNode, _ := node.findKey([]byte("Key1"))
	// // res := containingNode.items[index]

	// // fmt.Printf("the key is: %s, value is: %s", res.key, res.value)

	// // _ = dal.close()

	// options := &Options{
	// 	pageSize:       os.Getpagesize(),
	// 	MinFillPercent: 0.0125,
	// 	MaxFillPercent: 0.025,
	// }

	// dal, _ := newDal("./mainTest", options)

	// c := newCollection([]byte("collection1"), dal.root)

	// c.dal = dal
	// _ = c.Put([]byte("key1"), []byte("value1"))
	// _ = c.Put([]byte("key2"), []byte("value2"))
	// _ = c.Put([]byte("key3"), []byte("value3"))
	// _ = c.Put([]byte("key4"), []byte("value4"))
	// _ = c.Put([]byte("key5"), []byte("value5"))
	// _ = c.Put([]byte("key6"), []byte("value6"))
	// _ = c.Put([]byte("key7"), []byte("value7"))
	// _ = c.Put([]byte("key8"), []byte("value8"))
	// _ = c.Put([]byte("key9"), []byte("value9"))
	// _ = c.Put([]byte("key10"), []byte("value10"))
	// item, _ := c.Find([]byte("key1"))

	// fmt.Printf("key is: %s, value is: %s\n", item.key, item.value)

	// _ = c.Remove([]byte("key1"))
	// item, _ = c.Find([]byte("key1"))

	// dal.writeFreelist()
	// fmt.Printf("item is:%+v\n", item)

	// _ = dal.close()

	db, _ := Open("Demo7", &Options{MinFillPercent: 0.5, MaxFillPercent: 1.0})

	tx := db.WriteTx()
	collectionName := "Demo7Collection"
	createdCollection, _ := tx.CreateCollection([]byte(collectionName))

	newKey := []byte("key0")
	newVal := []byte("value0")
	_ = createdCollection.Put(newKey, newVal)

	_ = tx.Commit()
	_ = db.Close()

	db, _ = Open("Demo7", &Options{MinFillPercent: 0.5, MaxFillPercent: 1.0})
	tx = db.ReadTx()
	createdCollection, _ = tx.GetCollection([]byte(collectionName))

	item, _ := createdCollection.Find(newKey)

	_ = tx.Commit()
	_ = db.Close()

	fmt.Printf("key is: %s, value is: %s\n", item.key, item.value)

}
