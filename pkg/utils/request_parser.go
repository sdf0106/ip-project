package utils

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/dto"
	"io"
	"net/http"
)

func RequestBodyParser[T dto.DTO](r *http.Request, dto T) (T, error) {
	reqBodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		return dto, err
	}

	if err = json.Unmarshal(reqBodyBytes, &dto); err != nil {
		return dto, err
	}

	if err = dto.Validate(); err != nil {
		return dto, err
	}

	return dto, nil
}
