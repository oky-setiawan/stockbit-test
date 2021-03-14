package log

import (
	"context"
	"encoding/json"
	"github.com/oky-setiawan/stockbit-test/internal/entity"
	log "github.com/sirupsen/logrus"
	"sync"
)

const logActionQuery = `INSERT INTO log (action, method, request, response) VALUES (:action,:method,:request,:response)`

// LogAction will log any action and store in database
func (m *logRepository) LogAction(ctx context.Context, request *entity.LogActionRequest) (err error) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		bReq, errReq := json.Marshal(request.Request)
		if errReq != nil {
			log.Errorf("[LogAction] faile Marshal, err: %v", errReq.Error())
			return
		}
		request.RequestJSON = bReq
	}()

	go func() {
		defer wg.Done()
		bResp, errResp := json.Marshal(request.Response)
		if errResp != nil {
			log.Errorf("[LogAction] faile Marshal, err: %v", errResp.Error())
			return
		}
		request.ResponseJSON = bResp
	}()

	wg.Wait()

	_, err = m.db.GetMaster().NamedExecContext(ctx, logActionQuery, request)
	return err
}
