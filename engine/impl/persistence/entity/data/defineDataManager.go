package data

import (
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/errs"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

var defineDataManager DefineDataManager

type DefineDataManager struct {
	AbstractDataManager
}

func (DefineDataManager) SetDefineEntityManager(dataManager DefineDataManager) {
	defineDataManager = dataManager
}

func (DefineDataManager) GetDefineEntityManager() DefineDataManager {
	return defineDataManager
}

func (define DefineDataManager) FindDeployedProcessDefinitionByKey(key string) (Bytearry, error) {
	bytearries := Bytearry{}
	err := db.DB().Where("`key`=?", key).Where("deployment_id != 0").Order("version DESC", true).First(&bytearries).Error
	return bytearries, err
}

func (define DefineDataManager) GetBytearry(processDefineId int64) (Bytearry, error) {
	bytearries := Bytearry{}
	err := db.DB().Where("id=?", processDefineId).First(&bytearries).Error
	if err != nil {
		log.Infoln("Find bytearry by err", err)
		return bytearries, err
	}
	return bytearries, nil
}

func (define DefineDataManager) CreateByteArry(name string, key string, bytes string) error {
	bytearry, err := define.FindDeployedProcessDefinitionByKey(key)
	if err != nil {
		return err
	}
	var verion = 0
	verion = bytearry.Version
	verion++
	byteArry := Bytearry{Name: name, Bytes: bytes, Key: key, Version: verion}
	err = db.DB().Create(&byteArry).Error
	if err != nil {
		log.Infoln("Create bytearry err", err)
		return err
	}
	return nil
}

func (define DefineDataManager) FindProcessByTask(processInstanceId int64) (Bytearry, error) {
	bytearries := make([]Bytearry, 0)
	var sql = "SELECT b.* FROM bytearry b " +
		"LEFT JOIN process_instance p on b.id = p.process_define_id " +
		"WHERE p.id = ? "
	err := db.DB().Raw(sql, processInstanceId).Find(&bytearries).Error
	if err != nil {
		return Bytearry{}, err
	}
	if bytearries != nil && len(bytearries) > 0 {
		return bytearries[0], nil
	}
	return Bytearry{}, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}