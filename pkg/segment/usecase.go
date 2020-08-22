package segment

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Angry-Lab/api/pkg/entity"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/partyzanex/testutils"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type segmentCase struct {
	segments Repository
}

func (uc *segmentCase) Repository() Repository {
	return uc.segments
}
