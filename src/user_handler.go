package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUsers is ...
func GetUsers(c *gin.Context) {
	var users []user
	allUsers := selectAll(&users, c) // convert to interface

	c.JSON(200, allUsers)
}

//Find one user
func GetUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	user, _ := selectUserID(ID, c) //_ pra ignorar o valor que vier neste campo
	c.JSON(200, user)
}

// PostUser is used create an user
func PostUser(c *gin.Context) {
	var newUser user

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbConnect()
	db.Create(&newUser)

	c.JSON(201, newUser)
}

// PostUser is used to update an user
func PutUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		checkErr(err, 400, c)
		return
	}

	var updateUser user // vai receber o body

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, db := selectUserID(ID, c)

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.BirthDate = updateUser.BirthDate

	db.Save(&user) // nao precisei criar outra conexao com o banco (db.connect) reaprovei o retorno da função

	c.JSON(200, user)
}

//DeleteUser to delete a user
func DeleteUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		checkErr(err, 400, c)
		return
	}

	user, db := selectUserID(ID, c)

	if user.ID > 0 {
		db.Delete(user)
	}

	c.JSON(204, nil)
}
