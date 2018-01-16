package models

import (
	"errors"
	 "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	 "fmt"
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
	 session, err := mgo.Dial("localhost:27017")
	 if err != nil {
	 	fmt.Println("Could not connect to mongo: ", err.Error())
	
	 }
	 defer session.Close()
	 c := session.DB("test").C("Series")
	 err = c.Insert(&Serie {"WestWorld", 2016, "HBO"}, &Serie {"Peaky Blinders", 2010, "Netflix"})
	 if err != nil {
	 	fmt.Println("Error creating Serie: ", err.Error())
	
	 }
}

func AddOne(serie Serie) (serieName string) {
	/*Connection to database*/ 
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to mongo: ", err.Error())
   
	}
	defer session.Close()
	c := session.DB("test").C("Series")
	err = c.Insert(serie)
	if err != nil {
		fmt.Println("Error creating Serie: ", err.Error())
	}
	return serie.Name
}

func GetOne(SerieName string) (serie Serie, err error) {
	
	/*Connection to database*/ 
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to mongo: ", err.Error())
   
	}
	defer session.Close()
	c := session.DB("test").C("Series")
	result := Serie{}
	err = c.Find(bson.M{"name": SerieName}).One(&result)
	if err != nil {
		return Serie{}, errors.New("THis serie is not stored")
	}
	fmt.Println("Phone", result)
	return result, nil
}

func GetAll() []Serie {

	/*Connection to database*/ 
	 session, err := mgo.Dial("localhost:27017")
	 if err != nil {
	 	fmt.Println("Could not connect to mongo: ", err.Error())
	
	 }
	 defer session.Close()
	 c := session.DB("test").C("Series")

	 var results []Serie
	 err = c.Find(nil).All(&results)
 
	 if err != nil {
		 panic(err)
	 }
	 fmt.Println("Results All: ", results)
	return results
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

