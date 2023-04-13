package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailru/activerecord-cookbook/example/model/dictionary"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/promoperiods"
	"github.com/mailru/activerecord/pkg/activerecord"
)

const android = "android"

/*
type myLog struct {
}

func (l *myLog) GetLoggerFromContext(ctx context.Context) activerecord.LoggerInterface {
	return &myLog{}
}
func (l *myLog) SetLogLevel(level uint32) {
	log.Printf("SetLogLevel %d in mylog", level)
}
func (l *myLog) Fatal(args ...interface{}) {
	log.Print("Fatal in mylog", args)
}
func (l *myLog) Error(args ...interface{}) {
	log.Print("Error in mylog", args)
}
func (l *myLog) Warn(args ...interface{}) {
	log.Print("Warn in mylog", args)
}
func (l *myLog) Info(args ...interface{}) {
	log.Print("Info in mylog", args)
}
func (l *myLog) Debug(args ...interface{}) {
	log.Print("Debug in mylog", args)
}
*/

//nolint:gocognit
func main() { //nolint:gocyclo
	ctx := context.Background()
	//	requestId := "123456"
	//	ctx = context.WithValue(ctx, "requestId", requestId)

	activerecord.InitActiveRecord(
		activerecord.WithConfig(activerecord.NewDefaultConfigFromMap(map[string]interface{}{
			"box1.Master":   "127.0.0.1",
			"box1.Port":     "11011",
			"box1.PoolSize": 1,
			"box1.Timeout":  time.Millisecond * 200,
		})),
	)
	//arInst.Logger = &myLog{}
	activerecord.Logger().SetLogLevel(activerecord.DebugLoggerLevel)

	product, err := dictionary.GetProductByID(2)
	if err != nil {
		log.Fatal("Product not found")
	}

	platform := map[string]interface{}{"mobile": "ios"}

	promoperiod := promoperiods.New(ctx)

	if err = promoperiod.SetAction("10"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetCode("codecodecode"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetEmail("email1@domain.tld"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetFinish(product); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetID("3422b448-2460-4fd2-9183-8000de6f8342"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetPlatform(platform); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetStart(123123); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.Insert(ctx); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	logger := activerecord.Logger()

	selected, err := promoperiods.SelectByID(ctx, "3422b448-2460-4fd2-9183-8000de6f8342")
	if err != nil {
		logger.Error(ctx, "Error select ", err)
	} else {
		logger.Debug(ctx, selected.GetID())
		logger.Debug(ctx, selected.GetAction())
		logger.Debug(ctx, selected.GetPlatform())
		logger.Debug(ctx, "Flag", fmt.Sprintf("%0.32b", selected.GetPromobunch()))
		pl := selected.GetPlatform()

		if pl["mobile"] == android {
			pl["mobile"] = "ios"
		} else {
			pl["mobile"] = android
		}
		if selected.GetPromobunch()&(1<<3) == 1<<3 {
			logger.Debug(ctx, "Flag set 4")

			if err = selected.SetBitPromobunch(1 << 4); err != nil {
				log.Fatalf("Err Set %android", err)
			}

			if err = selected.ClearBitPromobunch(1 << 3); err != nil {
				log.Fatalf("Err Set %android", err)
			}
		} else {
			logger.Debug(ctx, "Flag set 3")

			if err = selected.SetBitPromobunch(1 << 3); err != nil {
				log.Fatalf("Err Set %android", err)
			}

			if err = selected.ClearBitPromobunch(1 << 4); err != nil {
				log.Fatalf("Err Set %android", err)
			}
		}
		logger.Debug(ctx, "Flag", fmt.Sprintf("%0.32b", selected.GetPromobunch()))

		if err = selected.SetPlatform(pl); err != nil {
			log.Fatalf("Err Set %android", err)
		}

		if err = selected.SetAction("11"); err != nil {
			log.Fatalf("Err Set %android", err)
		}

		if err = selected.Update(ctx); err != nil {
			log.Fatalf("Err Set %android", err)
		}

		logger.Debug(ctx, "Flag", fmt.Sprintf("%0.32b", selected.GetPromobunch()))
		//selected.Delete(ctx)
	}

	promoperiod = promoperiods.New(ctx)
	if err = promoperiod.SetAction("10"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetCode("codecodecode"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetEmail("email2@domain.tld"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetFinish(product); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetID("3422b448-2460-4fd2-9183-8000de6f8340"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetPlatform(platform); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetStart(123123); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.InsertOrReplace(ctx); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	promoperiod = promoperiods.New(ctx)
	if err = promoperiod.SetAction("10"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetCode("codecodecode"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetEmail("email3@domain.tld"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetFinish(product); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetID("3422b448-2460-4fd2-9183-8000de6f8350"); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetPlatform(platform); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.SetStart(123123); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	if err = promoperiod.InsertOrReplace(ctx); err != nil {
		log.Fatalf("Err Set %android", err)
	}

	selected, err = promoperiods.SelectByEmailAction(ctx, promoperiods.EmailActionIndexType{Email: "email3@domain.tld", Action: "10"})
	if selected != nil {
		logger.Debug(ctx, selected.GetEmail())
	} else if err != nil {
		logger.Error(ctx, "Select error %android", err)
	} else {
		logger.Error(ctx, "Select result is empty")
	}

	selectedm, _ := promoperiods.SelectByIDs(ctx, []string{"3422b448-2460-4fd2-9183-8000de6f8340", "3422b448-2460-4fd2-9183-8000de6f8350"})
	logger.Debug(ctx, selectedm[0].GetEmail())

	selectedm, _ = promoperiods.SelectByCode(ctx, "codecodecode", activerecord.NewLimiter(1))
	logger.Debug(ctx, selectedm[0].GetFinish())
	time.Sleep(time.Second * 10)
}
