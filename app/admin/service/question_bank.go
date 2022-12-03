package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
)

type QuestionBank struct {
	service.Service
}

func (e *QuestionBank) SetQuestionPage(c *dto.QuestionBankGetQuestionReq) (m []models.QuestionBank, err error) {
	//var list []models.QuestionBank
	err = e.getList(c, &m)

	return
}

// GetQuestionBankList 获取组织数据
func (e *QuestionBank) getList(c *dto.QuestionBankGetQuestionReq, list *[]models.QuestionBank) error {
	var err error
	var data models.QuestionBank

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *QuestionBank) Get(d *dto.QuestionBankGetReq, model *models.QuestionBank) error {
	var err error

	db := e.Orm.First(model, models.QuestionBank{Id: d.GetId().(int)})

	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建QuestionBank对象
func (e *QuestionBank) Insert(c *dto.QuestionBankInsertReq) error {
	var err error
	var data models.QuestionBank
	c.Generate(&data)
	// tx := e.Orm.Debug()
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *QuestionBank) Update(c *dto.QuestionBankUpdateReq) error {
	var err error
	var data models.QuestionBank

	// tx := e.Orm.Debug()
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&data, c.GetId())
	c.Generate(&data)
	db := tx.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateQuestionBank error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysDept
func (e *QuestionBank) Remove(d *dto.QuestionBankDeleteReq) error {
	var err error
	var data models.QuestionBank

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}


