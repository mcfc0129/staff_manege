package addStaff

import(
  "crypto/sha256"
)

func GetSHA256Binary(s string) []byte {
    r := sha256.Sum256([]byte(s))
    return r[:]
}
