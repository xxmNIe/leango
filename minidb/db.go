package minidb

import (
	"io"
	"os"
	"sync"
)

type MiniDB struct {
	indexes map[string]int64
	dbFile  *DBFile
	dirPath string
	mu      sync.RWMutex
}

func Open(dirPath string) (*MiniDB, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	dbFile, err := NewDBFile(dirPath)
	if err != nil {
		return nil, err
	}
	db := &MiniDB{
		dbFile:  dbFile,
		indexes: make(map[string]int64, 0),
		dirPath: dirPath,
	}
	db.loadIndexesFromFile()
	return db, nil
}
func (db *MiniDB) loadIndexesFromFile() error {
	if db.dbFile.Offset == 0 {
		return nil
	}
	var (
		validEntries []*Entry
		offset       int64
	)

	for {
		e, err := db.dbFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if off, ok := db.indexes[string(e.Key)]; ok && off == offset {
			validEntries = append(validEntries, e)
		}
		offset += e.GetSize()
	}

	if len(validEntries) > 0 {
		mergeDBFile, err := NewMegerFile(db.dirPath)
		if err != nil {
			return err
		}
		defer os.Remove(mergeDBFile.File.Name())

		for _, entry := range validEntries {
			writeOff := mergeDBFile.Offset
			err := mergeDBFile.Write(entry)
			if err != nil {
				return err
			}
			db.indexes[string(entry.Key)] = writeOff
		}

		dbFileName := db.dbFile.File.Name()

		db.dbFile.File.Close()
		os.Remove(dbFileName)

		mergeDBFileName := mergeDBFile.File.Name()
		mergeDBFile.File.Close()

		os.Rename(mergeDBFileName, db.dirPath+string(os.PathSeparator)+dbFileName)
		db.dbFile = mergeDBFile
	}
	return nil

}

func (db *MiniDB) PUT(key, value []byte) (err error) {
	if len(key) == 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	offset := db.dbFile.Offset
	entry := &Entry{Key: key, Value: value, Mark: PUT}

	err = db.dbFile.Write(entry)
	if err == nil {
		db.indexes[string(key)] = offset
	}
	return
}

func (db *MiniDB) GET(key []byte) (val []byte, err error) {
	if len(key) == 0 {
		return
	}
	offset, ok := db.indexes[string(key)]
	if !ok {
		return
	}
	var e *Entry
	e, err = db.dbFile.Read(offset)
	if err != nil && err != io.EOF {
		return
	}
	if e != nil {
		val = e.Value
	}
	return
}

func (db *MiniDB) Del(key []byte) (err error) {
	if len(key) == 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.indexes[string(key)]
	if !ok {
		return
	}
	e := &Entry{Key: key, Value: nil, Mark: DEL}

	err = db.dbFile.Write(e)
	if err != nil {
		return
	}
	delete(db.indexes, string(key))
	return
}
