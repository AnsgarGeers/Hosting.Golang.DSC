package dsc

import(
    "crypto/md5"
    "io"
    "log"
    "os"
)

// Apply
func Apply(res Resource) (bool, error) {
    return res.Apply()
}

// debug
func debug(res Resource, lvl int, msg string) {
    if LOG_LEVEL >= lvl {
        log.Println(res.Name(), "::", msg)
    }
}

// computeFileMd5
func computeFileMd5(filePath string) ([]byte, error) {
    var result []byte
    file, err := os.Open(filePath)
    if err != nil {
        return result, err
    }
    defer file.Close()
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return result, err
    }
    return hash.Sum(result), nil
}

// computeStringMd5
func computeStringMd5(content string) []byte {
    sum := md5.Sum([]byte(content))
    return sum[:]
}

// computeByteMd5
func computeByteMd5(content []byte) []byte {
    sum := md5.Sum(content)
    return sum[:]
}

// pathExists
func pathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}
