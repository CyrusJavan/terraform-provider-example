package pet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CyrusJavan/terraform-provider-example/example-server/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/prologic/bitcask"
)

type PetServer struct {
	DB *bitcask.Bitcask
}

func (s *PetServer) AddPet(p *models.Pet) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, pr runtime.Producer) {
		m, err := json.Marshal(p)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		err = s.DB.Put([]byte(strconv.FormatInt(p.ID, 10)), m)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func (s *PetServer) GetPetByID(id int64) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, pr runtime.Producer) {
		obj, err := s.DB.Get([]byte(strconv.FormatInt(id, 10)))
		if err == bitcask.ErrKeyNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(obj)
	})
}
