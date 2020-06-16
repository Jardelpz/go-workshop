package src

import "github.com/gin-gonic/gin"

//PostDebt to create a debt
func PostDebt(c *gin.Context) {
	var newDebt debt

	if err := c.ShouldBindJSON(&newDebt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, _ := selectUserID(newDebt.UserRefer, c) // pegar o usuario e enviar o id, userrefer esta retornando 0
	if user.ID == 0 {
		c.JSON(400, gin.H{"error": newDebt})
		return
	}
	// user.Debts = append(user.Debts, newDebt)
	// db2.Save(&user)

	db := dbConnect()

	db.Create(&newDebt)
	c.JSON(201, newDebt)

}

func GetDebt(c *gin.Context) {
	ID := c.Param("id")
	debt, _ := selecDebtId(ID, c)
	c.JSON(200, debt)
}

func GetDebts(c *gin.Context) {
	var debts []debt

	allDebts := selectAll(&debts, c)
	c.JSON(200, allDebts)
}

func PutDebt(c *gin.Context) {
	ID := c.Param("id")

	var debtUpdate debt
	if err := c.ShouldBindJSON(&debtUpdate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	debt, db := selecDebtId(ID, c)
	//
	user, _ := selectUserID(debt.UserRefer, c) // pegar o usuario e enviar o id
	if user.ID == 0 {
		c.JSON(400, gin.H{"error": "No user found"})
		return
	}
	//user.Debts = append(user.Debts, debt)

	//db2.Save(&user)
	//
	debt.UserRefer = debtUpdate.UserRefer
	debt.Company_name = debtUpdate.Company_name
	debt.Value = debtUpdate.Value
	debt.Date = debtUpdate.Date
	debt.Status = debtUpdate.Status

	db.Save(&debt)

	c.JSON(200, debt)
}

func DeleteDebt(c *gin.Context) {
	ID := c.Param("id")
	debt, db := selecDebtId(ID, c)
	// if debt.ID != nil {
	db.Delete(debt)
	// }

	c.JSON(204, nil)

}
