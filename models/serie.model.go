package models

import (
	"errors"
	// "gopkg.in/mgo.v2"
	// // "gopkg.in/mgo.v2/bson"
	// "fmt"
)

var (
	Array map[string]*Serie
)

type Serie struct {
	Name   string
	Year      int64
	Platform string
}



func init() {
	
	/*Connection to database*/ 
	// session, err := mgo.Dial("localhost:27017")
	// if err != nil {
	// 	fmt.Println("Could not connect to mongo: ", err.Error())
	
	// }
	// defer session.Close()
	// c := session.DB("test").C("Series")
	// err = c.Insert(&Serie {"WestWorld", 2016, "HBO"})
	// if err != nil {
	// 	fmt.Println("Error creating Serie: ", err.Error())
	
	// }


	Array = make(map[string]*Serie)
	Array["Westworld"] = &Serie{"Dexter", 2006, "Netflix"}
	Array["Dexter"] = &Serie{"Dexter", 2006, "Netflix"}
}

func AddOne(serie Serie) (serieName string) {
	Array[serie.Name] =&serie
	return serie.Name
}

func GetOne(SerieName string) (serie *Serie, err error) {
	if v, ok := Array[SerieName]; ok {
		return v, nil
	}
	return nil, errors.New("THis serie is not stored")
}

func GetAll() map[string]*Serie {
	return Array
}

// func Update(ObjectId string, Score int64) (err error) {
// 	if v, ok := Objects[ObjectId]; ok {
// 		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("ObjectId Not Exist")
// }

func Delete(SerieName string) {
	delete(Array, SerieName)
}

