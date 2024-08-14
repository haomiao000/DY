package middleware
import(
	"crypto/sha256"
	"encoding/hex"
)

func Gen_sha256(str string) string{
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hashStr := hasher.Sum(nil)
	return hex.EncodeToString(hashStr)
}