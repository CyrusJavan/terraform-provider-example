package pet

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/CyrusJavan/terraform-provider-example/example-server/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"git.mills.io/prologic/bitcask"
)

type PetServer struct {
	DB *bitcask.Bitcask
}

func (s *PetServer) AddPet(p *models.Pet) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, pr runtime.Producer) {
		if p.ID == 0 {
			p.ID = rand.Int63()
		}
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
		pb, _ := json.Marshal(p)
		_, _ = w.Write(pb)
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
		_, _ = w.Write(obj)
	})
}

func (s *PetServer) DeletePet(id int64) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, pr runtime.Producer) {
		err := s.DB.Delete([]byte(strconv.FormatInt(id, 10)))
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	})
}
