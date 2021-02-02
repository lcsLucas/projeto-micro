package migrations

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/lcslucas/projeto-micro/services/prova/model"
	"github.com/lcslucas/projeto-micro/utils"
	"gorm.io/gorm"
)

func ExecCreateDatabaseProva(ctx context.Context, database string, db *gorm.DB) error {
	var err error
	var result int

	tx := db.WithContext(ctx)

	err = tx.Debug().Raw("SELECT COUNT(*) FROM pg_catalog.pg_database WHERE lower(datname) = lower(?)", database).Scan(&result).Error
	if err != nil {
		return err
	}

	//return errors.New(fmt.Sprintf("****** Result: %s", result))

	if int(result) < 1 {
		err = tx.Debug().Exec("CREATE DATABASE " + database + " WITH OWNER 'postgres' ENCODING 'UTF8'").Error
		if err != nil {
			return err
		}

		log.Println(`Database ` + database + ` gerada com sucesso.`)
	}

	if err != nil {
		_ = tx.Debug().Exec("SET TIMEZONE TO 'America/Sao_Paulo'").Error
	}

	return nil
}

func ExecMigrationProva(ctx context.Context, database string, db *gorm.DB) error {
	var err error
	statusImport := false

	tx := db.WithContext(ctx).Begin()

	/* Resetar todas as tabelas no postgres */
	/*
		tx.Migrator().DropTable(model.Prova{})
		tx.Migrator().DropTable(model.ProvaExercicios{})
		tx.Migrator().DropTable(model.ProvaAluno{})

		tx.Commit()
		return errors.New("Teste")
	*/

	check := tx.Migrator().HasTable(model.Prova{})
	if !check {
		statusImport = true
		err = tx.Debug().Migrator().AutoMigrate(model.Prova{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "provas" gerada com sucesso.`)
		}
	}

	check = tx.Migrator().HasTable(model.ProvaExercicios{})
	if !check {
		err = tx.Debug().Migrator().AutoMigrate(model.ProvaExercicios{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "provas_exercicios" gerada com sucesso.`)
		}
	}

	check = tx.Migrator().HasTable(model.ProvaAluno{})
	if !check {
		err = tx.Debug().Migrator().AutoMigrate(model.ProvaAluno{})
		if err != nil {
			return err
		} else {
			log.Println(`Tabela de "provas_alunos" gerada com sucesso.`)
		}
	}

	if statusImport && err == nil {

		provas_aleatorias := []model.Prova{
			model.Prova{
				ID:           1,
				DataCadastro: time.Now(),
				DataInicio:   time.Now(),
				DataFinal:    time.Now(),
				Nome:         "Prova 1º Bimestre",
				Serie:        "3º Ano do ensino médio",
				Bimestre:     1,
				Finalizada:   false,
			},
			model.Prova{
				ID:           2,
				DataCadastro: time.Now(),
				DataInicio:   time.Now(),
				DataFinal:    time.Now(),
				Nome:         "Prova 1º Bimestre",
				Serie:        "2º Ano do ensino médio",
				Bimestre:     1,
				Finalizada:   false,
			},
			model.Prova{
				ID:           3,
				DataCadastro: time.Now(),
				DataInicio:   time.Now(),
				DataFinal:    time.Now(),
				Nome:         "Prova 1º Bimestre",
				Serie:        "1º Ano do ensino médio",
				Bimestre:     1,
				Finalizada:   false,
			},
			model.Prova{
				ID:           4,
				DataCadastro: time.Now(),
				DataInicio:   time.Now(),
				DataFinal:    time.Now(),
				Nome:         "Prova 1º Bimestre",
				Serie:        "9º Ano do ensino fundamental",
				Bimestre:     1,
				Finalizada:   false,
			},
			model.Prova{
				ID:           5,
				DataCadastro: time.Now(),
				DataInicio:   time.Now(),
				DataFinal:    time.Now(),
				Nome:         "Prova 1º Bimestre",
				Serie:        "8º Ano do ensino fundamental",
				Bimestre:     1,
				Finalizada:   false,
			},
		}

		for _, p := range provas_aleatorias {

			err := tx.Debug().Create(&p).Error

			if err != nil {
				tx.Rollback()
				return err
			}

			provas_alu := []model.ProvaAluno{
				model.ProvaAluno{
					ProvaID: p.ID,
					AlunoRA: "16.128.614-8",
				},
				model.ProvaAluno{
					ProvaID: p.ID,
					AlunoRA: "10.846.074-5",
				},
				model.ProvaAluno{
					ProvaID: p.ID,
					AlunoRA: "40.759.343-3",
				},
				model.ProvaAluno{
					ProvaID: p.ID,
					AlunoRA: "40.738.017-6",
				},
				model.ProvaAluno{
					ProvaID: p.ID,
					AlunoRA: "35.329.394-5",
				},
			}

			err = tx.Debug().Create(&provas_alu).Error

			if err != nil {
				tx.Rollback()
				return err
			}

			var prova_exers []model.ProvaExercicios

			for _, proAlu := range provas_alu {
				var num_sortiados []uint64
				var rID uint64

				for i := 0; i < 5; i++ {
					flag_ok := false

					for !flag_ok {
						rID = uint64(rand.Intn(9-1) + 1)

						if !utils.InSlice(uint64(rID), num_sortiados) {
							flag_ok = true
							num_sortiados = append(num_sortiados, rID)
						}

					}

					prova_exers = append(prova_exers, model.ProvaExercicios{
						ProvaAlunoID: proAlu.ID,
						ExercicioID:  uint64(rID),
					})

				}

			}

			err = tx.Debug().Create(&prova_exers).Error

			if err != nil {
				tx.Rollback()
				return err
			}

		}

	}

	tx.Commit()

	return nil
}
